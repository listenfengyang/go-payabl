package go_payabl

import (
	"encoding/json"
	"testing"
)

// Payabl 对 Apple Pay 交易新增 ccn_four_real(真实物理卡后4位),
// 回调需能解析出该字段(ccn_four 是苹果令牌后4位, ccn_four_real 是真实卡后4位)。
func TestCallback_Decodes_CcnFourReal(t *testing.T) {
	raw := `{"3dstatus":"VERIFIED","transactionid":"216093531","type":"capture",
	"errorcode":"0","timestamp":"1769501801","amount":"1.69","currency":"EUR",
	"bin":"424242","ccn_four":"1234","ccn_four_real":"5678",
	"cardholder":"John Doe","card_type":"VISA"}`

	var req PayablDepositCallbackReq
	if err := json.Unmarshal([]byte(raw), &req); err != nil {
		t.Fatalf("unmarshal err: %v", err)
	}

	if req.CcnFour != "1234" {
		t.Errorf("CcnFour = %q, want 1234", req.CcnFour)
	}
	if req.CcnFourReal != "5678" {
		t.Errorf("CcnFourReal = %q, want 5678 (真实卡后4位)", req.CcnFourReal)
	}
}
