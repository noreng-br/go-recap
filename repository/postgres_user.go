package repository

import (
	"context"
	"fmt"
  "database/sql"

  _ "github.com/jackc/pgx/v5/stdlib"

  "codeberg.org/noreng-br/models"
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

func (r *PostgresUserRepository) GetUserByUsername(ctx context.Context, username string) (models.User, error) {
    var user models.User

    db, err := sql.Open("pgx", r.connString)
    fmt.Println("In get user by username======================")
    fmt.Println(r.connString)
    fmt.Println("====================================")
    if err != nil {
      return user, fmt.Errorf("failed to open database: %w", err)
    }
    defer db.Close() // Ensure the connection pool is closed when the function exits

    selectSql := "SELECT * from users where username=$1";

    err = db.QueryRowContext(ctx, selectSql, username).Scan(
      &user.ID,
      &user.Username,
      &user.Email,
      &user.Password,
      &user.IsAdmin,
    )
    if err != nil {
      fmt.Println("===========================================")
      fmt.Println("Query error")
      fmt.Println(err.Error())
      fmt.Println("===========================================")
      return user, fmt.Errorf("failed to execute select query: %w", err)
    }
    return user, nil
}

func (r *PostgresUserRepository) GetUserByEmail(ctx context.Context, email string) (models.User, error) {
    var user models.User

    db, err := sql.Open("pgx", r.connString)
    fmt.Println("In get user by email======================")
    fmt.Println(r.connString)
    fmt.Println("====================================")
    if err != nil {
      return user, fmt.Errorf("failed to open database: %w", err)
    }
    defer db.Close() // Ensure the connection pool is closed when the function exits

    selectSql := "SELECT * from users where email=$1";

    err = db.QueryRowContext(ctx, selectSql, email).Scan(
      &user.ID,
      &user.Username,
      &user.Email,
      &user.Password,
      &user.IsAdmin,
    )
    if err != nil {
      fmt.Println("===========================================")
      fmt.Println("Query error")
      fmt.Println(err.Error())
      fmt.Println("===========================================")
      return user, fmt.Errorf("failed to execute select query: %w", err)
    }
    return user, nil
}

func (r *PostgresUserRepository) GetUsers(ctx context.Context) ([]models.User, error) {
    var users []models.User

    db, err := sql.Open("pgx", r.connString)
    fmt.Println("In get user by username======================")
    fmt.Println(r.connString)
    fmt.Println("====================================")
    if err != nil {
      return users, fmt.Errorf("failed to open database: %w", err)
    }
    defer db.Close() // Ensure the connection pool is closed when the function exits

    selectSql := "SELECT * from users";

    rows, err := db.Query(selectSql)
    if err != nil {
      fmt.Println("===========================================")
      fmt.Println("Query error")
      fmt.Println(err.Error())
      fmt.Println("===========================================")
      return users, fmt.Errorf("failed to execute select query: %w", err)
    }

    for rows.Next() {
      var user models.User

      rows.Scan(
        &user.ID,
        &user.Username,
        &user.Email,
        &user.Password,
        &user.IsAdmin,
      )

      users = append(users, user)
    }

    return users, nil
}

