package exception

type NewClientError struct {
	Message string
}

func (err NewClientError) Error() string {
	return err.Message
}