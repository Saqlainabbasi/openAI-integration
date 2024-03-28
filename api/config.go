package api

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// function to get load the api key data from the env
func envOPENAIKEY() string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln(err)
		log.Fatalln("Error While loading env")
	}
	return os.Getenv("YOUR_OPENAI_API_KEY")
}
