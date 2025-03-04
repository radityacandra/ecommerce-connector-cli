package model

type Product struct {
	Id          string
	Name        string `validate:"required" firestore:"name"`
	Description string `validate:"required" firestore:"description"`
	Category    string `validate:"required" firestore:"category"`
	Price       int    `validate:"required,gt=0" firestore:"price"`
	EanCode     string `validate:"required" firestore:"ean_code"`
	UserId      string `validate:"required" firestore:"user_id"`
}

func NewProduct(id, name, description, category string, price int, eanCode, userId string) *Product {
	return &Product{
		Id:          id,
		Name:        name,
		Description: description,
		Category:    category,
		Price:       price,
		EanCode:     eanCode,
		UserId:      userId,
	}
}
