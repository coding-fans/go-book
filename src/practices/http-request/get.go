/**
 * FileName:   get.go
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
    "net/http"
    "io/ioutil"
)

func main() {
    rsps, err := http.Get("http://example.com")
    if err != nil {
        fmt.Println("Request failed:", err)
        return
    }
    defer rsps.Body.Close()

    body, err := ioutil.ReadAll(rsps.Body)
    if err != nil {
        fmt.Println("Read body failed:", err)
        return
    }

    fmt.Println(string(body))
}
