package dao_system

type WechatPay struct {
	MchId         string `json:"mchId" p:"mchId" dc:"微信支付商户号"`
	MchKey        string `json:"mchKey" p:"mchKey" dc:"微信支付商户API密钥"`
	SerialNo      string `json:"serialNo" p:"serialNo" dc:"微信支付商户证书序列号"`
	ApiclientCert string `json:"apiclientCert" p:"apiclientCert" dc:"微信支付商户API证书"`
	ApiclientKey  string `json:"apiclientKey" p:"apiclientKey" dc:"商户API证书密钥"`
	PubKey        string `json:"pubKey" p:"pubKey" dc:"微信支付公钥文件"`
	PubId         string `json:"pubId" p:"pubId" dc:"微信支付公钥ID"`
}
