package main

import "fmt"

// Gitleaks will find this
var dbPassword = "sk_live_123456789abcdefgHIJKLMN"

func main() {
	fmt.Println(dbPassword)
}
