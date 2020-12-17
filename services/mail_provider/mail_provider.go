package mail_provider

type MailProvider struct {
}

func New() *MailProvider {
	return &MailProvider{}
}

// todo
func (m MailProvider) Send(email string, kind string, values map[string]string) error {
	return nil
}
