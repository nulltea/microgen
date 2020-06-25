package usecase

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/timoth-y/microgen/core/meta"
)

func (s operations) Init(param meta.OperationParam) error {
	if err := s.generateEntitiesLayer(); err != nil {
		return err
	}

	if err := s.generateUseCaseLayer(); err != nil {
		return err
	}

	if err := s.generateGatewayLayer(); err != nil {
		return err
	}

	return nil
}

func (s operations) generateEntitiesLayer() error {
	if err := os.MkdirAll(filepath.Dir("core"), 0755); err != nil {
		return fmt.Errorf("unable to create CORE dir: " + err.Error())
	}

	if err := os.MkdirAll(filepath.Dir("core/model"), 0755); err != nil {
		return fmt.Errorf("unable to create CORE/MODEL dir: " + err.Error())
	}

	if err := os.MkdirAll(filepath.Dir("core/service"), 0755); err != nil {
		return fmt.Errorf("unable to create CORE/SERVICE dir: " + err.Error())
	}

	if err := os.MkdirAll(filepath.Dir("core/repo"), 0755); err != nil {
		return fmt.Errorf("unable to create CORE/REPO dir: " + err.Error())
	}

	return nil
}

func (s operations) generateUseCaseLayer() error {
	if err := os.MkdirAll(filepath.Dir("usecase"), 0755); err != nil {
		return fmt.Errorf("unable to create USECASE dir: " + err.Error())
	}

	if err := os.MkdirAll(filepath.Dir("usecase/storage"), 0755); err != nil {
		return fmt.Errorf("unable to create USECASE/STORAGE dir: " + err.Error())
	}

	if err := os.MkdirAll(filepath.Dir("usecase/business"), 0755); err != nil {
		return fmt.Errorf("unable to create USECASE/BUSINESS dir: " + err.Error())
	}

	return nil
}

func (s operations) generateGatewayLayer() error {
	if err := os.MkdirAll(filepath.Dir("api"), 0755); err != nil {
		return fmt.Errorf("unable to create API dir: " + err.Error())
	}

	if err := os.MkdirAll(filepath.Dir("api/rest"), 0755); err != nil {
		return fmt.Errorf("unable to create API/REST dir: " + err.Error())
	}

	if err := os.MkdirAll(filepath.Dir("api/graphQL"), 0755); err != nil {
		return fmt.Errorf("unable to create API/GRAPHQL dir: " + err.Error())
	}

	return nil
}