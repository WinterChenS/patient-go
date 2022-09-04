package helpers

import (
	"fmt"
	"io"

	"winterchen.com/patient-go/patient/global"

	"github.com/minio/minio-go"
	"github.com/minio/minio-go/v6/pkg/policy"
)

// create bucket
func CreateMinoBuket(bucketName string) {
	location := "us-east-1"
	err := global.MinioClient.MakeBucket(bucketName, location)
	if err != nil {
		// check bucket exists
		exists, err := global.MinioClient.BucketExists(bucketName)
		fmt.Println(exists)
		if err == nil && exists {
			fmt.Printf("We already own %s\n", bucketName)
		} else {
			fmt.Println(err, exists)
			return
		}
	}

	err = global.MinioClient.SetBucketPolicy(bucketName, policy.BucketPolicyReadWrite)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Successfully created %s\n", bucketName)
}

// upload file
func UploadFile(bucketName, objectName string, reader io.Reader, objectSize int64) (ok bool) {
	n, err := global.MinioClient.PutObject(bucketName, objectName, reader, objectSize, minio.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		fmt.Println(err)
		return false
	}
	fmt.Println("Successfully uploaded bytes: ", n)
	return true
}

//  get file url
func GetFileUrl(bucketName string, fileName string) string {

	// reqParams := make(url.Values)
	// presignedURL, err := global.MinioClient.PresignedGetObject(bucketName, fileName, expires, reqParams)
	// if err != nil {
	// 	zap.L().Error(err.Error())
	// 	return ""
	// }
	// return fmt.Sprintf("%s", presignedURL)
	// return strings.Replace(fmt.Sprintf("%s", presignedURL), "http://"+global.Configs.Minio.Endpoint, global.Configs.Minio.Path, 1)

	return "http://" + global.Configs.Minio.Endpoint + "/" + bucketName + "/" + fileName
}
