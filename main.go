package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stripe/stripe-go/v79"
	"log"
	"os"
)

type Charge struct {
	Amount        int64  `json:"amount"`
	RecipientMail string `json:"recipient_mail"`
}

func main() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("error loading: ", err)
	}

	r := gin.Default()

	r.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{"message": "This is the customers page"})

	})

	r.POST("/api/charges", func(context *gin.Context) {

		var json Charge
		err := context.BindJSON(json)
		if err != nil {
			return
		}

		apiKey := os.Getenv("sk_test")
		stripe.Key = apiKey

	})

}
