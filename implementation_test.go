package lab2

import (
	"fmt"
	"testing"
	 . "gopkg.in/check.v1"
)

type PostfixToInfixSuite struct{}

var _ = Suite(&PostfixToInfixSuite{})

func (s *PostfixToInfixSuite) TestPostfixToInfix(c *C) {
	// звичайний вираз
	res, err := PostfixToInfix("2 3 +")
	c.Assert(err, IsNil)
	c.Assert(res, Equals, "2 + 3")

	// вираз з різними операціями
	res, err = PostfixToInfix("5 1 2 + 4 * + 3 -")
	c.Assert(err, IsNil)
	c.Assert(res, Equals, "5 + (1 + 2) * 4 - 3")

	// некоректні дані
	res, err = PostfixToInfix("")
	c.Assert(err, NotNil)

	res, err = PostfixToInfix("a b +")
	c.Assert(err, NotNil)

	// тест з повідомлення про помилку
	res, err = PostfixToInfix("4 2 - 3 * 5 +")
	c.Assert(err, IsNil)
	c.Assert(res, Equals, "(4 - 2) * 3 + 5")

	// недостатньо оперторів
	res, err = PostfixToInfix("2 3 + 4 5 -")
	c.Assert(err, NotNil);

	// некоректний символ
	res, err = PostfixToInfix("2 3 / 4 5 & -")
	c.Assert(err, NotNil);

	//більш складний вираз
	res, err = PostfixToInfix("1 2 + 3 4 - 5 * ^ 6 7 / /")
	c.Assert(err, IsNil)
	c.Assert(res, Equals, "((1 + 2) ^ ((3 - 4) * 5)) / (6 / 7)")
	
}
func Test(t *testing.T) { TestingT(t) }



func ExamplePostfixToInfix() {
	res, _ := PostfixToInfix("8 3 + 9 -")
	fmt.Println(res)

	// Output:
	// 8 + 3 - 9
}
