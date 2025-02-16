package main

import (
	"fmt"
	"net"
)

func getLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	fmt.Println("Addresses:", addrs)
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}

	for _, addr := range addrs {
		ipNet, ok := addr.(*net.IPNet)
		if ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
			return ipNet.IP.String()
		}
	}
	return "No Local IP found"
}

func main() {
	fmt.Println("Local IP Address:", getLocalIP())
}
