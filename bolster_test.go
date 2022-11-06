package bolster
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

	/*** Debug ************************************************************************************************/

	Printhr()
	Trace()
	// /home/rap/git/gobox/54-bolster.use/app.go main.main() [45]
	Trace("Monk", " ", "Parker")
	// /home/rap/git/gobox/54-bolster.use/app.go main.main() [47] Monk Parker
	Trace(9, 1, 3)
	// /home/rap/git/gobox/54-bolster.use/app.go main.main() [49] 9 1 3
	TimeTrace("Using log")
	// 2022/11/06 07:49:30 /home/rap/git/bolster/bolster_test.go github.com/rodolfoap/bolster.Test_lib() [51] Using log
	Log("Using simple log")
	// 2022/11/06 07:45:54 /home/rap/git/bolster/bolster_test.go github.com/rodolfoap/bolster.Test_lib() [53] Using simple log
	Logf("Formatted log: Type of '1': %T", 1) // Does not need CR/LF
	// 2022/11/06 07:47:27 /home/rap/git/bolster/bolster_test.go github.com/rodolfoap/bolster.Test_lib() [55] Formatted log: Type of '1': int
	Tracef("Formatted TRACE: Type of '1': %T", 1) // Does not need CR/LF
	// /home/rap/git/bolster/bolster_test.go github.com/rodolfoap/bolster.Test_lib() [57] Formatted TRACE: Type of '1': int

	_, err:=os.Open("I.dont.exist")
	Error(err)
	// 2022/11/05 17:09:24 /home/rap/git/gobox/54-bolster.use/app.go main.main() [61] ERROR: open I.dont.exist: no such file or directory

	_, err=os.Open("I.dont.exist")
	Fatal(err)
	// 2022/11/05 17:09:24 /home/rap/git/gobox/54-bolster.use/app.go main.main() [65] FATAL: open I.dont.exist: no such file or directory

	// Will not run, the FATAL error has exited.
	Printhr()
}
