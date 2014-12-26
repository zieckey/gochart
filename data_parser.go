package main

import (

	"github.com/zieckey/goini"
	"log"
	"os"
	"path/filepath"

)

const DataPrefix = "Data|"

type ChartIf interface {
	Parse(ini *goini.INI) (map[string]string, error)
	Template() string
}

type TemplateArgs struct {
	args map[string]string
	tmpl string
}

func Parse(file string) (tt TemplateArgs, err error) {
	ini := goini.New()
	err = ini.ParseFile(file)
	if err != nil {
		return tt, err
	}

	t, _ := ini.Get("ChartType")
	log.Printf("ini.Get ChartType=%v\n", t)
	if f, ok := ChartHandlers[t]; ok {
		log.Printf("f=%v ok=%v\n", f, ok)
		tt.args, err = f.Parse(ini)
		tt.tmpl = f.Template()
	}

	return tt, err
}

func LookupChartFiles(dir string) ([]string, error) {
	var files []string = make([]string, 0, 5)

	err := filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
		//log.Printf("lookup dir [%v] walk to path [%v] [%v]\n", dir, path, f.Name())
		if f == nil {
			return err
		}

		if f.IsDir() {
			return nil
		}

		if ok, err := filepath.Match("*.chart", f.Name()); err != nil {
			//log.Printf("NO a *.chart file [%v]\n", f.Name())
			return err
		} else if ok {
			log.Printf("Find a *.chart file [%v]\n", path)
			files = append(files, path)
		}
		return nil
	})

	return files, err
}
