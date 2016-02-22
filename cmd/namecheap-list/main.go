package main

import (
	"fmt"
	namecheap "github.com/billputer/go-namecheap"
	"os"
)

func main() {

	if os.Getenv("NAMECHEAPUSER") == "" {
		fmt.Println("Set environmental variables before running this program.")
		os.Exit(1)
	}
	if os.Getenv("NAMECHEAPKEY") == "" {
		fmt.Println("Set environmental variables before running this program.")
		os.Exit(1)
	}

	apiUser := os.Getenv("NAMECHEAPUSER")
	apiToken := os.Getenv("NAMECHEAPKEY")
	userName := os.Getenv("NAMECHEAPUSER")

	client := namecheap.NewClient(apiUser, apiToken, userName)

	domains, _ := client.DomainsGetList()
	for _, domain := range domains {
		fmt.Printf("%+v\n", domain.Name)
	}
}
