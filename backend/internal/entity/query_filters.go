package entity

import (
	"log"
	"net/url"
	"strconv"
	"time"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

// TODO: вынести из entity, этому тут не место
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
	WithRandom         bool
	CcalMin            uint64
	CcalMax            uint64
	CreatedBefore      time.Time
	CreatedAfter       time.Time
	CategoryUid        uuid.UUID
	Name               string
	HouseNumber        string
	LimitOnEach        uint64
	OffsetOnEach       uint64
	CookingTime        int64
	RecipeUid          uuid.UUID
	OrderUid           uuid.UUID
	UserUid            uuid.UUID
	UserUidForOrder    uuid.UUID
	CityUid            uuid.UUID
	CategoryUids       []uuid.UUID
	CardUid            uuid.UUID
	OrdersUids         []uuid.UUID
}

const defaultLimit = 10

const (
	QueryFieldLimit             = "limit"
	QueryFieldLimitOnCategories = "limit_on_categories"
	QueryFieldLimitOnProducts   = "limit_on_products"
	QueryFieldPage              = "page"
	QueryFieldWithCounts        = "with_counts"
	QueryFieldWithPhotos        = "with_photos"
	QueryFieldCcalMin           = "ccal_min"
	QueryFieldCcalMax           = "ccal_max"
	QueryFieldCreatedBefore     = "created_before"
	QueryFieldCreatedAfter      = "created_after"
	QueryFieldCategoryUid       = "category_uid"
	QueryFieldName              = "name"
	QueryFieldCookingTime       = "cooking_time"
	QueryFieldRecipeUid         = "recipe_uid"
	QueryFieldWithRandom        = "with_random"
	QueryFieldUserUid           = "user_uid"
	QueryFieldCityUid           = "city_uid"
	QueryFieldCategoryUids      = "category_uids"
	QueryFieldOrderUid          = "order_uid"
	QueryFieldCardUid           = "card_uid"
	QueryFieldOrdersUids        = "orders_uids"
	QueryFieldUserUidForOrder   = "user_uid_for_order"
	QueryFieldHouseNumber       = "house_number"
)

var queryFiltersFields = []string{
	QueryFieldLimit,
	QueryFieldLimitOnCategories,
	QueryFieldLimitOnProducts,
	QueryFieldPage,
	QueryFieldWithCounts,
	QueryFieldWithPhotos,
	QueryFieldCcalMin,
	QueryFieldCcalMax,
	QueryFieldCreatedBefore,
	QueryFieldCreatedAfter,
	QueryFieldCategoryUid,
	QueryFieldCategoryUids,
	QueryFieldName,
	QueryFieldCookingTime,
	QueryFieldWithRandom,
	QueryFieldRecipeUid,
	QueryFieldUserUid,
	QueryFieldOrderUid,
	QueryFieldCardUid,
	QueryFieldOrdersUids,
	QueryFieldUserUidForOrder,
	QueryFieldCityUid,
	QueryFieldHouseNumber,
}

type QueryFiltersParser struct {
	RequiredFields []string
	AllowedFields  map[string]struct{}
	fieldsParsers  map[string]func(url.Values, *QueryFilters) error
}

var fieldsParsers = map[string]func(url.Values, *QueryFilters) error{
	QueryFieldLimit:             parseLimit,
	QueryFieldLimitOnCategories: parseLimitOnCategories,
	QueryFieldLimitOnProducts:   parseLimitOnProducts,
	QueryFieldPage:              parsePage,
	QueryFieldWithCounts:        parseWithCounts,
	QueryFieldWithPhotos:        parseWithPhotos,
	QueryFieldCcalMin:           parseCcalMin,
	QueryFieldCcalMax:           parseCcalMax,
	QueryFieldCreatedBefore:     parseCreatedBefore,
	QueryFieldCreatedAfter:      parseCreatedAfter,
	QueryFieldCategoryUid:       parseCategoryUid,
	QueryFieldName:              parseName,
	QueryFieldCookingTime:       parseCookingTime,
	QueryFieldRecipeUid:         parseRecipeUid,
	QueryFieldWithRandom:        parseWithRandom,
	QueryFieldUserUidForOrder:   parseUserUidForOrder,
	QueryFieldCategoryUids:      parseCategoryUids,
	QueryFieldCityUid:           parseCityUid,
	QueryFieldHouseNumber:       parseHouseNumber,
	QueryFieldOrderUid:          parseOrderUid,
	QueryFieldCardUid:           parseCardUid,
	QueryFieldOrdersUids:        parseOrdersUids,
	QueryFieldUserUid:           parseUserUid,
}

