package leafModel

type (
	PagingParams struct {
		Page   int         `json:"page"`
		Limit  int         `json:"limit"`
		Sort   []string    `json:"sort"`
		Filter interface{} `json:"filter,omitempty"`
	}

	BasePagingResponse struct {
		Count       int          `json:"count"`
		CurrentPage int          `json:"currentPage"`
		TotalPage   int          `json:"totalPages"`
		Params      PagingParams `json:"params"`
	}

	BaseSimplePagingResponse struct {
		CurrentPage int          `json:"currentPage"`
		Next        bool         `json:"next"`
		Params      PagingParams `json:"params"`
	}

	PagingResponse struct {
		Count       int          `json:"count"`
		CurrentPage int          `json:"currentPage"`
		TotalPage   int          `json:"totalPages"`
		Params      PagingParams `json:"params"`
		Items       interface{}  `json:"items"`
	}

	SimplePagingResponse struct {
		CurrentPage int          `json:"currentPage"`
		Next        bool         `json:"next"`
		Params      PagingParams `json:"params"`
		Items       interface{}  `json:"items"`
	}
)
