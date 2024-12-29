package common

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func LoadClient(url string) *http.Response {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get credentials
	sessionCookie := os.Getenv("SESSION_COOKIE")

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.AddCookie(&http.Cookie{
		Name:  "session",
		Value: sessionCookie,
	})

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error fetching URL:", err)
		return nil
	}

	return resp
}
