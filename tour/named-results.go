/**
 * FileName:   named-results.go
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


 func split(sum int) (x, y int) {
     x = sum * 4 / 9
     y = sum - x
     return
 }


 func main() {
     fmt.Println(split(17))
 }
