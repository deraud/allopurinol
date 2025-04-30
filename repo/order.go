package repo

import (
	"context"
	"database/sql"
	"errors"
	"healthcare-allopurinol/constant"
	"healthcare-allopurinol/entity"

	"github.com/huandu/go-sqlbuilder"
)

type OrderRepoItf interface {
	InsertOrderGroup(ctx context.Context, o *entity.OrderGroup) (*entity.OrderGroup, error)
	InsertOrder(ctx context.Context, o *entity.Order) (*entity.Order, error)
	InsertOrderItemBulk(ctx context.Context, items []*entity.OrderItem) ([]*entity.OrderItem, error)
	UpdateOrderGroup(ctx context.Context, o *entity.OrderGroup) (*entity.OrderGroup, error)
	UpdateOrderGroupStatusProccessed(ctx context.Context, id int64) error
	IsOrderGroupIdExist(ctx context.Context, id int64) (bool, error)
	GetUserIdFromOrderGroup(ctx context.Context, id int64) (int64, error)
	GetPendingOrders(ctx context.Context, options *entity.PendingOrderOptions, userId int64) ([]*entity.PendingOrderGroup, error)
	GetPendingOrdersCount(ctx context.Context, userId int64) (int, error)
	UpdateOrder(ctx context.Context, o *entity.Order) (*entity.Order, error)
	GetPharmacyOrders(ctx context.Context, options *entity.PharmacyOrderOptions, pharmacyId int64) ([]*entity.Order, error)
	GetPharmacyOrdersCount(ctx context.Context, options *entity.PharmacyOrderOptions, pharmacyId int64) (int, error)
	GetOrder(ctx context.Context, id int64) (*entity.Order, error)
	GetOrderItems(ctx context.Context, orderId int64) ([]*entity.OrderItem, error)
	GetPharmacyIdByOrderId(ctx context.Context, orderId int64) (int64, error)
	IsIdExist(ctx context.Context, id int64) (bool, error)
	GetStatusIdByOrderId(ctx context.Context, orderId int64) (int64, error)
	GetOrders(ctx context.Context, options *entity.OrderOptions) ([]*entity.Order, error)
	GetOrdersCount(ctx context.Context, options *entity.OrderOptions) (int, error)
	GetUserIdFromOrder(ctx context.Context, id int64) (int64, error)
	GetUserOrders(ctx context.Context, options *entity.UserOrderOptions, userId int64) ([]*entity.Order, error)
	GetUserOrdersCount(ctx context.Context, options *entity.UserOrderOptions, userId int64) (int, error)
	UpdateOrderGroupStatusVerifying(ctx context.Context, id int64) error
}

type OrderRepoImpl struct {
	db *sql.DB
}

func NewOrderRepo(database *sql.DB) OrderRepoItf {
	return &OrderRepoImpl{
		db: database,
	}
}

func (r *OrderRepoImpl) IsOrderGroupIdExist(ctx context.Context, id int64) (bool, error) {
	var exist bool

	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("1").
		From("order_groups").
		Where(
			sb.Equal("id", id),
			sb.IsNull("deleted_at"),
		)

	query, args := sb.BuildWithFlavor(constant.DATABASE_FLAVOR)
	err := r.db.QueryRowContext(ctx, query, args...).Scan(&exist)
	if errors.Is(err, sql.ErrNoRows) {
		return false, nil
	}
	if err != nil {
		return false, err
	}

	return exist, nil
}

func (r *OrderRepoImpl) GetUserIdFromOrderGroup(ctx context.Context, id int64) (int64, error) {
	var userId int64

	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("user_id").
		From("order_groups").
		Where(
			sb.Equal("id", id),
			sb.IsNull("deleted_at"),
		)

	query, args := sb.BuildWithFlavor(constant.DATABASE_FLAVOR)
	err := r.db.QueryRowContext(ctx, query, args...).Scan(&userId)
	if err != nil {
		return 0, err
	}

	return userId, nil
}

func (r *OrderRepoImpl) GetPharmacyIdByOrderId(ctx context.Context, orderId int64) (int64, error) {
	var pharmacyId int64

	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("pharmacy_id").
		From("orders").
		Where(
			sb.Equal("id", orderId),
			sb.IsNull("deleted_at"),
		)

	query, args := sb.BuildWithFlavor(constant.DATABASE_FLAVOR)
	err := r.db.QueryRowContext(ctx, query, args...).Scan(&pharmacyId)
	if errors.Is(err, sql.ErrNoRows) {
		return 0, nil
	}
	if err != nil {
		return 0, err
	}

	return pharmacyId, nil
}

func (r *OrderRepoImpl) IsIdExist(ctx context.Context, id int64) (bool, error) {
	var exist bool

	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("1").
		From("orders").
		Where(
			sb.Equal("id", id),
			sb.IsNull("deleted_at"),
		)

	query, args := sb.BuildWithFlavor(constant.DATABASE_FLAVOR)
	err := r.db.QueryRowContext(ctx, query, args...).Scan(&exist)
	if errors.Is(err, sql.ErrNoRows) {
		return false, nil
	}
	if err != nil {
		return false, err
	}

	return exist, nil
}

func (r *OrderRepoImpl) GetStatusIdByOrderId(ctx context.Context, orderId int64) (int64, error) {
	var statusId int64

	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("status_id").
		From("orders").
		Where(
			sb.Equal("id", orderId),
			sb.IsNull("deleted_at"),
		)

	query, args := sb.BuildWithFlavor(constant.DATABASE_FLAVOR)
	err := r.db.QueryRowContext(ctx, query, args...).Scan(&statusId)
	if errors.Is(err, sql.ErrNoRows) {
		return 0, nil
	}
	if err != nil {
		return 0, err
	}

	return statusId, nil
}

func (r *OrderRepoImpl) GetUserIdFromOrder(ctx context.Context, id int64) (int64, error) {
	var userId int64

	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("user_id").
		From("orders").
		Where(
			sb.Equal("id", id),
			sb.IsNull("deleted_at"),
		)

	query, args := sb.BuildWithFlavor(constant.DATABASE_FLAVOR)
	err := r.db.QueryRowContext(ctx, query, args...).Scan(&userId)
	if err != nil {
		return 0, err
	}

	return userId, nil
}
