package main

import (
	"gopkg.in/yaml.v3"
	"io"
	"fmt"
	"log"
	"os"

	"github.com/davecgh/go-spew/spew"
	
)


func go_yaml_dump(input io.Reader, output io.Writer) (err error) {

	decoder := yaml.NewDecoder(input)
	for {
		var data interface{}
		err := decoder.Decode(&data)
		if err == io.EOF {
			return nil
		} else if err != nil {
			return err
		}
		spew.Fdump(output, data)
		fmt.Println("\n//----------------------")
	}
}

func main() {
	
	err := go_yaml_dump(os.Stdin, os.Stdout)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
}
