package servise

import (
	_ "GIN/models"
	"github.com/jordan-wright/email"
	"net/smtp"
)

type EmailServise struct{}

/**
 * 发送邮箱
 * @author lizibin
 */
func (EmailServise) SendEmail(registerEmail string,content string) int {
	//isHave := models.UrlInfo{}.SearchUrlByID(registerEmail,content)
	em := email.NewEmail()
	// 设置 sender 发送方 的邮箱 ， 此处可以填写自己的邮箱
	em.From = "2457665356 <2457665356@qq.com>"

	// 设置 receiver 接收方 的邮箱  此处也可以填写自己的邮箱， 就是自己发邮件给自己
	em.To = []string{registerEmail}

	// 设置主题
	em.Subject = "lizibinblog 邮件提醒"

	// 简单设置文件发送的内容，暂时设置成纯文本
	em.Text = []byte(content)

	//设置服务器相关的配置
	err := em.Send("smtp.qq.com:25", smtp.PlainAuth("", "lizibinblog@qq.com", "hqxgpsmvzqeadhjc", "smtp.qq.com"))
	if err != nil {
		return 1
	}else {
		return 0
	}

}
