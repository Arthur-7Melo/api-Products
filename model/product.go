package model

type Product struct {
	Id        int     `json:"product_id"`
	Name      string  `json:"name" binding:"required,min=3,max=50"`
	Price     float64 `json:"price" binding:"required,gt=0"`
	Categorie string  `json:"product_categorie" binding:"required,min=3,max=50"`
}