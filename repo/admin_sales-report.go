package repo

import (
	"context"
	"fmt"
	"healthcare-allopurinol/constant"
	"healthcare-allopurinol/entity"

	"github.com/huandu/go-sqlbuilder"
)

// select p.name, coalesce(SUM(c.price * oi.quantity),0), TO_CHAR(o.updated_at , 'DD/MM/YYYY') from order_items oi
// join orders o on o.id = oi.order_id
// join catalogs c on oi.catalog_id = c.id
// join products p on c.product_id = p.id
// where o.status_id = 4 and p.name ilike '%%'
// and o.created_at < '2025-02-1'::date and
// 	o.created_at >= '2024-01-1'::date
// group by p.name, TO_CHAR(o.updated_at , 'DD/MM/YYYY')

func (r *AdminRepoImpl) GetSalesByProductName(ctx context.Context, options *entity.SalesByNameOptions) error {
	sb := sqlbuilder.NewSelectBuilder()

	sb.Select("p.name, coalesce(SUM(c.price * oi.quantity),0)").
		From("order_items oi").
		Join("orders o", "o.id = oi.order_id").
		Join("catalogs c", "oi.catalog_id = c.id").
		Join("products p", "c.product_id = p.id").
		AddWhereClause(r.generateSalesByProductNameWhereClause(options)).
		GroupBy(fmt.Sprintf("p.name, TO_CHAR(o.updated_at , 'DD/MM/YYYY'"))
	return nil
}

func (r *AdminRepoImpl) generateSalesByProductNameWhereClause(options *entity.SalesByNameOptions) *sqlbuilder.WhereClause {
	whereClause := sqlbuilder.NewWhereClause()
	cond := sqlbuilder.NewCond()

	whereClause.AddWhereExpr(
		cond.Args,
		cond.Equal("o.status_id", 4),
		cond.ILike("p.name", fmt.Sprintf("%%%s%%", options.SearchValue)),
	)

	return whereClause
}

func (r *AdminRepoImpl) GetSalesReport(ctx context.Context, options *entity.ReportOptionsRequest, count int64) ([]*entity.Report, error) {
	reports := make([]*entity.Report, count)

	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("p.id", "p.name",
		"sum(case when status_id = 4 then o.total_price_product else 0 end) as sales").
		From("orders o").
		JoinWithOption(sqlbuilder.FullJoin, "pharmacies p", "o.pharmacy_id=p.id ").
		Where(sb.ILike("p.name", fmt.Sprintf("%%%s%%", options.SearchValue))).
		AddWhereClause(r.generateGetSalesReportWhereClause(options)).
		GroupBy("p.name, p.id").
		OrderBy(fmt.Sprintf("%s %s", "sales", options.SortOrder))

	query, args := sb.BuildWithFlavor(constant.DATABASE_FLAVOR)
	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for i := 0; rows.Next(); i++ {
		var report entity.Report
		err := rows.Scan(
			&report.PharmacyId,
			&report.PharmacyName,
			&report.Sales,
		)
		if err != nil {
			return nil, err
		}
		reports[i] = &report
		fmt.Println(&report)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return reports, nil
}

func (r *AdminRepoImpl) GetReportCountBySearchValue(ctx context.Context, options *entity.ReportOptionsRequest) (int64, error) {
	var count int64

	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("count(id)").
		From("pharmacies p").
		Where(sb.IsNull("deleted_at")).
		AddWhereClause(r.generateGetSalesReportWhereClause(options))

	query, args := sb.BuildWithFlavor(constant.DATABASE_FLAVOR)
	err := r.db.QueryRowContext(ctx, query, args...).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *AdminRepoImpl) generateGetSalesReportWhereClause(options *entity.ReportOptionsRequest) *sqlbuilder.WhereClause {
	whereClause := sqlbuilder.NewWhereClause()
	cond := sqlbuilder.NewCond()

	whereClause.AddWhereExpr(
		cond.Args,
		cond.ILike("p.name", fmt.Sprintf("%%%s%%", options.SearchValue)),
	)

	return whereClause
}
