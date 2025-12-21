package utils_amount

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/gogf/gf/v2/errors/gerror"
)

type AmountValidator struct{}

// ValidateAmount 验证金额格式
func (v *AmountValidator) ValidateAmount(amountStr string) error {
	if amountStr == "" {
		return gerror.New("金额不能为空")
	}

	// 1. 验证是否为有效数字
	_, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		return gerror.New("金额格式不正确")
	}

	// 2. 验证是否为正数
	amount, _ := strconv.ParseFloat(amountStr, 64)
	if amount <= 0 {
		return gerror.New("金额必须大于0")
	}

	// 3. 验证小数位数（最多2位）
	if !v.isValidDecimal(amountStr) {
		return gerror.New("金额最多保留两位小数")
	}

	// 4. 验证金额范围（可选）
	if amount > 1000000 { // 最大100万
		return gerror.New("金额不能超过1000000")
	}

	return nil
}

// isValidDecimal 验证小数位数
func (v *AmountValidator) isValidDecimal(amountStr string) bool {
	// 正则表达式：验证最多两位小数
	pattern := `^\d+(\.\d{1,2})?$`
	matched, _ := regexp.MatchString(pattern, amountStr)

	if !matched {
		return false
	}

	// 额外检查：确保小数部分不超过2位
	parts := strings.Split(amountStr, ".")
	if len(parts) == 2 {
		return len(parts[1]) <= 2
	}

	return true
}

// ValidateAndFormat 验证并格式化金额
func (v *AmountValidator) ValidateAndFormat(amountStr string) (float64, error) {
	if err := v.ValidateAmount(amountStr); err != nil {
		return 0, err
	}

	amount, _ := strconv.ParseFloat(amountStr, 64)
	// 格式化保留两位小数
	formatted := formatToTwoDecimal(amount)
	return formatted, nil
}

// formatToTwoDecimal 格式化为两位小数
func formatToTwoDecimal(amount float64) float64 {
	formatted, _ := strconv.ParseFloat(strconv.FormatFloat(amount, 'f', 2, 64), 64)
	return formatted
}
