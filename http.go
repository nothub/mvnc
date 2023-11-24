package mvnc

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

var HttpClient = http.DefaultClient

var jarMT = "application/java-archive"
var xmlMT = "application/xml"

var ua = "mvnc/0.1 (+github.com/nothub/mvnc)"
var auth = ""

func GetXml(url string, v any) error {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Accept", xmlMT)

	res, err := SendRequest(req)

	defer func(Body io.Closer) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(res.Body)

	err = xml.NewDecoder(res.Body).Decode(&v)
	if err != nil {
		return errors.New("xml decode error: " + err.Error())
	}

	return nil
}

func SendRequest(req *http.Request) (*http.Response, error) {
	req.Header.Set("User-Agent", ua)

	if auth != "" {
		req.Header.Set("Authorization", auth)
	}

	res, err := HttpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		return nil, errors.New("http status " + strconv.Itoa(res.StatusCode))
	}

	return res, nil
}

func GetMetadata(baseUrl string, groupId string, artifactId string) (*Metadata, error) {
	// baseUrl example: "https://maven.quiltmc.org/repository/release"
	// groupId example: "org.quiltmc"
	// artifactId example: "quilt-installer"

	groupId = strings.ReplaceAll(groupId, ".", "/")
	u := fmt.Sprintf("%s/%s/%s/%s", baseUrl, groupId, artifactId, metaFile)

	m := &Metadata{}
	err := GetXml(u, m)
	if err != nil {
		return nil, err
	}

	return m, nil
}

func GetArtifact() {
	// TODO: headers
	//   If-Modified-Since

}

func DeployArtifact() {
	// TODO: headers
	//   Content-Type
	//   Content-Length
	//   Content-Disposition
}
