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

	}

	return model.Response{}
}
