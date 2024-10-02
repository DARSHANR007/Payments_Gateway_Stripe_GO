package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/customer"
	"log"
	"os"
)

type Address struct {
	City        *string `json:"city"`
	Country     *string `json:"country"`
	Line1       *string `json:"line1"`
	Line2       *string `json:"line2"`
	Postal_code *string `json:"postal_code"`
	State       *string `json:"state"`
}

type createCustomers struct {
	Id       string   `json:"id"`
	Address  *Address `json:"address,omitempty"`
	Email    string   `json:"email"`
	Name     string   `json:"name"`
	Phone    string   `json:"phone"`
	Balance  string   `json:"balance"`
	Currency string   `json:"currency"`
}

func create_Customer(r *gin.Engine) {
	// Loading environment variables
	log.Println("Loading environment variables from api.env")
	err := godotenv.Load("api.env")
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}
	log.Println("Successfully loaded .env file")

	r.POST("/v1/customers", func(context *gin.Context) {
		log.Println("Received POST request to create customer")

		var json createCustomers
		if err := context.BindJSON(&json); err != nil {
			log.Println("Error binding JSON:", err)
			context.JSON(400, gin.H{"error": "Invalid request payload"})
			return
		}
		log.Printf("Request JSON parsed successfully: %+v\n", json)

		// Retrieve the Stripe API key from the environment
		apiKey := os.Getenv("sk_test")
		if apiKey == "" {
			log.Fatal("Stripe API key not found in environment variables")
		}
		stripe.Key = apiKey
		log.Println("Stripe API key successfully retrieved and set")

		// Create the Stripe Customer Params
		params := &stripe.CustomerParams{
			Email: stripe.String(json.Email),
			Name:  stripe.String(json.Name),
			Phone: stripe.String(json.Phone),
			Address: &stripe.AddressParams{
				City:       json.Address.City,
				Country:    json.Address.Country,
				Line1:      json.Address.Line1,
				Line2:      json.Address.Line2,
				PostalCode: json.Address.Postal_code,
				State:      json.Address.State,
			},
		}
		log.Println("Stripe customer parameters set")

		log.Println("Attempting to create customer with Stripe API")
		_, err := customer.New(params)

		if err != nil {
			log.Printf("Error creating customer: %v\n", err)
			context.JSON(500, gin.H{"error": "Failed to create customer"})
			return
		}

		log.Println("Customer successfully created in Stripe")
		context.JSON(200, gin.H{"message": "Account created successfully"})
	})
}
