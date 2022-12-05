package gx
import("fmt"; "log"; "os"; "runtime"; "strconv";)

// Default Log.Fatalf()
var logFatalf=log.Fatalf

// BDEBUG=={ 1 | 0 }
func isDebugMode() bool {
	ttable:=[]bool{false, true}
	index:=os.Getenv("BDEBUG")
	if len(index)==0 {
		return false
	} else {
		index, err:=strconv.Atoi(os.Getenv("BDEBUG"))
		Fatal(err)
		return ttable[index]
	}
}

// Produces a trace containing filename, function, line and a message, IF BDEBUG==1
// "/home/rap/git/gx/gx.go main.main() [9] 9 1 3" Adds spaces when neither is a string
func Trace(msgs ...interface{}) {
	pc:=make([]uintptr, 1)
	frame, _:=runtime.CallersFrames(pc[:runtime.Callers(2, pc)]).Next()
	if isDebugMode() {
		log.Printf("%s %s() [%d] %s\n", frame.File, frame.Function, frame.Line, fmt.Sprint(msgs...))
	}
}

// Produces a trace containing filename, function, line and a message, IF BDEBUG==1
// "/home/rap/git/gx/gx.go main.main() [9] 9 1 3" Adds spaces when neither is a string
func Tracef(format string, msgs ...interface{}) {
	pc:=make([]uintptr, 1)
	frame, _:=runtime.CallersFrames(pc[:runtime.Callers(2, pc)]).Next()
	if isDebugMode() {
		log.Printf("%s %s() [%d] %s\n", frame.File, frame.Function, frame.Line, fmt.Sprintf(format, msgs...))
	}
}

// Raises an error message, independently of BDEBUG
// Produces a trace containing date and time, filename, function, line and a message
// "2022/11/05 11:16:46 /home/rap/git/gx/gx.go main.main() [10] Using log"
func Log(msgs ...interface{}) {
	pc:=make([]uintptr, 1)
	frame, _:=runtime.CallersFrames(pc[:runtime.Callers(2, pc)]).Next()
	log.Printf("%s %s() [%d] %s\n", frame.File, frame.Function, frame.Line, fmt.Sprint(msgs...))
}

// Raises an error message, independently of BDEBUG
// Produces a trace containing date and time, filename, function, line and a message
// "2022/11/05 11:16:46 /home/rap/git/gx/gx.go main.main() [10] Using log"
func Logf(format string, msgs ...interface{}) {
	pc:=make([]uintptr, 1)
	frame, _:=runtime.CallersFrames(pc[:runtime.Callers(2, pc)]).Next()
	log.Printf("%s %s() [%d] %s\n", frame.File, frame.Function, frame.Line, fmt.Sprintf(format, msgs...))
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
		logFatalf("%s %s() [%d] FATAL: %v", frame.File, frame.Function, frame.Line, err)
	}
}
