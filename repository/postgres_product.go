package repository

import (
    "fmt"
    "context"
    "database/sql"

    "codeberg.org/noreng-br/models"
)

type PostgresProductRepository struct {
  connString string
}

func NewPostgresProductRepository(connString string) *PostgresProductRepository {
  return &PostgresProductRepository{connString: connString}
}

func (r *PostgresProductRepository) CreateProduct(ctx context.Context, product models.Product) (models.Product, error) {
  db, err := sql.Open("pgx", r.connString)
  fmt.Println("In createProduct======================")
  fmt.Println(r.connString)
  fmt.Println("====================================")
  if err != nil {
    return models.Product{}, fmt.Errorf("failed to open database: %w", err)
  }
  defer db.Close() // Ensure the connection pool is closed when the function exits

  // 2. SQL to insert data and return the generated ID.
  insertSQL := "INSERT INTO products (name, description, price, category_id) VALUES ($1, $2, $3) RETURNING product_id;"

  var newProductId int
  
  // 3. Execute the query using QueryRow, passing the name and age as arguments ($1, $2).
  // Scan the returned ID into the newUserID variable.
  err = db.QueryRow(insertSQL, product.Name, product.Description, product.Price).Scan(&newProductId)
  if err != nil {
    return models.Product{}, fmt.Errorf("failed to execute insert query: %w", err)
  }

  product.ProductID = fmt.Sprintf("%d", newProductId)

  return product, nil
}
