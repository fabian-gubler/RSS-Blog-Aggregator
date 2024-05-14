package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	PORT := os.Getenv("PORT")

	fmt.Println(PORT)
}