func NewQueryFiltersParser() *QueryFiltersParser {
	return &QueryFiltersParser{
		fieldsParsers: fieldsParsers,
	}
}

func (q *QueryFiltersParser) WithRequired(fields ...string) *QueryFiltersParser {
	q.RequiredFields = append(q.RequiredFields, fields...)
	return q
}

func (q *QueryFiltersParser) WithAllowed(fields ...string) *QueryFiltersParser {
	if len(q.AllowedFields) == 0 {
		q.AllowedFields = make(map[string]struct{}, len(fields))
	}

	for _, field := range fields {
		q.AllowedFields[field] = struct{}{}
	}
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

	for _, field := range queryFiltersFields {
		if len(q.AllowedFields) != 0 {
			if _, ok := q.AllowedFields[field]; !ok {
				continue
			}
		}

		if _, ok := parsed[field]; ok {
			continue
		}

		parseField, ok := q.fieldsParsers[field]
		if !ok {
			log.Printf("parser for query field %s not found", field)
			continue
		}

		parseField(query, &queryFilters)
		// if err := parseField(query, &queryFilters); err != nil {
		// 	// log.Printf("WARN: failed to parse field %s: %v", field, err)
		// }
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
	defer func() {
		if qFilters.Page == 0 {
			qFilters.Page = 1
		}
	}()

	if err := parseUint64(query, QueryFieldPage, &qFilters.Page); err != nil {
		return err
	}
	return nil
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

func parseRecipeUid(query url.Values, qFilters *QueryFilters) error {
	var err error
	qFilters.RecipeUid, err = uuid.FromString(query.Get(QueryFieldRecipeUid))
	if err != nil {
		return err
	}
	return nil
}

func parseCityUid(query url.Values, qFilters *QueryFilters) error {
	var err error
	qFilters.CityUid, err = uuid.FromString(query.Get(QueryFieldCityUid))
	if err != nil {
		return err
	}
	return nil
}

func parseName(query url.Values, qFilters *QueryFilters) error {
	qFilters.Name = query.Get(QueryFieldName)
	return nil
}

func parseHouseNumber(query url.Values, qFilters *QueryFilters) error {
	qFilters.HouseNumber = query.Get(QueryFieldHouseNumber)
	return nil
}

func parseCookingTime(query url.Values, qFilters *QueryFilters) error {
	return parseInt64(query, QueryFieldCookingTime, &qFilters.CookingTime)
}

func parseWithRandom(query url.Values, qFilters *QueryFilters) error {
	return parseBool(query, QueryFieldWithRandom, &qFilters.WithRandom)
}

func parseUserUidForOrder(query url.Values, qFilters *QueryFilters) error {
	var err error
	qFilters.UserUidForOrder, err = uuid.FromString(query.Get(QueryFieldUserUidForOrder))
	if err != nil {
		return err
	}
	return nil
}

func parseUserUid(query url.Values, qFilters *QueryFilters) error {
	var err error
	qFilters.UserUid, err = uuid.FromString(query.Get(QueryFieldUserUid))
	if err != nil {
		return err
	}
	return nil
}

func parseCategoryUids(query url.Values, qFilters *QueryFilters) error {
	uuidsStrings, ok := query[QueryFieldCategoryUids]

	if !ok {
		return errors.Errorf("field %s not found in query", QueryFieldCategoryUids)
	}

	for _, uuidString := range uuidsStrings {
		categoryUid, err := uuid.FromString(uuidString)

		if err != nil {
			return err
		}

		qFilters.CategoryUids = append(qFilters.CategoryUids, categoryUid)
	}

	return nil
}

func parseOrderUid(query url.Values, qFilters *QueryFilters) error {
	var err error
	qFilters.OrderUid, err = uuid.FromString(query.Get(QueryFieldOrderUid))
	if err != nil {
		return err
	}
	return nil
}

func parseOrdersUids(query url.Values, qFilters *QueryFilters) error {
	uuidsStrings, ok := query[QueryFieldOrdersUids]
	if !ok {
		return errors.Errorf("field %s not found in query", QueryFieldOrdersUids)
	}

	for _, uuidString := range uuidsStrings {
		orderUid, err := uuid.FromString(uuidString)
		if err != nil {
			return err
		}
		qFilters.OrdersUids = append(qFilters.OrdersUids, orderUid)
	}
	return nil
}

func parseCardUid(query url.Values, qFilters *QueryFilters) error {
	var err error
	qFilters.CardUid, err = uuid.FromString(query.Get(QueryFieldCardUid))
	if err != nil {
		return err
	}
	return nil
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
