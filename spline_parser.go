package main

import (
	"fmt"
	"github.com/bitly/go-simplejson"
	"github.com/zieckey/goini"
	"log"
	"strconv"
	"strings"
)

type SplineChart int
type ColumnChart struct {
	SplineChart
}

func (c *SplineChart) Parse(ini *goini.INI) (map[string]string, error) {
	log.Printf("c=%v ini=%v\n", c, ini)

	args := make(map[string]string)
	args["ChartType"], _ = ini.Get("ChartType")
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
