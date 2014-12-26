package main

import (
	"fmt"
	"github.com/bitly/go-simplejson"
	"github.com/zieckey/goini"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const DataPrefix = "Data|"

type ChartIf interface {
	Parse(ini *goini.INI) (map[string]string, error)
	Template() string
}

type SplineChart int

func (c *SplineChart) Parse(ini *goini.INI) (map[string]string, error) {
	log.Printf("c=%v ini=%v\n", c, ini)

	args := make(map[string]string)
	args["ChartType"] = "spline"
	args["Title"], _ = ini.Get("Title")
	args["SubTitle"], _ = ini.Get("SubTitle")
	args["YAxisText"], _ = ini.Get("YAxisText")
	args["XAxisNumbers"], _ = ini.Get("XAxisNumbers")
	args["ValueSuffix"], _ = ini.Get("ValueSuffix")

	datas := make([]interface{}, 0)

	kv, _ := ini.GetKvmap(goini.DefaultSection)
	for k, v := range kv {
		if !strings.HasPrefix(k, DataPrefix) {
			continue
		}

		dd := strings.Split(v, ", ")
		jd := make([]interface{}, 0)
		for _, d := range dd {
			
			val, err := strconv.ParseFloat(d, 64)
			if err == nil {
				jd = append(jd, val)
			}
			log.Printf("ParseFloat(%v) v=%v err=%v\n", d, val, err)
		}
		json := simplejson.New()
		json.Set("name", k[len(DataPrefix):])
		json.Set("data", jd)
		datas = append(datas, json)
	}

	json := simplejson.New()
	json.Set("DataArray", datas)
	b, _ := json.Get("DataArray").Encode()
	args["DataArray"] = string(b)

	fmt.Printf("DataArray:\n%v\n", string(b))
	fmt.Printf("=========================================================>>Args:\n%v\n", args)
	return args, nil
}

func (c *SplineChart) Template() string {
	return TemplateSplineHtml
}

type TemplateArgs struct {
	args map[string]string
	tmpl string
}

func Parse(file string) (ta TemplateArgs, err error) {
	ini := goini.New()
	err = ini.ParseFile(file)
	if err != nil {
		return ta, err
	}

	t, _ := ini.Get("ChartType")
	log.Printf("ini.Get ChartType=%v\n", t)
	if f, ok := ChartHandlers[t]; ok {
		log.Printf("f=%v ok=%v\n", f, ok)
		ta.args, err = f.Parse(ini)
		ta.tmpl = f.Template()
	}

	return ta, err
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
