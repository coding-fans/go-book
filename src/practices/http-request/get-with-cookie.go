/**
 * FileName:   get-with-cookie.go
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

    rqst, err := http.NewRequest("GET", "http://example.com", nil)
    if err != nil {
        fmt.Println("New request failed:", err)
        return
    }

    rqst.AddCookie(&http.Cookie{
        Name: "MyCookie",
        Value: "xxxx",
    })

    rsps, err := client.Do(rqst)
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
