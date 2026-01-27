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

// payment_method=1&type=capture&errormessage=&currency=EUR&amount=1.69&errorcode=0&security=5508c9364d229902488131089c67080a92f4b578b95fa9e7ca547e7a5cda3790&orderid=202601271015460906&related_token_ids=216093531&transactionid=216093531&timestamp=1769501801
func GenCallbackRequestDemo() string {
	return `{"payment_method":"1","type":"capture",
	"errormessage":"","currency":"EUR","amount":"1.69","errorcode":"0",
	"security":"5508c9364d229902488131089c67080a92f4b578b95fa9e7ca547e7a5cda3790",
	"orderid":"202601271015460906","related_token_ids":"216093531",
	"transactionid":"216093531","timestamp":"1769501801"}`

}
