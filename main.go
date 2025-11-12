package main

import 
(
  "fmt"
  "os"
  "log"
  "context"

  "github.com/joho/godotenv"
  "github.com/labstack/echo/v4"
  "github.com/labstack/echo/v4/middleware"
  "github.com/go-playground/validator/v10"

  "codeberg.org/noreng-br/handler"
  "codeberg.org/noreng-br/repository"
  "codeberg.org/noreng-br/service"
)

var connStr string
var ctx context.Context

func start() {
  fmt.Println("=======================================================")
  fmt.Println("=======================================================")
  fmt.Println("=======================================================")
  fmt.Println("=====================INIT========================")
  fmt.Println("=======================================================")
  fmt.Println("=======================================================")
  fmt.Println("=======================================================")
  // load env
  err := godotenv.Load()
  if err != nil {
    log.Fatalf("Error loading .env file: %v", err)
  }

  // Get individual components
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	dbname := os.Getenv("POSTGRES_DB")
	sslmode := os.Getenv("POSTGRES_SSL")

	// Construct the URI connection string
	connStr = fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		user,
		password,
		host,
		port,
		dbname,
		sslmode,
	)

  ctx = context.Background()

  fmt.Println(connStr, ctx)
  fmt.Println("THis will be an incredible application")


  jwtSecret := os.Getenv("JWT_SECRET")
  fmt.Println("================")
  fmt.Println(jwtSecret)
  fmt.Println("================")
}

func main() {
  e := echo.New()
  e.Validator = &CustomValidator{validator: validator.New()}
  start()

  rep, err := repository.NewRepositories(repository.Postgres, connStr)
  if err != nil {
    fmt.Println("An error ocurred when attempting to load repository")
  }
  s, err := service.NewService(*rep)
  if err != nil {
    fmt.Println("An error ocurred when attempting to load service")
  }

  h, err := handler.NewHandler(*s)
  if err != nil {
    fmt.Println("error in handler")
  }

	// Global Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Initialize and register all routes from the 'handler' package
	handler.InitRoutes(e, h)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
