package pkg

const (
	defaultPageSize    = 10
	DefaultMaxPageSize = 10000
	defaultPageNumber  = 1
)

type ParamPagination struct {
	PageSize   int `form:"page_size" json:"page_size"`
	PageNumber int `form:"page_number" json:"page_number"`
	Count      int `json:"-"`
}

func (p *ParamPagination) GetPageSize() int {
	pageSize := p.PageSize
	if pageSize == 0 {
		pageSize = defaultPageSize
	}

	return pageSize
}

func (p *ParamPagination) GetPageNumber() int {
	pageNumber := p.PageNumber
	if pageNumber == 0 {
		pageNumber = defaultPageNumber
	}

	return pageNumber
}

func (p *ParamPagination) GetOffer() int {
	offset := (p.GetPageNumber() - 1) * p.GetPageSize()
	return offset
}

type RespListWithPagination struct {
	PageSize   int         `json:"page_size"`
	PageNumber int         `json:"page_number"`
	Count      int         `json:"count"`
	Results    interface{} `json:"results" swaggertype:"array,object"`
}
