package golib
import("os/exec";)

func RunCommand(name string, command... string)string{
	cmd:=exec.Command(name, command...)
	std, err:=cmd.CombinedOutput()
	Fatal(err)
	return string(std)
}
