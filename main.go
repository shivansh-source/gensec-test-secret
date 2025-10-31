package main

import (
	"fmt"
	"os"
)

// Gitleaks will not find this
var dbPassword = os.Getenv("DB_PASSWORD")

func main() {
	fmt.Println(dbPassword)
}