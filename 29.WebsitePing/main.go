package main

import (
	"fmt"
	"net/http"
	"time"
)

func chechWebsiteStatus(url string) {
	clinet := http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := clinet.Get(url)
	if err != nil {
		fmt.Println("🔴", url, "is OFFLINE:", err)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		fmt.Println("🟢", url, "is ONLINE")
	} else {
		fmt.Println("🟠", url, "might be DOWN (Status:", resp.StatusCode, ")")
	}
	fmt.Println("Status:", resp.Status)
	fmt.Println("StatusCode:", resp.StatusCode)
}

// Check if the provided website is online or not!
func main() {

	var website string
	fmt.Print("Enter website URL: ")
	fmt.Scanln(&website)

	chechWebsiteStatus(website)
}
