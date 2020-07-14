package particeps

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
)

const (
	// AnonFiles is the constant for https://anonfiles.com/
	AnonFiles = iota
	// BayFiles is the constant for https://bayfiles.com/
	BayFiles = iota
	// Imgur is the constant for https://imgur.com/
	Imgur = iota
	// Filebin is the constant for https://filebin.net/
	Filebin = iota
)

// CheckFile checks if the filename exists and returns its size in pretty-print form
func CheckFile(filename string) (string, error) {
	fileInfo, err := os.Stat(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "particeps: error: could not find file \"%s\"\n", filename)
		return "", err
	}
	return prettySize(float64(fileInfo.Size())), nil
}

// AnonFilesUpload attemps to upload a file to AnonFiles and returns a success/failure string
func AnonFilesUpload(filename string) (UniversalResponse, error) {
	var finalRes UniversalResponse

	// Multi-part Body
	mpb := bytes.NewBuffer(nil)
	mw := multipart.NewWriter(mpb)

	partWriter, err := mw.CreateFormFile("file", filename)
	if err != nil {
		log.Fatalln(err) // TODO: return
	}

	fileReader, err := os.Open(filename)
	io.Copy(partWriter, fileReader)
	mw.Close()

	resp, err := http.Post("https://api.anonfiles.com/upload", mw.FormDataContentType(), mpb) // then we send the multipart body with the file to http.Post
	if err != nil {
		finalRes.Status = false
		return finalRes, err
	}

	//var result map[string]interface{}

	// json.NewDecoder(resp.Body).Decode(&result)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		finalRes.Status = false
		return finalRes, err
	}

	var successResponse AnonFilesSuccess
	err = json.Unmarshal(body, &successResponse)
	if err != nil {
		finalRes.Status = false
		return finalRes, err
	}

	finalRes.Status = successResponse.Status
	finalRes.FullURL = successResponse.Data.File.URL.Full
	finalRes.ShortURL = successResponse.Data.File.URL.Short
	return finalRes, nil
}

func round(val float64, roundOn float64, places int) (newVal float64) {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * val
	_, div := math.Modf(digit)
	if div >= roundOn {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	newVal = round / pow
	return
}

func prettySize(sizeInBytes float64) string {
	suffixes := [5]string{"B", "KB", "MB", "GB"}
	base := math.Log(sizeInBytes) / math.Log(1024)
	size := round(math.Pow(1024, base-math.Floor(base)), .5, 2)
	suffix := suffixes[int(math.Floor(base))]
	return strconv.FormatFloat(size, 'f', -1, 64) + " " + string(suffix)
}
