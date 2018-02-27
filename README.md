# Golang Pwned Passwords

A package to determine if a given password has been "pwned", meaning the password has been
compromised and may be used in a credential stuffing type attack. This package makes use of
the "pwned passwords" feature of "Have I Been Pwned"
https://haveibeenpwned.com/, which was created by Troy Hunt.

# Documentation

Documentation available on GoDoc

<a href="https://godoc.org/github.com/ecnepsnai/go-pwnedpassword"><img src="https://godoc.org/github.com/ecnepsnai/go-pwnedpassword?status.svg" alt="GoDoc"></a>

# Installation

```
go get "github.com/ecnepsnai/go-pwnedpassword/
```

# Usage

To check if a password has been compromised:

```golang
import "github.com/ecnepsnai/go-pwnedpassword"

password := "Your Users Password"
result, err := pwned.IsPwned(password)
if err != nil {
    // Something went wrong (probably couldn't contact the pwned password API)
}

if !result.Pwned {
    // Password hasn't been seen before. Doesn't mean it's safe, just lucky.
} else {
    count := result.TimesObserved
    // Password has been seen `count` times before.
}
```

If you want, you can also use `pwned.IsPwnedAsync` to check asynchronously:

```golang
import "github.com/ecnepsnai/go-pwnedpassword"

pwned.IsPwnedAsync(req.Password, func(result *pwned.Result, err error) {
    if err != nil {
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