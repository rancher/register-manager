package core

import (
	"net/http"
	"io/ioutil"
	"strings"
	"fmt"
	"time"
	"github.com/rancher/go-rancher/v2"
	"os"
)

const (
	Schema = "X-API-Schemas"
)

func ResolveUrl(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	if resp.StatusCode == 200 {
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		if strings.HasPrefix(string(data), "#!/bin/sh") {
			fmt.Println(url)
			return nil
		}
	}
	resp, err = http.Get(url)
	if err != nil {
		return err
	}
	if val, ok := resp.Header[Schema]; ok {
		url = val[0]
	}

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
	types := apiClient.GetTypes()
	if !inList("POST", types["registrationToken"].CollectionMethods) {
		projects, err := apiClient.Project.List(&client.ListOpts{
			Filters: map[string]interface{}{
				"uuid": "adminProject",
			},
		})
		if err != nil {
			return err
		}
		if len(projects.Data) == 0 {
			fmt.Println("Failed to find admin resource group")
		}
		apiClient, err = client.NewRancherClient(&client.ClientOpts{
			Timeout: time.Second * 30,
			Url: projects.Data[0].Links["schemas"],
			AccessKey: accessKey,
			SecretKey: secretKey,
		})
		if err != nil {
			return err
		}
	}
	tokens, err := apiClient.RegistrationToken.List(&client.ListOpts{
		Filters: map[string]interface{}{
			"state": "active",
		},
	})
	if err != nil {
		return err
	}
	token := &client.RegistrationToken{}
	if len(tokens.Data) == 0 {
		token, err = apiClient.RegistrationToken.Create(nil)
		if err != nil {
			return err
		}
	} else {
		token = &tokens.Data[0]
	}
	fmt.Println(token.RegistrationUrl)
	return nil
}

func inList(target string, list []string) bool {
	for _, item := range list {
		if item == target {
			return true
		}
	}
	return false
}