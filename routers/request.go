package routers

import "errors"

type Request struct {
	params map[string]any
	query  map[string]string
	Method string
}

func (this*Request) Get(key string) (any, error) {
	params, ok := this.params[key]
	if !ok {
		return "", errors.New("Param not found")
	}
	return params, nil
}
