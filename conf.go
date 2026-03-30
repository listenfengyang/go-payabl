package go_payabl

const (
	//"gateway_test"
	MERCHANT_ID = "gateway_test_3d"
	// "95e24f0e29c1e074abcf049c28ae1cde3d0adc67"
	// MERCHANT_ID = "gateway_test_3d"
	// MERCHANT_ID = "gateway_test_cb"
	// MERCHANT_ID = "95e24f0e29c1e074abcf049c28ae1cde3d0adc67"
	// "DJS3zI1Rh2" //
	SECRET = "b185"
	// DEPOSIT_URL      = "https://pay4.sandbox.payabl.com/pay/payment/init"
	DEPOSIT_URL      = "https://sandbox.payabl.com/pay/payment/init"
	WITHDRAW_URL     = "https://sandbox.payabl.com/pay/backoffice/payment_cft"
	NOTIFICATION_URL = "https://api-test.logtec.dev/fapi/payment/psp/public/payabl/deposit"
	//"https://api.cptmarkets.com/fapi/cpti/payment/psp/public/payabl/deposit/back"
	//"https://api-test.logtec.dev/fapi/payment/psp/public/payabl/deposit"
	RETURN_URL = "" //"https://portal.cptmarkets.com"

	// 获取sessionid
	GET_SESSION_ID_URL = "https://pay4.sandbox.payabl.com/pay/payment/get_payment_widget_session"
	// 移动端获取sessionid
	// https://pay4.payabl.com/pay/mobile/init
	MOBILE_GET_SESSION_ID_URL = "https://pay4.sandbox.payabl.com/pay/mobile/init"
)
