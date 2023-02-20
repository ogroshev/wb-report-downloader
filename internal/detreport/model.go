package detreport

import "reflect"

type DetailedReport struct {
	Data    reflect.Value
	IsEmpty bool
}

var repKinds = []any{ReportRowV1{}, ReportRowV2{}, ReportRowV3{}, ReportRowV4{}, ReportRowV5{}, ReportRowV6{}}

type ReportRowV6 struct {
	Number                                                            uint64  `db:"№"                                   xlsx:"№"`
	SupplyNumber                                                      uint64  `db:"Номер поставки"                      xlsx:"Номер поставки"`
	Subject                                                           string  `db:"Предмет"                             xlsx:"Предмет"`
	ItemCode                                                          uint64  `db:"Код номенклатуры"                    xlsx:"Код номенклатуры"`
	Brand                                                             string  `db:"Бренд"                               xlsx:"Бренд"`
	SuppliersArticle                                                  string  `db:"Артикул поставщика"                  xlsx:"Артикул поставщика"`
	Name                                                              string  `db:"Название"                            xlsx:"Название"`
	Size                                                              string  `db:"Размер"                              xlsx:"Размер"`
	Barcode                                                           uint64  `db:"Баркод"                              xlsx:"Баркод"`
	DocumentType                                                      string  `db:"Тип документа"                       xlsx:"Тип документа"`
	JustificationForPayment                                           string  `db:"Обоснование для оплаты"              xlsx:"Обоснование для оплаты"`
	DateOfTheOrderByTheBuyer                                          string  `db:"Дата заказа покупателем"             xlsx:"Дата заказа покупателем"`
	DateOfSale                                                        string  `db:"Дата продажи"                        xlsx:"Дата продажи"`
	Count                                                             uint64  `db:"Кол-во"                              xlsx:"Кол-во"`
	RetailPrice                                                       float64 `db:"Цена розничная"                      xlsx:"Цена розничная"`
	WildberriesSoldTheProductPr                                       float64 `db:"Вайлдберриз реализовал Товар (Пр)"   xlsx:"Вайлдберриз реализовал Товар (Пр)"`
	AgreedGroceryDiscount                                             float64 `db:"Согласованный продуктовый дискон"    xlsx:"Согласованный продуктовый дисконт, %"`
	Promocode                                                         float64 `db:"Промокод %"                          xlsx:"Промокод %"`
	FinalAgreedDiscount                                               float64 `db:"Итоговая согласованная скидка"       xlsx:"Итоговая согласованная скидка"`
	RetailPriceIncludingTheAgreedDiscount                             float64 `db:"Цена розничная с учетом согласова"   xlsx:"Цена розничная с учетом согласованной скидки"`
	RegularCustomerDiscountSPP                                        float64 `db:"Скидка постоянного Покупателя (СП"   xlsx:"Скидка постоянного Покупателя (СПП)"`
	KVVSizePercent                                                    float64 `db:"Размер кВВ, %"                       xlsx:"Размер кВВ, %"`
	KVVSizeWithoutVATBasicPercent                                     float64 `db:"Размер  кВВ без НДС, % Базовый"      xlsx:"Размер  кВВ без НДС, % Базовый"`
	FinalKVVWithoutVATPercent                                         float64 `db:"Итоговый кВВ без НДС, %"             xlsx:"Итоговый кВВ без НДС, %"`
	RemunerationFromSalesBeforeDeductionDfAttorneysServicesWithoutVAT float64 `db:"Вознаграждение с продаж до вычета "  xlsx:"Вознаграждение с продаж до вычета услуг поверенного, без НДС"`
	WildberriesRemunerationBBWithoutVAT                               float64 `db:"Вознаграждение Вайлдберриз (ВВ), б"  xlsx:"Вознаграждение Вайлдберриз (ВВ), без НДС"`
	VATOnWildberriesRemuneration                                      float64 `db:"НДС с Вознаграждения Вайлдберриз"    xlsx:"НДС с Вознаграждения Вайлдберриз"`
	ToTransferToTheSellerForTheSoldGoods                              float64 `db:"К перечислению Продавцу за реализ"   xlsx:"К перечислению Продавцу за реализованный Товар"`
	NumberOfDeliveries                                                uint64  `db:"Количество доставок"                 xlsx:"Количество доставок"`
	RefundAmount                                                      uint64  `db:"Количество возврата"                 xlsx:"Количество возврата"`
	ServicesForTheDeliveryOfGoodsToTheBuyer                           float64 `db:"Услуги по доставке товара покупат"   xlsx:"Услуги по доставке товара покупателю"`
	Fines                                                             float64 `db:"Общая сумма штрафов"                 xlsx:"Общая сумма штрафов"`
	Surcharges                                                        float64 `db:"Доплаты"                             xlsx:"Доплаты"`
	JustificationOfFinesAndSurcharges                                 string  `db:"Обоснование штрафов и доплат"        xlsx:"Виды логистики, штрафов и доплат"`
	StickerMP                                                         uint64  `db:"Стикер МП"                           xlsx:"Стикер МП"`
	OfficeNumber                                                      uint64  `db:"Номер офиса"                         xlsx:"Номер офиса"`
	NameOfTheDeliveryOffice                                           string  `db:"Наименование офиса доставки"         xlsx:"Наименование офиса доставки"`
	PartnersINN                                                       uint64  `db:"ИНН партнера"                        xlsx:"ИНН партнера"`
	Partner                                                           string  `db:"Партнер"                             xlsx:"Партнер"`
	Warehouse                                                         string  `db:"Склад"                               xlsx:"Склад"`
	Country                                                           string  `db:"Страна"                              xlsx:"Страна"`
	TypeOfBoxes                                                       string  `db:"Тип коробов"                         xlsx:"Тип коробов"`
	CustomsDeclarationNumber                                          string  `db:"Номер таможенной декларации"         xlsx:"Номер таможенной декларации"`
	ShK                                                               uint64  `db:"ШК"                                  xlsx:"ШК"`
	Rid                                                               uint64  `db:"Rid"                                 xlsx:"Rid"`
	Srid                                                              string  `db:"Srid"                                xlsx:"Srid"`
	RefundForTheDeliveryAndReturnOfGoodsToThePVZ                      float64 `db:"Возмещение за выдачу и возврат тов" xlsx:"Возмещение за выдачу и возврат товаров на ПВЗ"`
	ReimbursementOfAcquiringCosts                                     float64 `db:"Возмещение расходов по эквайрингу" xlsx:"Возмещение расходов по эквайрингу"`
	NameOfAcquiringBank                                               string  `db:"Наименование банка эквайринга" xlsx:"Наименование банка эквайринга"`
	MarkingCode                                                       string  `db:"Код маркировки"                      xlsx:"Код маркировки"`
	ReportID                                                          uint64  `db:"report_id"`
	SellerID                                                          uint64  `db:"seller_id"`
}
type ReportRowV5 struct {
	Number                                                            uint64  `db:"№"                                   xlsx:"№"`
	SupplyNumber                                                      uint64  `db:"Номер поставки"                      xlsx:"Номер поставки"`
	Subject                                                           string  `db:"Предмет"                             xlsx:"Предмет"`
	ItemCode                                                          uint64  `db:"Код номенклатуры"                    xlsx:"Код номенклатуры"`
	Brand                                                             string  `db:"Бренд"                               xlsx:"Бренд"`
	SuppliersArticle                                                  string  `db:"Артикул поставщика"                  xlsx:"Артикул поставщика"`
	Name                                                              string  `db:"Название"                            xlsx:"Название"`
	Size                                                              string  `db:"Размер"                              xlsx:"Размер"`
	Barcode                                                           uint64  `db:"Баркод"                              xlsx:"Баркод"`
	DocumentType                                                      string  `db:"Тип документа"                       xlsx:"Тип документа"`
	JustificationForPayment                                           string  `db:"Обоснование для оплаты"              xlsx:"Обоснование для оплаты"`
	DateOfTheOrderByTheBuyer                                          string  `db:"Дата заказа покупателем"             xlsx:"Дата заказа покупателем"`
	DateOfSale                                                        string  `db:"Дата продажи"                        xlsx:"Дата продажи"`
	Count                                                             uint64  `db:"Кол-во"                              xlsx:"Кол-во"`
	RetailPrice                                                       float64 `db:"Цена розничная"                      xlsx:"Цена розничная"`
	WildberriesSoldTheProductPr                                       float64 `db:"Вайлдберриз реализовал Товар (Пр)"   xlsx:"Вайлдберриз реализовал Товар (Пр)"`
	AgreedGroceryDiscount                                             float64 `db:"Согласованный продуктовый дискон"    xlsx:"Согласованный продуктовый дисконт, %"`
	Promocode                                                         float64 `db:"Промокод %"                          xlsx:"Промокод %"`
	FinalAgreedDiscount                                               float64 `db:"Итоговая согласованная скидка"       xlsx:"Итоговая согласованная скидка"`
	RetailPriceIncludingTheAgreedDiscount                             float64 `db:"Цена розничная с учетом согласова"   xlsx:"Цена розничная с учетом согласованной скидки"`
	RegularCustomerDiscountSPP                                        float64 `db:"Скидка постоянного Покупателя (СП"   xlsx:"Скидка постоянного Покупателя (СПП)"`
	KVVSizePercent                                                    float64 `db:"Размер кВВ, %"                       xlsx:"Размер кВВ, %"`
	KVVSizeWithoutVATBasicPercent                                     float64 `db:"Размер  кВВ без НДС, % Базовый"      xlsx:"Размер  кВВ без НДС, % Базовый"`
	FinalKVVWithoutVATPercent                                         float64 `db:"Итоговый кВВ без НДС, %"             xlsx:"Итоговый кВВ без НДС, %"`
	RemunerationFromSalesBeforeDeductionDfAttorneysServicesWithoutVAT float64 `db:"Вознаграждение с продаж до вычета "  xlsx:"Вознаграждение с продаж до вычета услуг поверенного, без НДС"`
	WildberriesRemunerationBBWithoutVAT                               float64 `db:"Вознаграждение Вайлдберриз (ВВ), б"  xlsx:"Вознаграждение Вайлдберриз (ВВ), без НДС"`
	VATOnWildberriesRemuneration                                      float64 `db:"НДС с Вознаграждения Вайлдберриз"    xlsx:"НДС с Вознаграждения Вайлдберриз"`
	ToTransferToTheSellerForTheSoldGoods                              float64 `db:"К перечислению Продавцу за реализ"   xlsx:"К перечислению Продавцу за реализованный Товар"`
	NumberOfDeliveries                                                uint64  `db:"Количество доставок"                 xlsx:"Количество доставок"`
	RefundAmount                                                      uint64  `db:"Количество возврата"                 xlsx:"Количество возврата"`
	ServicesForTheDeliveryOfGoodsToTheBuyer                           float64 `db:"Услуги по доставке товара покупат"   xlsx:"Услуги по доставке товара покупателю"`
	Fines                                                             float64 `db:"Общая сумма штрафов"                 xlsx:"Общая сумма штрафов"`
	Surcharges                                                        float64 `db:"Доплаты"                             xlsx:"Доплаты"`
	JustificationOfFinesAndSurcharges                                 string  `db:"Обоснование штрафов и доплат"        xlsx:"Обоснование штрафов и доплат"`
	StickerMP                                                         uint64  `db:"Стикер МП"                           xlsx:"Стикер МП"`
	OfficeNumber                                                      uint64  `db:"Номер офиса"                         xlsx:"Номер офиса"`
	NameOfTheDeliveryOffice                                           string  `db:"Наименование офиса доставки"         xlsx:"Наименование офиса доставки"`
	PartnersINN                                                       uint64  `db:"ИНН партнера"                        xlsx:"ИНН партнера"`
	Partner                                                           string  `db:"Партнер"                             xlsx:"Партнер"`
	Warehouse                                                         string  `db:"Склад"                               xlsx:"Склад"`
	Country                                                           string  `db:"Страна"                              xlsx:"Страна"`
	TypeOfBoxes                                                       string  `db:"Тип коробов"                         xlsx:"Тип коробов"`
	CustomsDeclarationNumber                                          string  `db:"Номер таможенной декларации"         xlsx:"Номер таможенной декларации"`
	ShK                                                               uint64  `db:"ШК"                                  xlsx:"ШК"`
	Rid                                                               uint64  `db:"Rid"                                 xlsx:"Rid"`
	Srid                                                              string  `db:"Srid"                                xlsx:"Srid"`
	RefundForTheDeliveryAndReturnOfGoodsToThePVZ                      float64 `db:"Возмещение за выдачу и возврат тов" xlsx:"Возмещение за выдачу и возврат товаров на ПВЗ"`
	ReimbursementOfAcquiringCosts                                     float64 `db:"Возмещение расходов по эквайрингу" xlsx:"Возмещение расходов по эквайрингу"`
	NameOfAcquiringBank                                               string  `db:"Наименование банка эквайринга" xlsx:"Наименование банка эквайринга"`
	MarkingCode                                                       string  `db:"Код маркировки"                      xlsx:"Код маркировки"`
	ReportID                                                          uint64  `db:"report_id"`
	SellerID                                                          uint64  `db:"seller_id"`
}

