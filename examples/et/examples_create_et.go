package etexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/et"
)

// CreateEtDcphy
func CreateEtDcphy() {
	client, err := et.NewClient("Your AK", "Your SK", "Your endpoint") // 初始化ak、sk和endpoint
	if err != nil {
		fmt.Printf("Failed to new et client, err: %v.\n", err)
		return
	}

	args := &et.CreateEtDcphyArgs{
		ClientToken: getClientToken(),      // 请求标识
		Name:        "Your Et name",        // 专线名称
		Description: "Your Et description", // 描述
		Isp:         "ISP_CMCC",            // 运营商
		IntfType:    "1G",                  // 物理端口规格
		ApType:      "SINGLE",              // 线路类型
		ApAddr:      "BJYZ",                // 接入点
		UserName:    "Your name",           // 用户名称
		UserPhone:   "Your Phone",          // 用户手机号码
		UserEmail:   "Your Email",          // 用户邮箱
		UserIdc:     "Your Idc",            // 对端地址
		LinkDelay:   100,                   // 端口延迟down时间，单位ms
		Billing: &et.Billing{ // 计费信息，可选
			PaymentTiming: "Prepaid",
			Reservation: &et.Reservation{
				ReservationLength:   1,
				ReservationTimeUnit: "month",
			},
		},
		AutoRenew: &et.Reservation{ // 自动续费信息，可选
			ReservationLength:   1,
			ReservationTimeUnit: "month",
		},
		Tags: []et.Tag{{
			TagKey:   "Your TagKey",
			TagValue: "Your TagValue",
		}}, // 标签
	}

	if _, err = client.CreateEtDcphy(args); err != nil {
		fmt.Printf("Failed to create a new et, err: %v.\n", err)
		return
	}
	fmt.Println("Successfully create a new et.")
}
