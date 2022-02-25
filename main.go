package main

import (
	"errors"
	"fmt"
	"log"
)

// Method two

var (
	ErrArticleNotExist     = errors.New("article does not exists")
	ErrArticleAlreadyExist = errors.New("article already exists")
)

type MyRealCustomError struct {
	Code    int
	Message string
	Detail  string
}

var (
	ErrInternal = MyRealCustomError{
		Code:    50000,
		Message: "error message for front-end to display",
		Detail:  "real custom error detail, for dev to debug",
	}
	ErrNotFound = MyRealCustomError{
		Code:    400004,
		Message: "not found bla bla, error message for front-end to display",
		Detail:  "not found bla bla, real custom error detail, for dev to debug",
	}
)

func (m *MyRealCustomError) Error() string {
	return m.Message
}

func (m *MyRealCustomError) Is(tgt error) bool {
	target, ok := tgt.(*MyRealCustomError)
	if !ok {
		return false
	}
	return m.Code == target.Code
}

type Article struct {
	Name string
}

func main() {
	article, err := RealControllerLayer()
	if err != nil {
		panic(err)
	}
	log.Println(article)
}

func RealControllerLayer() (*Article, error) {
	article, err := RealServiceLayer()
	if err != nil {
		return nil, RealErrorHandleMiddleware(err)
	}
	return article, err
}

func RealErrorHandleMiddleware(err error) error {
	if errors.Is(err, ErrArticleNotExist) {
		return err
	}
	if errors.Is(err, &ErrInternal) {
		log.Println(ErrInternal.Detail)
		return &ErrInternal
	}
	return err
}

func RealServiceLayer() (*Article, error) {
	article, err := RealRepositoryLayer()
	if err != nil {
		if errors.Is(err, ErrArticleNotExist) {
			return nil, fmt.Errorf("could not get detailed info, %w, id %v", err, "wuhu")
		}
		if errors.Is(err, ErrAlreadyExist) {
			return nil, fmt.Errorf("already exists, %w", err)
		}
		return nil, err
	}
	return article, nil
}

func RealRepositoryLayer() (*Article, error) {
	//return nil, fmt.Errorf("%w ", ErrArticleNotExist)
	//or
	//return nil, fmt.Errorf("%w ", &ErrInternal)
	//or
	return nil, fmt.Errorf("%w ", &MyRealCustomError{
		Code:    50000,
		Message: "i am message",
		Detail:  "i am detail",
	})
}
