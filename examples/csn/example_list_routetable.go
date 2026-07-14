package csnexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/csn"
)

func ListRouteTable() {
	client, err := csn.NewClient("Your AK", "Your SK", "csn.baidubce.com")
	if err != nil {
		fmt.Printf("Failed to new csn client, err: %v.\n", err)
		return
	}
	response, err := client.ListRouteTable("csnId", nil)
	if err != nil {
		fmt.Printf("Failed to list csn route table, err: %v.\n", err)
		return
	}
	fmt.Printf("%+v\n", *response)
}
