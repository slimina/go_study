package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

//客户端通过multipart.Write把文件的文本流写入一个缓存中，然后调用http的Post方法把缓存传到服务器。
func postFile(filename string, targetUrl string) error {

	bodyByte := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyByte)

	//可以调用bodyWriter.WriteField(fieldname, value)方法写很多其他类似的字段。

	fileWriter, err := bodyWriter.CreateFormFile("uploadfile", filename)
	if err != nil {
		return err
	}

	//打开文件句柄
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	//io copy
	_, err = io.Copy(fileWriter, f)
	if err != nil {
		return err
	}
	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()
	res, err := http.Post(targetUrl, contentType, bodyByte)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	res_body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	fmt.Println(res.Status)
	fmt.Println(string(res_body))
	return nil
}

func main() {

	targetUrl := "http://127.0.0.1:8080/upload"
	filename := "./test4.go"
	err := postFile(filename, targetUrl)
	if err != nil {
		fmt.Println(err)
	}
}
