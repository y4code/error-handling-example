# error-handling-example

use `errors` pkg to handle error in mutiple layered RESTful service


method one
use `errors.As` in top layer, handle the error created by `errors.New` in middle layers, as a result, the custom error will be "throw" from bottom layer to top layer


method two
implement `interface{ Is(error) bool }` interface for custom error, and use `errors.Is` in top layer to determine if the error stack contains custom error 
