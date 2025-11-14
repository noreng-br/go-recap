package repository

import (
    "context"
    "fmt"
    "time"
    "strconv"
    "errors"
    "log"
    "database/sql"

    "codeberg.org/noreng-br/models"
)

var ErrProductNotFound = errors.New("product not found")

type PostgresOrderRepository struct {
  connString string
}

func NewPostgresOrderRepository(connString string) *PostgresOrderRepository {
  return &PostgresOrderRepository{connString: connString}
}

func (r *PostgresOrderRepository) CreateOrder(ctx context.Context, req models.OrderRequest) (int, error) {
  log.Println(ctx)

  db, err := sql.Open("pgx", r.connString)
  fmt.Println("In createOrder======================")
  fmt.Println(r.connString)
  fmt.Println("====================================")
  if err != nil {
    return 0, fmt.Errorf("failed to open database: %w", err)
  }
  defer db.Close() // Ensure the connection pool is closed when the function exits

  // 1. Start Transaction
  tx, err := db.BeginTx(ctx, nil)
  if err != nil {
      return 0, err
  }
  // Defer rollback, ensuring it runs if Commit() isn't called due to an error.
  defer tx.Rollback() 

  // 2. Insert into the 'orders' table
  var newOrderID int
  insertOrderSQL := `
      INSERT INTO orders (user_id, ordered_date, status)
      VALUES ($1, $2, $3)
      RETURNING order_id;` // RETURNING gets the new PostgreSQL serial ID

  err = tx.QueryRowContext(ctx, insertOrderSQL, 
      req.UserID, time.Now(), req.Status).Scan(&newOrderID)
  if err != nil {
      return 0, err
  }

  // 3. Process each item and insert into 'order_product' junction table
  insertItemSQL := `
      INSERT INTO order_product (order_id, product_id, quantity, unit_price)
      VALUES ($1, $2, $3, $4);`
  
  getPriceSQL := `
      SELECT price FROM products WHERE product_id = $1;`

  for _, item := range req.Items {
      // A. Retrieve current unit price from 'products' table
      var unitPrice float64
      err := tx.QueryRowContext(ctx, getPriceSQL, item.ProductID).Scan(&unitPrice)
      
      if err != nil {
           // Handle product not found or query error
           if errors.Is(err, sql.ErrNoRows) {
               return 0, ErrProductNotFound
           }
           return 0, err
      }
      
      // B. Insert into the 'order_product' junction table
      _, err = tx.ExecContext(ctx, insertItemSQL, 
          newOrderID, item.ProductID, item.Quantity, unitPrice)
      if err != nil {
          return 0, err
      }
  }

  // 4. Commit Transaction
  err = tx.Commit()
  if err != nil {
      return 0, err
  }

  return newOrderID, nil
}


// GetUserOrders fetches all orders for a given user and populates their products.
func (r *PostgresOrderRepository) GetUserOrders(ctx context.Context, userID int) ([]models.OrderDetails, error) {
    log.Println(ctx)

    db, err := sql.Open("pgx", r.connString)
    fmt.Println("In createOrder======================")
    fmt.Println(r.connString)
    fmt.Println("====================================")
    if err != nil {
      return []models.OrderDetails{}, fmt.Errorf("failed to open database: %w", err)
    }
    defer db.Close() // Ensure the connection pool is closed when the function exits
    
    // STEP 1: Fetch all basic order details for the user
    ordersSQL := `
        SELECT order_id, ordered_date, delivered_date, status, user_id
        FROM orders
        WHERE user_id = $1
        ORDER BY ordered_date DESC;`
    
    rows, err := db.QueryContext(ctx, ordersSQL, userID)
    if err != nil {
        return []models.OrderDetails{}, err
    }
    defer rows.Close()

    // Map to store orders by ID for easy lookup
    ordersMap := make(map[int]models.OrderDetails)
    orderIDs := []int{}

    for rows.Next() {
        var order models.OrderDetails
        var deliveredDate sql.NullTime // Use sql.NullTime to handle NULL dates

        err := rows.Scan(
            &order.OrderID,
            &order.OrderedDate,
            &deliveredDate, // Scan into NullTime helper
            &order.Status,
            &order.UserID,
        )
        if err != nil {
            return nil, err
        }
        
        // Convert sql.NullTime to *time.Time
        if deliveredDate.Valid {
            order.DeliverDate = &deliveredDate.Time
        }

        ordersMap[order.OrderID] = order
        orderIDs = append(orderIDs, order.OrderID)
    }
    if rows.Err() != nil {
        return nil, rows.Err()
    }
    
    if len(orderIDs) == 0 {
        return []models.OrderDetails{}, nil // Return empty list if no orders found
    }

    // STEP 2: Fetch all order items for ALL fetched orders using the junction table
    // Uses the PostgreSQL ANY operator for efficiency with multiple IDs.
    itemsSQL := `
        SELECT 
            op.order_id, 
            p.product_id, 
            p.name, 
            op.quantity, 
            op.unit_price
        FROM order_product op
        JOIN products p ON op.product_id = p.product_id
        WHERE op.order_id = ANY($1::int[]);` // $1::int[] casts the slice to an array type
    
    // Convert slice of ints to a string array format for the ANY clause
    orderIDsArray := "{" + func (arr []int) string {
      s := ""
      for i, v := range arr {
          s += strconv.Itoa(v)
          if i < len(arr)-1 {
              s += ","
          }
      }
      return s
    }(orderIDs) + "}"
    
    itemRows, err := db.QueryContext(ctx, itemsSQL, orderIDsArray)
    if err != nil {
        return nil, err
    }
    defer itemRows.Close()

    // STEP 3: Map items back to their parent orders
    for itemRows.Next() {
        var item models.OrderItem
        var orderID int
        
        err := itemRows.Scan(
            &orderID,
            &item.ProductID,
            &item.Name,
            &item.Quantity,
            &item.UnitPrice,
        )
        if err != nil {
            return nil, err
        }

        // Retrieve the order from the map, append the item, and save it back
        order := ordersMap[orderID]
        order.Items = append(order.Items, item)
        ordersMap[orderID] = order // Update the map
    }
    if itemRows.Err() != nil {
        return nil, itemRows.Err()
    }

    // Convert the map values back to a slice for the final return
    finalOrders := make([]models.OrderDetails, 0, len(ordersMap))
    for _, order := range ordersMap {
        finalOrders = append(finalOrders, order)
    }

    return finalOrders, nil
}

