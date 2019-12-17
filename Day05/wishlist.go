package main

import "fmt"

func splitSwap(txt string) string {
	return txt[len(txt)/2:len(txt)] + txt[0:len(txt)/2]
}

func swapChars(txt string) string {
	chars := []rune(txt)
	for i := 0; i < len(chars)-1; i++ {
		tmp := chars[i+1]
		chars[i+1] = chars[i]
		chars[i] = tmp
	}
	return string(chars)
}

func main() {
	txt := "PonnistallHoppeslottTrommesett"
	fmt.Println(txt)
	fmt.Println(swapChars(splitSwap(txt)))
}
