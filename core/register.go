package core

import (
	"time"
	"github.com/rancher/go-rancher/v2"
	"os"
)

const (
	URL = "CATTLE_URL"
	AccessKey = "CATTLE_REGISTRATION_ACCESS_KEY"
	SecretKey = "CATTLE_REGISTRATION_SECRET_KEY"
)

func Register(token string) error {
	url := os.Getenv(URL)
	accessKey := os.Getenv(AccessKey)
	secretKey := os.Getenv(SecretKey)

	apiClient, err := client.NewRancherClient(&client.ClientOpts{
		Timeout: time.Second * 30,
		Url: url,
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
		list, err := apiClient.Register.List(&client.ListOpts{
			Filters: map[string]interface{}{
				"key": token,
			},
		})
		if err != nil {
			return err
		}
		print(list.Data[0].AccessKey, list.Data[0].SecretKey)
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
