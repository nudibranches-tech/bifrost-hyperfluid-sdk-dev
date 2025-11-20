package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println("🚀 Bifrost SDK - Fluent API Demo")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()

	if err := godotenv.Load("../.env"); err != nil {
		log.Printf("⚠️  Warning: .env not loaded: %v\n", err)
	}

	fmt.Println("Running fluent API examples...")
	fmt.Println()

	runFluentAPISimpleExample()
	runFluentAPIWithSelectExample()
	runFluentAPIComplexExample()
	runFluentAPICustomOrgExample()
	runFluentAPIMultipleChainsExample()

	fmt.Println()
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println("🎉 All examples completed!")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
}
