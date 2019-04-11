package main

// Example taken from https://github.com/paul-nelson-baker/pauls-toolbox/blob/master/cmd/check-pwnd/check-pwnd.go
import (
	"fmt"
	pwn "github.com/ecnepsnai/go-pwnedpassword"
	"golang.org/x/crypto/ssh/terminal"
	"syscall"
)

func main() {
	// Error handling omitted for brevity
	// Allows user to enter password without exposing it to the console
	fmt.Print("Enter password: ")
	passwordBytes, _ := terminal.ReadPassword(int(syscall.Stdin))
	// Pass the bytes taken from the console directly to the IsPwned service
	result, _ := pwn.IsPwned(&passwordBytes)
	if result.Pwned {
		fmt.Printf("Sorry, this has been pwnd %d times\n", result.TimesObserved)
	} else {
		fmt.Println("Cudos! This hasn't been pwnd, but be careful. It's not necessarily safe")
	}
	// Clear the bytes to ensure it's not left up to the GC to clean up and is done correctly
	for i := 0; i < len(passwordBytes); i++ {
		passwordBytes[i] = ' '
	}
}
