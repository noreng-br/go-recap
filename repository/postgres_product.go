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

func (r *PostgresProductRepository) GetProducts(ctx context.Context) ([]models.Product, error) {
  var products []models.Product

  db, err := sql.Open("pgx", r.connString)
  fmt.Println("In get products========================")
  if err != nil {
    return products, fmt.Errorf("failed to open database: %w", err)
  }

  defer db.Close()

  selectSql := "SELECT * from products";

  rows, err := db.Query(selectSql)
  if err != nil {
    fmt.Println("=====================================================")
    fmt.Println("Query error")
    fmt.Println(err.Error())
    fmt.Println("=====================================================")
    return products, fmt.Errorf("failed to execute select query: %w", err)
  }

  for rows.Next() {
    var product models.Product

    rows.Scan(
        &product.ProductID,
        &product.Name,
        &product.Description,
        &product.Price,
    )

    products = append(products, product)
  }

  // later on we need to get the categories given the product id

  return products, nil
}

func (r *PostgresProductRepository) GetProductById(ctx context.Context, productId string) (models.Product, error) {
  var product models.Product
  db, err := sql.Open("pgx", r.connString)
  fmt.Println("In Get product by Id =================================")
  fmt.Println("=============================================")
  if err != nil {
    return product, fmt.Errorf("failed to open database: %w", err)
  }

  defer db.Close()

  selectSql := "SELECT * from products where product_id=$1";

  err = db.QueryRowContext(ctx, selectSql, productId).Scan(
    &product.ProductID,
    &product.Name,
    &product.Description,
    &product.Price,
  )
  if err != nil {
    fmt.Println("==============================================")
    fmt.Println("Query error")
    fmt.Println(err.Error())
    fmt.Println("==============================================")
    return product, fmt.Errorf("failed to execute select query: %w", err)
  }

  return product, nil
}
