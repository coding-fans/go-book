/**
 * FileName:   demo.go
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
    "time"
)

func main() {
    now := time.Now()
    tokyoTZ, _ := time.LoadLocation("Asia/Tokyo")
    tokyoTime := now.In(tokyoTZ)
    fmt.Printf("Local time: %s\nTokyo time: %s\n", now, tokyoTime)

    _, err := http.Get("https://www.baidu.com/")
    if err == nil {
        fmt.Println("Baidu website is UP")
    } else {
        fmt.Printf("Baidu website is DOWN\nErr: %s\n", err.Error())
    }
}
