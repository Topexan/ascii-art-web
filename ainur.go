package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os/exec"
)

type ViewData struct {
	Input  string
	Fs     string
	Output string
}

func main() {
	http.HandleFunc("/index/", indexHandler)

	fmt.Println("Server is listening...")
	http.ListenAndServe(":8181", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	check := false
	fs := r.FormValue("fs")
	if string(fs) == "" {
		fs = "standard"
	}
	input := r.FormValue("input")
	fmt.Println(input)
	output, err := exec.Command("ascii-art-file/ascii-art-file.exe", input, fs).Output()
	runesInput := []rune(input)
	for i := range runesInput {
		if runesInput[i] < 32 || runesInput[i] > 126 {
			if runesInput[i] != 10 && runesInput[i] != 13 {
				check = true
			}
		}
	}
	if check {
		fmt.Fprintln(w, "400 bad request")
		return
	}
	for i := range runesInput {
		if runesInput[i] == 10 {
			fmt.Println("10")
		}
		if runesInput[i] == 13 {
			fmt.Println("13")
		}
	}
	if err != nil {
		fmt.Printf("%v\n%v\n%v", "Error", err, output)
	}
	if r.URL.Path != "/index/" {
		fmt.Fprintln(w, "404 page not found")
		return
	}
	if string(output) == "500 internal server error" {
		fmt.Fprintln(w, "500 internal server error")
		return
	}

	fmt.Println(string(output))
	data := ViewData{
		Input:  input,
		Fs:     fs,
		Output: string(output)}

	tmpl, _ := template.ParseFiles("index.html")
	tmpl.Execute(w, data)
}