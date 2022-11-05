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
	out, _:=c.cmd.StdoutPipe()
	err, _:=c.cmd.StderrPipe()
	c.Stdout=bufio.NewScanner(out)
	c.Stderr=bufio.NewScanner(err)
	c.Stdout.Split(bufio.ScanLines)
	c.Stderr.Split(bufio.ScanLines)
	return c
}

// The actual launching command
func (c *RunosStruc) RunosCommand() {
	stat:=c.cmd.Run()
	if errors.As(stat, &c.exerr) {
		c.Status=c.exerr.ExitCode()
	}
}
