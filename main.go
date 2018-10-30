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
	x := make(map[string][]string)

	for _, records := range dmi.Data {
		for _, record := range records {
			//fmt.Println(record)
			for k, v := range record {
				x[replacer.Replace(strings.ToLower(record["DMIName"]))] = append(x[replacer.Replace(strings.ToLower(record["DMIName"]))], fmt.Sprintf("%s=\"%s\"", replacer.Replace(strings.ToLower(k)), strings.Replace(v, "\"", "\\\"", -1)))
			}
		}
	}
	for k, v := range x {
		fmt.Printf("dmidecode_%s{%s} 1\n", k, strings.Join(v, ","))
	}
}
