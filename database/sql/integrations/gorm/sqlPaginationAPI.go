package leafGorm

import (
	"context"
	leafSql "github.com/enricodg/leaf-utilities/database/sql/sql"
	leafFunctions "github.com/paulusrobin/leaf-utilities/common/functions"
	leafModel "github.com/paulusrobin/leaf-utilities/common/model"
	"github.com/paulusrobin/leaf-utilities/encoding/json"
)

func (i *Impl) PaginateData(ctx context.Context, items interface{}, options leafSql.PaginateOptions) (leafModel.BasePagingResponse, error) {
	res := leafModel.BasePagingResponse{}

	paging := leafModel.PagingParams{
		Page:   options.Page,
		Limit:  options.Limit,
		Sort:   options.Sort,
		Filter: options.Filter,
	}
	res.Params = paging
	res.CurrentPage = paging.Page

	var (
		ormCount leafSql.ORM = i.newImpl(i.GormDB)
		count                = int64(res.Count)
	)
	if err := ormCount.
		WithContext(context.Background()).
		Count(ctx, &count).
		Error(); err != nil {
		return leafModel.BasePagingResponse{}, err
	}

	var orm leafSql.ORM = i.newImpl(i.GormDB)
	sortQuery, arrSort := i.sort(paging.Sort, options)
	if len(sortQuery) > 0 {
		orm = i.Order(sortQuery)
	}

	res.Count = int(count)
	res.TotalPage = leafFunctions.CalculateTotalPages(res.Count, paging.Limit)

	offset := (paging.Page - 1) * paging.Limit
	if err := orm.
		WithContext(context.Background()).
		Offset(offset).
		Limit(paging.Limit).
		Find(ctx, items).
		Error(); err != nil {
		return leafModel.BasePagingResponse{}, err
	}
	res.Params.Sort = arrSort
	return res, nil
}

func (i *Impl) SimplePaginateData(ctx context.Context, items interface{}, options leafSql.PaginateOptions) (leafModel.BaseSimplePagingResponse, error) {
	res := leafModel.BaseSimplePagingResponse{}

	paging := leafModel.PagingParams{
		Page:   options.Page,
		Limit:  options.Limit,
		Sort:   options.Sort,
		Filter: options.Filter,
	}
	res.Params = paging
	res.CurrentPage = paging.Page

	var orm leafSql.ORM = i.newImpl(i.GormDB)
	sortQuery, arrSort := i.sort(paging.Sort, options)
	res.Params.Sort = arrSort
	if len(sortQuery) > 0 {
		orm = i.Order(sortQuery)
	}

	offset := (paging.Page - 1) * paging.Limit
	if err := orm.
		WithContext(context.Background()).
		Offset(offset).
		Limit(paging.Limit+1).
		Find(ctx, items).
		Error(); err != nil {
		return leafModel.BaseSimplePagingResponse{}, err
	}
	result, err := i.interfaceSlice(items)
	if len(result) > paging.Limit && err == nil {
		result = result[:len(result)-1]
		byteItem, _ := json.Marshal(result)
		_ = json.Unmarshal(byteItem, &items)
		res.Next = true
	}
	return res, err
}
