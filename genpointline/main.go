package main

import (
	"fmt"
	"github.com/bitly/go-simplejson"
)

/*
[
   {
       name: 'Tokyo',
       data: [7.0, 6.9, 9.5, 14.5, 18.2, 21.5, 25.2, 26.5, 23.3, 18.3, 13.9, 9.6]
   },
   {
       name: 'New York',
       data: [-0.2, 0.8, 5.7, 11.3, 17.0, 22.0, 24.8, 24.1, 20.1, 14.1, 8.6, 2.5]
   },
   {
       name: 'Berlin',
       data: [-0.9, 0.6, 3.5, 8.4, 13.5, 17.0, 18.6, 17.9, 14.3, 9.0, 3.9, 1.0]
   },
   {
       name: 'London',
       data: [3.9, 4.2, 5.7, 8.5, 11.9, 15.2, 17.0, 16.6, 14.2, 10.3, 6.6, 4.8]
   }
]
*/

func main() {
	js := simplejson.New()
	js.Set("ChartType", "spline")
	js.Set("Title", "Monthly Average Temperature")
	js.Set("SubTitle", "Source: WorldClimate.com")
	js.Set("YAxisText", "Temperature (°C)")
	js.Set("XAxisNumbers", []interface{}{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12"})
	js.Set("ValueSuffix", "°C")
	
	j1 := simplejson.New()
	j1.Set("name", "Tokyo")
	j1.Set("data", []interface{}{7.0, 6.9, 9.5, 14.5, 18.2, 21.5, 25.2, 26.5, 23.3, 18.3, 13.9, 9.6})
	
	j2 := simplejson.New()
	j2.Set("name", "New York")
	j2.Set("data", []interface{}{-0.2, 0.8, 5.7, 11.3, 17.0, 22.0, 24.8, 24.1, 20.1, 14.1, 8.6, 2.5})
	
	j3 := simplejson.New()
	j3.Set("name", "Berlin")
	j3.Set("data", []interface{}{-0.9, 0.6, 3.5, 8.4, 13.5, 17.0, 18.6, 17.9, 14.3, 9.0, 3.9, 1.0})
	
	
	j4 := simplejson.New()
	j4.Set("name", "London")
	j4.Set("data", []interface{}{3.9, 4.2, 5.7, 8.5, 11.9, 15.2, 17.0, 16.6, 14.2, 10.3, 6.6, 4.8})
	
	
	js.Set("DataArray", []interface{}{j1, j2, j3, j4})
	
	//js.Set("DataArray", []interface{}{"1", "3.2", "5.1", 6.2, false})
	b, _ := js.EncodePretty()
	fmt.Printf("%v\n", string(b))
	fmt.Printf("============================================================\n")
	
	b, _ = js.Get("DataArray").EncodePretty()
	fmt.Printf("%v\n", string(b))
	
	fmt.Printf("============================================================\n")
	b, _ = js.Get("SubTitle").EncodePretty()
	fmt.Printf("%v\n", string(b))
	
}
