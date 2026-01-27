package go_payabl

import (
	"testing"
)

func TestDeposit(t *testing.T) {
	vLog := VLog{}
	//构造client
	cli := NewClient(vLog, &PayablInitParams{
		MerchantID:      MERCHANT_ID,
		NotificationURL: NOTIFICATION_URL,
		ReturnURL:       RETURN_URL,
		Secret:          SECRET,
		DepositURL:      DEPOSIT_URL,
	})

	//发请求
	resp, err := cli.Deposit(GenDepositRequestDemo())
	if err != nil {
		cli.logger.Errorf("err:%s\n", err.Error())
		return
	}
	cli.logger.Infof("resp:%+v\n", resp)
}

func GenDepositRequestDemo() PayablDepositReq {
	return PayablDepositReq{
		Currency:   "EUR",
		OrderId:    "22515161369",
		Amount:     "188000.00",
		Gender:     "M",
		FirstName:  "John",
		LastName:   "Doe",
		Email:      "john.doe@example.com",
		CustomerIp: "123.456.789.0",
	}
}
