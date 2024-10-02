package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/charge"
)

type ChargeJSON struct {
	Amount       int64  `json:"amount"`
	ReceiptEmail string `json:"receiptEmail"`
}

func main() {
	err := godotenv.Load("api.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		// Read the HTML file
		htmlContent, err := os.ReadFile("welcome.html")
		if err != nil {
			log.Println("Error reading welcome.html:", err)
			c.String(http.StatusInternalServerError, "Internal Server Error")
			return
		}

		c.Data(http.StatusOK, "text/html; charset=utf-8", htmlContent)
	})

	r.POST("/api/charges", func(c *gin.Context) {
		var json ChargeJSON
		if err := c.BindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		apiKey := os.Getenv("sk_test")
		stripe.Key = apiKey

		_, err := charge.New(&stripe.ChargeParams{
			Amount:       stripe.Int64(json.Amount),
			Currency:     stripe.String(string(stripe.CurrencyINR)),
			Source:       &stripe.SourceParams{Token: stripe.String("tok_visa")},
			ReceiptEmail: stripe.String(json.ReceiptEmail),
		})

		if err != nil {
			c.String(http.StatusBadRequest, "Request failed")
			return
		}

		c.String(http.StatusCreated, "Successfully charged")
	})

	r.Run(":8081")
}
