# Golang Pwned Passwords

[![Go Report Card](https://goreportcard.com/badge/github.com/ecnepsnai/go-pwnedpassword?style=flat-square)](https://goreportcard.com/report/github.com/ecnepsnai/go-pwnedpassword)
[![Godoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://godoc.org/github.com/ecnepsnai/go-pwnedpassword)
[![Releases](https://img.shields.io/github/release/ecnepsnai/go-pwnedpassword/all.svg?style=flat-square)](https://github.com/ecnepsnai/go-pwnedpassword/releases)
[![LICENSE](https://img.shields.io/github/license/ecnepsnai/go-pwnedpassword.svg?style=flat-square)](https://github.com/ecnepsnai/go-pwnedpassword/blob/master/LICENSE)

A package to determine if a given password has been "pwned", meaning the password has been
compromised and may be used in a credential stuffing type attack. This package makes use of
the "pwned passwords" feature of "Have I Been Pwned"
https://haveibeenpwned.com/, which was created by Troy Hunt.

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