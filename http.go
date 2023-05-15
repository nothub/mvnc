package mavengo

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

var HttpClient = http.DefaultClient

func GetXml(url string, v any) error {
	res, err := HttpClient.Get(url)
	if err != nil {
		return err
	}

	defer func(Body io.Closer) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(res.Body)

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		return errors.New("http status " + strconv.Itoa(res.StatusCode))
	}

	err = xml.NewDecoder(res.Body).Decode(&v)
	if err != nil {
		return errors.New("http status " + strconv.Itoa(res.StatusCode) + ": " + err.Error())
	}

	return nil
}
