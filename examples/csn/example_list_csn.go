package csnexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/csn"
)

func ListCsn() {
	client, err := csn.NewClient("Your AK", "Your SK", "csn.baidubce.com")
	if err != nil {
		fmt.Printf("Failed to new csn client, err: %v.\n", err)
		return
	}
	response, err := client.ListCsn(nil)
	if err != nil {
		fmt.Printf("Failed to list csn, err: %v.\n", err)
		return
	}
	fmt.Printf("%+v\n", *response)
	for _, item := range response.Csns {
		fmt.Printf("csnId: %s, createTime: %s\n", item.CsnId, item.CreateTime)
	}
}
