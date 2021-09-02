package main

import (
	"bytes"
	"fmt"
	"log"
	"syscall/js"
	"text/template"

	"github.com/Masterminds/sprig"
	"gopkg.in/yaml.v2"
)

type Values map[string]interface{}

type NestedValues struct {
	Values Values
}

func generateYaml(templateYaml string, valuesYaml string) (string, error) {
	valuesData := Values{}
	if err := yaml.Unmarshal([]byte(valuesYaml), &valuesData); err != nil {
		return "", err
	}

	asd := NestedValues{valuesData}

	var output bytes.Buffer

	t := template.Must(template.New("template").Funcs(sprig.TxtFuncMap()).Parse(templateYaml))
	if err := t.Execute(&output, asd); err != nil {
		log.Println("executing template:", err)
	}

	return output.String(), nil
}

func generateWrapper() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) != 2 {
			return "Invalid no of arguments passed"
		}
		templateYaml := args[0].String()
		valuesYaml := args[1].String()
		generatedYaml, err := generateYaml(templateYaml, valuesYaml)
		if err != nil {
			fmt.Printf("error from generate: %s\n", err)
			return err.Error()
		}
		return generatedYaml
	})
}

func main() {
	fmt.Println("Go Web Assembly")
	js.Global().Set("generateYaml", generateWrapper())
	<-make(chan bool)
}
