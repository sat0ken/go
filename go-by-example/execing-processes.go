package main

import (
    "os"
    "os/exec"
    "syscall"
)

func main() {
    binary, lookerr := exec.LookPath("ls")
    if lookerr != nil {
        panic(lookerr)
    }
    args := []string{"ls", "-a", "-l", "-h"}
    env := os.Environ()

    execErr := syscall.Exec(binary, args, env)
    if execErr != nil {
        panic(execErr)
    }
}
