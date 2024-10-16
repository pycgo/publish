// publish/internal/getTag.go

package getTag

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// GetTag returns the latest tag from a repository.
func GetTag(repoName, repoBranch string) (string, error) {
	url := fmt.Sprintf("https://reg.harbor.com/api/v2.0/projects/%s/repositories/%s/artifacts", repoName, repoBranch)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	req.SetBasicAuth("admin", "admin")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var result []Artifact
	err = json.Unmarshal(body, &result)
	if err != nil {
		return "", err
	}

	for _, artifact := range result {
		fmt.Println(artifact)
		if len(artifact.Tags) >= 2 {
			for _, tag := range artifact.Tags {
				if tag.Name != "latest" {
					return tag.Name, nil
				}
			}
		}
	}

	return "", fmt.Errorf("no suitable tag found")
}

type Artifact struct {
	Tags []Tag `json:"tags"`
}

type Tag struct {
	Name     string `json:"name"`
	PushTime string `json:"push_time"`
}
