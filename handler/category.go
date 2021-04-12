package handler

import (
	"context"
	"github.com/PonyWilliam/go-category/domain/model"
	"github.com/PonyWilliam/go-category/domain/service"
	categories "github.com/PonyWilliam/go-category/proto"
)

type Category struct{
	CategoryService service.ICategoryService
}
func(c *Category)CreateCategory(ctx context.Context,request *categories.Create_Category_Request,response *categories.Create_Category_Response) error{
	Category := &model.Category{
		CategoryName: request.CategoryName,
		CategoryDescription: request.CategoryDescription,
	}
	_,err := c.CategoryService.AddCategory(Category)
	if err!=nil{
		return err
	}
	response.Message = "添加成功"
	return nil
}
func(c *Category)DeleteCategory(ctx context.Context,request *categories.Delete_Category_Request,response *categories.Delete_Category_Response) error{
	err := c.CategoryService.DeleteCategoryByID(request.CategoryId)
	if err!=nil{
		return err
	}
	response.Message = "删除成功"
	return nil
}
func(c *Category)FindCategoryById(ctx context.Context,request *categories.FindCateGoryById_Request,response *categories.Category_Response) error{
	category,err := c.CategoryService.FindCategoryByID(request.Id)
	if err!=nil{
		response = &categories.Category_Response{}
		return nil
	}
	response.CategoryName = category.CategoryName
	response.CategoryDescription = category.CategoryDescription
	response.CategoryId = category.ID
	return nil
}
func(c *Category)FindCategoryByName(ctx context.Context,request *categories.Find_CategoryByName_Request,response *categories.Find_All_Response) error {
	category,err := c.CategoryService.FindCategoryByName(request.Name)
	if err!=nil{
		response = nil
		return nil
	}
	for _,v := range category {
		temp := &categories.Category_Response{}
		_ = Swap(v, temp)
		response.Category = append(response.Category,temp)
	}
	return nil
}
func(c *Category)FindAllCategory(ctx context.Context,request *categories.Find_All_Request,response *categories.Find_All_Response) error{
	category,err := c.CategoryService.FindAll()
	if err!=nil{
		response = nil
		return nil
	}
	for _,v := range category {
		temp := &categories.Category_Response{}
		_ = Swap(v, temp)
		response.Category = append(response.Category,temp)
	}
	return nil
}
func(c *Category)UpdateCategory(ctx context.Context,request *categories.Create_Category_Request,response *categories.Update_Category_Response)error{
	Category := &model.Category{
		CategoryName: request.CategoryName,
		CategoryDescription: request.CategoryDescription,
	}
	err := c.CategoryService.UpdateCategory(Category)
	if err != nil{
		return err
	}
	response.Message = "更新成功"
	return nil
}
func Swap(req model.Category,some *categories.Category_Response) error{
	some.CategoryId = req.ID
	some.CategoryDescription = req.CategoryDescription
	some.CategoryName = req.CategoryName
	return nil
}