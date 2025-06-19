package errors

func HandleError[T any](result T, err error) (T, error) {
	if err != nil {
		var zero T
		return zero, err
	}
	return result, nil
}
