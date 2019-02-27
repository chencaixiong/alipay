package alipay

// https://docs.alipay.com/pre-open/api_pre/alipay.user.agreement.page.sign
type AliPayUserAgreementPageSignIdentityParams struct {
	Username     string `json:"user_name"`
	CertNo       string `json:"cert_no"`
	IdentityHash string `json:"identity_hash"`
	SignUserId   string `json:"sign_user_id"`
}
type AliPayUserAgreementPageSign struct {
	AppAuthToken        string                                    `json:"-"` // 可选
	ReturnURL           string                                    `json:"-"`
	NotifyURL           string                                    `json:"notify_url"`
	PersonalProductCode string                                    `json:"personal_product_code"`
	ExternalLogonId     string                                    `json:"external_logon_id"` //用户在商户网站的登录账号，用于在签约页面展示，如果为空，则不展示
	IdentityParams      AliPayUserAgreementPageSignIdentityParams `json:"identity_params"`
}

func (this AliPayUserAgreementPageSign) APIName() string {
	return "alipay.user.agreement.page.sign"
}

func (this AliPayUserAgreementPageSign) Params() map[string]string {
	var m = make(map[string]string)
	m["app_auth_token"] = this.AppAuthToken
	m["notify_url"] = this.NotifyURL
	return m
}

func (this AliPayUserAgreementPageSign) ExtJSONParamName() string {
	return "biz_content"
}

func (this AliPayUserAgreementPageSign) ExtJSONParamValue() string {
	return marshal(this)
}

type AliPayUserAgreementPageSignResponse struct {
	Body struct {
		Code    string `json:"code"`
		Msg     string `json:"msg"`
		SubCode string `json:"sub_code"`
		SubMsg  string `json:"sub_msg"`

		//success
		ExternalAgreementNo string `json:"external_agreement_no"`
		PersonalProductCode string `json:"personal_product_code"`
		ValidTime           string `json:"valid_time"`
		SignScene           string `json:"sign_scene"`
		AgreementNo         string `json:"agreement_no"`
		ZmOpenId            string `json:"zm_open_id"`
		InvalidTime         string `json:"invalid_time"`
		SignTime            string `json:"sign_time"`
		AlipayUserId        string `json:"alipay_user_id"`
		Status              string `json:"status"`
		ForexEligible       string `json:"forex_eligible"`
		ExternalLogonId     string `json:"external_logon_id"`
		AlipayLogonId       string `json:"alipay_logon_id"`
	} `json:"alipay_user_agreement_page_sign_response"`
	ErrorResp struct {
		Code    string `json:"code"`
		Msg     string `json:"msg"`
		SubCode string `json:"sub_code"`
		SubMsg  string `json:"sub_msg"`
	} `json:"error_response"`
	Sign string `json:"sign"`
}

func (this *AliPayUserAgreementPageSignResponse) IsSuccess() (bool, string) {
	if this.Body.Code == K_SUCCESS_CODE {
		return true, ""
	}
	return false, marshal(this.ErrorResp)
}

////////////////////////////////////////////////////////////////////////////////
// https://docs.alipay.com/pre-open/api_pre/alipay.user.agreement.query
type AliPayUserAgreementQuery struct {
	AppAuthToken        string `json:"-"` // 可选
	PersonalProductCode string `json:"personal_product_code"`
	AlipayUserId        string `json:"alipay_user_id"`        //用户的支付宝账号对应 的支付宝唯一用户号，以 2088 开头的 16 位纯数字 组成
	AlipayLogonId       string `json:"alipay_logon_id"`       //用户的支付宝登录账号，支持邮箱或手机号码格式。本参数与alipay_user_id 不可同时为空
	SignScene           string `json:"sign_scene"`            //签约协议场景，商户和支付宝签约时确定，商户可咨询技术支持。
	ExternalAgreementNo string `json:"external_agreement_no"` //代扣协议中标示用户的唯一签约号(确保在商户系统中 唯一)
	ThirdPartyType      string `json:"third_party_type"`      //签约第三方主体类型
	AgreementNo         string `json:"agreement_no"`          //支付宝系统中用以唯一标识用户签约记录的编号（用户签约成功后的协议号 ） ，如果传了该参数，其他参数会被忽略
}

func (this AliPayUserAgreementQuery) APIName() string {
	return "alipay.user.agreement.query"
}

func (this AliPayUserAgreementQuery) Params() map[string]string {
	var m = make(map[string]string)
	m["app_auth_token"] = this.AppAuthToken
	return m
}

func (this AliPayUserAgreementQuery) ExtJSONParamName() string {
	return "biz_content"
}

func (this AliPayUserAgreementQuery) ExtJSONParamValue() string {
	return marshal(this)
}

type AliPayUserAgreementQueryResponse struct {
	Body struct {
		Code    string `json:"code"`
		Msg     string `json:"msg"`
		SubCode string `json:"sub_code"`
		SubMsg  string `json:"sub_msg"`

		//success
		//必填
		ValidTime           string `json:"valid_time"`
		AlipayLogonId       string `json:"alipay_logon_id"`
		InvalidTime         string `json:"invalid_time"`
		PricipalType        string `json:"pricipal_type"`
		DeviceId            string `json:"device_id"`
		PrincipalId         string `json:"principal_id"`
		SignScene           string `json:"sign_scene"`
		AgreementNo         string `json:"agreement_no"`
		ThirdPartyType      string `json:"third_party_type"` //签约第三方主体类型
		Status              string `json:"status"`
		SignTime            string `json:"sign_time"`
		PersonalProductCode string `json:"personal_product_code"`
		//选填
		ExternalAgreementNo string `json:"external_agreement_no"`
		ZmOpenId            string `json:"zm_open_id"`
		ExternalLogonId     string `json:"external_logon_id"`
	} `json:"alipay_user_agreement_query_response"`
	ErrorResp struct {
		Code    string `json:"code"`
		Msg     string `json:"msg"`
		SubCode string `json:"sub_code"`
		SubMsg  string `json:"sub_msg"`
	} `json:"error_response"`
	Sign string `json:"sign"`
}

func (this *AliPayUserAgreementQueryResponse) IsSuccess() (bool, string) {
	if this.Body.Code == K_SUCCESS_CODE {
		return true, ""
	}
	return false, marshal(this.ErrorResp)
}
