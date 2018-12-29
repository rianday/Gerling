package auth

import (
	"fmt"
	"libraries/models/sys"
)

func Register() {
	var tesst = sys.User{}

	names := []interface{}{1: "name"}

	err, result := tesst.Add(names)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("insert successful : ", result)
	}

}
