package controllers

import (
	"library-api-category/internal/commons/response"
	"library-api-category/internal/models"
	"library-api-category/internal/params"
	"library-api-category/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CategoryController interface {
	CreateCategory(ctx *gin.Context)
	GetDetailCategory(ctx *gin.Context)
	UpdateCategory(ctx *gin.Context)
	DeleteCategory(ctx *gin.Context)
	GetAllCategories(ctx *gin.Context)
	AddBookCategory(ctx *gin.Context)
	ListCategoryOfBook(ctx *gin.Context)
}

type CategoryControllerImpl struct {
	CategoryService services.CategoryService
}

func NewCategoryController(CategoryService services.CategoryService) CategoryController {
	return &CategoryControllerImpl{
		CategoryService: CategoryService,
	}
}

func (controller *CategoryControllerImpl) CreateCategory(ctx *gin.Context) {
	var req = new(params.CategoryRequest)

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err,
		})
		return
	}

	custErr := controller.CategoryService.CreateCategory(ctx, req)
	if custErr != nil {
		ctx.AbortWithStatusJSON(custErr.StatusCode, custErr)
		return
	}

	resp := response.CreatedSuccess()
	ctx.JSON(resp.StatusCode, resp)
}

func (controller *CategoryControllerImpl) GetDetailCategory(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err,
		})
		return
	}

	result, custErr := controller.CategoryService.GetDetailCategory(ctx, uint64(id))

	if custErr != nil {
		ctx.AbortWithStatusJSON(custErr.StatusCode, custErr)
		return
	}

	resp := response.GeneralSuccessCustomMessageAndPayload("Success get detail category", result)
	ctx.JSON(resp.StatusCode, resp)
}

func (controller *CategoryControllerImpl) UpdateCategory(ctx *gin.Context) {
	var req = new(params.CategoryRequest)

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err,
		})
		return
	}

	id, _ := strconv.Atoi(ctx.Param("id"))

	custErr := controller.CategoryService.UpdateCategory(ctx, uint64(id), req)
	if custErr != nil {
		ctx.AbortWithStatusJSON(custErr.StatusCode, custErr)
		return
	}

	resp := response.GeneralSuccessCustomMessageAndPayload("Success update data category", nil)
	ctx.JSON(resp.StatusCode, resp)
}

func (controller *CategoryControllerImpl) DeleteCategory(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err,
		})
		return
	}

	custErr := controller.CategoryService.DeleteCategory(ctx, uint64(id))
	if custErr != nil {
		ctx.AbortWithStatusJSON(custErr.StatusCode, custErr)
		return
	}

	resp := response.GeneralSuccessCustomMessageAndPayload("Success delete data category", nil)
	ctx.JSON(resp.StatusCode, resp)
}

func (controller *CategoryControllerImpl) GetAllCategories(ctx *gin.Context) {
	page := ctx.Query("page")
	limit := ctx.Query("limit")

	pageNum := 1
	limitSize := 5

	if page != "" {
		parsedPage, err := strconv.Atoi(page)
		if err == nil && parsedPage > 0 {
			pageNum = parsedPage
		}
	}

	if limit != "" {
		parsedLimit, err := strconv.Atoi(limit)
		if err == nil && parsedLimit > 0 {
			limitSize = parsedLimit
		}
	}

	pagination := models.Pagination{
		Page:     pageNum,
		Offset:   (pageNum - 1) * limitSize,
		PageSize: limitSize,
	}

	result, custErr := controller.CategoryService.GetAllCategories(ctx, &pagination)

	if custErr != nil {
		ctx.AbortWithStatusJSON(custErr.StatusCode, custErr)
		return
	}

	type Response struct {
		Categories interface{} `json:"categories"`
		Pagination interface{} `json:"pagination"`
	}

	var responses Response
	responses.Categories = result
	responses.Pagination = pagination

	resp := response.GeneralSuccessCustomMessageAndPayload("Success get data categories", responses)
	ctx.JSON(resp.StatusCode, resp)
}

func (controller *CategoryControllerImpl) AddBookCategory(ctx *gin.Context) {
	var req = new(params.BookCategoryRequest)

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err,
		})
		return
	}

	custErr := controller.CategoryService.AddBookCategory(ctx, req)
	if custErr != nil {
		ctx.AbortWithStatusJSON(custErr.StatusCode, custErr)
		return
	}

	resp := response.CreatedSuccess()
	ctx.JSON(resp.StatusCode, resp)
}

func (controller *CategoryControllerImpl) ListCategoryOfBook(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err,
		})
		return
	}

	result, custErr := controller.CategoryService.ListCategoryOfBook(ctx, uint64(id))

	if custErr != nil {
		ctx.AbortWithStatusJSON(custErr.StatusCode, custErr)
		return
	}
	resp := response.GeneralSuccessCustomMessageAndPayload("Success get data list book of categories", result)
	ctx.JSON(resp.StatusCode, resp)
}
