package repo

import (
	"context"
	"healthcare-allopurinol/constant"
	"healthcare-allopurinol/entity"

	"github.com/huandu/go-sqlbuilder"
)

func (r *OrderRepoImpl) UpdateOrderGroup(ctx context.Context, o *entity.OrderGroup) (*entity.OrderGroup, error) {
	orderGroup := *o

	ub := sqlbuilder.NewUpdateBuilder()
	ub.Update("order_groups").
		Set(
			ub.Assign("proof", o.Proof),
			ub.Assign("updated_at", "NOW()"),
		).
		Where(
			ub.Equal("id", o.Id),
			ub.IsNull("deleted_at"),
		)

	query, args := ub.BuildWithFlavor(constant.DATABASE_FLAVOR)
	_, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	return &orderGroup, nil
}

func (r *OrderRepoImpl) UpdateOrderGroupStatusProccessed(ctx context.Context, id int64) error {
	ub := sqlbuilder.NewUpdateBuilder()
	ub.Update("orders").
		Set(
			ub.Assign("status_id", constant.STATUS_PROCCESSED),
			ub.Assign("updated_at", "NOW()"),
		).
		Where(
			ub.Equal("order_group_id", id),
			ub.Equal("status_id", constant.STATUS_VERIFYING),
			ub.IsNull("deleted_at"),
		)

	query, args := ub.BuildWithFlavor(constant.DATABASE_FLAVOR)
	_, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *OrderRepoImpl) UpdateOrderGroupStatusVerifying(ctx context.Context, id int64) error {
	ub := sqlbuilder.NewUpdateBuilder()
	ub.Update("orders").
		Set(
			ub.Assign("status_id", constant.STATUS_VERIFYING),
			ub.Assign("updated_at", "NOW()"),
		).
		Where(
			ub.Equal("order_group_id", id),
			ub.Equal("status_id", constant.STATUS_WAITING_FOR_PAYMENT),
			ub.IsNull("deleted_at"),
		)

	query, args := ub.BuildWithFlavor(constant.DATABASE_FLAVOR)
	_, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *OrderRepoImpl) UpdateOrder(ctx context.Context, o *entity.Order) (*entity.Order, error) {
	order := *o

	ub := sqlbuilder.NewUpdateBuilder()
	ub.Update("orders").
		Set(
			ub.Assign("status_id", o.StatusId),
			ub.Assign("updated_at", "NOW()"),
		).
		Where(
			ub.Equal("id", o.Id),
			ub.IsNull("deleted_at"),
		)

	tx := extractTx(ctx)
	query, args := ub.BuildWithFlavor(constant.DATABASE_FLAVOR)
	var err error
	if tx != nil {
		_, err = tx.ExecContext(ctx, query, args...)
	} else {
		_, err = r.db.ExecContext(ctx, query, args...)
	}
	if err != nil {
		return nil, err
	}

	return &order, nil
}