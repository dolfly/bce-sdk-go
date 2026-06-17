package etexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/et"
)

// RemoveEtChannelUsers
func RemoveEtChannelUsers() {
	client, err := et.NewClient("Your AK", "Your SK", "bcc.bj.baidubce.com")
	if err != nil {
		fmt.Printf("Failed to new et client, err: %v.\n", err)
		return
	}

	args := &et.RemoveEtChannelUsersArgs{
		ClientToken:     getClientToken(),
		EtId:            "Your EtId",
		EtChannelId:     "Your EtChannelId",
		AuthorizedUsers: []string{"Your UserId"},
	}

	if err := client.RemoveEtChannelUsers(args); err != nil {
		fmt.Printf("Failed to remove et channel users, err: %v.\n", err)
		return
	}
	fmt.Println("Successfully remove et channel users.")
}
