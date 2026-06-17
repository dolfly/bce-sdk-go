package etexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/et"
)

// RemoveEtChannelRoutes
func RemoveEtChannelRoutes() {
	client, err := et.NewClient("Your AK", "Your SK", "bcc.bj.baidubce.com")
	if err != nil {
		fmt.Printf("Failed to new et client, err: %v.\n", err)
		return
	}

	args := &et.RemoveEtChannelRoutesArgs{
		ClientToken:  getClientToken(),
		EtId:         "Your EtId",
		EtChannelId:  "Your EtChannelId",
		RouteType:    "static-route",
		Networks:     []string{"192.168.0.0/16"},
		Ipv6Networks: []string{"2400:da00::/48"},
	}

	if err := client.RemoveEtChannelRoutes(args); err != nil {
		fmt.Printf("Failed to remove et channel routes, err: %v.\n", err)
		return
	}
	fmt.Println("Successfully remove et channel routes.")
}
