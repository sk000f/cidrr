package cidr_test

import (
	"testing"

	"github.com/sk000f/cidrr/int/cidr"
)

func TestParse(t *testing.T) {
	t.Run("192.168.1.0/32 is valid subnet", func(t *testing.T) {
		got := cidr.Parse("192.168.1.0/32")
		want := true

		if got != want {
			t.Errorf("got %v; want %v", got, want)
		}
	})
}
