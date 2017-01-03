package domain

type Pagination struct {
	Limit  int64 `validate:"neglect"`
	Page   int64 `validate:"neglect"`
	Offset int64
}
