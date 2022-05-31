package main

import (
	"fmt"
	"strings"
	"net/http"
	"io/ioutil"
)

func main() {

	url := "https://snyk.io/api/v1/org/$org/projects?="

	payload := strings.NewReader("{\n    \"filters\": {\n\t\t\t  \"name\": \"grafana/grafana\",\n        \"origin\": \"github-enterprise\"\n    }\n}\n")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "token $token ")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
 

	fmt.Println(string(body))

}
