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

	// Old fluent API examples (backward compatibility)
	fmt.Println("══════════════════════════════════════════════")
	fmt.Println("📚 SIMPLE FLUENT API (Catalog-first)")
	fmt.Println("══════════════════════════════════════════════")
	fmt.Println()

	runFluentAPISimpleExample()
	runFluentAPIWithSelectExample()
	runFluentAPIComplexExample()

	// NEW! Progressive fluent API examples
	fmt.Println()
	fmt.Println("══════════════════════════════════════════════")
	fmt.Println("✨ PROGRESSIVE FLUENT API (Type-safe)")
	fmt.Println("══════════════════════════════════════════════")
	fmt.Println()

	runProgressiveAPIExample1()
	runProgressiveAPIExample2()
	runProgressiveAPIExample3()
	runProgressiveAPIExample4()
	runProgressiveAPIExample5()
	runProgressiveAPIExample6()
	runProgressiveAPIListingExample()

	fmt.Println()
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println("🎉 All examples completed!")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
}
