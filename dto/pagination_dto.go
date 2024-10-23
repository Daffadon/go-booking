package dto

type (
	PaginationRequest struct {
		Page uint16 `form:"page" json:"page"`
	}
	PaginationResponse struct {
		Page      uint16 `json:"page"`
		PrevPage  uint16 `json:"prev_page"`
		NextPage  uint16 `json:"next_page"`
		TotalPage uint16 `json:"total_page"`
	}
)
