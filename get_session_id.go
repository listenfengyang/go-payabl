package go_payabl

import (
	"crypto/tls"
	"fmt"
	"net/url"

	jsoniter "github.com/json-iterator/go"
	"github.com/listenfengyang/go-payabl/utils"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/cast"
)

// sdk 获取sessionid
func (cli *Client) GetSessionId(req GetSessionIdReq) (*GetSessionIdRsp, error) {

	rawURL := cli.Params.GetSessionIdUrl

	var params map[string]string
	mapstructure.Decode(req, &params)

	//补充字段
	params["merchantid"] = cast.ToString(cli.Params.MerchantID)
	params["amount"] = req.Amount
	params["currency"] = req.Currency
	params["shop_url"] = cli.Params.ReturnURL
	params["firstname"] = req.Firstname
	params["lastname"] = req.Lastname
	params["email"] = req.Email
	params["customerip"] = req.Customerip
	signStr, _ := utils.Sign(params, cli.Params.Secret)
	params["signature"] = signStr
	fmt.Println(params)

	var result GetSessionIdRsp

	resp2, err := cli.ryClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
		SetCloseConnection(true).
		R().
		SetFormData(params).
		SetBody(params).
		SetHeaders(getHeaders()).
		SetDebug(cli.debugMode).
		SetResult(&result).
		SetError(&result).
		Post(rawURL)

	restLog, _ := jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(utils.GetRestyLog(resp2))
	cli.logger.Infof("PSPResty#payabl#getSessionId->%s", string(restLog))

	if err != nil {
		return nil, err
	}

	if resp2.StatusCode() != 200 {
		//反序列化错误会在此捕捉
		return nil, fmt.Errorf("status code: %d", resp2.StatusCode())
	}

	values, err := url.ParseQuery(resp2.String())
	if err != nil {
		return nil, fmt.Errorf("resp parse error: %w", err)
	}

	result.ErrorCode = values.Get("errorcode")
	result.ErrorMessage = values.Get("errormessage")
	result.OrderId = values.Get("orderid")
	result.SessionId = values.Get("session_id")
	result.Signature = values.Get("signature")
	result.TransactionId = values.Get("transactionid")

	// fmt.Printf("%+v\n", result)
	return &result, nil
}
