package domain

import "time"

type CodeforcesApi struct {
	Result []CodeforcesApiResult `json:"result"`
}

type CodeforcesApiResult struct {
	Name             string `json:"name"`
	DurationSeconds  int    `json:"durationSeconds,omitempty"`
	Phase            string `json:"phase"`
	StartTimeSeconds int    `json:"startTimeSeconds,omitempty"`
}

type CodeforcesInfo struct {
	Name      string
	StartTime time.Time
	EndTime   time.Time
}

func NewCodeforcesInfo(ca *CodeforcesApiResult) CodeforcesInfo {
	startTime := time.Unix(int64(ca.StartTimeSeconds), 0).In(jst)
	endTime := startTime.Add(time.Duration(ca.DurationSeconds) * time.Second).In(jst)
	return CodeforcesInfo{Name: ca.Name, StartTime: startTime, EndTime: endTime}
}
