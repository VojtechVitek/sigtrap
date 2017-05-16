sigtrap: Trap and print signals.
======

# Usage
```bash
$ sigtrap
PID: 54506
^Cinterrupt
^Cinterrupt
user defined signal 1
user defined signal 2
alarm clock
child exited
trace/BPT trap
continued
quit
interrupt
quit
Killed: 9
```

From another terminal
```bash
$ kill -l
 1) SIGHUP	 2) SIGINT	 3) SIGQUIT	 4) SIGILL	 5) SIGTRAP
 6) SIGABRT	 7) SIGEMT	 8) SIGFPE	 9) SIGKILL	10) SIGBUS
11) SIGSEGV	12) SIGSYS	13) SIGPIPE	14) SIGALRM	15) SIGTERM
16) SIGURG	17) SIGSTOP	18) SIGTSTP	19) SIGCONT	20) SIGCHLD
21) SIGTTIN	22) SIGTTOU	23) SIGIO	24) SIGXCPU	25) SIGXFSZ
26) SIGVTALRM	27) SIGPROF	28) SIGWINCH	29) SIGINFO	30) SIGUSR1
31) SIGUSR2

$ kill -30 54506
$ kill -31 54506
$ kill -14 54506
$ kill -20 54506
$ kill -5 54506
$ kill -19 54506
$ kill -3 54506
$ kill -2 54506
$ kill -3 54506
$ kill -9 54506
```

# Installation

```bash
go get -u github.com/VojtechVitek/sigtrap
```
*You might need to [download Go](https://golang.org/dl/) first.*

# License

Licensed under the [MIT License](./LICENSE).
