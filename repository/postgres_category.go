package repository

import (
    "fmt"
    "context"
    "database/sql"

    "codeberg.org/noreng-br/models"
  )

type PostgresCategoryRepository struct {
  connString string
}

func NewPostgresCategoryRepository(connString string) *PostgresCategoryRepository {
  return &PostgresCategoryRepository{connString: connString}
}

func (r *PostgresCategoryRepository) CreateCategory(ctx context.Context, name string) (models.Category, error) {
  fmt.Println(ctx)
  db, err := sql.Open("pgx", r.connString)
  fmt.Println("In createCategory======================")
  fmt.Println(r.connString)
  fmt.Println("====================================")
  if err != nil {
    return models.Category{}, fmt.Errorf("failed to open database: %w", err)
  }
  defer db.Close() // Ensure the connection pool is closed when the function exits

  // 2. SQL to insert data and return the generated ID.
  insertSQL := "INSERT INTO categories (name) VALUES ($1) RETURNING category_id;"

  var newCategoryId int
  
  // 3. Execute the query using QueryRow, passing the name and age as arguments ($1, $2).
  // Scan the returned ID into the newUserID variable.
  err = db.QueryRow(insertSQL, name).Scan(&newCategoryId)
  if err != nil {
    return models.Category{}, fmt.Errorf("failed to execute insert query: %w", err)
  }

  var category models.Category

  category.CategoryID = fmt.Sprintf("%d", newCategoryId)
  category.Name = name

  return category, nil
}
