package collect

import (
	"github.com/ziutek/mymysql/mysql"
	_ "github.com/ziutek/mymysql/native"

	"github.com/mysql-monitor/base"
	"github.com/mysql-monitor/model"
)

func ProcessNum(db mysql.Conn, cfg *base.Cfg) (*model.MetaData, error) {
	sql := "SHOW PROCESSLIST"
	rows, _, err := db.Query(sql)
	if err != nil {
		return nil, err
	}
	monitorName := "process_nums"
	threadNum := len(rows) - 1
	data := model.NewMetric(monitorName, cfg)
	data.SetValue(threadNum)
	return data, nil
}
