package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)
var Apikey string
func AiConnect()  {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error load env")
	}
	Apikey= fmt.Sprintf("apikey=%s",
	os.Getenv("API_KEY"),
	)
	if Apikey == "" {
		log.Fatal("GEMINI_API_KEY is not set in .env file")
	} 
}