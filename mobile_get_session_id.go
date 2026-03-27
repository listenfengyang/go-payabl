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
func (cli *Client) MobileGetSessionId(req MobileGetSessionIdReq) (*MobileGetSessionIdRsp, error) {

	rawURL := cli.Params.MobileGetSessionIdUrl

	var params map[string]string
	mapstructure.Decode(req, &params)

	//补充字段
	params["merchant_id"] = cast.ToString(cli.Params.MerchantID)
	params["amount"] = req.Amount
	params["currency"] = req.Currency
	params["email"] = req.Email
	params["order_id"] = req.OrderId
	params["app_bundle_id"] = req.AppBundleId
	params["notification_url"] = cli.Params.NotificationURL
	signStr, _ := utils.Sign(params, cli.Params.Secret)
	params["signature"] = signStr
	fmt.Println(params)

	var result MobileGetSessionIdRsp

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
	cli.logger.Infof("PSPResty#payabl#mobileGetSessionId->%s", string(restLog))

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

	result.Status = values.Get("status")
	result.EphemeralKey = values.Get("ephemeral_key")
	result.SessionId = values.Get("sessionid")
	result.TransactionId = values.Get("transactionid")
	result.ErrorCode = values.Get("errorcode")
	result.ErrorMessage = values.Get("errormessage")

	// fmt.Printf("%+v\n", result)
	return &result, nil
}
