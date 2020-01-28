package mailManager

import(
	"strings"
	"bytes"
	"net/smtp"
	"mime/quotedprintable"
)

var auth smtp.Auth = nil;

var from string = "rohandvivedi@gmail.com";

func InitMailClient(password string) {
	auth = smtp.PlainAuth("", from, password, "smtp.gmail.com")
}

func WriteEmail(dest []string, contentType string, subject string, bodyMessage string) string {

	header := make(map[string]string)
	header["From"] = from

	receipient := ""

	for _, user := range dest {
		receipient = receipient + user
	}

	header["To"] = receipient
	header["Subject"] = subject
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = contentType + "; charset=\"utf-8\""
	header["Content-Transfer-Encoding"] = "quoted-printable"
	header["Content-Disposition"] = "inline"

	message := ""

	for key, value := range header {
		message += key + ": " + value + "\r\n";
	}

	var encodedMessage bytes.Buffer

	finalMessage := quotedprintable.NewWriter(&encodedMessage)
	finalMessage.Write([]byte(bodyMessage))
	finalMessage.Close()

	message += "\r\n" + encodedMessage.String()

	return message
}

func WriteHTMLEmail(dest []string, subject, bodyMessage string) string {
	return WriteEmail(dest, "text/html", subject, bodyMessage)
}

func WritePlainEmail(dest []string, subject, bodyMessage string) string {
	return WriteEmail(dest, "text/plain", subject, bodyMessage)
}

func SendMail(Dest []string, Subject, bodyMessage string) (error) {

	msg := "From: " + from + "\n" +
		"To: " + strings.Join(Dest, ",") + "\n" +
		"Subject: " + Subject + "\n" + bodyMessage

	return smtp.SendMail("smtp.gmail.com:587", auth,
		from, Dest, []byte(msg))
}
