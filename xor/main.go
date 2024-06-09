package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"xor/cipherer"
)

var mode = flag.String("mode", "cipher", "Set to 'cipher' or 'decipher'. Default is 'cipher'")
var secretKey = flag.String("secret", "", "Secret key to use")

func main() {
	flag.Parse()
	
	if len(*secretKey) == 0 {
		fmt.Println("No secret is provided. Exit...")

	}

	switch *mode {

	case "cipher":
		plainText := getUserInput("Enter your text to cipher: ")
		cipheredText, err := cipherer.Cipher(plainText, *secretKey)
		if err != nil {
			fmt.Fprintf(os.Stderr, "An error occurred: %s\n", err)
			os.Exit(1)
		}
		fmt.Println(cipheredText)

	case "decipher":
		cipherText := getUserInput("Enter your text to decipher: ")
		desipheredText, err := cipherer.Decipher(cipherText, *secretKey)
		if err != nil {
			fmt.Fprintf(os.Stderr, "An error occurred: %s\n", err)
			os.Exit(1)
		}
		fmt.Println(desipheredText)

	default:
		fmt.Println("Invalid mode. Set to 'cipher' or 'decipher'")
		os.Exit(1)
	}

}

func getUserInput(msg string) string {

	fmt.Print(msg)

	reader := bufio.NewReader(os.Stdin)

	for {

		result, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("An error occurred", err)
			continue
		}
		return strings.TrimRight(result, "\n")
	}
}
