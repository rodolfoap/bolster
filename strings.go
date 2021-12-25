package main
import("bytes")

func Join(values... string) string{
	var b bytes.Buffer
	for _, v:=range values{ b.WriteString(v); }
	return b.String()
}
