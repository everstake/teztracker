package mailer

type FakeMailer struct {
}

func NewFakeMailer() *FakeMailer {
	return &FakeMailer{}
}

func (m FakeMailer) Send(email string, msgType string, values map[string]string) error {
	return nil
}
