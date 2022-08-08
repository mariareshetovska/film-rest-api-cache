package db

type ResTimeLogs struct {
	ID         int64  `json:"id"`
	Request    string `json:"request"`
	TimeDB     int64  `json:"time_db"`
	TimeRedis  int64  `json:"time_redis"`
	TimeMemory int64  `json:"time_memory"`
}

func GetAllLogs() ([]*ResTimeLogs, error) {
	con := Connect()
	defer con.Close()
	sql := "select * from response_time_log"
	rs, err := con.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rs.Close()
	var respLogs []*ResTimeLogs
	for rs.Next() {
		var respLog ResTimeLogs
		err := rs.Scan(&respLog.ID, &respLog.Request, &respLog.TimeDB, &respLog.TimeRedis, &respLog.TimeMemory)
		if err != nil {
			return nil, err
		}
		respLogs = append(respLogs, &respLog)
	}
	return respLogs, nil
}
