package main
import "os";import "os/exec";import "syscall";
func main() { if len(os.Args) == 1 { fork() } else { forkfork() }}
func fork() {
    cmd := exec.Command("/proc/self/exe", append([]string{"fork"})...)
    cmd.Stdin = os.Stdin; cmd.Stdout = os.Stdout; cmd.Stderr = os.Stderr
    cmd.SysProcAttr = &syscall.SysProcAttr{ Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID }
    cmd.Run()
}
func forkfork() {
    cmd := exec.Command("/bin/sh"); cmd.Stdin = os.Stdin; cmd.Stdout = os.Stdout; cmd.Stderr = os.Stderr
    syscall.Sethostname([]byte("majka-code16challenge"))
    syscall.Chroot("/tmp/cfs"); os.Chdir("/")
    syscall.Mount("proc", "proc", "proc", 0, "")
    cmd.Run()
}