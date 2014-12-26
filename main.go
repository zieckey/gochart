package main

import (
	"log"
)

const start = `version: 1.0
http://localhost:8000`

func main() {
	log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime)
	println(start)
	println(ListenAndServe(":8000").Error())
}
