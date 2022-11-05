package bolster
import("fmt"; "log"; "runtime";)

// Produces a trace containing filename, function, line and a message
// "/home/rap/git/bolster/bolster.go main.main() [9] 9 1 3" Adds spaces when neither is a string
func Trace(msgs ...interface{}) {
	pc:=make([]uintptr, 1)
	frame, _:=runtime.CallersFrames(pc[:runtime.Callers(2, pc)]).Next()
	fmt.Printf("%s %s() [%d] %s\n", frame.File, frame.Function, frame.Line, fmt.Sprint(msgs...))
}

// Produces a trace containing date and time, filename, function, line and a message
// "2022/11/05 11:16:46 /home/rap/git/bolster/bolster.go main.main() [10] Using log"
func TimeTrace(msgs ...interface{}) {
	pc:=make([]uintptr, 1)
	frame, _:=runtime.CallersFrames(pc[:runtime.Callers(2, pc)]).Next()
	log.Printf("%s %s() [%d] %s\n", frame.File, frame.Function, frame.Line, fmt.Sprint(msgs...))
}

// Raises an error message
func Error(err interface{}) {
	pc:=make([]uintptr, 1)
	frame, _:=runtime.CallersFrames(pc[:runtime.Callers(2, pc)]).Next()
	if err!=nil {
		log.Printf("%s %s() [%d] ERROR: %v", frame.File, frame.Function, frame.Line, err)
	}
}

// Raises an error message followed by os.Exit(1)
func Fatal(err interface{}) {
	pc:=make([]uintptr, 1)
	frame, _:=runtime.CallersFrames(pc[:runtime.Callers(2, pc)]).Next()
	if err!=nil {
		log.Fatalf("%s %s() [%d] FATAL: %v", frame.File, frame.Function, frame.Line, err)
	}
}
