package service

import (
	"github.com/PonyWilliam/go-category/domain/model"
	"github.com/PonyWilliam/go-category/domain/repository"
)

type ICategoryService interface {
	AddCategory(*model.Category)(int64,error)
	DeleteCategoryByID(int64) error
	UpdateCategory(category *model.Category) error
	FindCategoryByID(int64) (*model.Category,error)
	FindCategoryByName(string)([]model.Category,error)
	FindCategoryByLevel(int64)([]model.Category,error)
	FindAll()([]model.Category,error)
}
func NewCategoryService(categoryRepository repository.ICategoryRepository)ICategoryService{
	return &CategoryService{CategoryRepository: categoryRepository}
}
type CategoryService struct{
	CategoryRepository repository.ICategoryRepository
}
func(c *CategoryService) AddCategory(category *model.Category)(int64,error){
	return c.CategoryRepository.CreateCategory(category)
}
func(c *CategoryService) DeleteCategoryByID(id int64)error{
	return c.CategoryRepository.DeleteCategoryByID(id)
}
func(c *CategoryService) UpdateCategory(category *model.Category) error {
	return c.CategoryRepository.UpdateCategory(category)
}
func(c *CategoryService) FindCategoryByID(id int64)(*model.Category,error){
	return c.CategoryRepository.FindCategoryByID(id)
}
func(c *CategoryService) FindCategoryByName(name string)([]model.Category,error){
	return c.CategoryRepository.FindCategoryByName(name)
}
func(c *CategoryService) FindCategoryByLevel(level int64)([]model.Category,error){
	return c.CategoryRepository.FindCategoryByLevel(level)
}
func(c *CategoryService) FindAll()([]model.Category,error){
	return c.CategoryRepository.FindAll()
}