/**
 * FileName:   slices-of-slices.go
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
     "strings"
 )

 func main() {
     // Create a tic-tac-toe board.
     board := [][]string{
         []string{"-", "_", "_"},
         []string{"-", "_", "_"},
         []string{"-", "_", "_"},
     }

     // The players take turns.
     board[0][0] = "X"
     board[2][2] = "O"
     board[1][2] = "X"
     board[1][0] = "O"
     board[0][2] = "X"

     for i := 0; i < len(board); i++ {
         fmt.Printf("%s\n", strings.Join(board[i], " "))
     }
 }
