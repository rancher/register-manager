package core

import (
	"crypto/rand"
	"fmt"
	"github.com/pkg/errors"
	"github.com/rancher/go-rancher/v2"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

const (
	URL       = "CATTLE_URL"
	AccessKey = "CATTLE_REGISTRATION_ACCESS_KEY"
	SecretKey = "CATTLE_REGISTRATION_SECRET_KEY"
	tokenFile = "c:/registration_token"
)

func Register(token string, args ...string) error {
	url := os.Getenv(URL)
	accessKey := os.Getenv(AccessKey)
	secretKey := os.Getenv(SecretKey)
	if len(args) != 0 {
		url = args[0]
		accessKey = args[1]
		secretKey = args[2]
	}

	apiClient, err := client.NewRancherClient(&client.ClientOpts{
		Timeout:   time.Second * 30,
		Url:       url,
		AccessKey: accessKey,
		SecretKey: secretKey,
	})
	if err != nil {
		return err
	}
	resp, err := apiClient.Register.List(&client.ListOpts{
		Filters: map[string]interface{}{
			"key": token,
		},
	})
	if err != nil {
		return err
	}
	if len(resp.Data) == 0 {
		_, err := apiClient.Register.Create(&client.Register{
			Key: token,
		})
		if err != nil {
			return err
		}
		i := 0
		for {
			if i == 10 {
				return errors.New("Failed to genarate access key")
			}
			list, err := apiClient.Register.List(&client.ListOpts{
				Filters: map[string]interface{}{
					"key": token,
				},
			})
			if err != nil {
				return err
			}
			if len(list.Data) == 0 || list.Data[0].AccessKey == "" {
				time.Sleep(time.Second)
				i++
				continue
			}
			print(list.Data[0].AccessKey, list.Data[0].SecretKey)
			break
		}

	} else {
		list, err := apiClient.Register.List(&client.ListOpts{
			Filters: map[string]interface{}{
				"key": token,
			},
		})
		if err != nil {
			return err
		}
		print(list.Data[0].AccessKey, list.Data[0].SecretKey)
	}
	return nil
}

func RegisterWindows(args string) error {
	if _, err := os.Stat(tokenFile); err == nil {
		file, _ := os.Open(tokenFile)
		defer file.Close()
		data, err := ioutil.ReadAll(file)
		if err != nil {
			return err
		}
		token := string(data)
		parts := strings.Split(args, ",")
		if err := Register(token, parts...); err != nil {
			return err
		}
		return nil
	}
	file, err := os.Create(tokenFile)
	if err != nil {
		return err
	}
	defer file.Close()
	b := make([]byte, 64)
	rand.Read(b)
	token := fmt.Sprintf("%x", b)
	_, err = file.WriteString(token)
	if err != nil {
		return err
	}
	parts := strings.Split(args, ",")
	if err := Register(token, parts...); err != nil {
		return err
	}
	return nil
}
