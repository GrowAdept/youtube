// Strings
package main

import "fmt"

func main() {
	// A string is a sequence of bytes
	// A bytes is a unit of memory made up of 8 bits
	// A bit is a unit of memory with only 2 option, 1 or 0
	var a = "Hello there" // string literal encoded at UTF-8
	var b = "你好"
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println("H length:", len("H"))
	fmt.Println("你 length:", len("你"))
	/*
		for i := 0; i < len("H"); i++ {
			fmt.Printf(" H hexidecimal: %x decimal: %3d binary: %08b", "H"[i], "H"[i], "H"[i])
		}
		fmt.Print("\n")
	*/
	// fmt.Println(string(0x0048)) //using hexadecimal
	// fmt.Println(string(72))     //using decimal
	// fmt.Println(string(48))     //needs 0x infront so compiler knows it's hexadecimal
	/*
		for i := 0; i < len("你"); i++ {
			fmt.Printf("你 hexidecimal: %x decimal: %d binary: %08b\n", "你"[i], "你"[i], "你"[i]) // returns hex of each byte, not hex of code point
		}
	*/
	// fmt.Println(string(0x4F60))
	// fmt.Println(string(20320))

	// fmt.Print("\n", string(a[0]), "\n")
	// fmt.Println(string(b[0])) // doesn't work b/c this is 3 bytes long, not 1
	// fmt.Println(string(b[0:3]))
	// fmt.Println(string(b[3:6]))

	// fmt.Println("length of Hello there:", len("Hello there"))
	// fmt.Println("length of 你好:", len("你好"))
	/*
		for index, s := range "Hello there" {
			fmt.Printf("%3c is index number %3d decimal: %6d hexidecimal: %x binary: %08b\n", s, index, s, s, s)
		}
	*/
	// fmt.Print(string(0x0048), string(0x0065), string(0x006c), string(0x006c), string(0x006f), "\n")
	// fmt.Print(string(72), string(101), string(108), string(108), string(111), "\n")
	/*
		for index, s := range "你好" {
			fmt.Printf("%3c is index number %3d decimal: %6d hexidecimal: %x binary: %08b\n", s, index, s, s, s)
		}
	*/
	// fmt.Print(string(0x4F60), string(0x597d), "\n")
	// fmt.Print(string(20320), string(22909), "\n")

	// byteSlice1 := []byte{0x48, 0x65, 0x6c, 0x6c, 0x6f}
	// byteSlice1 := []byte{0x0048, 0x0065, 0x006c, 0x006c, 0x006f}
	// fmt.Println("string(byteSlice1):", string(byteSlice1))
	// byteSlice2 := []byte{0x4f60, 0x597d} // does not work, multi-byte character does not fit in single byte
	// fmt.Println(string(byteSlice2))      // does not work

	// runeSlice1 := []rune{0x0048, 0x0065, 0x006c, 0x006c, 0x006f}
	// fmt.Println("string(runeSlice1):", string(runeSlice1))
	// runeSlice2 := []rune{0x4f60, 0x597d}
	// fmt.Println("string(runeSlice2):", string(runeSlice2))

	/*
		fmt.Printf("type of byteSlice1: %T\n", byteSlice1)
		fmt.Printf("type of runeSlice1: %T\n", runeSlice1)
		fmt.Printf("type of runeSlice1: %T\n", runeSlice2)
	*/

	/*
		for i := 0; i < 3; i++ {
			fmt.Printf("你 hexidecimal: %x decimal: %d binary: %08b\n", "你"[i], "你"[i], "你"[i]) // returns hex of each byte, not hex of code point
		}
	*/
}
