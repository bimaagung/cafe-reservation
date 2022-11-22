package exception

import (
	"errors"
	"fmt"
)

func CheckError(err error) {
	if err != nil {
		fmt.Errorf(err.Error())
		errors.New(err.Error())
		//panic(err.Error())
	}
}