package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
)

var m map[string]interface{}

func main() {

	f, err := os.Open("conf.json")
	if err != nil {
		log.Fatalln(err)
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatalln(err)
	}

	if err := json.Unmarshal(b, &m); err != nil {
		log.Fatalln(err)
	}

	//fmt.Println(m["models"])

	bff := bytes.Buffer{}

	tmp, err := template.ParseFiles("f_main.tmpl")
	if err != nil {
		log.Fatalln(err)
	}
	err = tmp.ExecuteTemplate(&bff, "main", m)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(bff.String())

}
