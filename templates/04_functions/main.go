package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

var tpl *template.Template

// Create a FuncMap to register functions.
// "uc" is the ToUpper function from pachage strings but called uc in template.
// "ft" is a func declared below that slices string and returns first three characters.
var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

func init() {
	// Need to call the New function before a parse function so that the functions can be
	// attached before parsing, otherwise compiler will complain the functions don't exist.
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	if len(s) >= 3 {
		s = s[:3]
	}
	return s
}

func main() {

	sages := []string{"Buddha", "Gandhi", "Muhammad"}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", sages)
	if err != nil {
		log.Fatalln(err)
	}
}
