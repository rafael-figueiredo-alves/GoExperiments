package main

import "fmt"

func main() {
	var a int = 10
	var b float32 = 20.5
	var c string = "Hello, Go!"
	fmt.Print("Integer value:", a, " ", "Float value:", b, " ", "String value:", c, "\n") // Using Print QUe não pula linha, e não adiciona espaços (white spaces)
	fmt.Println("Integer value:", a, "Float value:", b, "String value:", c)               // Using Println que pula linha automaticamente
	fmt.Printf("Integer value: %d, Float value: %.2f, String value: %s\n", a, b, c)       // Using Printf for formatted output

	var i = 15.5
	var txt = "Hello World!"

	fmt.Printf("%v\n", i)   // default format
	fmt.Printf("%#v\n", i)  // Go-syntax representation
	fmt.Printf("%v%%\n", i) // print percentage sign
	fmt.Printf("%T\n", i)   // type of the variable

	fmt.Printf("%v\n", txt)  // default format
	fmt.Printf("%#v\n", txt) // Go-syntax representation
	fmt.Printf("%T\n", txt)  // type of the variable

	var f = 15

	fmt.Printf("%b\n", f)   // binary representation
	fmt.Printf("%d\n", f)   // decimal representation
	fmt.Printf("%+d\n", f)  // sign for positive numbers
	fmt.Printf("%o\n", f)   // octal representation
	fmt.Printf("%O\n", f)   // octal representation with 0o prefix
	fmt.Printf("%x\n", f)   // hexadecimal representation
	fmt.Printf("%X\n", f)   // hexadecimal representation (uppercase)
	fmt.Printf("%#x\n", f)  // hexadecimal representation with 0x prefix
	fmt.Printf("%4d\n", f)  // width of 4, right-aligned
	fmt.Printf("%-4d\n", f) // width of 4, left-aligned
	fmt.Printf("%04d\n", f) // width of 4, zero-padded

	txt = "Hello"

	fmt.Printf("%s\n", txt)   // string
	fmt.Printf("%q\n", txt)   // quoted string
	fmt.Printf("%8s\n", txt)  // width of 8, right-aligned
	fmt.Printf("%-8s\n", txt) // width of 8, left-aligned
	fmt.Printf("%x\n", txt)   // hexadecimal representation of string
	fmt.Printf("% x\n", txt)  // hexadecimal representation with spaces between bytes

	var g = true
	var j = false

	fmt.Printf("%t\n", g)
	fmt.Printf("%t\n", j)
}
