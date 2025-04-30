package service

import (
	"context"
	"healthcare-allopurinol/entity"
	"healthcare-allopurinol/sentinel"
)

func (s *OrderServiceImpl) GetPendingOrdersService(ctx context.Context, options *entity.PendingOrderOptions, credId int64) ([]*entity.PendingOrderGroup, error) {
	user, err := s.userRepo.GetUserByCredId(ctx, credId)
	if err != nil {
		return nil, err
	}

	count, err := s.orderRepo.GetPendingOrdersCount(ctx, user.Id)
	if err != nil {
		return nil, err
	}
	options.TotalRows = count

	canceledOrders, err := s.orderRepo.GetPendingOrders(ctx, options, user.Id)
	if err != nil {
		return nil, err
	}

	return canceledOrders, nil
}

func (s *OrderServiceImpl) GetPharmacyOrdersService(ctx context.Context, options *entity.PharmacyOrderOptions, pharmacistCredId int64) ([]*entity.Order, *entity.PharmacyOrderOptions, error) {
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

	count, err := s.orderRepo.GetPharmacyOrdersCount(ctx, options, pharmacyId)
	if err != nil {
		return nil, nil, err
	}
	options.TotalRows = count

	orders, err := s.orderRepo.GetPharmacyOrders(ctx, options, pharmacyId)
	if err != nil {
		return nil, nil, err
	}

	return orders, options, nil
}

func (s *OrderServiceImpl) GetPharmacyOrderService(ctx context.Context, id int64, pharmacistCredId int64) (*entity.Order, error) {
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

	exist, err := s.orderRepo.IsIdExist(ctx, id)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, sentinel.ErrOrderNotFound
	}

	pharmacyId, err := s.pharmacistRepo.GetPharmacyIdByCredId(ctx, pharmacistCredId)
	if err != nil {
		return nil, err
	}

	orderPharmacyId, err := s.orderRepo.GetPharmacyIdByOrderId(ctx, id)
	if err != nil {
		return nil, err
	}
	if pharmacyId != orderPharmacyId {
		return nil, sentinel.ErrPharmacistNoAccess
	}

	order, err := s.orderRepo.GetOrder(ctx, id)
	if err != nil {
		return nil, err
	}

	orderItems, err := s.orderRepo.GetOrderItems(ctx, id)
	if err != nil {
		return nil, err
	}

	order.OrderItems = orderItems
	return order, nil
}

func (s *OrderServiceImpl) GetOrdersService(ctx context.Context, options *entity.OrderOptions) ([]*entity.Order, *entity.OrderOptions, error) {
	count, err := s.orderRepo.GetOrdersCount(ctx, options)
	if err != nil {
		return nil, nil, err
	}
	options.TotalRows = count

	orders, err := s.orderRepo.GetOrders(ctx, options)
	if err != nil {
		return nil, nil, err
	}

	return orders, options, nil
}

func (s *OrderServiceImpl) GetUserOrdersService(ctx context.Context, options *entity.UserOrderOptions, userCredId int64) ([]*entity.Order, *entity.UserOrderOptions, error) {
	user, err := s.userRepo.GetUserByCredId(ctx, userCredId)
	if err != nil {
		return nil, nil, err
	}

	count, err := s.orderRepo.GetUserOrdersCount(ctx, options, user.Id)
	if err != nil {
		return nil, nil, err
	}
	options.TotalRows = count

	orders, err := s.orderRepo.GetUserOrders(ctx, options, user.Id)
	if err != nil {
		return nil, nil, err
	}

	return orders, options, nil
}

func (s *OrderServiceImpl) GetOrderService(ctx context.Context, id int64) (*entity.Order, error) {
	exist, err := s.orderRepo.IsIdExist(ctx, id)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, sentinel.ErrOrderNotFound
	}

	order, err := s.orderRepo.GetOrder(ctx, id)
	if err != nil {
		return nil, err
	}

	orderItems, err := s.orderRepo.GetOrderItems(ctx, id)
	if err != nil {
		return nil, err
	}

	order.OrderItems = orderItems
	return order, nil
}
