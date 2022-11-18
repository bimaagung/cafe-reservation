package exception

type ClientError struct {
	Message string
}

func (clientError ClientError) Error() string {
	return clientError.Message
}