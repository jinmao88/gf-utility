package curd

type Options interface {
	GetOptionLabel() string
	GetOptionValue() string
	GetOptionKey() string
	GetData() interface{}
}

type Option struct {
	Label string      `json:"label"`
	Value string      `json:"value"`
	Key   string      `json:"key"`
	Data  interface{} `json:"data"`
}

func BuildOptions[T Options](a []T) (res []Option) {

	for _, t := range a {
		res = append(res, Option{
			Label: t.GetOptionLabel(),
			Value: t.GetOptionKey(),
			Key:   t.GetOptionKey(),
			Data:  t.GetData(),
		})
	}
	return
}
