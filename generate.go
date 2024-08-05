package main

import (
	// "flag"
	"os"
	// "strings"
	"text/template"
	"time"
)

type TemplateValue struct {
	AfterJan2022 bool
	AfterMay2021 bool
	AfterAug2019 bool
	AfterAug2017 bool
	AfterAug2012 bool
	AfterJul2011 bool
	AfterAug2008 bool
}

func getDatetime(datetimeString string) time.Time {
	const shortForm = "2006-Jan-02"
	dt, _ := time.Parse(shortForm, datetimeString)
	return dt
}

func main() {
	// var titlePtr = flag.String("title", "{title}", "Document title")
	// var versionPtr = flag.String("version", "{default_version}", "Document version")
	// var templatePtr = flag.String("template", "{default_template}", "Path to template file")
	// var loremIpsumPtr = flag.Bool("loremipsum", false, "Include lorem ipsum text")
	// var pangramPtr = flag.Bool("pangram", false, "Include pangram")
	// var outputPtr = flag.String("output", "", "Output filename")
	// flag.Parse()

	now := time.Now()
	tenYearsAgo := now.AddDate(-10, 0, 0)
	tValue := TemplateValue{
		AfterJan2022: getDatetime("2022-Jan-01").After(tenYearsAgo),
		AfterMay2021: getDatetime("2021-May-01").After(tenYearsAgo),
		AfterAug2019: getDatetime("2019-Aug-01").After(tenYearsAgo),
	  AfterAug2017: getDatetime("2017-Aug-01").After(tenYearsAgo),
	  AfterAug2012: getDatetime("2012-Aug-01").After(tenYearsAgo),
	  AfterJul2011: getDatetime("2011-Jul-01").After(tenYearsAgo),
	  AfterAug2008: getDatetime("2008-Aug-01").After(tenYearsAgo),
	}

	var tFile = "master.md"
	// var tFileArray = strings.Split(tFile, ".")
	var outputFile = "recent.md" // tFileArray[0] + "-" + now.Format("2006-01-02") + "." + tFileArray[1]
	// if len(*outputPtr) > 0 {
	// 	outputFile = *outputPtr
	// }
	output, err := os.Create(outputFile)

	textTmpl, err := template.New(tFile).ParseFiles(tFile)
	if err != nil {
		panic(err)
	}
	err = textTmpl.Execute(output, tValue)
	if err != nil {
		panic(err)
	}
}