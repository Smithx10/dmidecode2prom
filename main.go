package main

import (
	"fmt"
	"strings"

	dmidecode "github.com/dselans/dmidecode"
)

func main() {
	// Create new dmidecode
	dmi := dmidecode.New()
	if err := dmi.Run(); err != nil {
		fmt.Printf("Unable to get dmidecode information. Error: %v\n", err)
	}

	replacer := strings.NewReplacer("-", "_", " ", "_")
	var returnString string
	var strSlice []string

	for _, records := range dmi.Data {
		for _, record := range records {
			for k, v := range record {
				strSlice = append(strSlice, fmt.Sprintf("%s=\"%s\"", replacer.Replace(strings.ToLower(k)), v))
			}
			returnString += fmt.Sprintf("dmidecode_%s{%s} 1\n", replacer.Replace(strings.ToLower(record["DMIName"])), fmt.Sprintf(strings.Join(strSlice, ",")))
		}
	}
	fmt.Println(returnString)
}