type ReportRowV4 struct {
	Number                                                            uint64  `db:"№"                                   xlsx:"№"`
	SupplyNumber                                                      uint64  `db:"Номер поставки"                      xlsx:"Номер поставки"`
	Subject                                                           string  `db:"Предмет"                             xlsx:"Предмет"`
	ItemCode                                                          uint64  `db:"Код номенклатуры"                    xlsx:"Код номенклатуры"`
	Brand                                                             string  `db:"Бренд"                               xlsx:"Бренд"`
	SuppliersArticle                                                  string  `db:"Артикул поставщика"                  xlsx:"Артикул поставщика"`
	Name                                                              string  `db:"Название"                            xlsx:"Название"`
	Size                                                              string  `db:"Размер"                              xlsx:"Размер"`
	Barcode                                                           uint64  `db:"Баркод"                              xlsx:"Баркод"`
	DocumentType                                                      string  `db:"Тип документа"                       xlsx:"Тип документа"`
	JustificationForPayment                                           string  `db:"Обоснование для оплаты"              xlsx:"Обоснование для оплаты"`
	DateOfTheOrderByTheBuyer                                          string  `db:"Дата заказа покупателем"             xlsx:"Дата заказа покупателем"`
	DateOfSale                                                        string  `db:"Дата продажи"                        xlsx:"Дата продажи"`
	Count                                                             uint64  `db:"Кол-во"                              xlsx:"Кол-во"`
	RetailPrice                                                       float64 `db:"Цена розничная"                      xlsx:"Цена розничная"`
	WildberriesSoldTheProductPr                                       float64 `db:"Вайлдберриз реализовал Товар (Пр)"   xlsx:"Вайлдберриз реализовал Товар (Пр)"`
	AgreedGroceryDiscount                                             float64 `db:"Согласованный продуктовый дискон"    xlsx:"Согласованный продуктовый дисконт, %"`
	Promocode                                                         float64 `db:"Промокод %"                          xlsx:"Промокод %"`
	FinalAgreedDiscount                                               float64 `db:"Итоговая согласованная скидка"       xlsx:"Итоговая согласованная скидка"`
	RetailPriceIncludingTheAgreedDiscount                             float64 `db:"Цена розничная с учетом согласова"   xlsx:"Цена розничная с учетом согласованной скидки"`
	RegularCustomerDiscountSPP                                        float64 `db:"Скидка постоянного Покупателя (СП"   xlsx:"Скидка постоянного Покупателя (СПП)"`
	KVVSizePercent                                                    float64 `db:"Размер кВВ, %"                       xlsx:"Размер кВВ, %"`
	KVVSizeWithoutVATBasicPercent                                     float64 `db:"Размер  кВВ без НДС, % Базовый"      xlsx:"Размер  кВВ без НДС, % Базовый"`
	FinalKVVWithoutVATPercent                                         float64 `db:"Итоговый кВВ без НДС, %"             xlsx:"Итоговый кВВ без НДС, %"`
	RemunerationFromSalesBeforeDeductionDfAttorneysServicesWithoutVAT float64 `db:"Вознаграждение с продаж до вычета "  xlsx:"Вознаграждение с продаж до вычета услуг поверенного, без НДС"`
	WildberriesRemunerationBBWithoutVAT                               float64 `db:"Вознаграждение Вайлдберриз (ВВ), б"  xlsx:"Вознаграждение Вайлдберриз (ВВ), без НДС"`
	VATOnWildberriesRemuneration                                      float64 `db:"НДС с Вознаграждения Вайлдберриз"    xlsx:"НДС с Вознаграждения Вайлдберриз"`
	ToTransferToTheSellerForTheSoldGoods                              float64 `db:"К перечислению Продавцу за реализ"   xlsx:"К перечислению Продавцу за реализованный Товар"`
	NumberOfDeliveries                                                uint64  `db:"Количество доставок"                 xlsx:"Количество доставок"`
	RefundAmount                                                      uint64  `db:"Количество возврата"                 xlsx:"Количество возврата"`
	ServicesForTheDeliveryOfGoodsToTheBuyer                           float64 `db:"Услуги по доставке товара покупат"   xlsx:"Услуги по доставке товара покупателю"`
	Fines                                                             float64 `db:"Общая сумма штрафов"                 xlsx:"Общая сумма штрафов"`
	Surcharges                                                        float64 `db:"Доплаты"                             xlsx:"Доплаты"`
	JustificationOfFinesAndSurcharges                                 string  `db:"Обоснование штрафов и доплат"        xlsx:"Обоснование штрафов и доплат"`
	StickerMP                                                         uint64  `db:"Стикер МП"                           xlsx:"Стикер МП"`
	OfficeNumber                                                      uint64  `db:"Номер офиса"                         xlsx:"Номер офиса"`
	NameOfTheDeliveryOffice                                           string  `db:"Наименование офиса доставки"         xlsx:"Наименование офиса доставки"`
	PartnersINN                                                       uint64  `db:"ИНН партнера"                        xlsx:"ИНН партнера"`
	Partner                                                           string  `db:"Партнер"                             xlsx:"Партнер"`
	Warehouse                                                         string  `db:"Склад"                               xlsx:"Склад"`
	Country                                                           string  `db:"Страна"                              xlsx:"Страна"`
	TypeOfBoxes                                                       string  `db:"Тип коробов"                         xlsx:"Тип коробов"`
	CustomsDeclarationNumber                                          string  `db:"Номер таможенной декларации"         xlsx:"Номер таможенной декларации"`
	ShK                                                               uint64  `db:"ШК"                                  xlsx:"ШК"`
	Rid                                                               uint64  `db:"Rid"                                 xlsx:"Rid"`
	Srid                                                              string  `db:"Srid"                                xlsx:"Srid"`
	RefundForTheDeliveryAndReturnOfGoodsToThePVZ                      float64 `db:"Возмещение за выдачу и возврат тов" xlsx:"Возмещение за выдачу и возврат товаров на ПВЗ"`
	ReimbursementOfAcquiringCosts                                     float64 `db:"Возмещение расходов по эквайрингу" xlsx:"Возмещение расходов по эквайрингу"`
	NameOfAcquiringBank                                               string  `db:"Наименование банка эквайринга" xlsx:"Наименование банка эквайринга"`
	ReportID                                                          uint64  `db:"report_id"`
	SellerID                                                          uint64  `db:"seller_id"`
}
type ReportRowV3 struct {
	Number                                                            uint64  `db:"№"                                   xlsx:"№"`
	SupplyNumber                                                      uint64  `db:"Номер поставки"                      xlsx:"Номер поставки"`
	Subject                                                           string  `db:"Предмет"                             xlsx:"Предмет"`
	ItemCode                                                          uint64  `db:"Код номенклатуры"                    xlsx:"Код номенклатуры"`
	Brand                                                             string  `db:"Бренд"                               xlsx:"Бренд"`
	SuppliersArticle                                                  string  `db:"Артикул поставщика"                  xlsx:"Артикул поставщика"`
	Name                                                              string  `db:"Название"                            xlsx:"Название"`
	Size                                                              string  `db:"Размер"                              xlsx:"Размер"`
	Barcode                                                           uint64  `db:"Баркод"                              xlsx:"Баркод"`
	DocumentType                                                      string  `db:"Тип документа"                       xlsx:"Тип документа"`
	JustificationForPayment                                           string  `db:"Обоснование для оплаты"              xlsx:"Обоснование для оплаты"`
	DateOfTheOrderByTheBuyer                                          string  `db:"Дата заказа покупателем"             xlsx:"Дата заказа покупателем"`
	DateOfSale                                                        string  `db:"Дата продажи"                        xlsx:"Дата продажи"`
	Count                                                             uint64  `db:"Кол-во"                              xlsx:"Кол-во"`
	RetailPrice                                                       float64 `db:"Цена розничная"                      xlsx:"Цена розничная"`
	WildberriesSoldTheProductPr                                       float64 `db:"Вайлдберриз реализовал Товар (Пр)"   xlsx:"Вайлдберриз реализовал Товар (Пр)"`
	AgreedGroceryDiscount                                             float64 `db:"Согласованный продуктовый дискон"    xlsx:"Согласованный продуктовый дисконт, %"`
	Promocode                                                         float64 `db:"Промокод %"                          xlsx:"Промокод %"`
	FinalAgreedDiscount                                               float64 `db:"Итоговая согласованная скидка"       xlsx:"Итоговая согласованная скидка"`
	RetailPriceIncludingTheAgreedDiscount                             float64 `db:"Цена розничная с учетом согласова"   xlsx:"Цена розничная с учетом согласованной скидки"`
	RegularCustomerDiscountSPP                                        float64 `db:"Скидка постоянного Покупателя (СП"   xlsx:"Скидка постоянного Покупателя (СПП)"`
	KVVSizePercent                                                    float64 `db:"Размер кВВ, %"                       xlsx:"Размер кВВ, %"`
	KVVSizeWithoutVATBasicPercent                                     float64 `db:"Размер  кВВ без НДС, % Базовый"      xlsx:"Размер  кВВ без НДС, % Базовый"`
	FinalKVVWithoutVATPercent                                         float64 `db:"Итоговый кВВ без НДС, %"             xlsx:"Итоговый кВВ без НДС, %"`
	RemunerationFromSalesBeforeDeductionDfAttorneysServicesWithoutVAT float64 `db:"Вознаграждение с продаж до вычета "  xlsx:"Вознаграждение с продаж до вычета услуг поверенного, без НДС"`
	WildberriesRemunerationBBWithoutVAT                               float64 `db:"Вознаграждение Вайлдберриз (ВВ), б"  xlsx:"Вознаграждение Вайлдберриз (ВВ), без НДС"`
	VATOnWildberriesRemuneration                                      float64 `db:"НДС с Вознаграждения Вайлдберриз"    xlsx:"НДС с Вознаграждения Вайлдберриз"`
	ToTransferToTheSellerForTheSoldGoods                              float64 `db:"К перечислению Продавцу за реализ"   xlsx:"К перечислению Продавцу за реализованный Товар"`
	NumberOfDeliveries                                                uint64  `db:"Количество доставок"                 xlsx:"Количество доставок"`
	RefundAmount                                                      uint64  `db:"Количество возврата"                 xlsx:"Количество возврата"`
	ServicesForTheDeliveryOfGoodsToTheBuyer                           float64 `db:"Услуги по доставке товара покупат"   xlsx:"Услуги по доставке товара покупателю"`
	Fines                                                             float64 `db:"Штрафы"                              xlsx:"Штрафы"`
	Surcharges                                                        float64 `db:"Доплаты"                             xlsx:"Доплаты"`
	JustificationOfFinesAndSurcharges                                 string  `db:"Обоснование штрафов и доплат"        xlsx:"Обоснование штрафов и доплат"`
	StickerMP                                                         uint64  `db:"Стикер МП"                           xlsx:"Стикер МП"`
	OfficeNumber                                                      uint64  `db:"Номер офиса"                         xlsx:"Номер офиса"`
	NameOfTheDeliveryOffice                                           string  `db:"Наименование офиса доставки"         xlsx:"Наименование офиса доставки"`
	PartnersINN                                                       uint64  `db:"ИНН партнера"                        xlsx:"ИНН партнера"`
	Partner                                                           string  `db:"Партнер"                             xlsx:"Партнер"`
	Warehouse                                                         string  `db:"Склад"                               xlsx:"Склад"`
	Country                                                           string  `db:"Страна"                              xlsx:"Страна"`
	TypeOfBoxes                                                       string  `db:"Тип коробов"                         xlsx:"Тип коробов"`
	CustomsDeclarationNumber                                          string  `db:"Номер таможенной декларации"         xlsx:"Номер таможенной декларации"`
	ShK                                                               uint64  `db:"ШК"                                  xlsx:"ШК"`
	Rid                                                               uint64  `db:"Rid"                                 xlsx:"Rid"`
	Srid                                                              string  `db:"Srid"                                xlsx:"Srid"`
	RefundForTheDeliveryAndReturnOfGoodsToThePVZ                      float64 `db:"Возмещение за выдачу и возврат тов" xlsx:"Возмещение за выдачу и возврат товаров на ПВЗ"`
	ReimbursementOfAcquiringCosts                                     float64 `db:"Возмещение расходов по эквайрингу" xlsx:"Возмещение расходов по эквайрингу"`
	NameOfAcquiringBank                                               string  `db:"Наименование банка эквайринга" xlsx:"Наименование банка эквайринга"`
	ReportID                                                          uint64  `db:"report_id"`
	SellerID                                                          uint64  `db:"seller_id"`
}

