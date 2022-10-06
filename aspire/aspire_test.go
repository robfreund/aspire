package aspire

import (
	"fmt"
	"testing"
)

func TestCapitalizeEveryThirdAlphanumericChar(t *testing.T) {
	tes := []struct{ in, expect string }{
		{"Aspiration.com", "asPirAtiOn.cOm"},
	}
	for _, ts := range tes {
		if e := CapitalizeEveryThirdAlphanumericChar(ts.in); e != ts.expect {
			t.Fatalf("expected: %s got: %s", ts.expect, e)
		}
	}
}

func TestMapString(t *testing.T) {
	tes := []struct{ in, expect string }{
		{"Aspiration.com", "asPirAtiOn.cOm"},
	}
	for _, ts := range tes {
		s := NewSkipString(3, ts.in)
		MapString(&s)
		e := fmt.Sprintf("%s", s)
		if e != ts.expect {
			t.Fatalf("expected: %s got: %s", ts.expect, e)
		}
	}
}
