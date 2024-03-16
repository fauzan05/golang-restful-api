package repository

import (
	"context"
	"database/sql"
	"errors"
	"golang-restful-api/helper"
	"golang-restful-api/model/domain"
)

type CategoryRepositoryImpl struct {
}

func NewCategoryRepository() CategoryRepository {
	return &CategoryRepositoryImpl{}
}

func (repository *CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	query := "INSERT INTO categories(name) VALUES(?)"
	result, err := tx.ExecContext(ctx, query, category.Name)
	helper.HandleErrorWithPanic(err)
	id, err := result.LastInsertId()
	helper.HandleErrorWithPanic(err)

	// cari kategori yang barusan di buat
	query = "SELECT * FROM categories WHERE id = ?"
	rows, _ := tx.QueryContext(ctx, query, id)
	category = domain.Category{}
	err = rows.Scan(&category.Id, &category.Name)
	helper.HandleErrorWithPanic(err)
	return category
}

func (repository *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	query := "UPDATE categories SET name = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, query, category.Name, category.Id)
	helper.HandleErrorWithPanic(err)
	return category
}

func (repository *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, categoryId int) {
	query := "DELETE FROM categories WHERE id = ?"
	_, err := tx.ExecContext(ctx, query, categoryId)
	helper.HandleErrorWithPanic(err)
}

func (repository *CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error) {
	query := "SELECT * FROM categories WHERE id = ?"
	rows, err := tx.QueryContext(ctx, query, categoryId)
	helper.HandleErrorWithPanic(err)
	category := domain.Category{}
	if rows.Next() {
		err := rows.Scan(&category.Id, &category.Name)
		helper.HandleErrorWithPanic(err)
		return category, nil
	} else {
		return category, errors.New("Category is not found")
	}
	 
}

func (repository *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	query := "SELECT * FROM categories"
	rows, err := tx.QueryContext(ctx, query)
	helper.HandleErrorWithPanic(err)

	var categories []domain.Category

	for rows.Next() {
		category := domain.Category{}
		err := rows.Scan(&category.Id, &category.Name)
		helper.HandleErrorWithPanic(err)
		categories = append(categories, category)
	}

	return categories
}