package helpers

import (
	"io"

	"winterchen.com/patient-go/patient/global"
	"go.uber.org/zap"
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
		if err == nil && exists {
			global.Log.Info("We already own " + bucketName)
		} else {
			global.Log.Error(err.Error())
			return
		}
	}

	err = global.MinioClient.SetBucketPolicy(bucketName, policy.BucketPolicyReadWrite)

	if err != nil {
		global.Log.Error(err.Error())
		return
	}
	global.Log.Info("Successfully created " + bucketName)
}

// upload file
func UploadFile(bucketName, objectName string, reader io.Reader, objectSize int64) (ok bool) {
	n, err := global.MinioClient.PutObject(bucketName, objectName, reader, objectSize, minio.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		global.Log.Error(err.Error())
		return false
	}
	global.Log.Info("Successfully uploaded bytes: ", zap.Any("n", n))
	return true
}

//  get file url
func GetFileUrl(bucketName string, fileName string) string {
	return "http://" + global.Configs.Minio.Endpoint + "/" + bucketName + "/" + fileName
}
