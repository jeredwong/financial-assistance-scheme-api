package dto

type PaginationQuery struct {
    Page     int 
    PageSize int 
}

type PaginatedResponse struct {
    Data       interface{} 
    TotalItems int64       
    TotalPages int         
    Page       int         
    PageSize   int         
}