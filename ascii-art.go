package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	base := ReadFile()
	art := GenerateArt(base)
	PrintArt(art)
}

func PrintArt(art []string) {
	for _, line := range art {
		fmt.Println(line)
	}
}

func GenerateArt(base []string) []string {
	s := os.Args[1]
	art := make([]string, 8)
	wordNum := 0
	for i := 0; i < len(s); i++ {
		if (s[i] == '\\' && i+1 < len(s) && s[i+1] == 'n') || s[i] == 13 && i+1 < len(s) && s[i+1] == 10 {
			newline := make([]string, 8)
			art = append(art, newline...)
			wordNum += 8
			i++
		} else if s[i] >= ' ' && s[i] <= '~' {
			line := (int(s[i]) - 32) * 9
			Convert(base, art, line, wordNum)
		}
	}
	return art
}

func Convert(base, art []string, index, wordNum int) {
	k := wordNum
	for i := index + 1; i <= index+8; i++ {
		art[k] += base[i]
		k++
	}
}

func ReadFile() []string {
	/*if len(os.Args) < 2 {
		os.Exit(3)
	}*/
	filename := "standard"
	if len(os.Args) == 3 {
		filename = os.Args[2]
	}
	file, err := os.Open(filename + ".txt")
	if err != nil {
		fmt.Print("500 internal server error")
		os.Exit(3)
	}
	var base []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		base = append(base, scanner.Text())
	}
	file.Close()
	return base
}
