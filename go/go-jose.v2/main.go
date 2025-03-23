package main

import (
	"flag"
	"fmt"
	"log"

	"gopkg.in/square/go-jose.v2"
)

func main() {
	// Define command-line flags for password and jwe token.
	password := flag.String("password", "", "Password for decryption")
	token := flag.String("jwe", "", "JWE token to decrypt")
	flag.Parse()

	// Check that both parameters were provided.
	if *password == "" || *token == "" {
		log.Fatal("Both -password and -jwe flags must be provided. Example: ./app -password=yourpass -jwe=yourtoken")
	}

	// Parse the JWE token.
	jweObject, err := jose.ParseEncrypted(*token)
	if err != nil {
		log.Fatalf("Failed to parse JWE token: %v", err)
	}

	// Decrypt the JWE token using the provided password as the decryption key.
	// In a real-world scenario, consider deriving a proper key from the password.
	plaintext, err := jweObject.Decrypt([]byte(*password))
	if err != nil {
		log.Fatalf("Failed to decrypt the JWE token: %v", err)
	}

	// Print the decrypted content.
	fmt.Println("Decrypted content:")
	fmt.Println(string(plaintext))
}