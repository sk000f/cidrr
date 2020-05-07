package cidr_test

import (
	"testing"

	"github.com/sk000f/cidrr/int/cidr"
)

var parseTests = []struct {
	name  string
	input string
	want  bool
	err   error
}{
	{
		"192.168.1.0/32 is a valid subnet",
		"192.168.1.0/32",
		true,
		nil,
	},
	{
		"256.256.256.256/256 is not a valid subnet",
		"256.256.256.256/256",
		false,
		nil,
	},
	{
		"255.255.255.250/32 is a valid subnet",
		"255.255.255.250/32",
		true,
		nil,
	},
	{
		"0.0.0.0/0 is a valid subnet",
		"0.0.0.0/0",
		true,
		nil,
	},
	{
		"0.0.0.0/33 is not a valid subnet",
		"0.0.0.0/33",
		false,
		nil,
	},
	{
		"0.0.0.0/-1 is not a valid subnet",
		"0.0.0.0/-1",
		false,
		nil,
	},
}

func TestValidate(t *testing.T) {
	for _, tt := range parseTests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := cidr.Validate(tt.input)

			if err != nil {

			}

			if got != tt.want {
				t.Errorf("got %v; want %v", got, tt.want)
			}
		})
	}

}

func TestParse(t *testing.T) {
	t.Run("192.168.1.0/26 returns correct network ID", func(t *testing.T) {
		got := cidr.Parse("192.168.1.0/26")
		want := "192.168.1.0"

		if got != want {
			t.Errorf("got %v; want %v", got, want)
		}
	})
}
