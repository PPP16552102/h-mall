package dto

type PageDto struct {
	PageNum  int64 `query:"page_num" validate:"min=1"`
	PageSize int64 `query:"page_size" validate:"min=1,max=100"`
}