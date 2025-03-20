package config

import (
	"os"
	"strconv"
)

type ConfigEnv struct {
	orderMinutesTTL              int64
	ordersCleanupIntervalMinutes int64
}

func NewConfigEnv() *ConfigEnv {

	orderTtl, err := strconv.Atoi(os.Getenv("ORDER_TTL_MINUTES"))
	if err != nil {
		panic(err)
	}
	ordersCleanupInterval, err := strconv.Atoi(os.Getenv("ORDERS_CLEANUP_INTERVAL_MINUTES"))
	if err != nil {
		panic(err)
	}

	return &ConfigEnv{
		orderMinutesTTL:              int64(orderTtl),
		ordersCleanupIntervalMinutes: int64(ordersCleanupInterval),
	}
}

func (c *ConfigEnv) OrderMinutesTTL() int64 {
	return c.orderMinutesTTL
}

func (c *ConfigEnv) OrdersCleanupIntervalMinutes() int64 {
	return c.ordersCleanupIntervalMinutes
}
