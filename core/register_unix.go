//+build !windows

package core

import "fmt"

func print(accessKey, secretKey string) {
	aKey := fmt.Sprintf("export CATTLE_ACCESS_KEY=%s", accessKey)
	sKey := fmt.Sprintf("export CATTLE_SECRET_KEY=%s", secretKey)
	fmt.Println(aKey)
	fmt.Println(sKey)
}
