package main

import (
	"flag"
	"log"
	"strings"

	"github.com/serialt/smail"
	"github.com/serialt/sugar/v3"
)

var (
	recvUser string
	subject  string
	body     string
	filename string
	mailtype string
)

type Config struct {
	Mailer smail.Mailer `yaml:"mailer"`
}

// Recver 邮件信息处理
func init() {
	flag.StringVar(&recvUser, "c", "", "收邮件地址,格式为 a11@qq.com,a22@gmail.com")
	flag.StringVar(&mailtype, "t", "", "邮件发送方式,g 群发邮件,s 一对一发邮件,默认是g")
	flag.StringVar(&subject, "s", "", "邮件主题")
	flag.StringVar(&body, "m", "", "邮件内容")
	flag.StringVar(&filename, "f", "", "添加的附件")
	flag.Parse()
}

func main() {
	var confg *Config
	err := sugar.LoadConfig("/Users/serialt/.smail.yaml", &confg)
	if err != nil {
		log.Fatalf("Read config failed: %v", err)
	}
	mail := confg.Mailer
	mail.MailTo = strings.Split(recvUser, ",")
	mail.Subject = subject
	mail.Body = body
	mail.Filename = filename

	switch mailtype {
	case "g", "":
		err := mail.SendMailByGroup()
		if err != nil {
			log.Println("Send mail failed", err)
			return
		}
		log.Println("Send mail successfully!")
	default:
		err := mail.SendMail()
		if err != nil {
			log.Println("Send mail failed", err)
			return
		}
		log.Println("Send mail successfully!")

	}

}
