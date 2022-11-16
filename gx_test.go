package gx
import("bufio"; "fmt"; "os"; "testing";)

var executionLog string

// A function that allows printing results in realtime
func readerr(scanner *bufio.Scanner, channel string) {
	for scanner.Scan() {
		executionLog+=scanner.Text()+"\n"
		fmt.Println(channel, scanner.Text())
	}
}

// Entrypoint
func Test_lib(t *testing.T) {
	/*** Strings **********************************************************************************************/

	Printhr()
	// "============..."
	fmt.Println(Join("Hello", ",", " ", "World", "!"))
	// "Hello, World!"

	/*** IO ***************************************************************************************************/

	Printhr()

	// FileToString(string) string
	s:=ReadFile("go.mod")
	fmt.Println(s)

	/*** EXEC *************************************************************************************************/

	Printhr()

	c:=RunosFactory("./runstatus.py", "10")
	go readerr(c.Stdout, "out:")
	go readerr(c.Stderr, "err:")
	c.RunosCommand()
	fmt.Println("EXITSTATUS:", c.Status)
	fmt.Println("LOG_LENGTH:", len(executionLog))

	/*** Misc *************************************************************************************************/

	Printhr()

	Log("Ternary operator: true: ", iff(1==2, false, true))
	Log("Ternary operator: 22: ", iff(2==3, 1, 22))
	Log("Ternary operator: d: ", iff("a"=="b", "c", "d"))

	/*** Debug ************************************************************************************************/

	Printhr()
	Trace()
	// 2022/11/06 08:02:26 /home/rap/git/gx/gx_test.go github.com/rodolfoap/gx.Test_lib() [45]
	Trace("Monk", " ", "Parker")
	// 2022/11/06 08:02:26 /home/rap/git/gx/gx_test.go github.com/rodolfoap/gx.Test_lib() [47] Monk Parker
	Trace(9, 1, 3)
	// 2022/11/06 08:02:26 /home/rap/git/gx/gx_test.go github.com/rodolfoap/gx.Test_lib() [49] 9 1 3
	Log("Using simple log")
	// 2022/11/06 08:02:26 /home/rap/git/gx/gx_test.go github.com/rodolfoap/gx.Test_lib() [51] Using simple log
	Logf("Formatted log: Type of '1': %T", 1) // Does not need CR/LF because the message is part of a log
	// 2022/11/06 08:02:26 /home/rap/git/gx/gx_test.go github.com/rodolfoap/gx.Test_lib() [53] Formatted log: Type of '1': int
	Tracef("Formatted TRACE: Type of '1': %T", 1) // Does not need CR/LF because the message is part of a log
	// 2022/11/06 08:02:26 /home/rap/git/gx/gx_test.go github.com/rodolfoap/gx.Test_lib() [55] Formatted TRACE: Type of '1': int

	_, err:=os.Open("I.dont.exist")
	Error(err)
	// 2022/11/06 08:02:26 /home/rap/git/gx/gx_test.go github.com/rodolfoap/gx.Test_lib() [59] ERROR: open I.dont.exist: no such file or directory

	_, err=os.Open("I.dont.exist")
	Fatal(err)
	// 2022/11/06 08:02:26 /home/rap/git/gx/gx_test.go github.com/rodolfoap/gx.Test_lib() [63] FATAL: open I.dont.exist: no such file or directory

	// Will not run, the FATAL error has exited.
	Printhr()
}
