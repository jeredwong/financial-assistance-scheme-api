package dto

type PaginationQuery struct {
    Page     int `form:"page" binding:"omitempty,min=1"`
    PageSize int `form:"page_size" binding:"omitempty,min=1,max=100"`
}

type PaginatedResponse struct {
    Data       interface{} `json:"data"`
    TotalItems int64       `json:"total_items"`
    TotalPages int         `json:"total_pages"`
    Page       int         `json:"page"`
    PageSize   int         `json:"page_size"`
}