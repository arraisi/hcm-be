package hmf

import (
	"context"

	"github.com/arraisi/hcm-be/internal/domain/dto/creditsimulation"
)

type HmfClient interface {
	GetBranches(ctx context.Context) ([]creditsimulation.BranchResponse, error)
}

type Service interface {
	GetBranches(ctx context.Context) ([]creditsimulation.BranchResponse, error)
}

type service struct {
	client HmfClient
}

func NewService(client HmfClient) Service {
	return &service{
		client: client,
	}
}

func (s *service) GetBranches(ctx context.Context) ([]creditsimulation.BranchResponse, error) {
	return s.client.GetBranches(ctx)
}
