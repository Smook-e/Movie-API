package data

type Filter struct {
	Page     int      `query:"page" default:"1" minimum:"1" maximum:"1000000"`
    PageSize int      `query:"page_size" default:"20" minimum:"1" maximum:"100"`
    Sort     string   `query:"sort" default:"id" enum:"id,title,year,runtime,-id,-title,-year,-runtime"`
}
type Metadata struct {
    CurrentPage int `json:"current_page,omitzero"`
    PageSize int `json:"page_size,omitzero"`
    FirstPage int `json:"first_page,omitzero"`
    LastPage int `json:"last_page,omitzero"`
    TotalRecords int `json:"total_records,omitzero"`
}

var sorts = map[string]string{
    "id":    "id ASC",
    "title": "title ASC",
    "year":  "year ASC",
    "runtime": "runtime ASC",
    "-id":   "id DESC",
    "-title": "title DESC",
    "-year":  "year DESC",
    "-runtime": "runtime DESC",
}

func (f *Filter) SortColumn() string {
    if sort, ok := sorts[f.Sort]; ok {
        return sort
    }
    return sorts["id"]
}
func (f *Filter) Offset() int {
    return (f.Page - 1) * f.PageSize
}
func (f *Filter) Limit() int {
    return f.PageSize
}

func calculateMetadata(totalRecords, page, pageSize int) Metadata {
    if totalRecords == 0 {
        return Metadata{}
    }
    return Metadata{
        CurrentPage: page,
        PageSize: pageSize,
        FirstPage: 1,
        LastPage: (totalRecords + pageSize - 1) / pageSize,
        TotalRecords: totalRecords,
    }
}