package main

import (
	"flag"
	"fmt"
	"html/template"
	"os"
)

var (
	f bool
)

func init() {
	flag.BoolVar(&f, "f", false, "force to overwrite file")
	flag.Usage = usage
}

func main() {
	flag.Parse()

	genFile := GenFile{
		TemplatePath: "cmd/gen/index.tmpl",
		OutputPath:   "cmd/gen/index.go",
	}
	genFile.Generate(Bind{
		Name: "world!",
	})

}

type GenFile struct {
	TemplatePath string
	OutputPath   string
}

func (g *GenFile) IsTemplateExist() bool {
	return fileExist(g.TemplatePath)
}
func (g *GenFile) IsOutputExist() bool {
	return fileExist(g.OutputPath)
}
func (g *GenFile) Generate(val interface{}) {
	if g.IsTemplateExist() {
		if g.IsOutputExist() && !f {
			fmt.Println("file exist:", g.OutputPath)
			return
		}
		tpl, err := template.ParseFiles(g.TemplatePath)
		FailIf(err)

		f, err := os.Create(g.OutputPath)
		defer f.Close()
		FailIf(err)

		//err = tpl.Execute(os.Stdout, val)
		err = tpl.Execute(f, val)
		FailIf(err)
	}
}

type Bind struct {
	Name string
}

func FailIf(err error) {
	if err != nil {
		panic(err)
	}
}

func fileExist(path string) bool {
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		return true
	}
	return false
}

func usage() {
	fmt.Fprintf(os.Stderr, `gen version: 0.0.1
Usage: gen [-f force]

Options:
`)
	flag.PrintDefaults()
}
