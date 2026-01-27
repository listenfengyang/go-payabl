package go_payabl

import (
	"encoding/json"
	"errors"

	"github.com/listenfengyang/go_payabl/utils"
	"github.com/mitchellh/mapstructure"
)

// 充值-成功回调
func (cli *Client) DepositCallback(req PayablDepositCallbackReq, processor func(PayablDepositCallbackReq) error) error {
	//验证签名
	var params map[string]string
	mapstructure.Decode(req, &params)

	flag := utils.VerifyCallback(params, cli.Params.Secret)
	if !flag {
		//签名校验失败
		reqJson, _ := json.Marshal(req)
		cli.logger.Errorf("Payabl deposit back verify fail, req: %s", string(reqJson))
		return errors.New("sign verify error")
	}

	//开始处理
	return processor(req)
}
