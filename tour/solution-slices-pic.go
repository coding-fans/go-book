/**
 * FileName:   solution-slices-pic.go
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

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
    s := make([][]uint8, dy)
    for y := 0; y < dy; y++ {
        sx := make([]uint8, dx)
        for x := 0; x < dx; x++ {
            sx[x] = uint8((x+y)/2)
            //sx[x] = uint8(x*y)
            //sx[x] = uint8(x^y)
        }
        s[y] = sx
    }
    return s
}

func main() {
    pic.Show(Pic)
}
