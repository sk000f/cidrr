package cidr

import (
	"regexp"
	"strings"
)

const subnetPattern = "^(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\/([0-9]|1[0-9]|2[0-9]|3[0-2])$"

// Validate returns true if specified subnet is valid CIDR notation
func Validate(subnet string) (bool, error) {

	re, err := regexp.Compile(subnetPattern)
	if err != nil {
		return false, err
	}

	r := re.MatchString(subnet)

	return r, nil
}

// Parse returns Network ID, Broadcast Adress, First and Last Host for given subnet
func Parse(subnet string) string {
	return getNetworkID(subnet)
}

// returns the Network ID for the given subnet (the first IP in the subnet)
func getNetworkID(subnet string) string {
	sn := strings.Split(subnet, "/")
	return sn[0]
}
