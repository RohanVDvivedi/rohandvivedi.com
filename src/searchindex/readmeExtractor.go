package searchindex

import (
	"net/http"
	"io"
	"os"
	"strconv"
	"errors"
)

const filestoreDirectory = "./db/index_data_assets"

func GetGithubRawFileDownloadLink(userName string, projectName string, fileName string) string {
	return "https://raw.githubusercontent.com/" + userName + "/" + projectName + "/master/" + fileName
}

func GetAndStoreGithubFile(userName string, projectName string, fileName string) error {
	GetApiRequestPath := GetGithubRawFileDownloadLink(userName, projectName, fileName)
	resp, err := http.Get(GetApiRequestPath)

	if(err != nil) {
		return err
	}

	if(resp.StatusCode == 200) {
		outputFileFolder := filestoreDirectory + "/" + userName + "/" + projectName
		err = os.MkdirAll(outputFileFolder, 0755)
		outFile, errFilecreate := os.Create(outputFileFolder + "/" + fileName)
		if(errFilecreate != nil) {
			return errFilecreate
		}

		buf := make([]byte, 1024)
		for {
			n, errInput := resp.Body.Read(buf)
			if errInput != nil && errInput != io.EOF {
				return errInput
			}
			if n == 0 {
				break
			}
			_, errOutput := outFile.Write(buf[:n])
        	if  errOutput != nil {
            	return errOutput
        	}
    	}
	} else {
		return errors.New("Github API failed with response code " + strconv.Itoa(resp.StatusCode))
	}
	return nil
}