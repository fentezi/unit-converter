package model

type Request struct {
	Input    string `json:"input"`
	UnitForm string `json:"unitForm"`
	UnitTo   string `json:"unitTo"`
}

type Response struct {
	Result float64 `json:"result"`
}
