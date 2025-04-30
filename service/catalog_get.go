package service

import (
	"context"
	"healthcare-allopurinol/constant"
	"healthcare-allopurinol/entity"
	"healthcare-allopurinol/sentinel"
)

func (s *CatalogServiceImpl) GetCatalogsService(ctx context.Context, options *entity.CatalogOptions, pharmacistCredId int64) ([]*entity.Catalog, *entity.CatalogOptions, error) {
	pharmacistId, err := s.pharmacistRepo.GetIdByCredId(ctx, pharmacistCredId)
	if err != nil {
		return nil, nil, err
	}

	assigned, err := s.pharmacistRepo.IsAssigned(ctx, pharmacistId)
	if err != nil {
		return nil, nil, err
	}
	if !assigned {
		return nil, nil, sentinel.ErrPharmacistNotAssigned
	}

	pharmacyId, err := s.pharmacistRepo.GetPharmacyIdByCredId(ctx, pharmacistCredId)
	if err != nil {
		return nil, nil, err
	}

	count, err := s.catalogRepo.GetCatalogsCount(ctx, options, pharmacyId)
	if err != nil {
		return nil, nil, err
	}
	options.TotalRows = count

	catalogs, err := s.catalogRepo.GetCatalogs(ctx, options, pharmacyId)
	if err != nil {
		return nil, nil, err
	}

	return catalogs, options, nil
}

func (s *CatalogServiceImpl) GetCatalogService(ctx context.Context, id int64, pharmacistCredId int64) (*entity.Catalog, error) {
	pharmacistId, err := s.pharmacistRepo.GetIdByCredId(ctx, pharmacistCredId)
	if err != nil {
		return nil, err
	}

	assigned, err := s.pharmacistRepo.IsAssigned(ctx, pharmacistId)
	if err != nil {
		return nil, err
	}
	if !assigned {
		return nil, sentinel.ErrPharmacistNotAssigned
	}

	exist, err := s.catalogRepo.IsIdExist(ctx, id)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, sentinel.ErrCatalogNotFound
	}

	pharmacyId, err := s.catalogRepo.GetPharmacyId(ctx, id)
	if err != nil {
		return nil, err
	}

	assigned, err = s.pharmacistRepo.IsAssignedToSpecificPharmacy(ctx, pharmacistCredId, pharmacyId)
	if err != nil {
		return nil, err
	}
	if !assigned {
		return nil, sentinel.ErrPharmacistNoAccess
	}

	catalog, err := s.catalogRepo.GetCatalog(ctx, id)
	if err != nil {
		return nil, err
	}

	return catalog, nil
}

func (s *CatalogServiceImpl) GetAvailableCatalogsService(ctx context.Context, options *entity.AvailableCatalogOptions, userCredId int64) ([]*entity.Catalog, *entity.AvailableCatalogOptions, error) {
	user, err := s.userRepo.GetUserByCredId(ctx, userCredId)
	if err != nil {
		return nil, nil, err
	}

	address := &entity.Address{
		Longitude: constant.DEFAULT_LOCATION_LONGITUDE,
		Latitude:  constant.DEFAULT_LOCATION_LATITUDE,
	}
	exist, err := s.addressRepo.IsUserActiveAddressExist(ctx, user.Id)
	if err != nil {
		return nil, nil, err
	}
	if exist {
		addressId, err := s.addressRepo.GetUserActiveAddressId(ctx, user.Id)
		if err != nil {
			return nil, nil, err
		}

		address, err = s.addressRepo.GetAddress(ctx, addressId)
		if err != nil {
			return nil, nil, err
		}
	}

	catalogs, count, err := s.catalogRepo.GetAvailableCatalogs(ctx, options, address)
	if err != nil {
		return nil, nil, err
	}
	options.TotalRows = count

	return catalogs, options, nil
}

func (s *CatalogServiceImpl) GetAvailableCatalogService(ctx context.Context, id int64) (*entity.Catalog, error) {
	exist, err := s.catalogRepo.IsIdExist(ctx, id)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, sentinel.ErrCatalogNotFound
	}
	
	catalog, err := s.catalogRepo.GetAvailableCatalog(ctx, id)
	if err != nil {
		return nil, err
	}

	return catalog, nil
}

