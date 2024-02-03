package models

// ini ditaro di entity dev 
type Product struct {
	Id          int64  `gorm:"PrimaryKey" json:"id"`
	NamaProduct string `gorm:"varchar(300)" json:"nama_product"`
	Deskripsi   string `gorm:"text" json:"deskripsi"`
}
