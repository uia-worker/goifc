// To avoid recursion in cases such as
//
// type X string
// func (x X) String() string { return Sprintf("<%s>", x) }
// convert the value before recurring:

// func (x X) String() string { return Sprintf("<%s>", string(x)) }

package main

import "fmt"

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
func (ip IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 88},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
	ip := IPAddr{127, 0, 0, 1}
	fmt.Println(ip)
}
