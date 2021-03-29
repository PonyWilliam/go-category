package repository

import (
	"github.com/PonyWilliam/go-category/domain/model"
	"github.com/jinzhu/gorm"
)

type ICategoryRepository interface{
	InitTable()	error
	CreateCategory(category *model.Category)(int64,error)
	DeleteCategoryByID(int64) error
	UpdateCategory(category *model.Category) error
	FindCategoryByID(int64) (*model.Category,error)
	FindCategoryByName(string)([]model.Category,error)
	FindCategoryByLevel(int64)([]model.Category,error)
	FindAll()([]model.Category,error)
}
func NewCategoryRepository(db *gorm.DB) ICategoryRepository{
	return &CategoryRepository{mysqlDb: db}
}
type CategoryRepository struct{
	mysqlDb *gorm.DB
}
func(c *CategoryRepository) InitTable() error{
	if c.mysqlDb.HasTable(&model.Category{}){
		return nil
	}
	return c.mysqlDb.CreateTable(&model.Category{}).Error
}
func(c *CategoryRepository) CreateCategory(category *model.Category)(int64,error){
	return category.ID,c.mysqlDb.Model(category).Create(&category).Error
}
func(c *CategoryRepository) DeleteCategoryByID(id int64)(error){
	return c.mysqlDb.Where("id = ?",id).Delete(&model.Category{}).Error
}
func(c *CategoryRepository) UpdateCategory(category *model.Category) error{
	return c.mysqlDb.Model(category).Update().Error
}
func(c *CategoryRepository) FindCategoryByID(id int64)(*model.Category,error){
	category := &model.Category{}
	return category,c.mysqlDb.First(category,id).Error
}
func(c *CategoryRepository) FindCategoryByName(name string)(category []model.Category,err error){
	return category,c.mysqlDb.Where("CategoryName = ?",name).Find(&category).Error
}
func(c *CategoryRepository) FindAll()(category []model.Category,err error){
	return category,c.mysqlDb.Find(&category).Error
}
func(c *CategoryRepository) FindCategoryByLevel(level int64)(category []model.Category,err error){
	return category,c.mysqlDb.Where("CategoryLevel = ?",level).Find(&category).Error
}