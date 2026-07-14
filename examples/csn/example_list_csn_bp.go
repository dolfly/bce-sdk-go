package csnexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/csn"
)

func ListCsnBp() {
	client, err := csn.NewClient("Your AK", "Your SK", "csn.baidubce.com")
	if err != nil {
		fmt.Printf("Failed to new csn client, err: %v.\n", err)
		return
	}
	response, err := client.ListCsnBp(nil)
	if err != nil {
		fmt.Printf("Failed to list csn bp, err: %v.\n", err)
		return
	}
	fmt.Printf("%+v\n", *response)
}
