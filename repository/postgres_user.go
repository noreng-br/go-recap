package repository

import (
	"context"
	"fmt"
  "database/sql"
  "codeberg.org/noreng-br/models"
  _ "github.com/jackc/pgx/v5/stdlib"
)

type PostgresUserRepository struct {
	// db *sql.DB connection here
	connString string
}

func NewPostgresUserRepository(connString string) *PostgresUserRepository {
	// Establish and verify PG connection
	return &PostgresUserRepository{connString: connString}
}

// Create implements the CategoryRepository interface
func (r *PostgresUserRepository) CreateUser(ctx context.Context, user models.User) (models.User, error) {
    db, err := sql.Open("pgx", r.connString)
    fmt.Println("In createUser======================")
    fmt.Println(r.connString)
    fmt.Println("====================================")
    if err != nil {
      return user, fmt.Errorf("failed to open database: %w", err)
    }
    defer db.Close() // Ensure the connection pool is closed when the function exits

    // 2. SQL to insert data and return the generated ID.
    insertSQL := "INSERT INTO users (username, email, password_hash, is_admin) VALUES ($1, $2, $3, $4) RETURNING id;"

    var newUserID int
    
    // 3. Execute the query using QueryRow, passing the name and age as arguments ($1, $2).
    // Scan the returned ID into the newUserID variable.
    err = db.QueryRow(insertSQL, user.Username, user.Email, user.Password, user.IsAdmin).Scan(&newUserID)
    if err != nil {
      return user, fmt.Errorf("failed to execute insert query: %w", err)
    }

    return user, nil
}
