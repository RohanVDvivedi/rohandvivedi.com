package githubsync

import (
	"net/http"
	"strconv"
	"errors"
	"io/ioutil"
)

func getGithubRawFileDownloadLink(userName string, projectName string, fileName string) string {
	return "https://raw.githubusercontent.com/" + userName + "/" + projectName + "/master/" + fileName
}

func GetGithubFile(userName string, projectName string, fileName string) (string, error) {
	GetApiRequestPath := getGithubRawFileDownloadLink(userName, projectName, fileName)

	resp, err := http.Get(GetApiRequestPath)
	defer resp.Body.Close()
	if(err != nil) {
		return "", err
	}

	if(resp.StatusCode == 200) {
		data, errRead := ioutil.ReadAll(resp.Body)
		return string(data), errRead
	}
	return "", errors.New("Github API failed with response code " + strconv.Itoa(resp.StatusCode))
}