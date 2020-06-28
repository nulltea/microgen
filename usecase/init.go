package usecase

import (
	"fmt"
	"os"

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

	fmt.Fprintf(os.Stdout, "Microservice structure generated successfully")
	return nil
}

func (s operations) generateEntitiesLayer() error {
	if err := os.MkdirAll("core", os.ModeDir); err != nil {
		return fmt.Errorf("unable to create CORE dir: " + err.Error())
	}

	if err := os.MkdirAll("core/model", os.ModePerm); err != nil {
		return fmt.Errorf("unable to create CORE/MODEL dir: " + err.Error())
	}

	if err := os.MkdirAll("core/service", os.ModePerm); err != nil {
		return fmt.Errorf("unable to create CORE/SERVICE dir: " + err.Error())
	}

	if err := os.MkdirAll("core/repo", os.ModePerm); err != nil {
		return fmt.Errorf("unable to create CORE/REPO dir: " + err.Error())
	}

	return nil
}

func (s operations) generateUseCaseLayer() error {
	if err := os.MkdirAll("usecase", os.ModePerm); err != nil {
		return fmt.Errorf("unable to create USECASE dir: " + err.Error())
	}

	if err := os.MkdirAll("usecase/storage", os.ModePerm); err != nil {
		return fmt.Errorf("unable to create USECASE/STORAGE dir: " + err.Error())
	}

	if err := os.MkdirAll("usecase/businesgen", os.ModePerm); err != nil {
		return fmt.Errorf("unable to create USECASE/BUSINESS dir: " + err.Error())
	}

	return nil
}

func (s operations) generateGatewayLayer() error {
	if err := os.MkdirAll("api", os.ModePerm); err != nil {
		return fmt.Errorf("unable to create API dir: " + err.Error())
	}

	if err := os.MkdirAll("api/rest", os.ModePerm); err != nil {
		return fmt.Errorf("unable to create API/REST dir: " + err.Error())
	}

	if err := os.MkdirAll("api/graphQL", os.ModePerm); err != nil {
		return fmt.Errorf("unable to create API/GRAPHQL dir: " + err.Error())
	}

	return nil
}
