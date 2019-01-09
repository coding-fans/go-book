/**
 * FileName:   echo-server.go
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
    "io/ioutil"
)

type EchoHandler struct {}

func (handler EchoHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
    // set response header
    writer.Header().Add("X-Data", "foo")

    // set response cookie
    http.SetCookie(writer, &http.Cookie{
        Name: "x-cookie",
        Value: "bar",
        MaxAge: 86400,
        Secure: true,
    })

    writer.WriteHeader(200)

    // echo network info
    fmt.Fprintln(writer, "===== Network =====")
    fmt.Fprintln(writer, "Remote Address:", request.RemoteAddr)
    fmt.Fprintln(writer)

    // echo request line info
    fmt.Fprintln(writer, "===== Request Line =====")
    fmt.Fprintln(writer, "Method: ", request.Method)
    fmt.Fprintln(writer, "URL: ", request.URL)
    fmt.Fprintln(writer, "Host: ", request.Host)
    //fmt.Fprintln(writer, "URI: ", request.RequestURI)
    fmt.Fprintf(writer, "Protocol: %v major=%v minor=%v\n", request.Proto,
        request.ProtoMajor, request.ProtoMinor)
    fmt.Fprintln(writer)

    // echo headers
    fmt.Fprintln(writer, "===== Header =====")
    for k, v := range request.Header {
        fmt.Fprintf(writer, "%v: %v\n", k, v)
    }
    fmt.Fprintln(writer)

    // echo body
    body, err := ioutil.ReadAll(request.Body)
    if err == nil && len(body) > 0 {
        fmt.Fprintln(writer, "===== Raw Body =====")
        fmt.Fprintln(writer, string(body))
    }
}

func main() {
    server := &http.Server{
        Addr: ":8080",
        Handler: EchoHandler{},
    }

    log.Println("Server starting...")
    log.Fatal(server.ListenAndServe())
}
