package service

import (
	"context"
	"github.com/shopspring/decimal"
	"healthcare-allopurinol/entity"
)

func calculateOfficial(s *CheckoutServiceImpl, ctx context.Context, shipping *entity.Shipping, method *entity.LogisticPartner) (*entity.ShippingCost, error) {
	distance, err := s.addressRepo.GetDistance(ctx, shipping.PharmacyLat, shipping.PharmacyLong, shipping.UserLat, shipping.UserLong)
	if err != nil {
		return nil, err
	}

	kmDist := distance.Div(decimal.NewFromInt(1000))
	cost := kmDist.Ceil().Mul(decimal.NewFromInt(*method.Rate))

	return &entity.ShippingCost{
		MethodId: method.Id,
		Method:   method.Name,
		Cost:     cost,
	}, nil
}

func getRajaOngkirId(s *CheckoutServiceImpl, ctx context.Context, postalCode string) (int64, error) {
	ro, err := s.rajaOngkirBridge.GetDataWithPostal(ctx, postalCode)
	if err != nil {
		return 0, err
	}

	return ro.Data[0].Id, nil
}

func getRajaOngkirCost(s *CheckoutServiceImpl, ctx context.Context, method *entity.LogisticPartner, userROId, pharmacyROId int64, weight decimal.Decimal) (decimal.Decimal, error) {
	ro, err := s.rajaOngkirBridge.GetShippingCost(ctx, pharmacyROId, userROId, weight, *method.Courier)
	if err != nil {
		return decimal.Zero, err
	}

	for _, data := range ro.Data {
		if data.Service == *method.Code {
			return decimal.NewFromInt(data.Cost), nil
		}
	}

	return decimal.Zero, nil
}

func getTotalWeight(catalogs []*entity.Catalog) decimal.Decimal {
	totalWeight := decimal.Decimal{}
	for _, catalog := range catalogs {
		totalWeight = totalWeight.Add(catalog.Product.Weight.Mul(decimal.NewFromInt(int64(catalog.Quantity))))
	}
	return totalWeight
}

func calculate3rdParty(s *CheckoutServiceImpl, ctx context.Context, shipping *entity.Shipping, method *entity.LogisticPartner) (*entity.ShippingCost, error) {
	userROId, err := s.rajaOngkirRepo.GetROId(ctx, shipping.UserPostal)
	if err != nil {
		return nil, err
	}
	if userROId == 0 {
		userROId, err = getRajaOngkirId(s, ctx, shipping.UserPostal)
		if err != nil {
			return nil, err
		}
		err = s.rajaOngkirRepo.InsertROId(ctx, shipping.UserPostal, userROId)
		if err != nil {
			return nil, err
		}
	}

	pharmacyROId, err := s.rajaOngkirRepo.GetROId(ctx, shipping.PharmacyPostal)
	if err != nil {
		return nil, err
	}
	if pharmacyROId == 0 {
		pharmacyROId, err = getRajaOngkirId(s, ctx, shipping.PharmacyPostal)
		if err != nil {
			return nil, err
		}
		err = s.rajaOngkirRepo.InsertROId(ctx, shipping.PharmacyPostal, pharmacyROId)
		if err != nil {
			return nil, err
		}
	}

	cost, err := getRajaOngkirCost(s, ctx, method, userROId, pharmacyROId, shipping.TotalWeight)
	if err != nil {
		return nil, err
	}

	if cost.IsZero() {
		return nil, nil
	}

	return &entity.ShippingCost{
		MethodId: method.Id,
		Method:   method.Name,
		Cost:     cost,
	}, nil
}

func (s *CheckoutServiceImpl) GetShippingCosts(ctx context.Context, shipping *entity.Shipping, credId, addressId int64) (*entity.Shipping, error) {
	methods, err := s.shippingRepo.GetLogisticPartnersByPharmacy(ctx, shipping.PharmacyId)
	if err != nil {
		return nil, err
	}

	user, err := s.userRepo.GetUserByCredId(ctx, credId)
	if err != nil {
		return nil, err
	}
	shipping.UserId = user.Id

	userAdr, err := s.addressRepo.GetAddress(ctx, addressId)
	if err != nil {
		return nil, err
	}
	shipping.UserPostal = userAdr.PostalCode
	shipping.UserLat = userAdr.Latitude
	shipping.UserLong = userAdr.Longitude

	pharmaAdr, err := s.pharmacyRepo.GetPharmacy(ctx, shipping.PharmacyId)
	if err != nil {
		return nil, err
	}
	shipping.PharmacyPostal = pharmaAdr.Address.PostalCode
	shipping.PharmacyLat = pharmaAdr.Address.Latitude
	shipping.PharmacyLong = pharmaAdr.Address.Longitude

	shipping.TotalWeight = getTotalWeight(shipping.Catalogs)

	for _, method := range methods {
		var cost *entity.ShippingCost
		if method.Rate != nil {
			cost, err = calculateOfficial(s, ctx, shipping, method)
			if err != nil {
				return nil, err
			}
		} else {
			cost, err = calculate3rdParty(s, ctx, shipping, method)
			if err != nil {
				return nil, err
			}
			if cost == nil {
				continue
			}
		}
		shipping.ShippingCosts = append(shipping.ShippingCosts, cost)
	}

	err = s.redisRepo.SetShippingData(shipping)
	if err != nil {
		return nil, err
	}

	return shipping, nil
}
