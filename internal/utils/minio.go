// coding: utf-8
// @Author : lryself
// @Date : 2022/4/3 23:18
// @Software: GoLand

package utils

import (
	"github.com/minio/minio-go"
	"github.com/spf13/viper"
	"io"
	"net/http"
)

type MinioClient struct {
	endpoint        string
	accessKeyID     string
	secretAccessKey string
	useSSL          bool
	bucketName      string
}

func GetMinioClient() MinioClient {
	var client MinioClient
	client.endpoint = viper.GetString("remote.minio.Endpoint")
	client.accessKeyID = viper.GetString("remote.minio.AccessKey")
	client.secretAccessKey = viper.GetString("remote.minio.SecretKey")
	client.useSSL = viper.GetBool("remote.minio.UseSSL")
	client.bucketName = viper.GetString("remote.minio.BucketName")
	return client
}

func (c MinioClient) initMinioClient() (*minio.Client, error) {
	// 初使化minio client对象。
	return minio.New(c.endpoint, c.accessKeyID, c.secretAccessKey, c.useSSL)
}

func (c MinioClient) PutObject(objectName, contentType string, object io.Reader, objectSize int64) error {
	minioClient, err := c.initMinioClient()
	if err != nil {
		return err
	}
	if contentType == "" {
		contentType, err = GetFileContentType(object)
		if err != nil {
			return err
		}
	}
	_, err = minioClient.PutObject(c.bucketName, objectName, object, objectSize, minio.PutObjectOptions{ContentType: contentType})
	return err
}

func GetFileContentType(out io.Reader) (string, error) {

	// Only the first 512 bytes are used to sniff the content type.
	buffer := make([]byte, 512)

	_, err := out.Read(buffer)
	if err != nil {
		return "", err
	}

	// Use the net/http package's handy DectectContentType function. Always returns a valid
	// content-type by returning "application/octet-stream" if no others seemed to match.
	contentType := http.DetectContentType(buffer)

	return contentType, nil
}
