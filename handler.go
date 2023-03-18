package lab2

import (
	"bytes"
	"io"
)

type ComputeHandler struct {
	Input  io.Reader
	Output io.Writer
}

func (ch *ComputeHandler) Compute() error {
    buf := new(bytes.Buffer)

    _, err := io.Copy(buf, ch.Input)
    if err != nil {
        return err
    }

    input := buf.Bytes()

    output, err := PostfixToInfix(string(input))
    if err != nil {
        return err
    }

    _, err = ch.Output.Write([]byte(output))
    if err != nil {
        return err
    }

    return nil
}
