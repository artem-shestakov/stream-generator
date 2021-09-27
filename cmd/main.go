package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/artem-shestakov/stream-generator/internal/models"
	"github.com/xuri/excelize/v2"
	"gopkg.in/yaml.v2"
)

func readXlsx(path string) *excelize.File {
	// Read *.xlsx file
	xlsxFile, err := excelize.OpenFile(path)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return xlsxFile
}

func createStreams(file *excelize.File) *models.Streams {
	var stream models.Stream
	var streams models.Streams
	start_name := 10100
	// Get all rows
	rows, err := file.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	// Generate streams from rows
	for _, row := range rows {
		// get ports from row
		ports := strings.Split(row[1], ",")
		// Generate stream for each port
		for _, port := range ports {
			var servers []string
			stream.Name = strconv.Itoa(start_name)
			for _, server := range strings.Split(row[0], ",") {
				servers = append(servers, server+":"+port)
			}
			stream.Servers = servers
			streams.Streams = append(streams.Streams, stream)
			start_name += 1
		}
	}
	return &streams
}

func writeFile(streams *models.Streams) error {
	// Prepare data
	data, err := yaml.Marshal(streams.Streams)
	if err != nil {
		return err
	}
	// Write to file
	err = ioutil.WriteFile("streams.yml", data, 0)
	if err != nil {
		return err
	}
	// Change file permission
	err = os.Chmod("streams.yml", 0644)
	if err != nil {
		return err
	}
	return nil
}

func readYaml(file string) []models.Playbook {
	var playbooks []models.Playbook
	yamlFile, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(yamlFile, &playbooks)
	if err != nil {
		log.Fatal(err)
	}

	return playbooks
}

func test(t interface{}) {
	switch reflect.TypeOf(t).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(t)
		for i := 0; i < s.Len(); i++ {
			fmt.Println(s.Index(i))
		}
	}
}

func main() {
	// f := readXlsx("streams.xlsx")
	// streams := createStreams(f)
	// err := writeFile(streams)
	// if err != nil {
	// 	fmt.Printf("Error to write file %s", err)
	// }
	playbooks := readYaml("./example/example.yml")
	test(playbooks[0].Roles[0].Vars["tcp_udp_nlb"])
}