type ReportRowV2 struct {
	Number                                                            uint64  `db:"№"                                   xlsx:"№"`
	SupplyNumber                                                      uint64  `db:"Номер поставки"                      xlsx:"Номер поставки"`
	Subject                                                           string  `db:"Предмет"                             xlsx:"Предмет"`
	ItemCode                                                          uint64  `db:"Код номенклатуры"                    xlsx:"Код номенклатуры"`
	Brand                                                             string  `db:"Бренд"                               xlsx:"Бренд"`
	SuppliersArticle                                                  string  `db:"Артикул поставщика"                  xlsx:"Артикул поставщика"`
	Name                                                              string  `db:"Название"                            xlsx:"Название"`
	Size                                                              string  `db:"Размер"                              xlsx:"Размер"`
	Barcode                                                           uint64  `db:"Баркод"                              xlsx:"Баркод"`
	DocumentType                                                      string  `db:"Тип документа"                       xlsx:"Тип документа"`
	JustificationForPayment                                           string  `db:"Обоснование для оплаты"              xlsx:"Обоснование для оплаты"`
	DateOfTheOrderByTheBuyer                                          string  `db:"Дата заказа покупателем"             xlsx:"Дата заказа покупателем"`
	DateOfSale                                                        string  `db:"Дата продажи"                        xlsx:"Дата продажи"`
	Count                                                             uint64  `db:"Кол-во"                              xlsx:"Кол-во"`
	RetailPrice                                                       float64 `db:"Цена розничная"                      xlsx:"Цена розничная"`
	WildberriesSoldTheProductPr                                       float64 `db:"Вайлдберриз реализовал Товар (Пр)"   xlsx:"Вайлдберриз реализовал Товар (Пр)"`
	AgreedGroceryDiscount                                             float64 `db:"Согласованный продуктовый дискон"    xlsx:"Согласованный продуктовый дисконт, %"`
	Promocode                                                         float64 `db:"Промокод %"                          xlsx:"Промокод %"`
	FinalAgreedDiscount                                               float64 `db:"Итоговая согласованная скидка"       xlsx:"Итоговая согласованная скидка"`
	RetailPriceIncludingTheAgreedDiscount                             float64 `db:"Цена розничная с учетом согласова"   xlsx:"Цена розничная с учетом согласованной скидки"`
	RegularCustomerDiscountSPP                                        float64 `db:"Скидка постоянного Покупателя (СП"   xlsx:"Скидка постоянного Покупателя (СПП)"`
	KVVSizePercent                                                    float64 `db:"Размер кВВ, %"                       xlsx:"Размер кВВ, %"`
	KVVSizeWithoutVATBasicPercent                                     float64 `db:"Размер  кВВ без НДС, % Базовый"      xlsx:"Размер  кВВ без НДС, % Базовый"`
	FinalKVVWithoutVATPercent                                         float64 `db:"Итоговый кВВ без НДС, %"             xlsx:"Итоговый кВВ без НДС, %"`
	RemunerationFromSalesBeforeDeductionDfAttorneysServicesWithoutVAT float64 `db:"Вознаграждение с продаж до вычета "  xlsx:"Вознаграждение с продаж до вычета услуг поверенного, без НДС"`
	ReimbursementOfAttorneysExpenses                                  uint64  `db:"Возмещение Расходов услуг поверен"   xlsx:"Возмещение Расходов услуг поверенного"`
	WildberriesRemunerationBBWithoutVAT                               float64 `db:"Вознаграждение Вайлдберриз (ВВ), б"  xlsx:"Вознаграждение Вайлдберриз (ВВ), без НДС"`
	VATOnWildberriesRemuneration                                      float64 `db:"НДС с Вознаграждения Вайлдберриз"    xlsx:"НДС с Вознаграждения Вайлдберриз"`
	ToTransferToTheSellerForTheSoldGoods                              float64 `db:"К перечислению Продавцу за реализ"   xlsx:"К перечислению Продавцу за реализованный Товар"`
	NumberOfDeliveries                                                uint64  `db:"Количество доставок"                 xlsx:"Количество доставок"`
	RefundAmount                                                      uint64  `db:"Количество возврата"                 xlsx:"Количество возврата"`
	ServicesForTheDeliveryOfGoodsToTheBuyer                           float64 `db:"Услуги по доставке товара покупат"   xlsx:"Услуги по доставке товара покупателю"`
	Fines                                                             float64 `db:"Штрафы"                              xlsx:"Штрафы"`
	Surcharges                                                        float64 `db:"Доплаты"                             xlsx:"Доплаты"`
	JustificationOfFinesAndSurcharges                                 string  `db:"Обоснование штрафов и доплат"        xlsx:"Обоснование штрафов и доплат"`
	StickerMP                                                         uint64  `db:"Стикер МП"                           xlsx:"Стикер МП"`
	OfficeNumber                                                      uint64  `db:"Номер офиса"                         xlsx:"Номер офиса"`
	NameOfTheDeliveryOffice                                           string  `db:"Наименование офиса доставки"         xlsx:"Наименование офиса доставки"`
	PartnersINN                                                       uint64  `db:"ИНН партнера"                        xlsx:"ИНН партнера"`
	Partner                                                           string  `db:"Партнер"                             xlsx:"Партнер"`
	Warehouse                                                         string  `db:"Склад"                               xlsx:"Склад"`
	Country                                                           string  `db:"Страна"                              xlsx:"Страна"`
	TypeOfBoxes                                                       string  `db:"Тип коробов"                         xlsx:"Тип коробов"`
	CustomsDeclarationNumber                                          string  `db:"Номер таможенной декларации"         xlsx:"Номер таможенной декларации"`
	ShK                                                               uint64  `db:"ШК"                                  xlsx:"ШК"`
	Rid                                                               uint64  `db:"Rid"                                 xlsx:"Rid"`
	Srid                                                              string  `db:"Srid"                                xlsx:"Srid"`
	ReportID                                                          uint64  `db:"report_id"`
	SellerID                                                          uint64  `db:"seller_id"`
}

