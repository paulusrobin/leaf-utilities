package leafHttpResponse

type (
	successResponse struct {
		Data interface{} `json:"data"`
	}

	PagingResponse struct {
		Count       int          `json:"count"`
		CurrentPage int          `json:"currentPage"`
		TotalPage   int          `json:"totalPages"`
		Params      PagingParams `json:"params"`
		Items       interface{}  `json:"items"`
	}

	BasePagingResponse struct {
		Count       int          `json:"count"`
		CurrentPage int          `json:"currentPage"`
		TotalPage   int          `json:"totalPages"`
		Params      PagingParams `json:"params"`
	}

	SimplePagingResponse struct {
		CurrentPage int          `json:"currentPage"`
		Next        bool         `json:"next"`
		Params      PagingParams `json:"params"`
	}

	PagingParams struct {
		Page   int         `json:"page"`
		Limit  int         `json:"limit"`
		Sort   []string    `json:"sort"`
		Filter interface{} `json:"filter,omitempty"`
	}

	ItemsResponse struct {
		Items []interface{} `json:"items"`
	}
)

func (p PagingResponse) Val() interface{} {
	return p
}

func (p BasePagingResponse) Val() interface{} {
	return p
}

func (s *successResponse) Val() interface{} {
	return s
}

func newSuccessResponse(data interface{}) Response {
	return &successResponse{
		Data: data,
	}
}
