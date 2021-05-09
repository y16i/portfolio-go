package tool

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strings"
)

type WpConfig struct {
	DbName     string
	DbUser     string
	DbPassword string
	DbHost     string
}

func ReadWpConfig(wpConfigPath string) *WpConfig {
	file, openErr := os.Open(wpConfigPath)
	if openErr != nil {
		log.Fatal(openErr)
	}
	defer file.Close()

	wpConfig := new(WpConfig)
	scanner := bufio.NewScanner(file) // char limit 65536 per line
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "'DB_NAME'") {
			wpConfig.DbName = getValue("DB_NAME", scanner.Text())
		} else if strings.Contains(scanner.Text(), "'DB_USER'") {
			wpConfig.DbUser = getValue("DB_USER", scanner.Text())
		} else if strings.Contains(scanner.Text(), "'DB_PASSWORD'") {
			wpConfig.DbPassword = getValue("DB_PASSWORD", scanner.Text())
		} else if strings.Contains(scanner.Text(), "'DB_HOST'") {
			wpConfig.DbHost = getValue("DB_HOST", scanner.Text())
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return wpConfig
}

func getValue(key string, line string) string {
	result := ""
	matcher := regexp.MustCompile(key + `.+'`)
	value := matcher.Find([]byte(line))
	trimmed := strings.ReplaceAll(string(value), " ", "")
	trimmed = strings.ReplaceAll(trimmed, "'", "")
	keyValue := strings.Split(trimmed, ",")

	if len(keyValue) > 0 {
		result = keyValue[1]
	}
	return result
}
