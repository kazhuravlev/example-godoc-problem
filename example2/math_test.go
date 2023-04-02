package math

import (
	"fmt"
)

func ExampleAdd() {
	{
		fmt.Println(Add(1, 2))
		// Output: 3
	}

	{
		_ = Add(1, 2)
	}
}
