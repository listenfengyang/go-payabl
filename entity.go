package go_payabl

type PayablInitParams struct {
	MerchantID       string `json:"merchantID" mapstructure:"merchantID" config:"merchantID"  yaml:"merchantID"`                         // merchantID
	NotificationURL  string `json:"notificationURL" mapstructure:"notificationURL" config:"notificationURL"  yaml:"notificationURL"`     // 通知URL
	ReturnURL        string `json:"returnURL" mapstructure:"returnURL" config:"returnURL"  yaml:"returnURL"`                             // 重定向URL
	DepositURL       string `json:"depositURL" mapstructure:"depositURL" config:"depositURL"  yaml:"depositURL"`                         // 入金URL
	WithdrawURL      string `json:"withdrawURL" mapstructure:"withdrawURL" config:"withdrawURL"  yaml:"withdrawURL"`                     // 出金URL
	WithdrawNotifURL string `json:"withdrawNotifURL" mapstructure:"withdrawNotifURL" config:"withdrawNotifURL"  yaml:"withdrawNotifURL"` // 出金回调URL
	Secret           string `json:"secret" mapstructure:"secret" config:"secret"  yaml:"secret"`                                         // 密钥

	// MerchantInfo `yaml:",inline" mapstructure:",squash"`

	// PayinAuthUrl        string `json:"payinAuthUrl" mapstructure:"payinAuthUrl" config:"payinAuthUrl"  yaml:"payinAuthUrl"`
	// PayinPreAuthUrl     string `json:"payinPreAuthUrl" mapstructure:"payinPreAuthUrl" config:"payinPreAuthUrl"  yaml:"payinPreAuthUrl"`
	// WithdrawUrl         string `json:"withdrawUrl" mapstructure:"withdrawUrl" config:"withdrawUrl"  yaml:"withdrawUrl"`
	// WithdrawResponseUrl string `json:"withdrawResponseUrl" mapstructure:"withdrawResponseUrl" config:"withdrawResponseUrl"  yaml:"withdrawResponseUrl"`
}

type MerchantInfo struct {
	MerchantID      string `json:"merchantID" mapstructure:"merchantID" config:"merchantID"  yaml:"merchantID"`                     // merchantID
	NotificationURL string `json:"notificationURL" mapstructure:"notificationURL" config:"notificationURL"  yaml:"notificationURL"` // 通知URL
	ReturnURL       string `json:"returnURL" mapstructure:"returnURL" config:"returnURL"  yaml:"returnURL"`                         // 重定向URL
	DepositURL      string `json:"depositURL" mapstructure:"depositURL" config:"depositURL"  yaml:"depositURL"`                     // 入金URL
	Secret          string `json:"secret" mapstructure:"secret" config:"secret"  yaml:"secret"`                                     // 密钥
}

//============================================================

// payabl入金
type PayablDepositReq struct {
	OrderId    string `json:"orderid" form:"orderid" mapstructure:"orderid"`          // 订单ID 为了商家方便
	Amount     string `json:"amount" form:"amount" mapstructure:"amount"`             // 金额
	Currency   string `json:"currency" form:"currency" mapstructure:"currency"`       // 币种
	Gender     string `json:"gender" form:"gender" mapstructure:"gender"`             // 客户性别
	FirstName  string `json:"firstname" form:"firstname" mapstructure:"firstname"`    // 客户姓
	LastName   string `json:"lastname" form:"lastname" mapstructure:"lastname"`       // 客户名
	Email      string `json:"email" form:"email" mapstructure:"email"`                // 客户邮箱
	CustomerIp string `json:"customerip" form:"customerip" mapstructure:"customerip"` // 客户IP

	// 以下SDK添加字段
	// MerchantId string `json:"merchantid" form:"merchantid" mapstructure:"merchantid"` // 商户ID
	// Signature string `json:"signature" form:"signature" mapstructure:"signature"` // 签名
	// ReturnUrl  string        `json:"url_return" form:"url_return" mapstructure:"url_return"` // 重定向URL
	// Language   string `json:"language" form:"language" mapstructure:"language"`     // 语言 zh-CN
	// En         string `json:"en" form:"en" mapstructure:"en"`                 // 语言 en
	// NotificationUrl string        `json:"notification_url" form:"notification_url" mapstructure:"notification_url"` // 通知URL
}

type PayablDepositRsp struct {
	ErrorCode     string `json:"errorcode" form:"errorcode" mapstructure:"errorcode"`             //错误码
	ErrorMessage  string `json:"errormessage" form:"errormessage" mapstructure:"errormessage"`    //错误信息
	SessionId     string `json:"sessionid" form:"sessionid" mapstructure:"sessionid"`             //会话ID
	TransactionId string `json:"transactionid" form:"transactionid" mapstructure:"transactionid"` //交易ID
	UserId        string `json:"user_id" form:"user_id" mapstructure:"user_id"`                   //用户ID
	StartUrl      string `json:"start_url" form:"start_url" mapstructure:"start_url"`             //跳转三方收银URL
}

