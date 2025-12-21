package dto_system

type Email struct {
	Host           string `p:"host" json:"host" v:"required#请输入邮箱服务器" dc:"邮箱服务器"`
	Port           int    `p:"port" json:"port" v:"required#请输入邮箱服务器端口号" dc:"邮箱服务器端口号"`
	User           string `p:"user" json:"user" v:"required#请输入邮箱服务器管理" dc:"邮箱服务器管理"`
	Email          string `p:"email" json:"email" v:"required#请输入邮箱服务器管理邮箱号" dc:"邮箱服务器管理邮箱号"`
	Pass           string `p:"pass" json:"pass" v:"required#请输入邮箱服务器密码" dc:"邮箱服务器密码"`
	MinSendTime    int    `p:"minSendTime" json:"minSendTime" v:"required#请输入最小发送间隔" dc:"最小发送间隔"`
	IpMaxSendCount int    `p:"ipMaxSendCount" json:"ipMaxSendCount" v:"required#请输入IP最大发送次数" dc:"IP最大发送次数"`
	CodeEndTime    int    `p:"codeEndTime" json:"codeEndTime" v:"required#请输入验证码有效期" dc:"验证码有效期"`
}
