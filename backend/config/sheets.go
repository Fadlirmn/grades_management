package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)
var (
	SheetId string
	CredentialPath string
) 
func SheetsConnect()  {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error load env")
	}
	SheetId = os.Getenv("SHEETS_ID")
	if SheetId == "" {
		log.Fatal("SHEETS_ID is not set in .env file")
	} 
	

	CredentialPath = os.Getenv("GOOGLE_CREDENTIALS_PATH")
	if CredentialPath == "" {
		CredentialPath="credentials.json"
	}
}