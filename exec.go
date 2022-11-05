package bolster
import("bufio"; "errors"; "os/exec";)

// This struct contains the execution context
type RunosStruc struct {
	Stdout *bufio.Scanner
	Stderr *bufio.Scanner
	exerr  *exec.ExitError
	cmd    *exec.Cmd
	Log    string
	Status int
}

// Runs an application, providing realtime execution logs
// This is required top setup the execution context
func RunosFactory(command string, args ...string) RunosStruc {
	c:=RunosStruc{Status: 0}
	c.cmd=exec.Command(command, args...)
	stdout, err:=c.cmd.StdoutPipe()
	Error(err)
	stderr, err:=c.cmd.StderrPipe()
	Error(err)
	c.Stdout=bufio.NewScanner(stdout)
	c.Stderr=bufio.NewScanner(stderr)
	c.Stdout.Split(bufio.ScanLines)
	c.Stderr.Split(bufio.ScanLines)
	return c
}

// The actual launching command
func (c *RunosStruc) RunosCommand() {
	stat:=c.cmd.Run()
	Error(stat)
	if errors.As(stat, &c.exerr) {
		c.Status=c.exerr.ExitCode()
	}
}
