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

	// Legacy API examples
	fmt.Println("═══════════════════════════════════════")
	fmt.Println("📚 LEGACY API EXAMPLES")
	fmt.Println("═══════════════════════════════════════")
	fmt.Println()

	runPostgresExample()
	runGraphQLExample()
	runOpenAPIExample()

	// New Fluent API examples
	fmt.Println()
	fmt.Println("═══════════════════════════════════════")
	fmt.Println("✨ NEW FLUENT API EXAMPLES")
	fmt.Println("═══════════════════════════════════════")
	fmt.Println()

	runFluentAPISimpleExample()
	runFluentAPIWithSelectExample()
	runFluentAPIComplexExample()
	runFluentAPICustomOrgExample()
	runFluentAPIMultipleChainsExample()
	runFluentAPIComparisonExample()

	fmt.Println()
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println("🎉 All examples completed!")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
}
