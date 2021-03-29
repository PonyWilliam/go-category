package model
type Category struct{
	ID int64 `gorm:"primary_key;not_null;auto_increment" json:"id"`
	CategoryName string `gorm:"unique_index,not_null" json:"category_name"`
	CategoryDescription string `json:"category_description"`
}
