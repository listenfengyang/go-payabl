package utils

import (
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"sort"
	"strings"

	"github.com/samber/lo"
	"github.com/spf13/cast"
)

func Sign(params map[string]string, accessKey string) (string, error) {

	// 1. 按key值字母顺序排序，k=v&k1=v2&...
	keys := lo.Keys(params)
	sort.Strings(keys)

	var sb strings.Builder
	for _, k := range keys {
		value := cast.ToString(params[k])
		if k != "signature" && value != "" {
			//只有非空才可以参与签名
			sb.WriteString(fmt.Sprintf("%s", value))
		}
	}

	// 2. 密钥加到字符末尾
	sb.WriteString(fmt.Sprintf("%s", accessKey))
	signStr := sb.String()

	fmt.Printf("[rawString]%s\n", signStr)

	// 3. 计算字符串的 SHA-1 十六进制值，哈希值必须是小写字母
	hash := sha1.Sum([]byte(signStr))
	signResult := hex.EncodeToString(hash[:])

	fmt.Printf("[rawUpString]%s\n", strings.ToLower(signResult))

	return strings.ToLower(signResult), nil
}

// 入金&出金回调-成功-验签
func VerifyCallback(params map[string]string, secret string) bool {
	signature := params["security"]

	// 1. 定义参与签名的key
	signKeyList := []string{"transactionid", "type", "errorcode", "timestamp"}

	// 2. Build sign string
	var sb strings.Builder
	for _, k := range signKeyList {
		v := cast.ToString(params[k])
		sb.WriteString(fmt.Sprintf("%s", v))
	}
	sb.WriteString(secret)
	signStr := sb.String()

	// fmt.Printf("[rawString]%s\n", signStr)
	hash := sha256.Sum256([]byte(signStr))
	// fmt.Printf("signature: %s\n", signature)
	fmt.Printf("sha256 sign: %s\n", strings.ToLower(hex.EncodeToString(hash[:])))

	return signature == strings.ToLower(hex.EncodeToString(hash[:]))
}

// 出金
func SignWithdraw(params map[string]string, accessKey string) (string, error) {

	// 1. 按key值字母顺序排序，k=v&k1=v2&...
	keys := lo.Keys(params)
	sort.Strings(keys)

	var sb strings.Builder
	for _, k := range keys {
		value := cast.ToString(params[k])
		if k != "signature" && value != "" {
			//只有非空才可以参与签名
			sb.WriteString(fmt.Sprintf("%s", value))
		}
	}

	// 2. 密钥加到字符末尾
	sb.WriteString(fmt.Sprintf("%s", accessKey))
	signStr := sb.String()

	fmt.Printf("[rawString]%s\n", signStr)

	// 3. 计算字符串的 SHA-1 十六进制值，哈希值必须是小写字母
	hash := sha1.Sum([]byte(signStr))
	signResult := hex.EncodeToString(hash[:])

	fmt.Printf("[rawUpString]%s\n", strings.ToLower(signResult))

	return strings.ToLower(signResult), nil
}

func VerifySignWithdraw(params map[string]string, accessKey string) (bool, error) {
	// Check if signature exists in params
	signature, exists := params["sign"]
	if !exists {
		return false, nil
	}

	// Remove signature from params for verification
	delete(params, "sign")

	// Generate current signature
	currentSignature, err := SignWithdraw(params, accessKey)
	if err != nil {
		return false, fmt.Errorf("signature generation failed: %w", err)
	}

	// Compare signatures
	return signature == currentSignature, nil
}
