package repositories

import (
	"context"
	"database/sql"
	"errors"
	"library-api-category/internal/models"
)

type CategoryRepository interface {
	CreateCategory(ctx context.Context, tx *sql.Tx, cate *models.Category) error
	FindCategoryByID(ctx context.Context, tx *sql.Tx, id uint64) (*models.Category, error)
	UpdateCategory(ctx context.Context, tx *sql.Tx, cate *models.Category) error
	DeleteCategory(ctx context.Context, tx *sql.Tx, id uint64) error
	GetAllCategories(ctx context.Context, tx *sql.Tx, pagination *models.Pagination) ([]*models.Category, error)
	AddBookCategory(ctx context.Context, tx *sql.Tx, bookCate *models.BookCategory) error
	ListCategoryOfBook(ctx context.Context, tx *sql.Tx, bookID uint64) ([]*models.Category, error)
}

type CategoryRepositoryImpl struct {
}

func NewCategoryRepository() CategoryRepository {
	return &CategoryRepositoryImpl{}
}

func (repository *CategoryRepositoryImpl) CreateCategory(ctx context.Context, tx *sql.Tx, cate *models.Category) error {
	query := `INSERT INTO categories (name, description, created_at, updated_at) VALUES ($1, $2, $3, $4)`
	response, err := tx.ExecContext(ctx, query, cate.Name, cate.Description, cate.CreatedAt, cate.UpdatedAt)
	if err != nil || response == nil {
		return errors.New("Failed to create a category, transaction rolled back. Reason: " + err.Error())
	}

	return nil
}

func (repository *CategoryRepositoryImpl) FindCategoryByID(ctx context.Context, tx *sql.Tx, id uint64) (*models.Category, error) {
	query := "SELECT id, name, description, created_at, updated_at FROM categories WHERE id = $1"
	rows, err := tx.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cate = models.Category{}
	if rows.Next() {
		err := rows.Scan(&cate.ID, &cate.Name, &cate.Description, &cate.CreatedAt, &cate.UpdatedAt)
		if err != nil {
			return nil, err
		}
		return &cate, nil
	} else {
		return nil, errors.New("category is not found")
	}
}

func (repository *CategoryRepositoryImpl) UpdateCategory(ctx context.Context, tx *sql.Tx, cate *models.Category) error {
	query := `UPDATE categories SET name = $1, description = $2, updated_at = $3 WHERE id = $4`

	_, err := tx.ExecContext(ctx, query,
		cate.Name,
		cate.Description,
		cate.UpdatedAt,
		cate.ID,
	)
	if err != nil {
		return errors.New("Failed to update a category, transaction rolled back. Reason: " + err.Error())
	}
	return nil
}

func (repository *CategoryRepositoryImpl) DeleteCategory(ctx context.Context, tx *sql.Tx, id uint64) error {
	SQL := `DELETE categories WHERE id = $1`

	_, err := tx.ExecContext(ctx, SQL, id)
	if err != nil {
		return errors.New("Failed to update a category, transaction rolled back. Reason: " + err.Error())
	}
	return nil
}

func (repository *CategoryRepositoryImpl) GetAllCategories(ctx context.Context, tx *sql.Tx, pagination *models.Pagination) ([]*models.Category, error) {
	query := `SELECT id, name, description, created_at, updated_at FROM categories ORDER BY updated_at DESC LIMIT $1 OFFSET $2`
	rows, err := tx.QueryContext(ctx, query, pagination.PageSize, pagination.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []*models.Category
	for rows.Next() {
		var cate models.Category
		err := rows.Scan(&cate.ID, &cate.Name, &cate.Description, &cate.CreatedAt, &cate.UpdatedAt)
		if err != nil {
			return nil, err
		}

		categories = append(categories, &cate)

	}
	return categories, nil
}

func (repository *CategoryRepositoryImpl) AddBookCategory(ctx context.Context, tx *sql.Tx, bookCate *models.BookCategory) error {
	query := `INSERT INTO book_categories (book_id, category_id) VALUES ($1, $2)`
	_, err := tx.ExecContext(ctx, query, bookCate.BookID, bookCate.CategoryID)
	if err != nil {
		return errors.New("Failed to create a book category, transaction rolled back. Reason: " + err.Error())
	}

	return nil
}

func (repository *CategoryRepositoryImpl) ListCategoryOfBook(ctx context.Context, tx *sql.Tx, bookID uint64) ([]*models.Category, error) {
	query := `
		SELECT c.id, c.name, c.description, c.created_at, c.updated_at 
		FROM book_categories bc
		JOIN categories c ON bc.book_id = c.id
		WHERE bc.book_id = $1`
	rows, err := tx.QueryContext(ctx, query, bookID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []*models.Category
	for rows.Next() {
		var cate models.Category
		err := rows.Scan(&cate.ID, &cate.Name, &cate.Description, &cate.CreatedAt, &cate.UpdatedAt)
		if err != nil {
			return nil, err
		}

		categories = append(categories, &cate)
	}
	return categories, nil
}
