package fetcher

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

func GetBodyAsString(url string) (string, error) {
	log.Println("getting xml via http")
	resp, err := http.Get(url)
	if err != nil {
		log.WithField("err", err).Error("error while getting sitemap")
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Error("wrong status", resp.Status)
		return "", failedToDownload()
	}

	log.Println("reading data from response")
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error("err while reading body", err)
		return "", failedToDownload()
	}

	return string(body), nil
}

func failedToDownload() error {
	return errors.New("fail to download sitemap")
}
