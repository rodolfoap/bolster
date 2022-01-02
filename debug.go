package golib
import("log"; "fmt"; "runtime";)

func Println(msgs ...interface{}) {
	pc:=make([]uintptr, 1)
	frame, _:=runtime.CallersFrames(pc[:runtime.Callers(2, pc)]).Next()
	fmt.Println(frame.File, frame.Function, "["+string(frame.Line)+"]:", fmt.Sprintln(msgs...))
}
func Printf(format string, msgs ...interface{}) {
	pc:=make([]uintptr, 1)
	frame, _:=runtime.CallersFrames(pc[:runtime.Callers(2, pc)]).Next()
	fmt.Printf("%s %s()[%d]: %s\n", frame.File, frame.Function, frame.Line, fmt.Sprintf(format, msgs...))
}
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
func Fatal(err interface{}){
	pc:=make([]uintptr, 1)
	frame, _:=runtime.CallersFrames(pc[:runtime.Callers(2, pc)]).Next()
	if err!=nil{
		log.Fatalf("%s %s()[%d] FATAL: %v", frame.File, frame.Function, frame.Line, err)
	}
}
