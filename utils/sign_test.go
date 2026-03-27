package utils

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"testing"
)

// 计算字符串的 SHA-1 十六进制值，哈希值必须是小写字母
func TestSign(t *testing.T) {

	plain := `2020.00AED10.17.24.157ahmad7mohammadj@gmail.comAhmad Mohammad Abdulla MohammadenAljaasmi95e24f0e29c1e074abcf049c28ae1cde3d0adc67https://api.cptmarkets.com/fapi/cpti/payment/psp/public/payabl/deposit/back202603242034560382`
	secret := "DJS3zI1Rh2"

	final := plain + secret

	sum := sha1.Sum([]byte(final))
	fmt.Println(hex.EncodeToString(sum[:])) // 小写hex

	// signStr := "2020.00AED10.17.24.157ahmad7mohammadj@gmail.comAhmad Mohammad Abdulla MohammadenAljaasmi95e24f0e29c1e074abcf049c28ae1cde3d0adc67https://api.cptmarkets.com/fapi/cpti/payment/psp/public/payabl/deposit/back202603242034560382DJS3zI1Rh2"

	// fmt.Println(signStr)

	// hash := sha1.Sum([]byte(signStr))
	// signResult := hex.EncodeToString(hash[:])

	// fmt.Printf("[rawUpString]%s\n", strings.ToLower(signResult))

}
