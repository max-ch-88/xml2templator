package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

func getxml(filename string) ZabbixExport {
	var zabbixExport ZabbixExport

	xmlFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}

	defer xmlFile.Close()

	byteValue, _ := ioutil.ReadAll(xmlFile)

	err = xml.Unmarshal(byteValue, &zabbixExport)
	if err != nil {
		fmt.Println(err)
	}

	return zabbixExport
}

func main() {

	fileName := "Template_App_MS_SQL_-_Windows_Server_Failover_Cluster.xml"
	xmlIn := getxml(fileName)

	xmlOut, err := xml.MarshalIndent(xmlIn, "", "    ")
	if err != nil {
		fmt.Printf("Error marshaling XML file: %s\n", err)
	}

	// re := regexp.MustCompile(filepath.Ext(fileName) + `$`)
	// fileName = re.ReplaceAllLiteralString(fileName, ".xml")

	err = ioutil.WriteFile("fileName.xml", xmlOut, os.ModePerm)
	if err != nil {
		fmt.Printf("Error writing XML file: %s\n", err)
		return
	}

	// var ymlIn Value_maps

	// for i := 0; i < len(xmlIn.Value_maps.Value_map); i++ {
	// 	ymlIn.Value_map[i] = xmlIn.Value_maps.Value_map[i].Name
	// 	var b []Mapping
	// 	for j := 0; j < len(xmlIn.Value_maps.Value_map[i].Mappings.Mapping); j++ {
	// 		// fmt.Println(xmlIn.Value_maps.Value_map[i].Mappings.Mapping[j])
	// 		b = append(b, Mapping{xmlIn.Value_maps.Value_map[i].Mappings.Mapping[j].Value,
	// 			xmlIn.Value_maps.Value_map[i].Mappings.Mapping[j].Newvalue})
	// 	}

	// 	fmt.Println(a, b)
	// 	// ymlIn.Value_maps = append(ymlIn.Value_maps, struct{a,b})
	// }

	ymlOut, err := yaml.Marshal(xmlIn)
	if err != nil {
		fmt.Printf("Error marshaling YAML file: %s\n", err)
	}

	err = ioutil.WriteFile("fileName.yml", ymlOut, os.ModePerm)
	if err != nil {
		fmt.Printf("Error writing YAML file: %s\n", err)
		return
	}

}
