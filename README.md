# yamlspew

Simple program to parse YAML with Golang and output the parsed tree with 'spew':[https://godoc.org/github.com/davecgh/go-spew/spew], a deep pretty printer for Go data structures.

yamlspew reads from standard input and outputs every document as 'spew' rendition of the parse objects. Documents are seperated by a // comment line.

Example usage:

```
$ echo '{ int: 2, float: 3.4e5, bool: true, nothing: ~ , sequence: [1,2], base64: !!binary WUVT } ' | ./yamlspew
(map[interface {}]interface {}) (len=6) {
 (string) (len=8) "sequence": ([]interface {}) (len=2 cap=2) {
  (int) 1,
  (int) 2
 },
 (string) (len=6) "base64": (string) (len=3) "YES",
 (string) (len=3) "int": (int) 2,
 (string) (len=5) "float": (float64) 340000,
 (string) (len=4) "bool": (bool) true,
 (string) (len=7) "nothing": (interface {}) <nil>
}

```

yamlok uses the Go language YAML parser "gopkg.in/yaml.v3"
