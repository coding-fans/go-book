/**
 * FileName:   exercise-fibonacci-closure.go
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
}

func main() {
    f := fibonacci()
    for i := 0; i < 10; i++ {
        fmt.Println(f())
    }
}
