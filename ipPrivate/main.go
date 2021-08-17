package main

import (
	"fmt"
	"net"
)

func main() {
	cases := []string{
		"192.0.0.1",
		"192.168.0.1",
		"255.255.255.0",
		"10.0.0.1",
		"172.16.0.1",
	}
	for _, v := range cases {
		ipv4 := net.ParseIP(v)
		fmt.Printf("%s is %t\n", v, ipv4.IsPrivate())
	}
}
