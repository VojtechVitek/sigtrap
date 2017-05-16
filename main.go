package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, signals...)

	fmt.Printf("PID: %v\n", os.Getpid())

	go func() {
		for {
			fmt.Printf("%s\n", <-sig)
		}
	}()

	select {}
}

var signals = []os.Signal{
	syscall.SIGABRT,
	syscall.SIGALRM,
	syscall.SIGBUS,
	syscall.SIGCHLD,
	//syscall.SIGCLD,
	syscall.SIGCONT,
	syscall.SIGFPE,
	syscall.SIGHUP,
	syscall.SIGILL,
	syscall.SIGINT,
	syscall.SIGIO,
	syscall.SIGIOT,
	syscall.SIGKILL,
	syscall.SIGPIPE,
	//syscall.SIGPOLL,
	syscall.SIGPROF,
	//syscall.SIGPWR,
	syscall.SIGQUIT,
	syscall.SIGSEGV,
	//syscall.SIGSTKFLT,
	syscall.SIGSTOP,
	syscall.SIGSYS,
	syscall.SIGTERM,
	syscall.SIGTRAP,
	syscall.SIGTSTP,
	syscall.SIGTTIN,
	syscall.SIGTTOU,
	//syscall.SIGUNUSED,
	syscall.SIGURG,
	syscall.SIGUSR1,
	syscall.SIGUSR2,
	syscall.SIGVTALRM,
	syscall.SIGWINCH,
	syscall.SIGXCPU,
	syscall.SIGXFSZ,
}
