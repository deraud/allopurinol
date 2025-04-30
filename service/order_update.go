package service

import (
	"context"
	"healthcare-allopurinol/constant"
	"healthcare-allopurinol/entity"
	"healthcare-allopurinol/sentinel"
	"healthcare-allopurinol/utility/logger"

	"github.com/robfig/cron"
)

func (s *OrderServiceImpl) UpdateOrderGroupService(ctx context.Context, o *entity.OrderGroup, userCredId int64) (*entity.OrderGroup, error) {
	user, err := s.userRepo.GetUserByCredId(ctx, userCredId)
	if err != nil {
		return nil, err
	}

	exist, err := s.orderRepo.IsOrderGroupIdExist(ctx, o.Id)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, sentinel.ErrOrderGroupNotFound
	}

	groupOrderUserId, err := s.orderRepo.GetUserIdFromOrderGroup(ctx, o.Id)
	if err != nil {
		return nil, err
	}
	if user.Id != groupOrderUserId {
		return nil, sentinel.ErrOrderGroupNotAssociated
	}
	o.UserId = user.Id

	orderGroup, err := s.orderRepo.UpdateOrderGroup(ctx, o)
	if err != nil {
		return nil, err
	}

	err = s.orderRepo.UpdateOrderGroupStatusVerifying(ctx, o.Id)
	if err != nil {
		return nil, err
	}

	c := cron.New()
	c.AddFunc("0 * * * * *", func() {
		err = s.orderRepo.UpdateOrderGroupStatusProccessed(ctx, o.Id)
		if err != nil {
			logger.Log.Error(err)
		}
		c.Stop()
	})
	c.Start()

	return orderGroup, nil
}

func (s *OrderServiceImpl) UpdateUserOrderCanceledService(ctx context.Context, o *entity.Order, userCredId int64) (*entity.Order, error) {
	user, err := s.userRepo.GetUserByCredId(ctx, userCredId)
	if err != nil {
		return nil, err
	}

	exist, err := s.orderRepo.IsIdExist(ctx, o.Id)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, sentinel.ErrOrderNotFound
	}

	userId, err := s.orderRepo.GetUserIdFromOrder(ctx, o.Id)
	if err != nil {
		return nil, err
	}
	if user.Id != userId {
		return nil, sentinel.ErrOrderUserNotAssociated
	}

	statusId, err := s.orderRepo.GetStatusIdByOrderId(ctx, o.Id)
	if err != nil {
		return nil, err
	}
	if statusId > constant.STATUS_WAITING_FOR_PAYMENT {
		return nil, sentinel.ErrOrderCancelUser
	}

	var order *entity.Order
	err = s.transactorRepo.WithTransaction(ctx, func(txCtx context.Context) error {
		o.StatusId = constant.STATUS_CANCELED
		order, err = s.orderRepo.UpdateOrder(txCtx, o)
		if err != nil {
			return err
		}

		orderItems, err := s.orderRepo.GetOrderItems(txCtx, o.Id)
		if err != nil {
			return err
		}

		for _, item := range orderItems {
			err = s.catalogRepo.UpdateStock(txCtx, item.Catalog.Id, item.Quantity)
			if err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (s *OrderServiceImpl) UpdateUserOrderConfirmedService(ctx context.Context, o *entity.Order, userCredId int64) (*entity.Order, error) {
	user, err := s.userRepo.GetUserByCredId(ctx, userCredId)
	if err != nil {
		return nil, err
	}

	exist, err := s.orderRepo.IsIdExist(ctx, o.Id)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, sentinel.ErrOrderNotFound
	}

	userId, err := s.orderRepo.GetUserIdFromOrder(ctx, o.Id)
	if err != nil {
		return nil, err
	}
	if user.Id != userId {
		return nil, sentinel.ErrOrderUserNotAssociated
	}

	statusId, err := s.orderRepo.GetStatusIdByOrderId(ctx, o.Id)
	if err != nil {
		return nil, err
	}
	if statusId != constant.STATUS_SENT {
		return nil, sentinel.ErrOrderConfirmUser
	}

	o.StatusId = constant.STATUS_CONFIRMED
	order, err := s.orderRepo.UpdateOrder(ctx, o)
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (s *OrderServiceImpl) UpdatePharmacyOrderShippedService(ctx context.Context, o *entity.Order, pharmacistCredId int64) (*entity.Order, error) {
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

	exist, err := s.orderRepo.IsIdExist(ctx, o.Id)
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

	orderPharmacyId, err := s.orderRepo.GetPharmacyIdByOrderId(ctx, o.Id)
	if err != nil {
		return nil, err
	}
	if pharmacyId != orderPharmacyId {
		return nil, sentinel.ErrPharmacistNoAccess
	}

	statusId, err := s.orderRepo.GetStatusIdByOrderId(ctx, o.Id)
	if err != nil {
		return nil, err
	}
	if statusId >= constant.STATUS_SENT {
		return nil, sentinel.ErrStatusBackward
	}

	o.StatusId = constant.STATUS_SENT
	order, err := s.orderRepo.UpdateOrder(ctx, o)
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (s *OrderServiceImpl) UpdatePharmacyOrderCanceledService(ctx context.Context, o *entity.Order, pharmacistCredId int64) (*entity.Order, error) {
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

	exist, err := s.orderRepo.IsIdExist(ctx, o.Id)
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

	orderPharmacyId, err := s.orderRepo.GetPharmacyIdByOrderId(ctx, o.Id)
	if err != nil {
		return nil, err
	}
	if pharmacyId != orderPharmacyId {
		return nil, sentinel.ErrPharmacistNoAccess
	}

	statusId, err := s.orderRepo.GetStatusIdByOrderId(ctx, o.Id)
	if err != nil {
		return nil, err
	}
	if statusId >= constant.STATUS_SENT {
		return nil, sentinel.ErrOrderCancel
	}

	var order *entity.Order
	err = s.transactorRepo.WithTransaction(ctx, func(txCtx context.Context) error {
		o.StatusId = constant.STATUS_CANCELED
		order, err = s.orderRepo.UpdateOrder(txCtx, o)
		if err != nil {
			return err
		}

		orderItems, err := s.orderRepo.GetOrderItems(txCtx, o.Id)
		if err != nil {
			return err
		}

		for _, item := range orderItems {
			err = s.catalogRepo.UpdateStock(txCtx, item.Catalog.Id, item.Quantity)
			if err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return order, nil
}
