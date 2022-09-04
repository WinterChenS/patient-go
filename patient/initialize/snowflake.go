package initialize

import (
	"time"

	sf "github.com/bwmarrin/snowflake"
	"winterchen.com/patient-go/patient/global"
)

func InitSnowflake(startTime time.Time, machineID int64) (err error) {
	sf.Epoch = startTime.UnixNano() / 1000000
	global.Snowflake, err = sf.NewNode(machineID)
	return
}
