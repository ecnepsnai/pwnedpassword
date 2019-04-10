// Package pwned A package to determine if a given password has been "pwned", meaning the password has been compromised
// and may be used in a credential stuffing type attack. This package makes use of the "pwned passwords" feature of
// "Have I Been Pwned" https://haveibeenpwned.com/Passwords, which was created by Troy Hunt.
package pwned

import (
	"crypto/sha1"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

// Result describes a result from the Pwned Password service.
type Result struct {
	// Pwned has the password been seen at least once. A value of false doesn't mean the password is any good though.
	Pwned bool
	// TimesObserved the number of times this password has been seen by the pwned password service.
	TimesObserved uint64
}

type pwnedHash struct {
	Hash  string
	Range string
}

// IsPwnedAsync will asynchronously check if the provided password has been pwned. Calls `cb` with the result when finished.
func IsPwnedAsync(password []byte, cb func(*Result, error)) {
	go func() {
		cb(IsPwned(password))
	}()
}

// IsPwned will synchronously check if the provided password has been pwned.
func IsPwned(password []byte) (*Result, error) {
	if len(password) == 0 {
		return nil, fmt.Errorf("empty password provided")
	}

	hash, err := getHash(password)
	if err != nil {
		return nil, err
	}

	resp, err := http.Get("https://api.pwnedpasswords.com/range/" + hash.Range)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = resp.Body.Close()
	}()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(body), "\r\n")
	for _, line := range lines {
		components := strings.Split(line, ":")
		if len(components) != 2 {
			return nil, fmt.Errorf("invalid response from pwned password API")
		}

		resultHash := components[0]
		countStr := components[1]

		if hash.Range+resultHash == hash.Hash {
			count, err := strconv.ParseUint(countStr, 10, 64)
			if err != nil {
				return nil, err
			}

			ret := Result{
				Pwned:         true,
				TimesObserved: count,
			}
			return &ret, nil
		}
	}

	ret := Result{
		Pwned:         false,
		TimesObserved: 0,
	}
	return &ret, nil
}

func getHash(password []byte) (*pwnedHash, error) {
	h := sha1.New()
	_, err := h.Write(password)
	if err != nil {
		return nil, err
	}
	hash := fmt.Sprintf("%x", h.Sum(nil))
	hash = strings.ToUpper(hash)
	if len(hash) < 5 {
		return nil, fmt.Errorf("unable to hash password")
	}

	result := pwnedHash{
		Hash:  hash,
		Range: hash[0:5],
	}

	return &result, nil
}
