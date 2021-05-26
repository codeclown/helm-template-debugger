package main

import (
	"bufio"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"

	"github.com/Masterminds/sprig"
	"gopkg.in/yaml.v2"
)

type Values map[string]interface{}

type NestedValues struct {
	Values Values
}

func GenerateHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("error reading body: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}

	valuesString := ""
	templateString := ""
	current := ""

	scanner := bufio.NewScanner(strings.NewReader(string(body)))
	for scanner.Scan() {
		line := scanner.Text()
		if line == "### VALUES ###" {
			current = "values"
		} else if line == "### TEMPLATE ###" {
			current = "template"
		} else if current == "values" {
			valuesString += line + "\n"
		} else if current == "template" {
			templateString += line + "\n"
		}
	}
	if err := scanner.Err(); err != nil {
		log.Printf("error processing body: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}

	valuesData := Values{}
	err = yaml.Unmarshal([]byte(valuesString), &valuesData)
	if err := scanner.Err(); err != nil {
		log.Printf("error parsing values: %v", err)
		http.Error(w, "error parsing values", http.StatusBadRequest)
		return
	}

	asd := NestedValues{valuesData}

	t := template.Must(template.New("template").Funcs(sprig.TxtFuncMap()).Parse(templateString))
	if err := t.Execute(w, asd); err != nil {
		log.Println("executing template:", err)
	}

	// fmt.Fprintf(w, "--- Values: %s\n\n--- Template: %s", valuesString, templateString)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/generate", GenerateHandler)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
