package params

type Club struct {
	ClubName string `json:"clubname"`
	Point    uint64 `json:"point"`
}

type RecordGame struct {
	ClubHomeName string `json:"clubhomename"`
	ClubAwayName string `json:"clubawayname"`
	Score        string `json:"score"`
}

type ClubStandings struct {
	ClubName string `json:"clubname"`
	Standing int    `json:"standing"`
}
