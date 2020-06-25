package usecase

import (
	"github.com/timoth-y/microgen/core/service"
)

type operations struct {

}

func NewOperationService() service.OperationService {
	return &operations{

	}
}
