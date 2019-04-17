package main

import (
	"log"
	"net/smtp"
	"strings"
)

func SendEmail(email string) string {
	//邮箱账号、授权码、发送方式
	auth := smtp.PlainAuth("","","","smtp.qq.com")
	to := []string{email}
	//发件人名称
	nickname := ""
	//发件人邮箱
	user := ""
	//邮箱主题
	subject := ""
	//发送格式
	content_type := "Content-Type: text/plain; charset=UTF-8"
	//邮件内容
	body := ""
	msg := []byte("To: " + strings.Join(to, ",") + "\r\nFrom: " + nickname +
		"<" + user + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	emailerr := smtp.SendMail("smtp.qq.com:25", auth, user, to, msg)
	if emailerr != nil {
		log.Fatal("Error：",emailerr)
	}
	return "发送成功"
}

func main()  {
	//对方邮箱账号
	SendEmail("")
}
