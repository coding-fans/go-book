/**
 * FileName:   functions.go
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


func add(x int, y int) int {
    return x + y
}


func main() {
    fmt.Println(add(43, 13))
}
