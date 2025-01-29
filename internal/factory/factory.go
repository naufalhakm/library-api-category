package factory

import (
	"database/sql"
	"library-api-category/internal/controllers"
	"library-api-category/internal/repositories"
	"library-api-category/internal/services"
)

type Provider struct {
	CategoryProvider controllers.CategoryController
}

func InitFactory(db *sql.DB) *Provider {

	cateRepo := repositories.NewCategoryRepository()
	cateService := services.NewCategoryService(db, cateRepo)
	cateController := controllers.NewCategoryController(cateService)

	return &Provider{
		CategoryProvider: cateController,
	}
}
