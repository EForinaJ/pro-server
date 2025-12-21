package dto_system

type WechatPay struct {
	MchId         string `v:"required#请输入微信支付商户号" json:"mchId" p:"mchId" dc:"微信支付商户号"`
	MchKey        string `v:"required#请输入微信支付商户API密钥" json:"mchKey" p:"mchKey" dc:"微信支付商户API密钥"`
	SerialNo      string `v:"required#请输入微信支付商户证书序列号" json:"serialNo" p:"serialNo" dc:"微信支付商户证书序列号"`
	ApiclientCert string `v:"required#请输入微信支付商户API证书" json:"apiclientCert" p:"apiclientCert" dc:"微信支付商户API证书"`
	ApiclientKey  string `v:"required#请输入微信支付商户API证书密钥" json:"apiclientKey" p:"apiclientKey" dc:"商户API证书密钥"`
	PubKey        string `v:"required#请输入微信支付公钥文件" json:"pubKey" p:"pubKey" dc:"微信支付公钥文件"`
	PubId         string `v:"required#请输入微信支付公钥ID" json:"pubId" p:"pubId" dc:"微信支付公钥ID"`
}
