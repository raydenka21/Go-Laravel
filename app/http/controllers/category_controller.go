package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"goravel/app/helper"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
	"goravel/app/http/response"
	"goravel/app/models"
)

type CategoryController struct {
	rdb     *redis.Client
	rKey    string
	rExpire time.Duration
}

func NewCategoryController(rdb *redis.Client) *CategoryController {
	return &CategoryController{
		rdb:     rdb,
		rKey:    "category",
		rExpire: 60 * time.Second,
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
	var createCategories []response.CategoryResponse
	var key string

	key = r.rKey + "_all_start_" + strconv.Itoa(start) + "_limit_" + strconv.Itoa(limit)
	valRedis, errRedis := helper.RedisGet(ctx, r.rdb, key)

	if errRedis != nil {
		facades.Log().Debug(errRedis)
		return response.ApiResponse(ctx, 500, createCategories, "")
	} else if valRedis != "empty" {
		json.Unmarshal([]byte(valRedis), &createCategories)
		return response.ApiResponse(ctx, 200, createCategories, "Get All Category")
	}

	err := facades.Orm().Query().Offset(start).Limit(limit).Get(&category)

	if err != nil {
		facades.Log().Debug(err)
		return response.ApiResponse(ctx, 500, createCategories, "")
	}
	detailResponse := response.ToCategoryResponseAll(category)
	resRedis, _ := json.Marshal(detailResponse)
	setRedisErr := helper.RedisSet(ctx, r.rdb, key, resRedis, r.rExpire)
	if setRedisErr != nil {
		facades.Log().Debug(err)
		return response.ApiResponse(ctx, 500, createCategories, "")
	}

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
