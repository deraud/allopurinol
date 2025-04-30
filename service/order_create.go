package service

import (
	"context"
	"healthcare-allopurinol/constant"
	"healthcare-allopurinol/entity"

	"github.com/shopspring/decimal"
)

func (s *OrderServiceImpl) CreateOrderGroupService(ctx context.Context, pharmacies []*entity.Pharmacy, catalogs map[int64][]*entity.Catalog, details *entity.CheckoutDetails, userCredId int64) (*entity.OrderGroup, error) {
	user, err := s.userRepo.GetUserByCredId(ctx, userCredId)
	if err != nil {
		return nil, err
	}

	//TODO validation
	// up to this point pharmacies is validated to be from cart
	// payment method (exist)
	// pharmacy match with logistic partner

	pharmacyLogisticPartner := make(map[int64]int64)
	for _, order := range details.OrderDetails {
		pharmacyLogisticPartner[order.PharmacyId] = order.LogisticPartnerId
	}

	orderGroup := &entity.OrderGroup{
		UserId: user.Id,
	}

	err = s.transactorRepo.WithTransaction(ctx, func(txCtx context.Context) error {
		orderGroup, err = s.orderRepo.InsertOrderGroup(txCtx, orderGroup)
		if err != nil {
			return err
		}

		orders := make([]*entity.Order, len(pharmacies))
		for i, p := range pharmacies {
			shippingData, err := s.redisRepo.GetShippingData(user.Id, p.Id)
			if err != nil {
				return err
			}

			var shippingCost decimal.Decimal
			for _, data := range shippingData.ShippingCosts {
				if data.MethodId == pharmacyLogisticPartner[p.Id] {
					shippingCost = data.Cost
					break
				}
			}

			productCost := decimal.NewFromInt(0)
			for _, c := range catalogs[p.Id] {
				sum := c.Price.Mul(decimal.NewFromInt(int64(c.Quantity)))
				productCost = productCost.Add(sum)
			}

			order := entity.Order{
				UserId:             user.Id,
				AddressId:          details.AddressId,
				StatusId:           constant.STATUS_WAITING_FOR_PAYMENT,
				PaymentMethodId:    details.PaymentMethodId,
				PharmacyId:         p.Id,
				LogisticPartnerId:  pharmacyLogisticPartner[p.Id],
				OrderGroupId:       orderGroup.Id,
				TotalPriceProduct:  productCost,
				TotalPriceShipping: shippingCost,
			}

			newOrder, err := s.orderRepo.InsertOrder(txCtx, &order)
			if err != nil {
				return err
			}

			var orderItems []*entity.OrderItem
			for _, c := range catalogs[p.Id] {
				orderItem := entity.OrderItem{
					OrderId:   newOrder.Id,
					CatalogId: c.Id,
					Quantity:  c.Quantity,
				}
				orderItems = append(orderItems, &orderItem)
				err = s.catalogRepo.UpdateStock(txCtx, c.Id, -1*c.Quantity)
				if err != nil {
					return err
				}
			}

			newOrderItems, err := s.orderRepo.InsertOrderItemBulk(txCtx, orderItems)
			if err != nil {
				return err
			}

			newOrder.OrderItems = newOrderItems
			orders[i] = newOrder

			err = s.cartRepo.DeleteCartItemByUserId(txCtx, user.Id)
			if err != nil {
				return err
			}
		}
		orderGroup.Orders = orders
		return nil
	})
	if err != nil {
		return nil, err
	}

	return orderGroup, nil
}
