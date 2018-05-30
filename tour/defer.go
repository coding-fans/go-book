/**
 * FileName:   defer.go
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

import "fmt"


func main() {
    defer fmt.Println("world")

    fmt.Println("hello")
}
