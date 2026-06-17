package etexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/et"
)

// AddEtChannelUsers
func AddEtChannelUsers() {
	client, err := et.NewClient("Your AK", "Your SK", "bcc.bj.baidubce.com")
	if err != nil {
		fmt.Printf("Failed to new et client, err: %v.\n", err)
		return
	}

	args := &et.AddEtChannelUsersArgs{
		ClientToken:     getClientToken(),
		EtId:            "Your EtId",
		EtChannelId:     "Your EtChannelId",
		AuthorizedUsers: []string{"Your UserId"},
	}

	if err := client.AddEtChannelUsers(args); err != nil {
		fmt.Printf("Failed to add et channel users, err: %v.\n", err)
		return
	}
	fmt.Println("Successfully add et channel users.")
}
