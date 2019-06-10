package main

import "fmt"

func ShowcaseVariables() {

	/*
		Variables
				Normal/Short Declaration
				Block scope outer cant access inner, inner can acccess outer
				---

				Тhese are the following rules for naming a Golang variable:

			    A name must begin with a letter, and can have any number of additional letters and numbers.
			    A variable name cannot start with a number.
			    A variable name cannot contain spaces.
			    If the name of a variable begins with a lower-case letter, it can only be accessed within the current package this is considered as unexported variables.
			    If the name of a variable begins with a capital letter, it can be accessed from packages outside the current package one this is considered as exported variables.
			    If a name consists of multiple words, each word after the first should be capitalized like this: empName, EmpAddress, etc.
			    Variable names are case-sensitive (car, Car and CAR are three different variables).
	*/

	//Constants:
	const x1 string = "Hello, World"
	const (
		x = 10
		y = 20
		z = 30
	)

	var a8 uint8   // Unsigned 8-bit integers (0 to 255)
	var a16 uint16 // Unsigned 16-bit integers (0 to 65535)
	a16 = 65535
	var a32 uint32 // Unsigned 32-bit integers (0 to 4294967295)
	a32 = 1231212312
	var a64 uint64 // Unsigned 64-bit integers (0 to 18446744073709551615)
	a64 = 12412412323312

	fmt.Println(a8, a16, a32, a64)

	var b8 int8 //Signed 8-bit integers (-128 to 127)
	b8 = -5
	var b16 int16 // Signed 32-bit integers (-2147483648 to 2147483647)
	b16 = -12312
	var b32 int32 // Signed 64-bit integers (-9223372036854775808 to 9223372036854775807)
	b32 = -1231231212
	var b64 int64
	b64 = -312312331312312

	fmt.Println(b8, b16, b32, b32, b64)

	var f32 float32 = 1234.23445       // IEEE-754 32-bit floating-point numbers
	var f64 float64 = 123123214.412312 // IEEE-754 64-bit floating-point numbers

	fmt.Println(f32, f64)

	var c64 complex64 = 8789579877.721   // Complex numbers with float32 real and imaginary parts
	var c128 complex128 = 985214745.7216 // Complex numbers with float64 real and imaginary parts

	fmt.Println(c64, c128)

	var corr, incorr bool
	corr = true
	incorr = false
	fmt.Println(corr, incorr)

	/*
		Strings in Go are a sequence of UTF-8 characters enclosed in double quotes (") "Hello, World" or backticks `Hello, World`.
		If you use the single quote (') you mean one character (encoded in UTF-8) — which is not a string in Go.
		Once assigned to a variable the string can not be changed: strings in Go are immutable

		indexing a string yields its bytes, not its characters

		`` raw string
		"" regular string - can contain escape characters

		 Go source code is UTF-8, so the source code for the string literal is UTF-8 text.
		 If that string literal contains no escape sequences, which a raw string cannot,
		 the constructed string will hold exactly the source text between the quotes.
		 Thus by definition and by construction the raw string will always contain a valid UTF-8 representation of its contents.
		 larly, unless it contains UTF-8-breaking escapes like those from the previous section, a regular string literal
		 will also always contain valid UTF-8.

		  Some people think Go strings are always UTF-8, but they are not: only string literals are UTF-8.
		  As we showed in the previous section, string values can contain arbitrary bytes; as we showed in this one, string literals always contain UTF-8 text as long as they have no byte-level escapes.


		Notice how we don't need to set string as type
	*/
	var city, country = "Sofia", "Bulgaria"
	var c rune = 'a' // or var c = 'a' rune is int32 alias representing bytes which may or may not be utf-8 encoded
	// In this case, however, we need to
	var city2 string
	city2 = "Plovdiv"

	fmt.Println(city, city2, c, country)
}
