package core

import (
	"net/http"
	"bufio"
	"strings"
	"fmt"
	"os"
)

func Load(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "CATTLE_REGISTRATION_ACCESS_KEY") {
			str := strings.Split(line, "=")[1]
			value := fmt.Sprintf("[Environment]::SetEnvironmentVariable(\"CATTLE_REGISTRATION_ACCESS_KEY\", \"%s\", \"Machine\")", str[1:len(str)-1])
			fmt.Println(value)
		} else if strings.Contains(line, "CATTLE_REGISTRATION_SECRET_KEY") {
			str := strings.Split(line, "=")[1]
			value := fmt.Sprintf("[Environment]::SetEnvironmentVariable(\"CATTLE_REGISTRATION_SECRET_KEY\", \"%s\", \"Machine\")", str[1:len(str)-1])
			fmt.Println(value)
		} else if strings.Contains(line, "CATTLE_URL") {
			str := strings.Split(line, "=")[1]
			value := fmt.Sprintf("[Environment]::SetEnvironmentVariable(\"CATTLE_URL\", \"%s\", \"Machine\")", str[1:len(str)-1])
			fmt.Println(value)
		} else if strings.Contains(line, "DETECTED_CATTLE_AGENT_IP") {
			if os.Getenv("CATTLE_AGENT_IP") != "" {
				str := strings.Split(line, "=")[1]
				value := fmt.Sprintf("[Environment]::SetEnvironmentVariable(\"CATTLE_AGENT_IP\", \"%s\", \"Machine\")", str[1:len(str)-1])
				fmt.Println(value)
			}
		}
	}
	return nil
}
