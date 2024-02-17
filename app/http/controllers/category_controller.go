package controllers

import (
	"fmt"
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"goravel/app/http/response"
	"goravel/app/models"
)

type CategoryController struct {
	//Dependent services
}

func NewCategoryController() *CategoryController {
	return &CategoryController{
		//Inject services
	}
}

func (r *CategoryController) Index(ctx http.Context) http.Response {
	start := ctx.Request().QueryInt("start")
	limit := ctx.Request().QueryInt("limit")
	if limit > 100 {
		limit = 100
	} else if limit < 1 {
		limit = 10
	}
	var category []models.Category

	err := facades.Orm().Query().Offset(start).Limit(limit).Get(&category)
	//fmt.Println(total)
	var createCategories []response.CategoryResponse
	if err != nil {
		facades.Log().Debug(err)
		return response.ApiResponse(ctx, 500, createCategories, "")
	}
	detailResponse := response.ToCategoryResponseAll(category)
	return response.ApiResponse(ctx, 200, detailResponse, "Get All Category")
}

func (r *CategoryController) Show(ctx http.Context) http.Response {
	id := ctx.Request().Route("id")
	var category models.Category
	err := facades.Orm().Query().FindOrFail(&category, id)
	if err != nil {
		facades.Log().Debug(err)
		return response.ApiResponse(ctx, 404, category, "")
	}
	detailResponse := response.ToCategoryResponseDetail(category)
	return response.ApiResponse(ctx, 200, detailResponse, "")
}

func (r *CategoryController) Store(ctx http.Context) http.Response {
	validator, _ := ctx.Request().Validate(map[string]string{
		"name": "required|max_len:50",
	})
	if validator.Fails() {
		messages := validator.Errors().All()
		return response.ApiResponse(ctx, 400, messages, "")
	}
	name := ctx.Request().Input("name")
	category := models.Category{Name: name}
	err := facades.Orm().Query().Create(&category)
	if err != nil {
		facades.Log().Debug(err)
		return response.ApiResponse(ctx, 500, category, "")
	}
	detailResponse := response.ToCategoryResponseDetail(category)
	return response.ApiResponse(ctx, 201, detailResponse, "")
}

func (r *CategoryController) Update(ctx http.Context) http.Response {
	validator, _ := ctx.Request().Validate(map[string]string{
		"name": "required|max_len:50",
	})
	if validator.Fails() {
		messages := validator.Errors().All()
		return response.ApiResponse(ctx, 400, messages, "")
	}
	id := ctx.Request().Route("id")
	var category models.Category
	err := facades.Orm().Query().FindOrFail(&category, id)
	if err != nil {
		facades.Log().Debug(err)
		return response.ApiResponse(ctx, 404, category, "")
	}
	fmt.Println(category)
	name := ctx.Request().Input("name")
	_, err = facades.Orm().Query().Model(&category).Where("id", id).Update(models.Category{Name: name})
	if err != nil {
		facades.Log().Debug(err)
		return response.ApiResponse(ctx, 404, category, "")
	}
	detailResponse := response.ToCategoryResponseDetail(category)
	return response.ApiResponse(ctx, 200, detailResponse, "")
}

func (r *CategoryController) Destroy(ctx http.Context) http.Response {
	id := ctx.Request().Route("id")
	var category models.Category
	err := facades.Orm().Query().FindOrFail(&category, id)
	if err != nil {
		facades.Log().Debug(err)
		return response.ApiResponse(ctx, 404, nil, "")
	}
	_, err = facades.Orm().Query().Delete(&category, id)
	if err != nil {
		facades.Log().Debug(err)
		return response.ApiResponse(ctx, 500, nil, "")
	}
	return response.ApiResponse(ctx, 200, nil, "")
}
