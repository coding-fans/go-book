/**
 * FileName:   post.go
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
    rsps, err := http.Post("http://example.com/", "plain/text", nil)
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
