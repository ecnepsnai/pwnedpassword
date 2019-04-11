package main

// Example taken from https://github.com/paul-nelson-baker/pauls-toolbox/blob/master/cmd/check-pwnd/check-pwnd.go

func main() {
	// Error handling omitted for brevity
	// Allows user to enter password without exposing it to the console
	fmt.Print("Enter password: ")
	password, _ := terminal.ReadPassword(int(syscall.Stdin))
	// Pass the bytes taken from the console directly to the IsPwned service
	result, _ := pwn.IsPwned(password)
	if result.Pwned {
		fmt.Printf("Sorry, this has been pwnd %d times\n", result.TimesObserved)
	} else {
		fmt.Println("Cudos! This hasn't been pwnd, but be careful it's not necessarily safe")
	}
}
