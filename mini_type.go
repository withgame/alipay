package alipay

////////////////////////////////////////////////////////////////////////////////
// https://docs.alipay.com/mini/api/templatemessage

type MiniTemplateMessage struct {
	ToUserId       string `json:"to_user_id"`
	FormId         string `json:"form_id"`
	UserTemplateId string `json:"user_template_id"`
	Page           string `json:"page"`
	Data           string `json:"data"`
}

func (this MiniTemplateMessage) APIName() string {
	return "alipay.open.app.mini.templatemessage.send"
}

func (this MiniTemplateMessage) Params() map[string]string {
	var m = make(map[string]string)
	m["to_user_id"] = this.ToUserId
	m["form_id"] = this.FormId
	m["user_template_id"] = this.UserTemplateId
	m["page"] = this.Page
	m["data"] = this.Data
	return m
}

type MiniTemplateMessageRsp struct {
	Content struct {
		Code    string `json:"code"`
		Msg     string `json:"msg"`
		SubCode string `json:"sub_code"`
		SubMsg  string `json:"sub_msg"`
	} `json:"alipay_open_app_mini_templatemessage_send_response"`
	Sign string `json:"sign"`
}

////////////////////////////////////////////////////////////////////////////////
// https://docs.alipay.com/mini/api/openapi-qrcode

type MiniOpenApiQRcode struct {
	UrlParam   string `json:"url_param"`
	QueryParam string `json:"query_param"`
	Describe   string `json:"describe"`
}

func (this MiniOpenApiQRcode) APIName() string {
	return "alipay.open.app.qrcode.create"
}

func (this MiniOpenApiQRcode) Params() map[string]string {
	var m = make(map[string]string)
	m["url_param"] = this.UrlParam
	m["query_param"] = this.QueryParam
	m["describe"] = this.Describe
	return m
}

type MiniOpenApiQRcodeRsp struct {
	Content struct {
		Code      string `json:"code"`
		Msg       string `json:"msg"`
		SubCode   string `json:"sub_code"`
		SubMsg    string `json:"sub_msg"`
		QRCodeUrl string `json:"qr_code_url"`
	} `json:"alipay_open_app_qrcode_create_response"`
	Sign string `json:"sign"`
}

////////////////////////////////////////////////////////////////////////////////
// https://docs.open.alipay.com/api_49/alipay.security.risk.content.detect

type SecurityRiskContentDetect struct {
	Content string `json:"content"`
}

func (this SecurityRiskContentDetect) APIName() string {
	return "alipay.security.risk.content.detect"
}

func (this SecurityRiskContentDetect) Params() map[string]string {
	var m = make(map[string]string)
	m["content"] = this.Content
	return m
}

type SecurityRiskContentDetectRsp struct {
	Content struct {
		Code     string   `json:"code"`
		Msg      string   `json:"msg"`
		SubCode  string   `json:"sub_code"`
		SubMsg   string   `json:"sub_msg"`
		Action   string   `json:"action"`
		Keywords []string `json:"keywords"`
		UniqueId string   `json:"unique_id"`
	} `json:"alipay_security_risk_content_detect_response"`
	Sign string `json:"sign"`
}
