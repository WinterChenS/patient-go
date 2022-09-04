package config

type Config struct {
	Mysql    string      `mapstructure:"mysql"`    //mysql url
	Port     int         `mapstructure:"port"`     //server addr
	Name     string      `mapstructure:"name"`     //server name
	LogsPath string      `mapstructure:"logsPath"` //server logsPath
	Minio    MinioConfig `mapstructure:"minio"`    //minio config
}

type MinioConfig struct {
	Endpoint        string `mapstructure:"endpoint"`
	AccessKeyID     string `mapstructure:"accessKeyID"`
	SecretAccessKey string `mapstructure:"secretAccessKey"`
	BucketName      string `mapstructure:"bucketName"`
	Path            string `mapstructure:"path"`
}
