// coding: utf-8
// @Author : lryself
// @Date : 2021/5/13 1:31
// @Software: GoLand

package snowflake

import (
	"gitee.com/lryself/go-utils/snowflake"
	"github.com/spf13/viper"
	"strconv"
	"sync"
	"tem_go_project/internal/globals"
)

var (
	worker     *snowflake.Worker
	workerOnce sync.Once
	log        = globals.GetLogger()
)

func GetSnowflakeID() string {
	workerOnce.Do(func() {
		workerID := viper.GetInt64("system.WorkerID")
		dataCenterID := viper.GetInt64("system.DataCenterID")
		worker = snowflake.NewWorker(workerID, dataCenterID)
		log.Infoln("雪花算法ID生成初始化服务完成!")
	})
	id, err := worker.NextID()
	if err != nil {
		log.Errorln("生成id出错 ", err)
	}
	return strconv.FormatUint(id, 10)
}
