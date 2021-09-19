package contests

import "time"

type CodeforcesApi struct {
	Name             string `json:"name"`
	Type             string `json:"type"`
	DurationSeconds  int    `json:"durationSeconds"`
	StartTimeSeconds int    `json:"startTimeSeconds"`
}

type CodeforcesInfo struct {
	Name      string
	Type      string
	StartTime time.Time
	EndTime   time.Time
}

func NewCodeforcesInfo(ca *CodeforcesApi) CodeforcesInfo {
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	startTime := time.Unix(int64(ca.StartTimeSeconds), 0).In(jst)
	endTime := startTime.Add(time.Duration(ca.DurationSeconds) * time.Second).In(jst)
	return CodeforcesInfo{Name: ca.Name, Type: ca.Type, StartTime: startTime, EndTime: endTime}
}
