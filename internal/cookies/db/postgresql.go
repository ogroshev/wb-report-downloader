package cookiesdb

import (
	"context"
	"wb-report-downloader/internal/cookies"
	"wb-report-downloader/pkg/client/postgresql"
)

type repository struct {
	client postgresql.Client
}

func (r *repository) GetCookies (ctx context.Context, sellerID uint64) (cookies.Cookies, error) {
	q := `
		SELECT 
		u.x_supplier_id,
		ua.token_value
	FROM sellers s
	JOIN seller_user u ON s.id = u.seller_id
	JOIN seller_user_auth ua ON u.id = ua.user_id
	JOIN (SELECT u.seller_id, max(sua.update_date) as newest
		FROM seller_user u
		JOIN seller_user_auth sua ON u.id = sua.user_id
		WHERE token_type = 'seller_token'
		GROUP BY u.seller_id
		) s1 on s1.newest = ua.update_date
	WHERE token_type = 'seller_token'
	AND u.seller_id = $1
	;
	`

	var c cookies.Cookies
	err := r.client.QueryRow(ctx, q, sellerID).Scan(&c.XSupplierID, &c.WBToken)
	if err != nil {
		return cookies.Cookies{}, err
	}
	return c, nil
}

func NewRepository(client postgresql.Client) cookies.Repository {
	return &repository{
		client: client,
	}
}
