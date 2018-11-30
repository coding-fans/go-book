/**
 * FileName:   counter-server.go
 * Author:     Fasion Chan
 * @contact:   fasionchan@gmail.com
 * @version:   $Id$
 *
 * Description:
 *
 * Changelog:
 *
 **/

package main

import (
    "fmt"
    "log"
    "net/http"
)

var counter = 0

type CounterHandler struct {}

func (handler CounterHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
    counter += 1
    fmt.Fprintln(writer, counter)
}

func main() {
    server := &http.Server{
        Addr: ":8080",
        Handler: CounterHandler{},
    }

    log.Fatal(server.ListenAndServe())
}
