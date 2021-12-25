package golib
import("log"; "fmt"; "runtime";)

func Tracef(format string, msgs ...interface{}) {
	pc:=make([]uintptr, 1)
	frame, _:=runtime.CallersFrames(pc[:runtime.Callers(2, pc)]).Next()
	log.Printf("%s %s()[%d]: %s\n", frame.File, frame.Function, frame.Line, fmt.Sprintf(format, msgs...))
}
func Trace() {
	pc:=make([]uintptr, 1)
	frame, _:=runtime.CallersFrames(pc[:runtime.Callers(2, pc)]).Next()
	log.Printf("%s %s()[%d]\n", frame.File, frame.Function, frame.Line)
}
