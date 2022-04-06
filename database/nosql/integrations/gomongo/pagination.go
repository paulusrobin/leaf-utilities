package leafGoMongo

import (
	"context"
	leafFunctions "github.com/paulusrobin/leaf-utilities/common/functions"
	leafModel "github.com/paulusrobin/leaf-utilities/common/model"
	leafNoSql "github.com/paulusrobin/leaf-utilities/database/nosql/nosql"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strings"
)

func (i implementation) Paginate(ctx context.Context, items interface{}, paginateOptions leafNoSql.PaginateOptions) (leafModel.PagingResponse, error) {
	res := leafModel.PagingResponse{}

	paging := leafModel.PagingParams{
		Page:   paginateOptions.Page,
		Limit:  paginateOptions.Limit,
		Sort:   paginateOptions.Sort,
		Filter: paginateOptions.Filter,
	}
	res.CurrentPage = paging.Page
	res.Params = paging

	count, err := i.CountWithFilter(ctx, paginateOptions.Collection, paginateOptions.Filter, options.Count())
	if err != nil {
		return leafModel.PagingResponse{}, err
	}
	res.Count = int(count)

	res.TotalPage = leafFunctions.CalculateTotalPages(res.Count, paging.Limit)

	offset := int64((paging.Page - 1) * paging.Limit)
	limit := int64(paginateOptions.Limit)
	if err := i.FindAll(ctx, paginateOptions.Collection, paginateOptions.Filter, items, &options.FindOptions{
		Limit: &limit,
		Skip:  &offset,
		Sort:  i.sort(paging.Sort, paginateOptions),
	}); err != nil {
		return leafModel.PagingResponse{}, err
	}

	res.Items = items
	return res, nil
}

func (i implementation) sort(sort []string, paginateOptions leafNoSql.PaginateOptions) primitive.M {
	var mongoSort primitive.M

	for _, s := range sort {
		s = strings.TrimSpace(s)
		if len(s) < 1 {
			continue
		}

		sortDirection := 1
		if s[0] == '-' {
			sortDirection = -1
			s = s[1:]
		}

		s = i.getMappedField(s, paginateOptions)
		if len(s) < 1 {
			continue
		}

		if mongoSort == nil {
			mongoSort = primitive.M{}
		}
		mongoSort[s] = sortDirection
	}

	return mongoSort
}

func (i implementation) getMappedField(s string, paginateOptions leafNoSql.PaginateOptions) string {
	if paginateOptions.FieldMap == nil {
		return s
	}

	if paginateOptions.MapOrDefault {
		if mapped := paginateOptions.FieldMap[s]; len(mapped) < 1 {
			return mapped
		}

		return s
	}

	if mapped := paginateOptions.FieldMap[s]; len(mapped) > 0 {
		return mapped
	}

	return ""
}
