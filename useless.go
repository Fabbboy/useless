package main

import "syscall"

func main() {
	syscall.Kill(syscall.Getpid(), syscall.SIGTERM)
}
