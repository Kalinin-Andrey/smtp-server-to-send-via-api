package email

const ()

type Email struct {
	ID    uint
	Email string
	Phone uint
}

func (e Email) Validate() error {
	return nil
}
