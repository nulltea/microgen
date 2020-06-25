package service

import "github.com/timoth-y/microgen/core/meta"

type OperationService interface {
	Init(param meta.OperationParam) error
	Generate(param meta.OperationParam) error
}