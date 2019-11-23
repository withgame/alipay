package alipay

const (
	kProductionPublicAppAuthorize = "https://openauth.alipay.com/oauth2/publicAppAuthorize.htm"
	kSandboxPublicAppAuthorize    = "https://openauth.alipaydev.com/oauth2/publicAppAuthorize.htm"
)

const (
	kProductionAppToAppAuth = "https://openauth.alipay.com/oauth2/appToAppAuth.htm"
	kSandboxAppToAppAuth    = "https://openauth.alipaydev.com/oauth2/appToAppAuth.htm"
)

////////////////////////////////////////////////////////////////////////////////
// https://docs.open.alipay.com/api_9/alipay.system.oauth.token
type SystemOauthToken struct {
	AppAuthToken string `json:"-"` // 可选
	GrantType    string `json:"-"` // 值为 authorization_code 时，代表用code换取；值为refresh_token时，代表用refresh_token换取
	Code         string `json:"-"`
	RefreshToken string `json:"-"`
}

func (this SystemOauthToken) APIName() string {
	return "alipay.system.oauth.token"
}

func (this SystemOauthToken) Params() map[string]string {
	var m = make(map[string]string)
	m["app_auth_token"] = this.AppAuthToken
	m["grant_type"] = this.GrantType
	if this.Code != "" {
		m["code"] = this.Code
	}
	if this.RefreshToken != "" {
		m["refresh_token"] = this.RefreshToken
	}
	return m
}

type SystemOauthTokenRsp struct {
	Content struct {
		Code         string `json:"code"`
		Msg          string `json:"msg"`
		SubCode      string `json:"sub_code"`
		SubMsg       string `json:"sub_msg"`
		UserId       string `json:"user_id"`
		AccessToken  string `json:"access_token"`
		ExpiresIn    int64  `json:"expires_in"`
		RefreshToken string `json:"refresh_token"`
		ReExpiresIn  int64  `json:"re_expires_in"`
	} `json:"alipay_system_oauth_token_response"`
	Error *struct {
		Code    string `json:"code"`
		Msg     string `json:"msg"`
		SubCode string `json:"sub_code"`
		SubMsg  string `json:"sub_msg"`
	} `json:"error_response"` // 不要访问此结构体
	Sign string `json:"sign"`
}

////////////////////////////////////////////////////////////////////////////////
// https://docs.open.alipay.com/api_2/alipay.user.info.share
type UserInfoShare struct {
	AppAuthToken string `json:"-"` // 可选
	AuthToken    string `json:"-"` // 是
}

func (this UserInfoShare) APIName() string {
	return "alipay.user.info.share"
}

func (this UserInfoShare) Params() map[string]string {
	var m = make(map[string]string)
	m["app_auth_token"] = this.AppAuthToken
	m["auth_token"] = this.AuthToken
	return m
}

type UserInfoShareRsp struct {
	Content struct {
		Code               string `json:"code"`
		Msg                string `json:"msg"`
		SubCode            string `json:"sub_code"`
		SubMsg             string `json:"sub_msg"`
		AuthNo             string `json:"auth_no"`
		UserId             string `json:"user_id"`
		Avatar             string `json:"avatar"`
		Province           string `json:"province"`
		City               string `json:"city"`
		NickName           string `json:"nick_name"`
		IsStudentCertified string `json:"is_student_certified"`
		UserType           string `json:"user_type"`
		UserStatus         string `json:"user_status"`
		IsCertified        string `json:"is_certified"`
		Gender             string `json:"gender"`
	} `json:"alipay_user_info_share_response"`
	Sign string `json:"sign"`
}

////////////////////////////////////////////////////////////////////////////////
// https://docs.open.alipay.com/api_9/alipay.open.auth.token.app
type OpenAuthTokenApp struct {
	GrantType    string `json:"grant_type"` // 值为 authorization_code 时，代表用code换取；值为refresh_token时，代表用refresh_token换取
	Code         string `json:"code"`
	RefreshToken string `json:"refresh_token"`
}

func (this OpenAuthTokenApp) APIName() string {
	return "alipay.open.auth.token.app"
}

func (this OpenAuthTokenApp) Params() map[string]string {
	var m = make(map[string]string)
	m["grant_type"] = this.GrantType
	if this.Code != "" {
		m["code"] = this.Code
	}
	if this.RefreshToken != "" {
		m["refresh_token"] = this.RefreshToken
	}
	return m
}

type OpenAuthTokenAppRsp struct {
	Content struct {
		Code            string `json:"code"`
		Msg             string `json:"msg"`
		SubCode         string `json:"sub_code"`
		SubMsg          string `json:"sub_msg"`
		AppAuthToken    string `json:"app_auth_token"`
		UserId          string `json:"user_id"`
		AuthAppId       string `json:"auth_app_id"`
		ExpiresIn       int64  `json:"expires_in"`
		ReExpiresIn     int64  `json:"re_expires_in"`
		AppRefreshToken string `json:"app_refresh_token"`
	} `json:"alipay_open_auth_token_app_response"`
	Sign string `json:"sign"`
}

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
