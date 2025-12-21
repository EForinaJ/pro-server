package utils_desensitize

func DesensitizeIDCard(idCard string) string {
	length := len(idCard)

	// 根据身份证长度决定脱敏策略
	switch {
	case length == 18: // 18位身份证
		return idCard[:6] + "********" + idCard[14:]
	case length == 15: // 15位身份证
		return idCard[:6] + "******" + idCard[12:]
	default: // 其他长度，保守处理
		if length <= 8 {
			return idCard
		}
		// 显示前4位和后4位，中间用*代替
		showLen := 4
		if length < 12 {
			showLen = 2
		}
		return idCard[:showLen] + "****" + idCard[length-showLen:]
	}
}
