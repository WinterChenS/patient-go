package global

import (
	sf "github.com/bwmarrin/snowflake"
	ut "github.com/go-playground/universal-translator"
	"github.com/jinzhu/gorm"
	"github.com/minio/minio-go"
	"go.uber.org/zap"
	"winterchen.com/patient-go/patient/config"
)

var (
	Configs     config.Config
	Trans       ut.Translator
	DB          *gorm.DB
	Log         *zap.Logger
	MinioClient *minio.Client
	Snowflake   *sf.Node
)
