// namecheap-list prints all your domain names, one per line.
package main

import (
	"fmt"
	"os"

	namecheap "github.com/billputer/go-namecheap"
)

func main() {

	if os.Getenv("NAMECHEAPUSER") == "" || os.Getenv("NAMECHEAPKEY") == "" {
		fmt.Println("Set NAMECHEAPUSER and NAMECHEAPKEY environmental variables before running this program.")
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
