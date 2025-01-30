// package mq 消息队列初始化模块
package mq

import (
	"OJ/pkg/global"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"github.com/gofiber/fiber/v2/log"
)

// InitRocketMQ 初始化消息队列生产者
func InitRocketMQ() {
	// 配置生产者（根据实际情况修改配置参数）
	p, err := rocketmq.NewProducer(
		// 设置 NameServer 地址（必填）
		producer.WithNameServer([]string{"127.0.0.1:9876"}),
		// 设置生产者组（必填）
		producer.WithGroupName("Your_Producer_Group"),
		// 设置重试次数（可选）
		producer.WithRetry(2),
		// 设置命名空间（可选，根据服务端配置）
		// rocketmq.WithNamespace("Your_Namespace"),
	)

	if err != nil {
		log.Errorf("创建 RocketMQ 生产者失败: %v", err)
		panic("RocketMQ 生产者初始化失败") // 根据需求决定是否终止程序
	}
	global.RocketMQ = p
}
