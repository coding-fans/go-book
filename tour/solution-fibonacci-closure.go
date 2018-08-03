/**
 * FileName:   solution-fibonacci-closure.go
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

// fabonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
    a, b := 0, 1
    return func() int {
        value := a
        a, b = b, a+b
        return value
    }
}

func main() {
    f := fibonacci()
    for i := 0; i < 10; i++ {
        fmt.Println(f())
    }
}
