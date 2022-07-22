package lib

import (
	"fmt"
	"net/http"
)

var I int

func Test() {
	I = 1
}

func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello world, %s", request.URL.Path)
}
