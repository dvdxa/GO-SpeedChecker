package main

import (
	"bytes"
	"io"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

func checkUploadSpeed(url string, data []byte) (*float64, error) {
	startTime := time.Now()

	response, err := http.Post(url, "application/octet-stream", bytes.NewReader(data))
	if err != nil {
		return nil, errors.Wrapf(err, "uploadSpeed: failed to make request")
	}
	defer response.Body.Close()

	elapsedTime := time.Since(startTime).Seconds()

	uploadSpeed := float64(len(data) / int(elapsedTime))

	return &uploadSpeed, nil
}

func checkDownloadSpeed(url string) (*float64, error) {
	startTime := time.Now()

	response, err := http.Get(url)
	if err != nil {
		return nil, errors.Wrapf(err, "downloadSpeed: failed to make request")
	}

	defer response.Body.Close()

	_, err = io.Copy(io.Discard, response.Body)
	if err != nil {
		return nil, errors.Wrapf(err, "downloadSpeed: failed to copy body")
	}

	elapsedTime := time.Since(startTime).Seconds()

	downloadSpeed := float64(response.ContentLength) / elapsedTime

	return &downloadSpeed, nil
}
