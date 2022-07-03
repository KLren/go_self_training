package ptr_practice

import "fmt"

func LibTest(a int) {
	fmt.Println("Input value is: ", a)
}

// PtrPractice is used to practice pointer function
func PtrPractice(n *int) {
	*n++
	fmt.Println("\naddress to n: ", &n, "\nn point to: ", n, "\nvalue of n point to: ", *n)
}
