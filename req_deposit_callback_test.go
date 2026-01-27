package go_payabl

import (
	"encoding/json"
	"fmt"
	"testing"
)

type VLog struct {
}

func (l VLog) Debugf(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}
func (l VLog) Infof(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}
func (l VLog) Warnf(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}
func (l VLog) Errorf(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}

func TestCallback(t *testing.T) {
	vLog := VLog{}
	//构造client
	cli := NewClient(vLog, &PayablInitParams{
		MerchantID:      MERCHANT_ID,
		NotificationURL: NOTIFICATION_URL,
		ReturnURL:       RETURN_URL,
		Secret:          SECRET,
		DepositURL:      DEPOSIT_URL,
	})

	//1. 获取请求
	req := GenCallbackRequestDemo() //提现的返回
	var backReq PayablDepositCallbackReq
	err := json.Unmarshal([]byte(req), &backReq)
	if err != nil {
		cli.logger.Errorf("Error:%s", err.Error())
		return
	}

	//2. 处理请求
	err = cli.DepositCallback(backReq, func(PayablDepositCallbackReq) error { return nil })
	if err != nil {
		cli.logger.Errorf("Error:%s", err.Error())
		return
	}
	cli.logger.Infof("resp:%+v\n", backReq)
}

// 3dstatus=VERIFIED&timestamp=1722343778&orderid=ORDER-00005&errorcode=0&type=capture&related_token_ids=&
// errormessage=&3dauthentication_flow=challenge_flow&3dProtocolVersion=2.2.0&transactionid=105937169&
// 3dauthentication_status=Y&security=4f77617275d09fe3f13a74d52f6900b1bd124eb52fb0b9fa99c78a0969179057
func GenCallbackRequestDemo() string {
	return `{"3dstatus":"VERIFIED","timestamp":"1722343778",
	"orderid":"ORDER-00005","errorcode":"0","type":"capture","related_token_ids":"",
	"errormessage":"","3dauthentication_flow":"challenge_flow","3dProtocolVersion":"2.2.0",
	"transactionid":"105937169","3dauthentication_status":"Y","security":"4f77617275d09fe3f13a74d52f6900b1bd124eb52fb0b9fa99c78a0969179057"}`

}
