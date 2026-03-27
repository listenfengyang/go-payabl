package go_payabl

import (
	"testing"
)

func TestGetSessionIdTest(t *testing.T) {
	vLog := VLog{}
	//构造client
	cli := NewClient(vLog, &PayablInitParams{
		MerchantID:      MERCHANT_ID,
		NotificationURL: NOTIFICATION_URL,
		ReturnURL:       RETURN_URL,
		Secret:          SECRET,
		DepositURL:      DEPOSIT_URL,
		GetSessionIdUrl: GET_SESSION_ID_URL,
	})

	//发请求
	resp, err := cli.GetSessionId(GenGetSessionIdRequestDemo())
	if err != nil {
		cli.logger.Errorf("err433:%s\n", err.Error())
		return
	}
	cli.logger.Infof("resp:%+v\n", resp)
}

func GenGetSessionIdRequestDemo() GetSessionIdReq {
	return GetSessionIdReq{
		Amount:     "11",
		Currency:   "AED",
		ShopUrl:    "http://127.0.0.1",
		Firstname:  "li",
		Lastname:   "ya",
		Email:      "",
		Customerip: "",
	}
}
