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

	var resultString string
	for _, records := range dmi.Data {
		for _, record := range records {
			for k, v := range record {
				resultString += fmt.Sprintf("dmidecode_%s{%s=\"%s\"} 1 \n", strings.Replace(strings.ToLower(record["DMIName"]), " ", "_", 5), strings.Replace(strings.ToLower(k), " ", "_", 5), v)
			}
		}
	}
	fmt.Println(resultString)
}
