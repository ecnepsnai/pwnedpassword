# Golang Pwned Passwords

A package to determine if a given password has been "pwned", meaning the password has been
compromised and may be used in a credential stuffing type attack. This package makes use of
the "pwned passwords" feature of "Have I Been Pwned"
(https://haveibeenpwned.com/)[https://haveibeenpwned.com/], which was created by Troy Hunt.

# Installation

```
go get "github.com/ecnepsnai/go-pwnedpassword/
```

# Usage

To check if a password has been compromised:

```golang
import "github.com/ecnepsnai/go-pwnedpassword"

...
password := "Your Users Password"
result, err := pwned.IsPwned(password)
if err != nil {
    // Something went wrong (probably couldn't read the pwned password API)
}

if !result.Pwned {
    // Password hasn't been seen before. Doesn't mean it's safe, just lucky.
} else {
    count := result.TimesObserved
    // Password has been seen `count` times before.
}
```

If you want, you can also use `pwned.IsPwnedAsync` to check asynchronously.
