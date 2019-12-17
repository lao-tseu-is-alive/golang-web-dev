package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseFiles("one.gmao")
	if err != nil {
		log.Fatalln(err)
	}
	// now the template "one.gmao" is executed and the content goes to os.Stdout
	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}
	//the same tpl is used to parse others templates
	tpl, err = tpl.ParseFiles("two.gmao", "vespa.gmao")
	if err != nil {
		log.Fatalln(err)
	}
	// now a couple of templates are ready to be used in the tpl, let's run the "vespa.gmao"
	err = tpl.ExecuteTemplate(os.Stdout, "vespa.gmao", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "two.gmao", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "one.gmao", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
