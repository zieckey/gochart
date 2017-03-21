package main

import (
	"strconv"
	"strings"

	"github.com/bitly/go-simplejson"
	"github.com/zieckey/goini"
)

type SplineChart int

func (c *SplineChart) Parse(ini *goini.INI, file string) (map[string]string, error) {
	args := make(map[string]string)
	args["ChartType"], _ = ini.Get("ChartType")
	args["Title"], _ = ini.Get("Title")
	args["SubTitle"], _ = ini.Get("SubTitle")
	args["YAxisText"], _ = ini.Get("YAxisText")
	args["XAxisNumbers"], _ = ini.Get("XAxisNumbers")
	args["ValueSuffix"], _ = ini.Get("ValueSuffix")
	args["Height"], _ = ini.Get("Height")

	datas := make([]interface{}, 0)

	mapkeys, kvmap, err := LoadConfGetOrderMap(file)
	if err != nil {
		return nil, err
	}

	for _, key := range mapkeys {
		if !strings.HasPrefix(key, DataPrefix) {
			continue
		}

		dd := strings.Split(kvmap[key], ",")
		jd := make([]interface{}, 0)
		for _, d := range dd {
			d = strings.TrimSpace(d)
			val, err := strconv.ParseFloat(d, 64)
			if err == nil {
				jd = append(jd, val)
			}
		}
		json := simplejson.New()
		json.Set("name", key[len(DataPrefix):])
		json.Set("data", jd)
		datas = append(datas, json)
	}

	json := simplejson.New()
	json.Set("DataArray", datas)
	b, _ := json.Get("DataArray").Encode()
	args["DataArray"] = string(b)
	//log.Println(args)
	return args, nil
}

func (c *SplineChart) Template() string {
	return TemplateSplineHtml
}
