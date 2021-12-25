package golib
import("fmt"; "strings"; "bytes"; "golang.org/x/term";)

func Join(values... string) string{
	var b bytes.Buffer
	for _, v:=range values{ b.WriteString(v); }
	return b.String()
}

func Printhr(){
	fmt.Println(HR())
}

func HR()string{
        if term.IsTerminal(0) {
                width, _, _:=term.GetSize(0)
                return strings.Repeat("=", width)
        }
        return ""
}
