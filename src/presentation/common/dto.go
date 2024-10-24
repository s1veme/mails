package common

type Pagination struct {
	AllCount int `json:"all_count"`
	Limit    int `json:"limit"`
	Offset   int `json:"offset"`
	Page     int `json:"page"`
}

type ApiResponse[T any] struct {
	Data     *T                  `json:"data"`
	Errors   map[string][]string `json:"errors"`
	Messages []string            `json:"messages"`
}

func NewSuccessResponse[T any](data *T, messages ...string) ApiResponse[T] {
	return ApiResponse[T]{
		Data:     data,
		Errors:   nil,
		Messages: messages,
	}
}

func NewErrorResponse[T any](errors map[string][]string, messages ...string) ApiResponse[T] {
	return ApiResponse[T]{
		Data:     nil,
		Errors:   errors,
		Messages: messages,
	}
}
