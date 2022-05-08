package mail

import (
	"fmt"
	"github.com/imroc/req/v3"
	sysConfig "hanya-go/pkg/config"
	"hanya-go/pkg/helpers/str"
	"hanya-go/pkg/logger"
)

type SendcloudResponseBody struct {
	Result     bool   `json:"result"`
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
	Info       struct {
		EmailIdList []string `json:"emailIdList"`
	} `json:"info"`
}

type SENDCLOUD struct {
}

// Send 实现 email.Driver interface 的 Send 方法
func (s *SENDCLOUD) Send(email Email, config map[string]string) bool {

	logger.DebugJSON("SENDCLOUD发送邮件", "EMAIL对象原始信息", email)

	var respBody SendcloudResponseBody

	postData := map[string]string{
		"apiUser": sysConfig.Get("mail.sendcloud.apiUser"),
		"apiKey":  sysConfig.Get("mail.sendcloud.apiKey"),
		"from":    fmt.Sprintf("%v <%v>", email.FromName, email.FromAddress),
		"to":      str.SliceSplit(email.To, ";"),
		"subject": email.Subject,
		"html":    str.ByteToString(email.HTML),
	}

	logger.DebugJSON("SENDCLOUD发送邮件", "发件详情", postData)

	client := req.C()
	resp, err := client.R().SetFormData(postData).SetResult(&respBody).Post("https://api.sendcloud.net/apiv2/mail/send")

	if err != nil {
		logger.ErrorString("SENDCLOUD发送邮件", "发送失败", err.Error())
		return false
	}

	if resp.IsSuccess() && respBody.Result {
		return true
	} else {
		return false
	}
}
