package main

import (
	"fmt"
	s "strings"
	"unicode"
)

var f = fmt.Printf

func main() {
	upper := s.ToUpper("Hello There")
	f("To Upper: %s\n", upper)
	f("To Lower: %s\n", s.ToLower("Hello THERE"))
	f("%s\n", s.Title("tHis wiLL be A title!"))
	f("EqualFold: %v\n", s.EqualFold("Mihalis", "MIHAlis")) //true
	f("EqualFold: %v\n", s.EqualFold("Mihalis", "MIHAli"))  //false

	f("Prefix: %v\n", s.HasPrefix("Mihalis", "Mi")) //true
	f("Prefix: %v\n", s.HasPrefix("Mihalis", "mi")) //false
	f("Suffix: %v\n", s.HasSuffix("Mihalis", "is")) //true
	f("Suffix: %v\n", s.HasSuffix("Mihalis", "IS")) //false

	f("Index: %v\n", s.Index("Mihalis", "ha")) //1
	f("Index: %v\n", s.Index("Mihalis", "Ha")) //-1
	f("Count: %v\n", s.Count("Mihalis", "i"))  //2
	f("Count: %v\n", s.Count("Mihalis", "I"))  //0
	f("Repeat: %s\n", s.Repeat("ab", 5))

	f("Compare: %v\n", s.Compare("Mihalis", "MIHALIS")) //1 a < b
	f("Compare: %v\n", s.Compare("Mihalis", "Mihalis")) //0 a = b
	f("Compare: %v\n", s.Compare("MIHALIS", "MIHalis")) //- 1 a > b
	f("Fields: %v\n", s.Fields("This is a string!"))
	f("Fields: %v\n", s.Fields("Thisis\na\tstring!"))
	f("%s\n", s.Split("abcd efg", ""))

	f("%s\n", s.Replace("abcd efg", "", "_", -1))
	f("%s\n", s.Replace("abcd efg", "", "_", 4))
	f("%s\n", s.Replace("abcd efg", "", "_", 2))
	lines := []string{"Line 1", "Line 2", "Line 3"}
	f("Join: %s\n", s.Join(lines, "+++"))
	f("SplitAfter: %s\n", s.SplitAfter("123++432++", "++"))
	trimFunction := func(c rune) bool {
		return !unicode.IsLetter(c)
	}
	f("TrimFunc: %s\n", s.TrimFunc("123 abc ABC \t .", trimFunction))
}