func (s *CatalogServiceImpl) GetCheckoutCatalogsService(ctx context.Context, userCredId int64, addressId int64) ([]*entity.Pharmacy, map[int64][]*entity.Catalog, []*entity.Catalog, error) {
	user, err := s.userRepo.GetUserByCredId(ctx, userCredId)
	if err != nil {
		return nil, nil, nil, err
	}

	exist, err := s.addressRepo.IsAddressExist(ctx, addressId)
	if err != nil {
		return nil, nil, nil, err
	}
	if !exist {
		return nil, nil, nil, sentinel.ErrAddressNotFound
	}

	addressUserId, err := s.addressRepo.GetUserId(ctx, addressId)
	if err != nil {
		return nil, nil, nil, err
	}
	if addressUserId == nil || user.Id != *addressUserId {
		return nil, nil, nil, sentinel.ErrAddressNotAssociated
	}

	address, err := s.addressRepo.GetAddress(ctx, addressId)
	if err != nil {
		return nil, nil, nil, err
	}

	cartItems, err := s.cartRepo.GetCartItems(ctx, user.Id)
	if err != nil {
		return nil, nil, nil, err
	}
	if len(cartItems) == 0 {
		return make([]*entity.Pharmacy, 0), make(map[int64][]*entity.Catalog), make([]*entity.Catalog, 0), nil
	}

	catalogs, err := s.catalogRepo.GetCheckoutCatalogs(ctx, cartItems, address)
	if err != nil {
		return nil, nil, nil, err
	}

	quantity := make(map[int64]int)
	for _, item := range cartItems {
		quantity[item.ProductId] = item.Quantity
	}

	pharmacies := make(map[int64]*entity.Pharmacy)
	catalogItems := make(map[int64][]*entity.Catalog)
	products := make(map[int64]bool)
	pharmacyIds := make([]int64, 0)
	for _, c := range catalogs {
		c.Quantity = quantity[c.Product.Id]
		pharmacyId := c.Pharmacy.Id

		pharmacyIds = append(pharmacyIds, pharmacyId)
		pharmacies[pharmacyId] = c.Pharmacy
		catalogItems[pharmacyId] = append(catalogItems[pharmacyId], c)
		products[c.Product.Id] = true
	}

	var pharmaciesLogisticPartner []*entity.Pharmacy
	if len(pharmacyIds) > 0 {
		pharmaciesLogisticPartner, err = s.shippingRepo.GetLogisticPartnersByPharmacyBulk(ctx, pharmacyIds)
		if err != nil {
			return nil, nil, nil, err
		}
		for _, pharmacy := range pharmaciesLogisticPartner {
			pharmacies[pharmacy.Id].LogisticPartners = pharmacy.LogisticPartners
		}
	}

	var (
		unavailableProductIds []int64
		unavailableProducts   []*entity.Product
	)
	for _, item := range cartItems {
		if exist := products[item.ProductId]; !exist {
			unavailableProductIds = append(unavailableProductIds, item.ProductId)
		}
	}
	if len(unavailableProductIds) > 0 {
		unavailableProducts, err = s.productRepo.GetProductByIdBulk(ctx, unavailableProductIds)
		if err != nil {
			return nil, nil, nil, err
		}
	}

	unavailableCatalogs := make([]*entity.Catalog, len(unavailableProducts))
	for i, p := range unavailableProducts {
		var catalog entity.Catalog
		catalog.Product = p
		catalog.Quantity = quantity[p.Id]
		unavailableCatalogs[i] = &catalog
	}

	result := make([]*entity.Pharmacy, len(pharmacies))
	i := 0
	for _, v := range pharmacies {
		result[i] = v
		i++
	}
	return result, catalogItems, unavailableCatalogs, nil
}

func (s *CatalogServiceImpl) GetMostBoughtCatalogsService(ctx context.Context) ([]*entity.Catalog, error) {
	productIds, err := s.productRepo.GetMostBoughtProductIdsToday(ctx)
	if err != nil {
		return nil, err
	}

	if len(productIds) == 0 {
		productIds, err = s.productRepo.GetMostBoughtProductIdsAllTime(ctx)
		if err != nil {
			return nil, err
		}
	}

	catalogs, err := s.catalogRepo.GetCatalogsByProductIds(ctx, productIds)
	if err != nil {
		return nil, err
	}

	return catalogs, nil
}
