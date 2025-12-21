package dao_system

type Email struct {
	Host           string `json:"host" dc:"邮箱服务器"`
	Port           int    `json:"port" dc:"邮箱服务器端口"`
	User           string `json:"user" dc:"邮箱服务器管理"`
	Email          string `json:"email" dc:"邮箱服务器管理邮箱号"`
	Pass           string `json:"pass" dc:"邮箱服务器密码"`
	MinSendTime    int    `json:"minSendTime" dc:"最小发送间隔时间"`
	IPMaxSendCount int    `json:"ipMaxSendCount" dc:"IP最大发送次数"`
	CodeEndTime    int    `json:"codeEndTime" dc:"验证码过期时间"`
}
