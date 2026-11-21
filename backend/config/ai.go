package config

import (
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
	val := os.Getenv("API_KEY")
	if val == "" {
		log.Fatal("GEMINI_API_KEY is not set in .env file")
	} 
	Apikey = val
}