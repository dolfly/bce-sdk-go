package vpcexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/vpc"
)

// 以下为示例代码，实际开发中请根据需要进行修改和补充

func UpdateNatReleaseProtectionSwitch() {
	ak, sk, endpoint := "Your AK", "Your SK", "bcc.bj.baidubce.com"

	natClient, _ := vpc.NewClient(ak, sk, endpoint) // 初始化client

	NatID := "Your nat's id"

	args := &vpc.UpdateNatReleaseProtectionSwitchArgs{
		// true: 开启释放保护，false: 关闭释放保护
		DeleteProtect: true,
	}

	if err := natClient.UpdateNatReleaseProtectionSwitch(NatID, args); err != nil {
		fmt.Println("update nat release protection switch error: ", err)
		return
	}

	fmt.Printf("update nat release protection switch for %s success.\n", NatID)
}
