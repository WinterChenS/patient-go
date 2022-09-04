package helpers

import "winterchen.com/patient-go/patient/global"

func GenID() int64 {
	return global.Snowflake.Generate().Int64()
}
