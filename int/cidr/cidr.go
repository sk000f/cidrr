package cidr

import (
	"regexp"
)

const subnetPattern = "^(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\/\\d+$"

// Parse returns true if specified subnet is valid CIDR notation
func Parse(subnet string) (bool, error) {

	re, err := regexp.Compile(subnetPattern)
	if err != nil {
		return false, err
	}

	r := re.MatchString(subnet)

	return r, nil
}
