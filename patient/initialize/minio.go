package initialize

import (
	"log"

	"winterchen.com/patient-go/patient/global"
	"winterchen.com/patient-go/patient/helpers"

	"github.com/minio/minio-go"
	_ "github.com/minio/minio-go/pkg/encrypt"
)

func InitMinIO() {
	minioInfo := global.Configs.Minio
	// init minio client object.
	minioClient, err := minio.New(minioInfo.Endpoint, minioInfo.AccessKeyID, minioInfo.SecretAccessKey, false)
	if err != nil {
		log.Fatalln(err)
	}
	// set global minio client object.
	global.MinioClient = minioClient
	// create bucket.
	helpers.CreateMinoBuket(global.Configs.Minio.BucketName)
}
