package exception

type NewNotFoundError struct {
	Message string
}

func (err NewNotFoundError) Error() string {
	return err.Message
}