package smail

import (
	"crypto/tls"

	"gopkg.in/gomail.v2"
)

// Mailer 邮件配置
type Mailer struct {
	User     string   `yaml:"user"`
	Pass     string   `yaml:"pass"`
	Smtp     string   `yaml:"smtp"`
	Port     int      `yaml:"port"`
	MailTo   []string `yaml:"mailTo"`
	Subject  string
	Body     string
	Filename string
}

// // NewMailer new一个对象
// func New(cfg string) *Mailer {

// }

// SendMail send mail by group
func (s *Mailer) SendMailByGroup() error {

	m := gomail.NewMessage()
	m.SetHeader("From", s.User)
	m.SetHeader("To", s.MailTo...)
	m.SetHeader("Subject", s.Subject)
	m.SetBody("text/html", s.Body)
	if len(s.Filename) > 0 {
		m.Attach(s.Filename)
	}
	d := gomail.NewDialer(s.Smtp, s.Port, s.User, s.Pass)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	err := d.DialAndSend(m)
	return err
}

func (s *Mailer) SendMail() error {
	var err error
	for _, v := range s.MailTo {
		m := gomail.NewMessage()
		m.SetHeader("From", s.User)
		m.SetHeader("To", v)
		m.SetHeader("Subject", s.Subject)
		m.SetBody("text/html", s.Body)
		if len(s.Filename) > 0 {
			m.Attach(s.Filename)
		}
		d := gomail.NewDialer(s.Smtp, s.Port, s.User, s.Pass)
		d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
		err = d.DialAndSend(m)
	}
	return err
}
