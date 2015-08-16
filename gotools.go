//a go library of tools i use
//Dana Urban
//04/01/2015

package gotools

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//get input from the command line
//i am sure their are better ways but this is my way
func Stdin_line_read(a string) string {
	//a is used to pass in a prompt
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(a + "\n")
	b, _ := reader.ReadString('\n')
	return strings.TrimSpace(b)
}
func Slr(a string) string { // same as a above.  shorter name
	//a is used to pass in a prompt
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(a + "\n")
	b, _ := reader.ReadString('\n')
	return strings.TrimSpace(b)
}