func (r *PostgresOrderRepository) ListOrders(ctx context.Context) ([]models.OrderDetails, error) {
    db, err := sql.Open("pgx", r.connString)
    fmt.Println("In List all Orders======================")
    fmt.Println(r.connString)
    fmt.Println("====================================")
    if err != nil {
      return []models.OrderDetails{}, fmt.Errorf("failed to open database: %w", err)
    }
    defer db.Close() // Ensure the connection pool is closed when the function exits
    // STEP 1: Fetch all basic order details
    ordersSQL := `
        SELECT order_id, ordered_date, delivered_date, status, user_id
        FROM orders
        ORDER BY ordered_date DESC;` // Fetch all orders, no WHERE clause
    
    rows, err := db.QueryContext(ctx, ordersSQL)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    ordersMap := make(map[int]models.OrderDetails)
    orderIDs := []int{}

    for rows.Next() {
        var order models.OrderDetails
        var deliveredDate sql.NullTime // Handles NULL dates from the DB

        err := rows.Scan(
            &order.OrderID,
            &order.OrderedDate,
            &deliveredDate,
            &order.Status,
            &order.UserID,
        )
        if err != nil {
            return nil, err
        }
        
        // Convert sql.NullTime to *time.Time
        if deliveredDate.Valid {
            order.DeliverDate = &deliveredDate.Time
        }

        ordersMap[order.OrderID] = order
        orderIDs = append(orderIDs, order.OrderID)
    }
    if rows.Err() != nil {
        return nil, rows.Err()
    }
    
    if len(orderIDs) == 0 {
        return []models.OrderDetails{}, nil // Return empty list
    }

    // STEP 2: Fetch all order items for ALL fetched orders
    // This query is highly efficient as it hits the DB only once for all items.
    itemsSQL := `
        SELECT 
            op.order_id, 
            p.product_id, 
            p.name, 
            op.quantity, 
            op.unit_price
        FROM order_product op
        JOIN products p ON op.product_id = p.product_id
        WHERE op.order_id = ANY($1::int[]);`
    
    // Prepare the array of IDs for the PostgreSQL ANY operator
    orderIDsArray := "{" + func (arr []int) string {
      s := ""
      for i, v := range arr {
          s += strconv.Itoa(v)
          if i < len(arr)-1 {
              s += ","
          }
      }
      return s
    }(orderIDs) + "}"
    
    itemRows, err := db.QueryContext(ctx, itemsSQL, orderIDsArray)
    if err != nil {
        return nil, err
    }
    defer itemRows.Close()

    // STEP 3: Map items back to their parent orders
    for itemRows.Next() {
        var item models.OrderItem
        var orderID int
        
        err := itemRows.Scan(
            &orderID,
            &item.ProductID,
            &item.Name,
            &item.Quantity,
            &item.UnitPrice,
        )
        if err != nil {
            return nil, err
        }

        // Retrieve the order, append the item, and save it back to the map
        order := ordersMap[orderID]
        order.Items = append(order.Items, item)
        ordersMap[orderID] = order 
    }
    if itemRows.Err() != nil {
        return nil, itemRows.Err()
    }

    // STEP 4: Convert the map into a slice for the final return
    finalOrders := make([]models.OrderDetails, 0, len(ordersMap))
    for _, order := range ordersMap {
        finalOrders = append(finalOrders, order)
    }

    return finalOrders, nil
}
