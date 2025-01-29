package services

import (
	"context"
	"database/sql"
	"library-api-category/internal/commons/response"
	"library-api-category/internal/models"
	"library-api-category/internal/params"
	"library-api-category/internal/repositories"
	"time"
)

type CategoryService interface {
	CreateCategory(ctx context.Context, req *params.CategoryRequest) *response.CustomError
	GetDetailCategory(ctx context.Context, id uint64) (*params.CategoryResponse, *response.CustomError)
	UpdateCategory(ctx context.Context, id uint64, req *params.CategoryRequest) *response.CustomError
	DeleteCategory(ctx context.Context, id uint64) *response.CustomError
	GetAllCategories(ctx context.Context, pagination *models.Pagination) ([]*params.CategoryResponse, *response.CustomError)
	AddBookCategory(ctx context.Context, req *params.BookCategoryRequest) *response.CustomError
	ListCategoryOfBook(ctx context.Context, bookID uint64) ([]*params.CategoryResponse, *response.CustomError)
}

type CategoryServiceImpl struct {
	DB                 *sql.DB
	CategoryRepository repositories.CategoryRepository
}

func NewCategoryService(db *sql.DB, CategoryRepository repositories.CategoryRepository) CategoryService {
	return &CategoryServiceImpl{
		DB:                 db,
		CategoryRepository: CategoryRepository,
	}
}

func (service *CategoryServiceImpl) CreateCategory(ctx context.Context, req *params.CategoryRequest) *response.CustomError {
	tx, err := service.DB.Begin()
	if err != nil {
		return response.GeneralError("Failed Connection to database errors: " + err.Error())
	}
	defer func() {
		err := recover()
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	var cate = models.Category{
		Name:        req.Name,
		Description: req.Description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	err = service.CategoryRepository.CreateCategory(ctx, tx, &cate)

	if err != nil {
		return response.GeneralError(err.Error())
	}

	return nil
}

func (service *CategoryServiceImpl) GetDetailCategory(ctx context.Context, id uint64) (*params.CategoryResponse, *response.CustomError) {
	tx, err := service.DB.Begin()
	if err != nil {
		return nil, response.GeneralError("Failed Connection to database errors: " + err.Error())
	}
	defer func() {
		err := recover()
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	cate, err := service.CategoryRepository.FindCategoryByID(ctx, tx, id)
	if err != nil {
		return nil, response.NotFoundError("Category not found")
	}

	cateResponse := &params.CategoryResponse{
		ID:          cate.ID,
		Name:        cate.Name,
		Description: cate.Description,
		CreatedAt:   cate.CreatedAt,
		UpdatedAt:   cate.UpdatedAt,
	}

	return cateResponse, nil
}

func (service *CategoryServiceImpl) UpdateCategory(ctx context.Context, id uint64, req *params.CategoryRequest) *response.CustomError {
	tx, err := service.DB.Begin()
	if err != nil {
		return response.GeneralError("Failed to connect to the database: " + err.Error())
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	book := models.Category{
		ID:          id,
		Name:        req.Name,
		Description: req.Description,
		UpdatedAt:   time.Now(),
	}

	err = service.CategoryRepository.UpdateCategory(ctx, tx, &book)
	if err != nil {
		return response.GeneralError("Failed to update category: " + err.Error())
	}

	return nil
}

func (service *CategoryServiceImpl) DeleteCategory(ctx context.Context, id uint64) *response.CustomError {
	tx, err := service.DB.Begin()
	if err != nil {
		return response.GeneralError("Failed to connect to the database: " + err.Error())
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	err = service.CategoryRepository.DeleteCategory(ctx, tx, id)
	if err != nil {
		return response.GeneralError("Failed to delete category: " + err.Error())
	}

	return nil
}

func (service *CategoryServiceImpl) GetAllCategories(ctx context.Context, pagination *models.Pagination) ([]*params.CategoryResponse, *response.CustomError) {
	tx, err := service.DB.Begin()
	if err != nil {
		return nil, response.GeneralError("Failed to connect to the database: " + err.Error())
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	pagination.Offset = (pagination.Page - 1) * pagination.PageSize

	categories, err := service.CategoryRepository.GetAllCategories(ctx, tx, pagination)
	if err != nil {
		return nil, response.GeneralError("Failed to fetch categories: " + err.Error())
	}

	cateResponses := make([]*params.CategoryResponse, len(categories))
	for i, cate := range categories {
		cateResponses[i] = &params.CategoryResponse{
			ID:          cate.ID,
			Name:        cate.Name,
			Description: cate.Description,
			CreatedAt:   cate.CreatedAt,
			UpdatedAt:   cate.UpdatedAt,
		}
	}

	pagination.PageCount = (pagination.TotalCount + pagination.PageSize - 1) / pagination.PageSize

	return cateResponses, nil
}

func (service *CategoryServiceImpl) AddBookCategory(ctx context.Context, req *params.BookCategoryRequest) *response.CustomError {
	tx, err := service.DB.Begin()
	if err != nil {
		return response.GeneralError("Failed Connection to database errors: " + err.Error())
	}
	defer func() {
		err := recover()
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	var bookCate = models.BookCategory{
		CategoryID: req.CategoryID,
		BookID:     req.BookID,
	}

	err = service.CategoryRepository.AddBookCategory(ctx, tx, &bookCate)

	if err != nil {
		return response.GeneralError(err.Error())
	}

	return nil
}

func (service *CategoryServiceImpl) ListCategoryOfBook(ctx context.Context, bookID uint64) ([]*params.CategoryResponse, *response.CustomError) {
	tx, err := service.DB.Begin()
	if err != nil {
		return nil, response.GeneralError("Failed to connect to the database: " + err.Error())
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	categories, err := service.CategoryRepository.ListCategoryOfBook(ctx, tx, bookID)
	if err != nil {
		return nil, response.GeneralError("Failed to fetch list book categories: " + err.Error())
	}

	cateResponses := make([]*params.CategoryResponse, len(categories))
	for i, cate := range categories {
		cateResponses[i] = &params.CategoryResponse{
			ID:          cate.ID,
			Name:        cate.Name,
			Description: cate.Description,
			CreatedAt:   cate.CreatedAt,
			UpdatedAt:   cate.UpdatedAt,
		}
	}

	return cateResponses, nil
}
