package math

import "fmt"

func ExampleAdd() {
	// Code block
	{
		fmt.Println(Add(1, 2))
		// Output: 3
	} // <--- This bracket will not present in output HTML

	// Code block (second one). This block will not present in HTML at all
	{
		fmt.Println(Add(1, 2))
	}
}
