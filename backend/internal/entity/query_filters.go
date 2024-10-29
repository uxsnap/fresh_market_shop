package entity

import (
	"log"
	"net/url"
	"strconv"
	"time"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

type QueryFilters struct {
	Limit              uint64
	LimitOnCategories  uint64
	LimitOnProducts    uint64
	Page               uint64
	Offset             uint64
	OffsetOnCategories uint64
	OffsetOnProducts   uint64
	ProductsWithCount  bool
	ProductsWithPhotos bool
	WithCounts         bool
	WithPhotos         bool
	CcalMin            uint64
	CcalMax            uint64
	CreatedBefore      time.Time
	CreatedAfter       time.Time
	CategoryUid        uuid.UUID
	Name               string
	LimitOnEach        uint64
	OffsetOnEach       uint64
	CookingTime        int64
}

const defaultLimit = 10

const (
	QueryFieldLimit              = "limit"
	QueryFieldLimitOnCategories  = "limit_on_categories"
	QueryFieldLimitOnProducts    = "limit_on_products"
	QueryFieldPage               = "page"
	QueryFieldProductsWithCounts = "products_with_counts"
	QueryFieldProductsWithPhotos = "products_with_photos"
	QueryFieldWithCounts         = "with_counts"
	QueryFieldWithPhotos         = "with_photos"
	QueryFieldCcalMin            = "ccal_min"
	QueryFieldCcalMax            = "ccal_max"
	QueryFieldCreatedBefore      = "created_before"
	QueryFieldCreatedAfter       = "created_after"
	QueryFieldCategoryUid        = "category_uid"
	QueryFieldName               = "name"
	QueryFieldCookingTime        = "cooking_time"
)

type QueryFiltersParser struct {
	RequiredFields []string
	fieldsParsers  map[string]func(url.Values, *QueryFilters) error
}

func NewQueryFiltersParser() *QueryFiltersParser {
	return &QueryFiltersParser{
		fieldsParsers: map[string]func(url.Values, *QueryFilters) error{
			QueryFieldLimit:              parseLimit,
			QueryFieldLimitOnCategories:  parseLimitOnCategories,
			QueryFieldLimitOnProducts:    parseLimitOnProducts,
			QueryFieldPage:               parsePage,
			QueryFieldProductsWithCounts: parseProductsWithCounts,
			QueryFieldProductsWithPhotos: parseProductsWithPhotos,
			QueryFieldWithCounts:         parseWithCounts,
			QueryFieldWithPhotos:         parseWithPhotos,
			QueryFieldCcalMin:            parseCcalMin,
			QueryFieldCcalMax:            parseCcalMax,
			QueryFieldCreatedBefore:      parseCreatedBefore,
			QueryFieldCreatedAfter:       parseCreatedAfter,
			QueryFieldCategoryUid:        parseCategoryUid,
			QueryFieldName:               parseName,
			QueryFieldCookingTime:        parseCookingTime,
		},
	}
}

func (q *QueryFiltersParser) WithRequired(fields ...string) *QueryFiltersParser {
	q.RequiredFields = append(q.RequiredFields, fields...)
	return q
}

func (q *QueryFiltersParser) ParseQuery(query url.Values) (QueryFilters, error) {
	parsed := make(map[string]struct{}, 15)
	queryFilters := QueryFilters{}

	for _, requiredField := range q.RequiredFields {
		parseField, ok := q.fieldsParsers[requiredField]
		if !ok {
			return QueryFilters{}, errors.Errorf("unknow required query field %s", requiredField)
		}

		if err := parseField(query, &queryFilters); err != nil {
			return QueryFilters{}, errors.Wrapf(err, "cant parse required field %s", requiredField)
		}
		parsed[requiredField] = struct{}{}
	}

	queryFiltersFields := []string{
		QueryFieldLimit,
		QueryFieldLimitOnCategories,
		QueryFieldLimitOnProducts,
		QueryFieldPage,
		QueryFieldProductsWithCounts,
		QueryFieldProductsWithPhotos,
		QueryFieldWithCounts,
		QueryFieldWithPhotos,
		QueryFieldCcalMin,
		QueryFieldCcalMax,
		QueryFieldCreatedBefore,
		QueryFieldCreatedAfter,
		QueryFieldCategoryUid,
		QueryFieldName,
		QueryFieldCookingTime,
	}

	for _, field := range queryFiltersFields {
		if _, ok := parsed[field]; ok {
			continue
		}

		parseField, ok := q.fieldsParsers[field]
		if !ok {
			log.Printf("parser for query field %s not found", field)
			continue
		}

		if err := parseField(query, &queryFilters); err != nil {
			log.Printf("WARN: failed to parse field %s: %v", field, err)
		}
	}

	queryFilters.Offset = (queryFilters.Page - 1) * queryFilters.Limit
	queryFilters.OffsetOnCategories = (queryFilters.Page - 1) * queryFilters.LimitOnCategories
	queryFilters.OffsetOnProducts = (queryFilters.Page - 1) * queryFilters.LimitOnProducts

	return queryFilters, nil
}

func parseLimit(query url.Values, qFilters *QueryFilters) error {
	if err := parseUint64(query, QueryFieldLimit, &qFilters.Limit); err != nil {
		return err
	}
	if qFilters.Limit == 0 {
		qFilters.Limit = defaultLimit
	}
	return nil
}

func parseLimitOnCategories(query url.Values, qFilters *QueryFilters) error {
	if err := parseUint64(query, QueryFieldLimitOnCategories, &qFilters.LimitOnCategories); err != nil {
		return err
	}
	if qFilters.LimitOnCategories == 0 {
		qFilters.LimitOnCategories = defaultLimit
	}
	return nil
}

func parseLimitOnProducts(query url.Values, qFilters *QueryFilters) error {
	if err := parseUint64(query, QueryFieldLimitOnProducts, &qFilters.LimitOnProducts); err != nil {
		return err
	}
	if qFilters.LimitOnProducts == 0 {
		qFilters.LimitOnProducts = defaultLimit
	}
	return nil
}

func parsePage(query url.Values, qFilters *QueryFilters) error {
	if err := parseUint64(query, QueryFieldPage, &qFilters.Page); err != nil {
		return err
	}
	if qFilters.Page == 0 {
		qFilters.Page = 1
	}
	return nil
}

func parseProductsWithCounts(query url.Values, qFilters *QueryFilters) error {
	return parseBool(query, QueryFieldProductsWithCounts, &qFilters.ProductsWithCount)
}

func parseProductsWithPhotos(query url.Values, qFilters *QueryFilters) error {
	return parseBool(query, QueryFieldProductsWithPhotos, &qFilters.ProductsWithPhotos)
}

func parseWithCounts(query url.Values, qFilters *QueryFilters) error {
	return parseBool(query, QueryFieldWithCounts, &qFilters.WithCounts)
}

func parseWithPhotos(query url.Values, qFilters *QueryFilters) error {
	return parseBool(query, QueryFieldWithPhotos, &qFilters.WithPhotos)
}

func parseCcalMin(query url.Values, qFilters *QueryFilters) error {
	return parseUint64(query, QueryFieldCcalMin, &qFilters.CcalMin)
}

func parseCcalMax(query url.Values, qFilters *QueryFilters) error {
	return parseUint64(query, QueryFieldCcalMax, &qFilters.CcalMax)
}

func parseCreatedBefore(query url.Values, qFilters *QueryFilters) error {
	return parseDate(query, QueryFieldCreatedBefore, &qFilters.CreatedBefore)
}

func parseCreatedAfter(query url.Values, qFilters *QueryFilters) error {
	return parseDate(query, QueryFieldCreatedAfter, &qFilters.CreatedAfter)
}

func parseCategoryUid(query url.Values, qFilters *QueryFilters) error {
	var err error
	qFilters.CategoryUid, err = uuid.FromString(query.Get(QueryFieldCategoryUid))
	if err != nil {
		return err
	}
	return nil
}

func parseName(query url.Values, qFilters *QueryFilters) error {
	qFilters.Name = query.Get(QueryFieldName)
	return nil
}

func parseCookingTime(query url.Values, qFilters *QueryFilters) error {
	return parseInt64(query, QueryFieldCookingTime, &qFilters.CookingTime)
}

//-----COMMON-----//

func parseBool(query url.Values, fieldName string, dest *bool) error {
	reqField := query.Get(fieldName)
	if len(reqField) == 0 {
		return errors.Errorf("field %s not found in query", fieldName)
	}
	var err error
	*dest, err = strconv.ParseBool(reqField)
	if err != nil {
		return err
	}
	return nil
}

func parseUint64(query url.Values, fieldName string, dest *uint64) error {
	reqField := query.Get(fieldName)
	if len(reqField) == 0 {
		return errors.Errorf("field %s not found in query", fieldName)
	}
	var err error
	*dest, err = strconv.ParseUint(reqField, 10, 64)
	if err != nil {
		return err
	}
	return nil
}

func parseInt64(query url.Values, fieldName string, dest *int64) error {
	reqField := query.Get(fieldName)
	if len(reqField) == 0 {
		return errors.Errorf("field %s not found in query", fieldName)
	}
	var err error
	*dest, err = strconv.ParseInt(reqField, 10, 64)
	if err != nil {
		return err
	}
	return nil
}

func parseDate(query url.Values, fieldName string, dest *time.Time) error {
	reqField := query.Get(fieldName)
	if len(reqField) == 0 {
		return errors.Errorf("field %s not found in query", fieldName)
	}
	var err error
	*dest, err = time.Parse("2006-01-02T15:04:05", reqField)
	if err != nil {
		return err
	}
	return nil
}
