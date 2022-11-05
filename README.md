# Github library example

## Create the library:

Project in `github.com/rodolfoap/bolster`:

#### `go.mod`:

```
module github.com/rodolfoap/bolster

go 1.15
```

#### `strings.go`:

```
package bolster
import("bytes")

func Join(values... string) string{
	var b bytes.Buffer
	for _, v:=range values{ b.WriteString(v); }
	return b.String()
}
```
#### Create a tag:

```
$ git tag v0.6.0
$ git push origin v0.6.0
```

## Use it:
```
package main
import("fmt"; "github.com/rodolfoap/bolster")

func main() {
	fmt.Println(Join("Hello", ",", " ", "World", "!"))
}
```

## Useful commands:
```
# echo $GOROOT
go clean -modcache
go get -u github.com/rodolfoap/bolster@v0.7.0
```
