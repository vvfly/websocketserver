package aproto

type GetChargeRoomUserListReq struct {
	AnchorID   string `json:"anchorId"`
	LiveID     string `json:"liveId"`
	PageSize   int    `json:"pageSize"`
	PageNumber int    `json:"pageNumber"`
	Sort       string `json:"sort"` // up 升序;down 降序
}

type GetChargeRoomUserListResp struct {
	AnchorID   string `json:"anchorId"`
	LiveID     string `json:"liveId"`
	PageSize   int    `json:"pageSize"`
	PageNumber int    `json:"pageNumber"`
	Sort       string `json:"sort"` // up 升序;down 降序
}
