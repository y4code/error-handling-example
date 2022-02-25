package main

import (
	"errors"
	"fmt"
	"log"
)

//Method one

var (
	ErrPostNotExist = errors.New("post does not exists")
	ErrAlreadyExist = errors.New("post already exists")
)

type MyCustomError struct {
	Message string
	Detail  string
}

func (m *MyCustomError) Error() string {
	return m.Message
}

type Post struct {
	Name string
}

//func main() {
//	post, err := ControllerLayer()
//	if err != nil {
//		panic(err)
//	}
//	log.Println(post)
//}

func ControllerLayer() (*Post, error) {
	post, err := ServiceLayer()
	if err != nil {
		return nil, ErrorHandleMiddleware(err)
	}
	return post, err
}

func ErrorHandleMiddleware(err error) error {
	var (
		myCustomError *MyCustomError
	)
	if errors.Is(err, ErrPostNotExist) {
		return err
	}
	if errors.As(err, &myCustomError) {
		log.Println(myCustomError.Detail)
		return myCustomError
	}
	return err
}

func ServiceLayer() (*Post, error) {
	post, err := RepositoryLayer()
	if err != nil {
		// middle layer, here is where to add ingredients
		if errors.Is(err, ErrPostNotExist) {
			return nil, fmt.Errorf("could not get detailed info, %w from service layer", err)
		}
		if errors.Is(err, ErrAlreadyExist) {
			return nil, fmt.Errorf("already exists, %w from service layer", err)
		}
		return nil, err
	}
	return post, nil
}

func RepositoryLayer() (*Post, error) {
	//return nil, fmt.Errorf("%w from repository layer", ErrPostNotExist)
	return nil, fmt.Errorf("%w ", &MyCustomError{
		Message: "i am message",
		Detail:  "i am detail",
	})
}
