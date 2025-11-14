// handler/handler.go
package handler

import (
  "os"
  "net/http"

  "github.com/golang-jwt/jwt/v5"
  echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"

  "codeberg.org/noreng-br/models"
)

// AdminOnly is an Echo middleware that checks if the user's role claim matches the required role.
func AdminOnly(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
        // 1. Get the validated JWT object from the context (set by echo-jwt)
        user := c.Get("user").(*jwt.Token)
        
        // 2. Extract the custom claims
        claims := user.Claims.(*models.JWTCustomClaims)

        // 3. Check the user's role
        if claims.Role != "admin" {
            // Role is not "admin", deny access.
            return c.JSON(http.StatusForbidden, echo.Map{
                "message": "Access Denied: Admin privileges required.",
            })
        }

        // 4. Role check passed, proceed to the next handler
        return next(c)
    }
}

func restricted(c echo.Context) error {
    // The JWT middleware successfully verified the token.
    // The token object is stored in the context under the key "user" by default.
    user := c.Get("user").(*jwt.Token)
    claims := user.Claims.(*models.JWTCustomClaims)

    // You can now access the data from the token's payload
    userID := claims.UserID
    userRole := claims.Role

    return c.JSON(http.StatusOK, echo.Map{
        "message":  "Welcome to the restricted area!",
        "user_id":  userID,
        "user_role": userRole,
    })
}


// InitRoutes sets up all API routes on the provided Echo instance.
func InitRoutes(e *echo.Echo, h *Handler) {
	// Group routes under a common prefix, like /v1
	api := e.Group("/api")

	// --- User Routes ---
	// Since GetUsers and GetUserByID are in the same package (handler),
	// they can be called directly without any prefix.
	api.POST("/auth", h.Login)
	api.POST("/users", h.CreateUser)
  api.GET("/products", h.GetProducts)
  api.GET("/product", h.GetProduct)

  // 2. Restricted Group
  r := e.Group("/api/user")

  secretKey := os.Getenv("JWT_SECRET")

  // Configure the JWT Middleware
  config := echojwt.Config{
      NewClaimsFunc: func(c echo.Context) jwt.Claims {
          // Tells the middleware which struct to use for claims
          return new(models.JWTCustomClaims) 
      },
      SigningKey: []byte(secretKey),
  }

  jwtMiddleware := echojwt.WithConfig(config)

  // Apply the JWT middleware to the restricted group
  r.Use(jwtMiddleware)
  
  // Protected route
  r.POST("/order", restricted)

  adminGroup := e.Group("/api/admin")
  adminGroup.Use(jwtMiddleware) // First, validate the token
  adminGroup.Use(AdminOnly)

  /*Only Admin can perform*/
	adminGroup.GET("/users", h.GetUsers)
  adminGroup.POST("/product", h.CreateProduct)
  adminGroup.POST("/category", h.CreateCategory)
  adminGroup.PATCH("/product-categories", h.AddProductCategories)
  adminGroup.PUT("/product", h.UpdateProduct)
  adminGroup.GET("/categories", h.GetCategories)
  adminGroup.DELETE("/category", h.DeleteCategory)
}
