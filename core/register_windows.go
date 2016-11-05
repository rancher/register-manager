package core

import "fmt"

func print(accessKey, secretKey string) {
	aKey := fmt.Sprintf("[Environment]::SetEnvironmentVariable(\"CATTLE_ACCESS_KEY\", \"%s\", \"Machine\")", accessKey)
	sKey := fmt.Sprintf("[Environment]::SetEnvironmentVariable(\"CATTLE_SECRET_KEY\", \"%s\", \"Machine\")", secretKey)
	fmt.Println(aKey)
	fmt.Println(sKey)
}