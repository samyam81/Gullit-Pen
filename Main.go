package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Define a struct for Order
type Order struct {
	ID             uint   `gorm:"primary_key"`
	FirstName      string
	LastName       string
	PhoneNumber    string
	BillingAddress string
	Email          string
	ProductQuantity int
	ProductColor   string
	PaymentMethod  string
}

func main() {
	// Initialize Gin
	r := gin.Default()

	// Setup database connection
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=your_postgres_username dbname=ruud_pen password=your_postgres_password sslmode=disable")
	if err != nil {
		panic("Failed to connect to database")
	}
	defer db.Close()

	// Auto migrate the Order struct
	db.AutoMigrate(&Order{})

	// Define routes for serving static files
	r.StaticFS("/", http.Dir("./static"))

	// Route to handle form submission and store data
	r.POST("/api/orders/submit", func(c *gin.Context) {
		var order Order
		if err := c.BindJSON(&order); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON payload"})
			return
		}

		// Create a new order record in the database
		if err := db.Create(&order).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save order"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Order submitted successfully"})
	})

	// Serve static HTML files for different routes
	http.HandleFunc("/GullitPen", serveStaticFile("./static/seven.html"))
	http.HandleFunc("/seven20", serveStaticFile("./static/seven20.html"))
	http.HandleFunc("/seven2", serveStaticFile("./static/seven2.html"))
	http.HandleFunc("/seven", serveStaticFile("./static/seven.html"))
	http.HandleFunc("/seven0", serveStaticFile("./static/seven0.html"))
	http.HandleFunc("/seven10", serveStaticFile("./static/seven10.html"))
	http.HandleFunc("/seven1", serveStaticFile("./static/seven1.html"))

	// Run the HTTP server
	go func() {
		fmt.Println("Server is running on http://localhost:8080")
		if err := http.ListenAndServe(":8080", nil); err != nil {
			fmt.Printf("Error starting server: %s\n", err)
		}
	}()

	// Run the Gin server
	r.Run(":8081")
}

// Helper function to serve static files
func serveStaticFile(path string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, path)
	}
}
