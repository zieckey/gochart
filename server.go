package main

import (
	//"fmt"
	"net/http"
	"text/template"
)

var (
	ChartHandlers = make(map[string]ChartIf)
	ChartFiles    []string
	Index         int
)

func handler(w http.ResponseWriter, r *http.Request) {
	tt, err := Parse(ChartFiles[Index])
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	Index++
	Index = Index % len(ChartFiles)

	if t, err := template.New("foo").Parse(tt.tmpl); err != nil {
		w.Write([]byte(err.Error()))
	} else {
		if err = t.ExecuteTemplate(w, "T", tt.args); err != nil {
			w.Write([]byte(err.Error()))
		}
	}
}

func ListenAndServe(addr string) error {
	http.HandleFunc("/", handler)
	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {})

	var err error
	ChartFiles, err = LookupChartFiles(".")
	if err != nil {
		return err
	}

	// Register chart handlders
	ChartHandlers["spline"] = new(SplineChart)
	ChartHandlers["column"] = new(SplineChart)
	ChartHandlers["area"] = new(SplineChart)
	ChartHandlers["bar"] = new(SplineChart)
	ChartHandlers["line"] = new(SplineChart)
	ChartHandlers["pie"] = new(PieChart)

	return http.ListenAndServe(addr, nil)
}
