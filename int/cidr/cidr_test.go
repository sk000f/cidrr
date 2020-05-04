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
		"192.168.1.0/32 is a valid IP and subnet",
		"192.168.1.0/32",
		true,
		nil,
	},
	{
		"256.256.256.256/256 is not a valid IP and subnet",
		"256.256.256.256/256",
		false,
		nil,
	},
	{
		"255.255.255.250/3 is a valid IP and subnet",
		"255.255.255.250/3",
		true,
		nil,
	},
}

func TestParse(t *testing.T) {
	for _, tt := range parseTests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := cidr.Parse(tt.input)

			if err != nil {

			}

			if got != tt.want {
				t.Errorf("got %v; want %v", got, tt.want)
			}
		})
	}

}
