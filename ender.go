package main

import (
	"crypto/aes"
	"crypto/cipher"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func banner() {

	banner := `
              ______  _____     ___  _____   __  
        /\   |  ____|/ ____|   |__ \| ____| / /  
       /  \  | |__  | (___ ______ ) | |__  / /_  
      / /\ \ |  __|  \___ \______/ /|___ \| '_ \ 
     / ____ \| |____ ____) |    / /_ ___) | (_) |
    /_/    \_\______|_____/    |____|____/ \___/ 
    ----------------------------------------------
    Ender - AES-256 file encryption/decryption tool																						
												 `
	fmt.Println(banner)

}

func main() {
	var key string
	var iv string
	version := flag.Bool("v", false, "Print version")
	enBool := flag.Bool("e", false, "Encrypt AES-256")
	deBool := flag.Bool("d", false, "Decrypt AES-256")
	filepath := flag.String("f", "", "File path")
	flag.Parse()
	banner()

	if *version == true {
		fmt.Println("                  Version: 1.0.0")
		fmt.Println("                 Author: @spix-777")
		os.Exit(0)
	}

	if *filepath == "" {
		fmt.Println("Please specify a file path")
		os.Exit(0)
	}

	inputFile := *filepath

	if *enBool == false && *deBool == false {
		fmt.Println("Please specify -e or -d")
		return
	} else if *enBool == true && *deBool == true {
		fmt.Println("Please specify -e or -d")
		return
	} else if *enBool == true && *deBool == false {
		fmt.Println("Encrypting file...")

		// Generate a random key
		characters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ01234567890123456789"
		randomCharString := generateRandomString(32, characters)
		randomDigitString := generateRandomString(16, characters)
		key = randomCharString
		iv = randomDigitString

		fmt.Println("Key:", key)
		fmt.Println("IV:", iv)
		fmt.Println("File path:", inputFile, "\n")

		err := encryptFile([]byte(key), []byte(iv), inputFile)
		if err != nil {
			fmt.Println("Error encrypting file:", err)
			return
		}
		fmt.Println("File encrypted successfully.")
	} else if *enBool == false && *deBool == true {
		fmt.Println("Decrypting file...")
		// Generate key
		fmt.Println("Please enter key and iv")
		fmt.Printf("Key: ")
		fmt.Scanf("%s", &key)
		fmt.Printf("IV: ")
		fmt.Scanf("%s", &iv)

		fmt.Println("File path:", inputFile, "\n")

		err := decryptFile([]byte(key), []byte(iv), inputFile)
		if err != nil {
			fmt.Println("Error decrypting file:", err)
			return
		}
		fmt.Println("File decrypted successfully.")
	}
}

func generateRandomString(length int, chars string) string {
	rand.Seed(time.Now().UnixNano())
	result := make([]byte, length)
	for i := range result {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}

func encryptFile(key, iv []byte, inputFile string) error {
	// Read the content of the input file
	plaintext, err := os.ReadFile(inputFile)
	if err != nil {
		return err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	ciphertext := make([]byte, len(plaintext))
	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(ciphertext, plaintext)

	// Write the encrypted content to the output file
	err = os.WriteFile(inputFile, ciphertext, 0644)
	if err != nil {
		return err
	}

	return nil
}

func decryptFile(key, iv []byte, inputFile string) error {
	ciphertext, err := os.ReadFile(inputFile)
	if err != nil {
		return err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	plaintext := make([]byte, len(ciphertext))
	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(plaintext, ciphertext)

	err = os.WriteFile(inputFile, plaintext, 0644)
	if err != nil {
		return err
	}

	return nil
}
