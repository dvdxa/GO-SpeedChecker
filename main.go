package main

import "fmt"

func main() {
	url := "put here your url"
	dataSize := 1 * 1024 * 1024

	downloadSpeed, err := checkDownloadSpeed(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	uploadSpeed, err := checkUploadSpeed(url, make([]byte, dataSize))
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Download speed: %.2f Mbps\n", *downloadSpeed/1024/1024)
	fmt.Printf("Upload speed: %.2f Mbps\n", *uploadSpeed/1024/1024)
}
