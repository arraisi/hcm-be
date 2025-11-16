package customer

import (
	"context"
	"database/sql"
	"errors"
	"github.com/arraisi/hcm-be/internal/domain"
	"github.com/arraisi/hcm-be/internal/domain/dto/customer"
	errorx "github.com/arraisi/hcm-be/pkg/errors"
	"github.com/arraisi/hcm-be/pkg/utils"
)

func (s *service) InquiryCustomer(ctx context.Context, request customer.CustomerInquiryRequest) (domain.Customer, error) {
	params := customer.GetCustomerRequest{}
	if *request.FlagNoHp {
		params.PhoneNumber = utils.ToValue(request.NoHp)
	} else {
		params.KTPNumber = utils.ToValue(request.NIK)
	}

	customerData, err := s.repo.GetCustomer(ctx, params)
	if errors.Is(err, sql.ErrNoRows) {
		return domain.Customer{}, errorx.ErrCustomerNotFound
	}
	if err != nil {
		return domain.Customer{}, err
	}
	return customerData, nil
}
