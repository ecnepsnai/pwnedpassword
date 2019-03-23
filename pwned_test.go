package pwned

import (
	"crypto/rand"
	"encoding/base64"
	"sync"
	"testing"
)

func TestPositiveMatch(t *testing.T) {
	t.Parallel()

	rt, err := IsPwned("password")
	if err != nil {
		t.Errorf("IsPwned() error = %s", err.Error())
		return
	}

	if !rt.Pwned {
		t.Errorf("TestIsPwned fail: want Pwned: true got, Pwned: false")
	}

	if rt.TimesObserved == 0 {
		t.Errorf("TestIsPwned fail: want TimesObserved: 0 got, TimesObserved: %d", rt.TimesObserved)
	}
}

func TestNegativeMatch(t *testing.T) {
	t.Parallel()

	// Generate Random Password
	randB := make([]byte, 32)
	rand.Read(randB)
	password := base64.StdEncoding.EncodeToString(randB)

	rt, err := IsPwned(password)
	if err != nil {
		t.Errorf("IsPwned() error = %s", err.Error())
		return
	}

	if rt.Pwned {
		t.Errorf("TestIsPwned fail: want Pwned: false got, Pwned: true")
	}

	if rt.TimesObserved > 0 {
		t.Errorf("TestIsPwned fail: want TimesObserved: >0 got, TimesObserved: 0")
	}
}

func TestAsync(t *testing.T) {
	t.Parallel()

	wg := sync.WaitGroup{}
	wg.Add(1)

	var result *Result
	var err error

	IsPwnedAsync("password", func(r *Result, e error) {
		result = r
		err = e
		wg.Done()
	})

	wg.Wait()

	if err != nil {
		t.Errorf("IsPwned() error = %s", err.Error())
		return
	}

	if !result.Pwned {
		t.Errorf("TestIsPwned fail: want Pwned: true got, Pwned: false")
	}

	if result.TimesObserved == 0 {
		t.Errorf("TestIsPwned fail: want TimesObserved: 0 got, TimesObserved: %d", result.TimesObserved)
	}
}

func TestEmptyPassword(t *testing.T) {
	t.Parallel()

	_, err := IsPwned("")
	if err == nil {
		t.Errorf("No error seen when one expected")
		return
	}
}
