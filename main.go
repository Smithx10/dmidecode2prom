package main

import (
	"fmt"
	"strings"

	dmidecode "github.com/dselans/dmidecode"
)

func main() {

	// Define Result String
	var resultString string
	// Create new dmidecode
	dmi := dmidecode.New()
	if err := dmi.Run(); err != nil {
		fmt.Printf("Unable to get dmidecode information. Error: %v\n", err)
	}

	replacer := strings.NewReplacer(":", "_", "-", "_", " ", "_", "\"", "\\\"")
	x := make(map[string][]string)

	// Iterate on Top of all the DMI Data
	for _, records := range dmi.Data {
		// Iterate on Each Record Type
		for _, record := range records {
			// Expand And Model Data Most Appropriate for each Type
			//fmt.Println(record["DMIName"])

			// Clean Empty Values
			for k, v := range record {
				if v == " " {
					delete(record, k)
				}
			}

			// Memory Device
			if record["DMIName"] == "Memory Device" {
				// Set KeyName
				keyName := replacer.Replace(strings.ToLower(record["DMIName"]))
				// Create Temporary Array for Storing ne KVs
				var memdevarray []string
				// Iterate Over them and Populate array
				for k, v := range record {
					memdevarray = append(memdevarray, fmt.Sprintf("%s=\"%s\"", replacer.Replace(strings.ToLower(k)), strings.Replace(v, "\"", "\\\"", -1)))
				}
				// Append string for return
				resultString += fmt.Sprintf("dmidecode_%s{%s} 1\n", keyName, strings.Join(memdevarray, ","))

			}

			// Bios Information
			//if record["DMIName"] == "BIOS Information" {
			//keyName := replacer.Replace(strings.ToLower(record["DMIName"]))
			//// Get Characteristics Array
			//formatedCharacteristics := strings.ToLower(record["Characteristics"])
			//characteristics := strings.Split(formatedCharacteristics, "\t\t")
			//// Flatten Chars
			//for _, v := range characteristics {
			//record["characteristics"+"_"+fmt.Sprintf(v)] = "true"

			//}
			//delete(record, "Characteristics")

			//for k, v := range record {
			//x[keyName] = append(x[keyName], fmt.Sprintf("%s=\"%s\"", replacer.Replace(strings.ToLower(k)), v))
			//}

			//}
			// Process the rest without special formatting
			if record["DMIName"] != "Memory Device" {
				for k, v := range record {

					x[replacer.Replace(strings.ToLower(record["DMIName"]))] = append(x[replacer.Replace(strings.ToLower(record["DMIName"]))], fmt.Sprintf("%s=\"%s\"", replacer.Replace(strings.ToLower(k)), replacer.Replace(v)))
				}
			}
		}
	}
	for k, v := range x {
		resultString += fmt.Sprintf("dmidecode_%s{%s} 1\n", k, strings.Join(v, ","))
	}
	fmt.Println(resultString)
}
