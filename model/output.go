package model

type Output struct{
	Success bool `json:"success"`
	Error string `json:"error"`
}

func NewOuput(sucess bool, error string) Output{
	return Output{Success: sucess, Error: error}
}
