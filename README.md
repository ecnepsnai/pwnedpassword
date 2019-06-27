# Golang Pwned Passwords

[![Go Report Card](https://goreportcard.com/badge/github.com/paul-nelson-baker/go-pwnedpassword?style=flat-square)](https://goreportcard.com/report/github.com/paul-nelson-baker/go-pwnedpassword)
[![Godoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://godoc.org/github.com/paul-nelson-baker/go-pwnedpassword)
[![Releases](https://img.shields.io/github/release/paul-nelson-baker/go-pwnedpassword/all.svg?style=flat-square)](https://github.com/paul-nelson-baker/go-pwnedpassword/releases)
[![LICENSE](https://img.shields.io/github/license/paul-nelson-baker/go-pwnedpassword.svg?style=flat-square)](https://github.com/paul-nelson-baker/go-pwnedpassword/blob/master/LICENSE)

This package makes use of the "pwned passwords" API of ["Have I Been Pwned" https://haveibeenpwned.com/](https://haveibeenpwned.com/), which was created by [Troy Hunt](https://haveibeenpwned.com/About).

This is a forked version of [ecnepsnai/go-pwnedpassword](https://github.com/ecnepsnai/go-pwnedpassword), which utilizes bytes slices instead of strings. The reason for this choice is for the safety of the password itself. We can clear a slice of bytes ourselves, but a string may accidently live longer than we intend it to causing the password to be leaked logically if we're not careful. Using byte slices mitigates this risk to some extent.

# Installation

```
go get "github.com/paul-nelson-baker/go-pwnedpassword
```

# Usage

To check if a password has been compromised:

```golang
import (
    "golang.org/x/crypto/ssh/terminal"
    "syscall"
    pwn "github.com/paul-nelson-baker/go-pwnedpassword"
)

fmt.Print("Enter password: ")
passwordBytes, _ := terminal.ReadPassword(int(syscall.Stdin))

result, err := pwn.IsPwned(password)
if err != nil || result == nil {
    // Something went wrong (probably couldn't contact the pwned password API)
}

if result.Pwned {
    fmt.Printf("Sorry, this has been pwnd %d times\n", result.TimesObserved)
} else {
    fmt.Println("Cudos! This hasn't been pwnd, but be careful. It's not necessarily safe")
}
```

If you want, you can also use `pwned.IsPwnedAsync` to check asynchronously:

```golang
import "github.com/ecnepsnai/go-pwnedpassword"

pwned.IsPwnedAsync(req.Password, func(result *pwned.Result, err error) {
    if err != nil || result == nil {
        // Something went wrong (probably couldn't contact the pwned password API)
    }

    if !result.Pwned {
        // Password hasn't been seen before. Doesn't mean it's safe, just lucky.
    } else {
        count := result.TimesObserved
        // Password has been seen `count` times before.
    }
})
```

# License

MIT

go-pwnedpassword is not endorsed or affiliated with Troy Hunt, Have I Been Pwned, or Pwned Passwords.