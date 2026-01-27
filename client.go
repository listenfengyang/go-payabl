package go_payabl

import (
	"github.com/go-resty/resty/v2"
	"github.com/listenfengyang/go_payabl/utils"
)

type Client struct {
	Params *PayablInitParams

	ryClient  *resty.Client
	debugMode bool //是否调试模式
	logger    utils.Logger
}

func NewClient(logger utils.Logger, params *PayablInitParams) *Client {
	return &Client{
		Params: params,

		ryClient:  resty.New(), //client实例
		debugMode: false,
		logger:    logger,
	}
}

func (cli *Client) SetDebugModel(debugModel bool) {
	cli.debugMode = debugModel
}

func (cli *Client) SetMerchantInfo(merchant MerchantInfo) {
	cli.Params.MerchantID = merchant.MerchantID
}