// 入金回调
type PayablDepositCallbackReq struct {
	Status               string `json:"3dstatus" form:"3dstatus" mapstructure:"3dstatus"`                                              // 请求状态：VERIFIED=成功 400=请求失败
	Timestamp            string `json:"timestamp" form:"timestamp" mapstructure:"timestamp"`                                           // 时间戳
	OrderId              string `json:"orderid" form:"orderid" mapstructure:"orderid"`                                                 // 订单ID
	ErrorCode            string `json:"errorcode" form:"errorcode" mapstructure:"errorcode"`                                           // 错误码
	Type                 string `json:"type" form:"type" mapstructure:"type"`                                                          // 类型
	RelatedTokenIds      string `json:"related_token_ids" form:"related_token_ids" mapstructure:"related_token_ids"`                   // 相关tokenID
	ErrorMessage         string `json:"errormessage" form:"errormessage" mapstructure:"errormessage"`                                  // 错误信息
	AuthenticationFlow   string `json:"3dauthentication_flow" form:"3dauthentication_flow" mapstructure:"3dauthentication_flow"`       // 3D认证流程
	ProtocolVersion      string `json:"3dProtocolVersion" form:"3dProtocolVersion" mapstructure:"3dProtocolVersion"`                   // 3D协议版本
	TransactionId        string `json:"transactionid" form:"transactionid" mapstructure:"transactionid"`                               // 交易ID
	AuthenticationStatus string `json:"3dauthentication_status" form:"3dauthentication_status" mapstructure:"3dauthentication_status"` // 3D认证状态
	Security             string `json:"security" form:"security" mapstructure:"security"`                                              // 安全校验
}

// payabl出金
type PayablWithdrawReq struct {
	MerchantId     string `json:"merchantid" form:"merchantid" mapstructure:"merchantid"`                // 商户ID
	Amount         string `json:"amount" form:"amount" mapstructure:"amount"`                            // 金额
	Currency       string `json:"currency" form:"currency" mapstructure:"currency"`                      // 币种
	PaymentMethod  string `json:"payment_method" form:"payment_method" mapstructure:"payment_method"`    // 出金方式
	Signature      string `json:"signature" form:"signature" mapstructure:"signature"`                   // 签名
	Ccn            string `json:"ccn" form:"ccn" mapstructure:"ccn"`                                     // 出金账号
	ExpMonth       string `json:"exp_month" form:"exp_month" mapstructure:"exp_month"`                   // 过期月
	ExpYear        string `json:"exp_year" form:"exp_year" mapstructure:"exp_year"`                      // 过期年
	CardholderName string `json:"cardholder_name" form:"cardholder_name" mapstructure:"cardholder_name"` // 持卡人姓名
	TransactionId  string `json:"transactionid" form:"transactionid" mapstructure:"transactionid"`       // 出金订单号
}

type PayablWithdrawRsp struct {
	TransactionId string `json:"transaction_id" mapstructure:"transaction_id"` // 出金订单号
	TransId       string `json:"transid" mapstructure:"transid"`
	Status        string `json:"status" mapstructure:"status"`             //0=成功，-6001=失败...
	ErrorMessage  string `json:"errormessage" mapstructure:"errormessage"` //错误信息
	ErrMsg        string `json:"errmsg" mapstructure:"errmsg"`             //错误信息
	Amount        string `json:"amount" mapstructure:"amount"`             // 金额
	Price         string `json:"price" mapstructure:"price"`               // 金额
	Currency      string `json:"currency" mapstructure:"currency"`         // 币种
	OrderId       string `json:"orderid" mapstructure:"orderid"`           // 订单ID
}

// 出金回调
type PayablWithdrawCallbackReq struct {
	Status             string `json:"status" form:"status" mapstructure:"status"` //200=成功，400=失败...
	Message            string `json:"message" form:"message" mapstructure:"message"`
	StatusCode         string `json:"status_code" form:"status_code" mapstructure:"status_code"`                            //出金状态：20001=出金成功 20002=出金失败
	Amount             string `json:"amount" form:"amount" mapstructure:"amount"`                                           //出金金额
	AfterChargesAmount string `json:"after_charges_amount" form:"after_charges_amount" mapstructure:"after_charges_amount"` //出金金额（包含手续费）
	TransactionId      string `json:"transaction_id" form:"transaction_id" mapstructure:"transaction_id"`                   //出金订单号
	Currency           string `json:"currency" form:"currency" mapstructure:"currency"`                                     //币种
	ReferenceCode      string `json:"reference_code" form:"reference_code" mapstructure:"reference_code"`                   //引用订单号
	CreatedAt          string `json:"created_at" form:"created_at" mapstructure:"created_at"`                               //创建时间
	UpdatedAt          string `json:"updated_at" form:"updated_at" mapstructure:"updated_at"`                               //更新时间
	Timestamp          string `json:"timestamp" form:"timestamp" mapstructure:"timestamp"`                                  //时间戳
	Signature          string `json:"signature" form:"signature" mapstructure:"signature"`                                  //签名
}

type ZPayWithdrawCallbackRsp struct {
	Status  int32  `json:"status" mapstructure:"status"` //请求状态：200=请求成功 400=请求失败
	Message string `json:"message" mapstructure:"message"`
}
