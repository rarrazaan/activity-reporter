package cons

type (
	ContentEmailType string
	SubjectEmail     string
)

const (
	SmtpAuthAddress   = "smtp.gmail.com"
	SmtpServerAddress = "smtp.gmail.com:587"

	SubjectVerificationEmail SubjectEmail     = "Mini Socmed Verifiction Email"
	VerificationEmailContent ContentEmailType = `
	<h2><b>Welcome to Mini Socmed!</h2>
	<p><Please click on the link below to complete your verifiction process:
	<br>
	<a href="%s">Verifivation email link</a>/p>
	<p>Thank you</p>
	`
)
