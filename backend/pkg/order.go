package pkg

import (
	"angrymiao-ai/pkg/snowflake"
	"strconv"
)

// GenerateOrderNo 生成唯一订单号
func GenerateOrderNo() string {
	now := UTC8Now().Format("20060102")
	snowFlakeID := snowflake.GetWorker().GetID()

	return now + strconv.FormatInt(snowFlakeID, 10)
}
