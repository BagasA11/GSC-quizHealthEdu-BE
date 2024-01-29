package dto

type ItemCreate struct {
	Name     string `json:"name" binding:"required,min=10"`
	Price    int64  `json:"price" binding:"required"`
	Stock    int16  `json:"stock" binding:"required"`
	Unit     string `json:"unit" binding:"max=7"`
	RetailID uint
}

type ItemCart struct {
	Id    uint   `json:"id" binding:"required"`
	Price int64  `json:"price" binding:"required"`
	Qty   uint16 `json:"quantity" binding:"required"`
}
