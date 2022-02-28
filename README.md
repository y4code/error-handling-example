# error-handling-example

Use `errors` pkg to handle error in mutiple layered RESTful service

Inpired by https://stackoverflow.com/questions/63306404/overriding-errors-is-not-working-with-custom-errors


# method one
`method_one.go`

Use `errors.As` in top layer, handle the error created by `errors.New` in middle layers, as a result, the custom error will be "throw" from bottom layer to top layer


# method two
`main.go`

Implement `interface{ Is(error) bool }` interface for custom error, and use `errors.Is` in top layer to determine if the error stack contains custom error 

> Suggest method two
