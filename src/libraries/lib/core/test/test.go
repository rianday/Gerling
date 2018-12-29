package test

import (
	"fmt"
	"libraries/models/test"
)

func RunCrudExample() {
	var tesst = test.Mahasiswa{}

	result, err := tesst.Delete()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("delete successful : ", result)
	}

	result, err = tesst.Create()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("insert successful : ", result)
	}

	result, err = tesst.Update()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("update successful : ", result)
	}

	result2 := tesst.Read()
	for i, data := range result2 {
		fmt.Printf("index (%d) : %s\n", i, data.Name)
	}
	// fmt.Println(result2[0].Name)
}
