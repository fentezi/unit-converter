package service

import (
	"log/slog"

	"github.com/fentezi/unit-converter/model"
)

type Service struct {
	log *slog.Logger
}

func NewService(log *slog.Logger) *Service {
	return &Service{
		log: log,
	}
}

func (s *Service) Convert(request model.Request) model.Response {
	switch request.UnitForm {
	case "m":
		switch request.UnitTo {
		case "km":
			return model.Response{Result: request.Input / 1000}
		case "ft":
			return model.Response{Result: request.Input * 3.28084}
		}
	case "km":
		switch request.UnitTo {
		case "m":
			return model.Response{Result: request.Input * 1000}
		case "ft":
			return model.Response{Result: request.Input * 3280.84}
		}
	case "ft":
		switch request.UnitTo {
		case "m":
			return model.Response{Result: request.Input / 3.28084}
		case "km":
			return model.Response{Result: request.Input / 3280.84}
		}

	case "mm":
		switch request.UnitTo {
		case "cm":
			return model.Response{Result: request.Input / 10}
		case "km":
			return model.Response{Result: request.Input / 10000000}

		}
	case "cm":
		switch request.UnitTo {
		case "mm":
			return model.Response{Result: request.Input * 10}
		case "km":
			return model.Response{Result: request.Input / 100000}
		}

	case "kg":
		switch request.UnitTo {
		case "g":
			return model.Response{Result: request.Input * 1000}
		case "lb":
			return model.Response{Result: request.Input * 2.205}

		}

	case "g":
		switch request.UnitTo {
		case "kg":
			return model.Response{Result: request.Input / 1000}
		case "lb":
			return model.Response{Result: request.Input / 453.592}

		}

	case "lb":
		switch request.UnitTo {
		case "kg":
			return model.Response{Result: request.Input / 2.205}
		case "g":
			return model.Response{Result: request.Input * 453.592}

		}

	case "C":
		switch request.UnitTo {
		case "F":
			return model.Response{Result: request.Input*1.8 + 32}
		}

	case "F":
		switch request.UnitTo {
		case "C":
			return model.Response{Result: (request.Input - 32) / 1.8}
		}

	case "K":
		switch request.UnitTo {
		case "C":
			return model.Response{Result: request.Input - 273.15}
		}
	}

	return model.Response{}
}
