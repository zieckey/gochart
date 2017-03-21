package main

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/zieckey/goini"
)

const DataPrefix = "Data|"
const emptyRunes = " \r\t\v"

type ChartIf interface {
	Parse(ini *goini.INI, file string) (map[string]string, error)
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
	if f, ok := ChartHandlers[t]; ok {
		tt.args, err = f.Parse(ini, file)
		tt.tmpl = f.Template()
	}

	return tt, err
}

func LookupChartFiles(dir string) ([]string, error) {
	var files []string = make([]string, 0, 5)

	err := filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}

		if f.IsDir() {
			return nil
		}

		if ok, err := filepath.Match("*.chart", f.Name()); err != nil {
			return err
		} else if ok {
			log.Printf("Add chart file %v", path)
			files = append(files, path)
		}
		return nil
	})

	if len(files) == 0 {
		return files, errors.New("Not found any *.chart files")
	}

	return files, err
}

/**
 * 读取配置文件获取有序map
 *
 * @return order map
 * @param {[type]} ini string [description]
 */
func LoadConfGetOrderMap(configFile string) ([]string, map[string]string, error) {
	stream, err := ioutil.ReadFile(configFile)
	if err != nil {
		return nil, nil, errors.New("cannot load config file")
	}
	content := string(stream)

	confMap := make(map[string]string)
	mapkeys := make([]string, 0)

	lines := strings.Split(content, "\n")
	for _, line := range lines {
		line = strings.Trim(line, emptyRunes)
		//过滤注释 //过滤非DataPrefix
		if line == "" || line[0] == '#' || !strings.HasPrefix(line, DataPrefix) {
			continue
		}
		//过滤
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			key := strings.Trim(parts[0], emptyRunes)
			value := strings.Trim(parts[1], emptyRunes)
			mapkeys = append(mapkeys, key)
			confMap[key] = value
		}
	}
	return mapkeys, confMap, nil
}
