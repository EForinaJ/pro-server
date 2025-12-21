package utils_desensitize

import "unicode/utf8"

func DesensitizeRealName(name string) string {
	// 处理空字符串
	if name == "" {
		return ""
	}

	// 计算UTF-8字符数量（处理中文）
	runeCount := utf8.RuneCountInString(name)

	switch runeCount {
	case 1: // 单字姓名
		return name + "*"
	case 2: // 两字姓名
		runes := []rune(name)
		return string(runes[0]) + "*"
	default: // 三字及以上姓名
		runes := []rune(name)
		// 显示第一个字，中间用*代替，显示最后一个字
		hidden := ""
		for i := 0; i < runeCount-2; i++ {
			hidden += "*"
		}
		return string(runes[0]) + hidden + string(runes[runeCount-1])
	}
}
