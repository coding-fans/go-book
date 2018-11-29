/**
 * FileName:   variables-with-initializers.go
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

var i, j int = 1, 2


func main() {
    var c, python, java = true, false, "no!"
    fmt.Println(i, j, c, python, java)
}
