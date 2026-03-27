package go_payabl

import (
	"testing"
)

func TestMobileGetSessionIdTest(t *testing.T) {
	vLog := VLog{}
	//构造client
	cli := NewClient(vLog, &PayablInitParams{
		MerchantID:            MERCHANT_ID,
		NotificationURL:       NOTIFICATION_URL,
		ReturnURL:             RETURN_URL,
		Secret:                SECRET,
		DepositURL:            DEPOSIT_URL,
		MobileGetSessionIdUrl: MOBILE_GET_SESSION_ID_URL,
	})

	//发请求
	resp, err := cli.MobileGetSessionId(GenMobileGetSessionIdRequestDemo())
	if err != nil {
		cli.logger.Errorf("err433:%s\n", err.Error())
		return
	}
	cli.logger.Infof("resp:%+v\n", resp)
}

func GenMobileGetSessionIdRequestDemo() MobileGetSessionIdReq {
	return MobileGetSessionIdReq{
		Amount:      "45",
		Currency:    "EUR",
		OrderId:     "2025664646790",
		AppBundleId: "338as7df61l32k0a9ufdag9659as8dff", // 安卓：com.logtec.cpt  ios: com.cptmarkets.x
		Email:       "ji4253@gmail.com",
	}
}
