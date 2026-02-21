package model

type Output struct{
	Sucess bool `json:"sucess"`
	Error string `json:"error"`
}

func NewOuput(sucess bool, error string) Output{
	return Output{Sucess: sucess, Error: error}
}
