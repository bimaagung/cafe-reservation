package exception

import "fmt"

type Unathorized struct {
	Message string
}

func (unauthorized Unathorized) Error() string {
	fmt.Println(unauthorized.Message)
	return unauthorized.Message
}