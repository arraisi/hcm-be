package customer

import (
	"context"
	"github.com/arraisi/hcm-be/internal/domain"
	"github.com/arraisi/hcm-be/internal/domain/dto/customer"
)

func (s *service) InquiryCustomer(ctx context.Context, request customer.CustomerInquiryRequest) (domain.Customer, error) {
	return domain.Customer{}, nil
}
