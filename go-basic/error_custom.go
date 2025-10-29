package main

import "fmt"

type validationError struct {
	Message string
}

func (v *validationError) Error() string {
	return v.Message
}

type notFoundError struct {
	Message string
}

func (n *notFoundError) Error() string {
	return n.Message
}

func SaveData(id string, data any) error {
	if id == "" {
		return &validationError{Message: "validation error"}
	}

	if id != "Aidil" {
		return &notFoundError{Message: "data not found"}
	}

	return nil // func ok
}

func main() {
	err := SaveData("Aidil", nil)

	if err != nil {
		// terjadi error
		// if validationErr, ok := err.(*validationError); ok {
		// 	fmt.Println("validation error :", validationErr.Error())
		// } else if notFoundErr, ok := err.(*notFoundError); ok {
		// 	fmt.Println("not found :", notFoundErr.Error())
		// } else {
		// 	fmt.Println("unkown error :", err.Error())
		// }

		switch finalError := err.(type) {
		case *validationError:
			fmt.Println("validation error :", finalError.Error())
		case *notFoundError:
			fmt.Println("validation error :", finalError.Error())
		default :
			fmt.Println("unkown error :", err.Error())
		}
	} else {
		// success
		fmt.Println("Sukses")
	}
}