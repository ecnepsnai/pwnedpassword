package pwned

import (
	"crypto/rand"
	"encoding/base64"
	"testing"
)

func TestIsPwned(t *testing.T) {
	t.Run("Test Positive Match", func(t *testing.T) {
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
	})

	t.Run("Test Negative Match", func(t *testing.T) {
		// Generate Random Password
		randB := make([]byte, 16)
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
	})
}
