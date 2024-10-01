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

// ChargeJSON incoming data for Stripe API
type ChargeJSON struct {
	Amount       int64  `json:"amount"`
	ReceiptEmail string `json:"receiptEmail"`
}

func main() {
	// load .env file
	err := godotenv.Load("api.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// set up server
	r := gin.Default()

	// basic hello world GET route
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	// our basic charge API route
	r.POST("/api/charges", func(c *gin.Context) {
		// we will bind our JSON body to the `json` var
		var json ChargeJSON
		c.BindJSON(&json)

		// Set Stripe API key
		apiKey := os.Getenv("sk_test")
		stripe.Key = apiKey

		_, err := charge.New(&stripe.ChargeParams{
			Amount:       stripe.Int64(json.Amount),
			Currency:     stripe.String(string(stripe.CurrencyINR)),
			Source:       &stripe.SourceParams{Token: stripe.String("tok_visa")},
			ReceiptEmail: stripe.String(json.ReceiptEmail)})

		if err != nil {
			// Handle any errors from attempt to charge
			c.String(http.StatusBadRequest, "Request failed")
			return
		}

		c.String(http.StatusCreated, "Successfully charged")
	})

	r.Run(":8080")
}
