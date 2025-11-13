package repository

import (
    "fmt"
    "context"
    "strconv"
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

func (r *PostgresProductRepository) GetProducts(ctx context.Context) ([]models.ProductWithCategories, error) {
  var products []models.ProductWithCategories

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
    var product models.ProductWithCategories

    rows.Scan(
        &product.ProductID,
        &product.Name,
        &product.Description,
        &product.Price,
    )

    productId, err := strconv.Atoi(product.ProductID)
    if err != nil {
      return products, err
    }

    categories, err := r.GetCategoriesByProductID(ctx, productId)
    if err != nil {
      return products, err
    }

    product.Categories = func(cat []models.Category) []string {
      var categoryNames []string

      for _, c := range cat {
        categoryNames = append(categoryNames, c.Name)
      }

      return categoryNames
    }(categories)

    products = append(products, product)
  }

  // later on we need to get the categories given the product id

  return products, nil
}

func (r *PostgresProductRepository) GetProductById(ctx context.Context, productId string) (models.ProductWithCategories, error) {
  var product models.ProductWithCategories
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

  id, err := strconv.Atoi(productId)
  if err != nil {
    return product, err
  }

  categories, err := r.GetCategoriesByProductID(ctx, id)
  if err != nil {
    return product, err
  }

  product.Categories = func(cat []models.Category) []string {
    var categoryNames []string

    for _, c := range cat {
      categoryNames = append(categoryNames, c.Name)
    }

    return categoryNames
  }(categories)

  return product, nil
}

func (r *PostgresProductRepository) AddCategoriesToProduct(ctx context.Context, productID int, categoryIDs []int) error {
    
    db, err := sql.Open("pgx", r.connString)
    fmt.Println("In get user by username======================")
    fmt.Println(r.connString)
    fmt.Println("====================================")
    if err != nil {
      return fmt.Errorf("failed to open database: %w", err)
    }
    defer db.Close() // Ensure the connection pool is closed when the function exits

    tx, err := db.BeginTx(ctx, nil)
    if err != nil {
        return err
    }
    // Defer the rollback in case of an error.
    defer tx.Rollback()

    // SQL statement to insert the link into the junction table
    sqlStatement := `
        INSERT INTO product_category (product_id, category_id)
        VALUES ($1, $2)
        ON CONFLICT (product_id, category_id) DO NOTHING;
    `
    // ON CONFLICT DO NOTHING ensures we don't fail if the link already exists.

    for _, categoryID := range categoryIDs {
        // Prepare the statement within the transaction
        _, err := tx.ExecContext(ctx, sqlStatement, productID, categoryID)
        
        if err != nil {
            // Rollback the transaction on the first error
            return err 
        }
    }

    // Commit the transaction to make all insertions permanent.
    return tx.Commit()
}

// GetCategoriesByProductID fetches all categories linked to a specific product.
func (r *PostgresProductRepository) GetCategoriesByProductID(ctx context.Context, productID int) ([]models.Category, error) {
    db, err := sql.Open("pgx", r.connString)
    fmt.Println("In get categories by product id======================")
    fmt.Println(r.connString)
    fmt.Println("====================================")
    if err != nil {
      return []models.Category{}, fmt.Errorf("failed to open database: %w", err)
    }
    defer db.Close() // Ensure the connection pool is closed when the function exits

    sqlStatement := `
        SELECT c.category_id, c.name
        FROM categories c
        JOIN product_category pc ON c.category_id = pc.category_id
        WHERE pc.product_id = $1;
    `
    rows, err := db.QueryContext(ctx, sqlStatement, productID)
    if err != nil {
        return []models.Category{}, err
    }
    defer rows.Close()

    categories := []models.Category{}
    for rows.Next() {
        var category models.Category
        // Scan the resulting columns into the Category struct fields
        if err := rows.Scan(&category.CategoryID, &category.Name); err != nil {
            return nil, err
        }
        categories = append(categories, category)
    }
    
    return categories, rows.Err()
}
