package alipay

import "errors"

// SendTemplateMessage 小程序通过openapi给用户触达消息 https://docs.alipay.com/mini/api/templatemessage
func (this *Client) SendTemplateMessage(param MiniTemplateMessage) (result *MiniTemplateMessageRsp, err error) {
	err = this.doRequest("POST", param, &result)
	return result, err
}

// CreateAppQrcode 小程序生成推广二维码接口 https://docs.alipay.com/mini/api/openapi-qrcode
func (this *Client) CreateAppQrcode(param MiniOpenApiQRcode) (result *MiniOpenApiQRcodeRsp, err error) {
	err = this.doRequest("POST", param, &result)
	return result, err
}

// DetectRiskContent 小程序内容风险检测服务 https://docs.open.alipay.com/api_49/alipay.security.risk.content.detect
func (this *Client) DetectRiskContent(param SecurityRiskContentDetect) (result *SecurityRiskContentDetectRsp, err error) {
	contentArr := []rune(param.Content)
	if len(contentArr) > 2000 {
		err = errors.New("content length > 2000")
		return
	}
	err = this.doRequest("POST", param, &result)
	return result, err
}
