package dto

type (
	PaginationRequest struct {
		Page uint16 `form:"page"`
	}
	PaginationResponse struct {
		Page      uint16 `json:"page"`
		NextPage  uint16 `json:"next_page"`
		TotalPage uint16 `json:"total_page"`
	}
)
