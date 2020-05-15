package file

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
)

func init() {
	//todo flag in config
	log.SetOutput(os.Stdout)
}

func Save(data string, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		log.Error("err while creating file", err)
		return errors.New("failed to write file")
	}
	defer file.Close()
	_, err = file.WriteString(data)

	if err != nil {
		log.Error("err while writing file", err)
	}

	return nil
}

func Open(name string) (string, error) {
	file, err := os.Open(name)

	if err != nil {
		return "", err
	}

	defer file.Close()

	data, err := ioutil.ReadAll(file)

	if err != nil {
		return "", err
	}

	return string(data), nil
}
