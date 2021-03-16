package handler

import (
	"context"
	"github.com/PonyWilliam/go-category/domain/model"
	"github.com/PonyWilliam/go-category/domain/service"
	categories "github.com/PonyWilliam/go-category/proto"
	"github.com/PonyWilliam/go-common"
)

type Category struct{
	CategoryService service.ICategoryService
}
func(c *Category)CreateCategory(ctx context.Context,request *categories.Create_Category_Request,response *categories.Create_Category_Response) error{
	Category := &model.Category{
		CategoryName: request.CategoryName,
		CategoryImage: request.CategoryImage,
		CategoryDescription: request.CategoryDescription,
		CategoryLevel: request.CategoryLevel,
		CategoryParent: request.CategoryParent,
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
	rsp := &categories.Category_Response{}
	if err != nil{
		return err
	}
	_ = common.SwapTo(category,rsp)
	return nil
}
func(c *Category)FindCategoryByName(ctx context.Context,request *categories.Find_CategoryByName_Request,response *categories.Category_Response) error {
	return nil
}
func(c *Category)FindCategoryByLevel(ctx context.Context,request *categories.Find_CategoryByLevel_Request,response *categories.Category_Response) error{
	return nil
}
func(c *Category)FindAllCategory(ctx context.Context,request *categories.Find_All_Request,response *categories.Find_All_Response) error{
	return nil
}
func(c *Category)UpdateCategory(ctx context.Context,request *categories.Create_Category_Request,response *categories.Update_Category_Response)error{
	Category := &model.Category{
		CategoryName: request.CategoryName,
		CategoryImage: request.CategoryImage,
		CategoryDescription: request.CategoryDescription,
		CategoryLevel: request.CategoryLevel,
		CategoryParent: request.CategoryParent,
	}
	err := c.CategoryService.UpdateCategory(Category)
	if err != nil{
		return err
	}
	response.Message = "更新成功"
	return nil
}