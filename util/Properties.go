package util

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type AppConfigProperties map[string]string

func readPropertiesFile() (AppConfigProperties, error) {
	config := AppConfigProperties{}

	file, err := os.Open("application.properties")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if equal := strings.Index(line, "="); equal >= 0 {
			if key := strings.TrimSpace(line[:equal]); len(key) > 0 {
				value := ""
				if len(line) > equal {
					value = strings.TrimSpace(line[equal+1:])
				}
				config[key] = value
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return config, nil
}

func GetStringProperties(key string) string {

	props, _ := readPropertiesFile()

	return props[key]
}

func GetIntProperties(key string) int {

	props, _ := readPropertiesFile()

	ret, _ := strconv.Atoi(props[key])

	return ret
}
