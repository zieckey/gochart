package main

import (
	"fmt"
	"github.com/zieckey/goini"
	"log"
	"strings"
)

type PieChart int

func (c *PieChart) Parse(ini *goini.INI) (map[string]string, error) {
	log.Printf("c=%v ini=%v\n", c, ini)

	args := make(map[string]string)
	args["ChartType"], _ = ini.Get("ChartType")
	args["Title"], _ = ini.Get("Title")
	args["SubTitle"], _ = ini.Get("SubTitle")
	args["SeriesName"], _ = ini.Get("SeriesName")
	
	DataArray := "[\n"
	
	kv, _ := ini.GetKvmap(goini.DefaultSection)
	for k, v := range kv {
		if !strings.HasPrefix(k, DataPrefix) {
			continue
		}
		
		key := k[len(DataPrefix):]
		DataArray = DataArray + "['" + key + "' , " + v + "],\n"
	}

	DataArray = DataArray + "]"
		
	args["DataArray"] = DataArray

	fmt.Printf("=========================================================>>Args:\n%v\n", args)
	return args, nil
}

func (c *PieChart) Template() string {
	return TemplatePieHtml
}

