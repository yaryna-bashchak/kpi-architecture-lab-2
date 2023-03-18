package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	lab2 "github.com/yaryna-bashchak/kpi-architecture-lab-2"
)

var (
	inputExpression = flag.String("e", "", "Expression to compute")
	inputFromFile = flag.String("f", "", "Source file")
	outputToFile = flag.String("o", "", "Output to the file")
)

func main() {
	flag.Parse()

	var input io.Reader
	var output = os.Stdout

	switch {
    case *inputExpression != "":
        input = strings.NewReader(*inputExpression)

    case *inputFromFile != "":
        f, err := os.Open(*inputFromFile)
        if err != nil {
            fmt.Fprintln(os.Stderr, "Error:", err)
            return
        }
        defer f.Close()
        input = f

    default:
        fmt.Fprintln(os.Stderr, "Input expression is empty")
        return
    }

    if *outputToFile != "" {
        f, err := os.Create(*outputToFile)

        if err != nil {
            fmt.Fprintln(os.Stderr, "Error:", err)
            return
        }

        defer f.Close()
        output = f
    }

    handler := &lab2.ComputeHandler{
        Input:  input,
        Output: output,
    }

    if err := handler.Compute(); err != nil {
        fmt.Fprintln(os.Stderr, "Error:", err)
    }
}
