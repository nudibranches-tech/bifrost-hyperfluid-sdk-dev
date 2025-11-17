package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println("🚀 Bifrost SDK - Examples Demo")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()

	if err := godotenv.Load("../.env"); err != nil {
		log.Printf("⚠️  Warning: .env not loaded: %v\n", err)
	}

	fmt.Println("Running all examples...")
	fmt.Println()

	runPostgresExample()
	runGraphQLExample()
	runOpenAPIExample()

	fmt.Println()
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println("🎉 All examples completed!")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
}
