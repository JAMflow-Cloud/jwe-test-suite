package main

import (
	"flag"
	"fmt"
	"log"

	jose "github.com/dvsekhvalnov/jose2go"
)

func main() {
	// Define command-line flags for password and JWE token.
	password := flag.String("password", "", "Password for decryption")
	token := flag.String("jwe", "", "JWE token to decrypt")
	flag.Parse()

	// Check that both parameters were provided.
	if *password == "" || *token == "" {
		log.Fatal("Both -password and -jwe flags must be provided. Example: ./cli -password=yourpass -jwe=yourtoken")
	}

	// Decrypt the JWE token using the provided password.
	plaintext, _, err := jose.Decode(*token, *password)
	if err != nil {
		log.Fatalf("Failed to decrypt the JWE token: %v", err)
	}

	// Print the decrypted content.
	fmt.Println("Decrypted content:")
	fmt.Println(string(plaintext))
}
