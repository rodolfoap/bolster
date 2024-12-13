package gx
import("bufio"; "fmt"; "os"; "testing"; "github.com/stretchr/testify/assert";)

var executionLog string

// A function that allows printing results in realtime
func readerr(scanner *bufio.Scanner, channel string) {
	for scanner.Scan() {
		executionLog+=scanner.Text()+"\n"
		fmt.Println(channel, scanner.Text())
	}
}

// Entrypoint
func TestGXLib(t *testing.T) {
	/*** Strings **********************************************************************************************/

	Printhr() // strings.go::Printhr()
	// "============..."
	fmt.Println(Join("Hello", ",", " ", "World", "!")) // strings.go::Join()
	// "Hello, World!"

	/*** IO ***************************************************************************************************/

	Printhr()

	s:=ReadFile("go.mod") // io.go::ReadFile()
	fmt.Println(s)

	/*** EXEC *************************************************************************************************/

	Printhr()

	c:=RunosFactory("./runstatus.py", "10") // exec.go::RunosFactory()
	go readerr(c.Stdout, "out:")
	go readerr(c.Stderr, "err:")
	c.RunosCommand() // exec.go::RunosCommand()
	fmt.Println("EXITSTATUS:", c.Status)
	fmt.Println("LOG_LENGTH:", len(executionLog))

	/*** Misc *************************************************************************************************/

	Printhr()

	Log("Ternary operator: true: ", Iff(1==2, false, true))
	Log("Ternary operator: 22: ", Iff(2==3, 1, 22))
	Log("Ternary operator: d: ", Iff("a"=="b", "c", "d"))

	m:=ProgressiveMean{} // Mean base parameters can be predefined if required
	fmt.Printf("Count: %v; Average: %v\n", m.C, m.Avg)

	for i:=10; i<17; i++ {
		m.UpdateMean(float64(i)) // type to average must be float64
		fmt.Printf("New element: %v; Count: %v; Average: %v\n", i, m.C, m.Avg)
	}
	assert.Equal(t, m.Avg, 13.0)

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

	// Fatal(err) is tested separately
}

// TestFatal is used to do tests which are supposed to be fatal
func TestFatal(t *testing.T) {
	origLogFatalf:=logFatalf
	defer func() { logFatalf=origLogFatalf }()

	errors:=[]string{}
	logFatalf=func(format string, args ...interface{}) {
		errors=append(errors, format)
		fmt.Printf("Log.Fatalf() called: %#v\n", errors)
	}
	_, err:=os.Open("I.dont.exist")
	Fatal(err)
	// 2022/11/06 08:02:26 /home/rap/git/gx/gx_test.go github.com/rodolfoap/gx.Test_lib() [63] FATAL: open I.dont.exist: no such file or directory
	assert.Greater(t, len(errors), 0)
}
