# Github library example

## Create the library:

Project in `github.com/rodolfoap/golib`:

#### `go.mod`:

```
module github.com/rodolfoap/golib

go 1.15
```

#### `strings.go`:

```
package golib
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
import("fmt"; "github.com/rodolfoap/golib")

func main() {
	fmt.Println(Join("Hello", ",", " ", "World", "!"))
}
```

## Useful commands:
```
# GOROOT=/usr/local/go
go clean -modcache
sudo find /usr/local/go -name cache -print -exec rm -r {} \;
sudo find /usr/local/go -name golib -print -exec rm -r {} \;
sudo find /usr/local/go -name rodolfoap -print -exec rm -r {} \;
sed -i '/require github.com.rodolfoap.golib/d' go.mod
go get -u github.com/rodolfoap/golib@v0.7.0
```
