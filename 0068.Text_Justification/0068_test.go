package leetcode

import "fmt"

func Example_case1() {
	ans := fullJustify(
		[]string{"This", "is", "an", "example", "of", "text", "justification."},
		16,
	)
	for _, str := range ans {
		fmt.Printf("%q\n", str)
	}
	// Output:
	// "This    is    an"
	// "example  of text"
	// "justification.  "
}

func Example_case2() {
	ans := fullJustify(
		[]string{"What", "must", "be", "acknowledgment", "shall", "be"},
		16,
	)
	for _, str := range ans {
		fmt.Printf("%q\n", str)
	}
	// Output:
	// "What   must   be"
	// "acknowledgment  "
	// "shall be        "
}

func Example_case3() {
	ans := fullJustify(
		[]string{"Science", "is", "what", "we", "understand", "well", "enough",
			"to", "explain", "to", "a", "computer.", "Art", "is", "everything",
			"else", "we", "do"},
		20,
	)
	for _, str := range ans {
		fmt.Printf("%q\n", str)
	}
	// Output:
	// "Science  is  what we"
	// "understand      well"
	// "enough to explain to"
	// "a  computer.  Art is"
	// "everything  else  we"
	// "do                  "
}
