/**
 * FileName:   get-with-timeout.go
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
    "time"
)

func main() {
    var client = &http.Client{
        Timeout: time.Second * 5,
    }

    rsps, err := client.Get("http://example.com")
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
