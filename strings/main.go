package strings

import (
	"fmt"
	"unicode/utf8"
)

func Main() {
	func() {
		name := "Hello World"
		fmt.Println(name)
	}()

	func() {
		printBytes := func(s string) {
			fmt.Printf("Bytes: ")
			for i := 0; i < len(s); i++ {
				fmt.Printf("%x ", s[i])
			}
		}

		name := "Hello World"
		fmt.Printf("String: %s\n", name)
		printBytes(name)
	}()

	func() {
		printBytes := func(s string) {
			fmt.Printf("Bytes: ")
			for i := 0; i < len(s); i++ {
				fmt.Printf("%x ", s[i])
			}
		}

		printChars := func(s string) {
			fmt.Printf("Characters: ")
			for i := 0; i < len(s); i++ {
				fmt.Printf("%c ", s[i])
			}
		}

		name := "Hello World"
		fmt.Printf("String: %s\n", name)
		printChars(name)
		fmt.Printf("\n")
		printBytes(name)
		fmt.Printf("\n\n")
		name = "Señor"
		fmt.Printf("String: %s\n", name)
		printChars(name)
		fmt.Printf("\n")
		printBytes(name)
	}()

	func() {
		charsAndBytePosition := func(s string) {
			for index, rune := range s {
				fmt.Printf("%c starts at byte %d\n", rune, index)
			}
		}

		name := "Señor"
		charsAndBytePosition(name)
	}()

	func() {
		byteSlice := []byte{0x43, 0x61, 0x66, 0xC3, 0xA9}
		str := string(byteSlice)
		fmt.Println(str)
	}()

	func() {
		byteSlice := []byte{67, 97, 102, 195, 169} //decimal equivalent of {'\x43', '\x61', '\x66', '\xC3', '\xA9'}
		str := string(byteSlice)
		fmt.Println(str)
	}()

	func() {
		runeSlice := []rune{0x0053, 0x0065, 0x00f1, 0x006f, 0x0072}
		str := string(runeSlice)
		fmt.Println(str)
	}()

	func() {
		word1 := "Señor"
		fmt.Printf("String: %s\n", word1)
		fmt.Printf("Length: %d\n", utf8.RuneCountInString(word1))
		fmt.Printf("Number of bytes: %d\n", len(word1))

		fmt.Printf("\n")
		word2 := "Pets"
		fmt.Printf("String: %s\n", word2)
		fmt.Printf("Length: %d\n", utf8.RuneCountInString(word2))
		fmt.Printf("Number of bytes: %d\n", len(word2))
	}()

	func() {
		compareStrings := func(str1 string, str2 string) {
			if str1 == str2 {
				fmt.Printf("%s and %s are equal\n", str1, str2)
				return
			}
			fmt.Printf("%s and %s are not equal\n", str1, str2)
		}

		string1 := "Go"
		string2 := "Go"
		compareStrings(string1, string2)

		string3 := "hello"
		string4 := "world"
		compareStrings(string3, string4)
	}()

	func() {
		string1 := "Go"
		string2 := "is awesome"
		result := string1 + " " + string2
		fmt.Println(result)
		result2 := fmt.Sprintf("%s %s", string1, string2)
		fmt.Println(result2)
	}()

	func() {
		/*mutate := func(s string) string {
			s[0] = 'a' //any valid unicode character within single quote is a rune
			cannot assign to s[0]
			return s
		}*/

		mutate := func(s []rune) string {
			s[0] = 'a'
			return string(s)
		}

		h := "hello"
		fmt.Println(mutate([]rune(h)))
	}()
}
