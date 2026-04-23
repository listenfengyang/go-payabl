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
		ApplePayMerchantID:    APPLEPAY_MERCHANT_ID,
		ApplePaySecret:        APPLEPAY_SECRET,
	})

	//发请求
	payType := "apple_pay"
	resp, err := cli.MobileGetSessionId(payType, GenMobileGetSessionIdRequestDemo())
	if err != nil {
		cli.logger.Errorf("err433:%s\n", err.Error())
		return
	}
	cli.logger.Infof("resp:%+v\n", resp)
}

func GenMobileGetSessionIdRequestDemo() MobileGetSessionIdReq {
	return MobileGetSessionIdReq{
		Amount:   "42",
		Currency: "EUR",
		OrderId:  "20263753253266",
		// Merchantid:  "merchant_user_test",
		// Signature:   "1a29075414d8061aa1e9ef6eb4f20a69dc0f2f36",
		AppBundleId: "com.logtec.cpt", // 安卓：com.logtec.cpt  ios: com.cptmarkets.x
		Email:       "ajsf@gmail.com",
	}
}
