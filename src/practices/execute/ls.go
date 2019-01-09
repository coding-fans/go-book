/**
 * FileName:   ls.go
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
    "bytes"
    "fmt"
    "os/exec"
)

func main() {
    cmdName := "ls"

    path, err := exec.LookPath(cmdName)
    if err != nil {
        fmt.Println("Can't find command:", cmdName)
        return
    }

    cmd := exec.Command(path, "-l")

    var stdout bytes.Buffer
    cmd.Stdout = &stdout

    err = cmd.Run()
    if err != nil {
        fmt.Println("Run command error:", err)
        return
    }

    fmt.Println(stdout.String())
}
