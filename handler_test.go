package lab2

import (
	"bytes"
	"strings"
	. "gopkg.in/check.v1"
)

func (s *PostfixToInfixSuite) TestComputeHandler(c *C) {
    var buf bytes.Buffer

    handler := ComputeHandler{
        Input:  strings.NewReader("2 3 - 4 5 * ^"),
        Output: &buf,
    }

    c.Assert(handler.Compute(), IsNil)
    c.Assert(buf.String(), Equals, "(2 - 3) ^ (4 * 5)")
}

func (s *PostfixToInfixSuite) TestComputeHandlerWithError(c *C) {
    var buf bytes.Buffer

    handler := ComputeHandler{
        Input:  strings.NewReader("4 5 * a -"),
        Output: &buf,
    }

    c.Assert(handler.Compute(), NotNil)
}
