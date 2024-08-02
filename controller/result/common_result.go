package result

type CommonResult[T any] struct {
	ExtInfo string
	Success bool
	ErrorMessage string
	Data T
}

func NewCommonResult[T any](data T) *CommonResult[T] {
	return &CommonResult[T]{
		Data: data,
		Success: true,
		ExtInfo: "",
		ErrorMessage: "",
	}
}
