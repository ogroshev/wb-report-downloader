package db

import (
	"context"
	"fmt"

	"wb-report-downloader/internal/seller"
	"wb-report-downloader/pkg/client/postgresql"
)

type repository struct {
	client postgresql.Client
}

func (r *repository) GetSellers(ctx context.Context) ([]seller.Seller, error) {
	const(
		kTokenType = "seller_token"
	)

	q := fmt.Sprintf(`
		SELECT 
		u.seller_id,
		s.name,
		u.id,
		u.user_name,
		u.x_supplier_id,
		u.x_user_id,
		ua.token_value
	FROM sellers s
	JOIN seller_user u ON s.id = u.seller_id
	JOIN seller_user_auth ua ON u.id = ua.user_id
	JOIN (SELECT u.seller_id, max(sua.update_date) as newest
		FROM seller_user u
		JOIN seller_user_auth sua ON u.id = sua.user_id
		WHERE token_type = '%s'
		GROUP BY u.seller_id
		) s1 on s1.newest = ua.update_date
	WHERE token_type = '%s'
	 `, kTokenType, kTokenType)

	rows, err := r.client.Query(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sellers []seller.Seller
	for rows.Next() {
		var s seller.Seller
		err := rows.Scan(&s.ID, &s.Name, &s.UserID, &s.UserName, &s.XSupplierID, &s.XUserID, &s.WBToken)
		if err != nil {
			return nil, err
		}
		sellers = append(sellers, s)
	}
	return sellers, nil
}


func NewRepository(client postgresql.Client) seller.Repository {
	return &repository{
		client: client,
	}
}
