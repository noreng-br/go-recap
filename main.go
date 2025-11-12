package main

import 
(
  "fmt"
  "time"
  "os"
  "log"
  "context"

  "github.com/joho/godotenv"

  "codeberg.org/noreng-br/models"
  "codeberg.org/noreng-br/repository"
  "codeberg.org/noreng-br/service"
)

func main() {

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
	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		user,
		password,
		host,
		port,
		dbname,
		sslmode,
	)

  ctx := context.Background()

  for {
    fmt.Println("hi")
    fmt.Println("THis will be an incredible application")
    rep, err := repository.NewRepositories(repository.Postgres, connStr)
    if err != nil {
      fmt.Println("An error ocurred when attempting to load repository")
    }
    fmt.Println(rep)
    s, err := service.NewService(*rep)
    if err != nil {
      fmt.Println("An error ocurred when attempting to load service")
    }
    fmt.Println(s)
    value, err := s.CreateUser(ctx, models.User{
      Username: "batata",
      Email: "batata@bat.com.br",
      Password: "banana",
      IsAdmin: false ,
    })
    if err != nil {
      log.Println("===================================")
      log.Println(value)
      log.Println("===================================")
    }
    fmt.Println(value)
    time.Sleep(1 * time.Second)
    fmt.Println(models.User{})
  }
}