type ReportRowV1 struct {
	SupplyNumber                            uint64  `db:"Номер поставки"                      xlsx:"Номер поставки"`
	Subject                                 string  `db:"Предмет"                             xlsx:"Предмет"`
	ItemCode                                uint64  `db:"Код номенклатуры"                    xlsx:"Код номенклатуры"`
	Brand                                   string  `db:"Бренд"                               xlsx:"Бренд"`
	SuppliersArticle                        string  `db:"Артикул поставщика"                  xlsx:"Артикул поставщика"`
	Size                                    string  `db:"Размер"                              xlsx:"Размер"`
	Barcode                                 uint64  `db:"Баркод"                              xlsx:"Баркод"`
	DocumentType                            string  `db:"Тип документа"                       xlsx:"Тип документа"`
	JustificationForPayment                 string  `db:"Обоснование для оплаты"              xlsx:"Обоснование для оплаты"`
	DateOfTheOrderByTheBuyer                string  `db:"Дата заказа покупателем"             xlsx:"Дата заказа покупателем"`
	DateOfSale                              string  `db:"Дата продажи"                        xlsx:"Дата продажи"`
	Count                                   uint64  `db:"Кол-во"                              xlsx:"Кол-во"`
	RetailPrice                             float64 `db:"Цена розничная"                      xlsx:"Цена розничная"`
	WildberriesSoldTheProductPr             float64 `db:"Вайлдберриз реализовал Товар (Пр)"   xlsx:"Вайлдберриз реализовал Товар (Пр)"`
	AgreedGroceryDiscount                   float64 `db:"Согласованный продуктовый дискон"    xlsx:"Согласованный продуктовый дисконт, %"`
	Promocode                               float64 `db:"Промокод %"                          xlsx:"Промокод %"`
	FinalAgreedDiscount                     float64 `db:"Итоговая согласованная скидка"       xlsx:"Итоговая согласованная скидка"`
	RetailPriceIncludingTheAgreedDiscount   float64 `db:"Цена розничная с учетом согласова"   xlsx:"Цена розничная с учетом согласованной скидки"`
	KVVSizePercent                          float64 `db:"Размер кВВ, %"                       xlsx:"Размер кВВ, %"`
	NumberOfDeliveries                      uint64  `db:"Количество доставок"                 xlsx:"Количество доставок"`
	RefundAmount                            uint64  `db:"Количество возврата"                 xlsx:"Количество возврата"`
	ServicesForTheDeliveryOfGoodsToTheBuyer float64 `db:"Услуги по доставке товара покупат"   xlsx:"Услуги по доставке товара покупателю"`
	Fines                                   float64 `db:"Штрафы"                              xlsx:"Штрафы"`
	Surcharges                              float64 `db:"Доплаты"                             xlsx:"Доплаты"`
	JustificationOfFinesAndSurcharges       string  `db:"Обоснование штрафов и доплат"        xlsx:"Обоснование штрафов и доплат"`
	StickerMP                               uint64  `db:"Стикер МП"                           xlsx:"Стикер МП"`
	Warehouse                               string  `db:"Склад"                               xlsx:"Склад"`
	Country                                 string  `db:"Страна"                              xlsx:"Страна"`
	TypeOfBoxes                             string  `db:"Тип коробов"                         xlsx:"Тип коробов"`
	CustomsDeclarationNumber                string  `db:"Номер таможенной декларации"         xlsx:"Номер таможенной декларации"`
	ShK                                     uint64  `db:"ШК"                                  xlsx:"ШК"`
	VATrate                                 float64 `db:"Ставка НДС"                          xlsx:"Ставка НДС"`
	CostPriceAmount                         float64 `db:"Себестоимость Сумма"                 xlsx:"Себестоимость Сумма"`
	RemunerationToTheBuyer                  float64 `db:"Вознаграждение покупателю"           xlsx:"Вознаграждение покупателю"`
	RemunerationToTheSupplier               float64 `db:"Вознаграждение поставщику"           xlsx:"Вознаграждение поставщику"`
	TheAmountOfTheSalesCommission           float64 `db:"Сумма комиссии продаж"               xlsx:"Сумма комиссии продаж"`
	ToTransferForTheProduct                 float64 `db:"К перечислению за товар"             xlsx:"К перечислению за товар"`
	ToTransferForGoodsVAT                   float64 `db:"К перечислению за товар, НДС"        xlsx:"К перечислению за товар, НДС"`
	SPP                                     float64 `db:"СПП"                                 xlsx:"СПП"`
	ReportID                                uint64  `db:"report_id"`
	SellerID                                uint64  `db:"seller_id"`
}
