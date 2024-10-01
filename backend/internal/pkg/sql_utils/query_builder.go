package sqlUtils

type QueryBuilder struct {
	Filters []QueryFilter
	Order   QueryOrder
	Limit   uint64
	Offset  uint64
}

type QueryFilter struct {
	Field string
	Op    string
	Value interface{}
}

type QueryOrder struct {
	Field string
	Asc   bool
}
