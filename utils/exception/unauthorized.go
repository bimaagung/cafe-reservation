package exception

type NewUnauthorized struct{}

func (err NewUnauthorized) Error() string {
	return ""
}