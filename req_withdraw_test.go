package go_payabl

import "testing"

func TestWithdraw(t *testing.T) {
	vLog := VLog{}
	//构造client
	cli := NewClient(vLog, &PayablInitParams{
		MerchantID:      MERCHANT_ID,
		NotificationURL: NOTIFICATION_URL,
		ReturnURL:       RETURN_URL,
		Secret:          SECRET,
		DepositURL:      DEPOSIT_URL,
		WithdrawURL:     WITHDRAW_URL,
	})

	//发请求
	resp, err := cli.WithdrawReq(GenWithdrawRequestDemo())
	if err != nil {
		cli.logger.Errorf("err:%s\n", err.Error())
		return
	}
	cli.logger.Infof("resp:%+v\n", resp)
}

func GenWithdrawRequestDemo() PayablWithdrawReq {
	return PayablWithdrawReq{
		Currency:       "EUR",
		PaymentMethod:  "1", // 1=信用卡
		Amount:         "1.42",
		Ccn:            "5546989999990033",
		ExpMonth:       "07",
		ExpYear:        "2020",
		CardholderName: "Max Mustermann",
		// TransactionId:  "20230824152000001",
	}
}
