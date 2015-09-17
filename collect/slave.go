package collect

import (
	"fmt"

	"github.com/ziutek/mymysql/mysql"
	_ "github.com/ziutek/mymysql/native"

	"github.com/mysql-monitor/base"
	"github.com/mysql-monitor/model"
)

var SlaveStatusToSend = []string{
	"Exec_Master_Log_Pos",
	"Read_Master_Log_Pos",
	"Relay_Log_Pos",
	"Seconds_Behind_Master",
	"Slave_IO_Running",
	"Slave_SQL_Running",
}

func SlaveStatus(db mysql.Conn, cfg *base.Cfg) ([]*model.MetaData, error) {
	isSlave := model.NewMetric("Is_slave", cfg)
	row, res, err := db.QueryFirst("SHOW SLAVE STATUS")
	if err != nil {
		return nil, err
	}

	// be master
	if row == nil {
		isSlave.SetValue(0)
		isSlave.Tags = fmt.Sprintf("%s,role=%s", isSlave.Tags, "master")
		return []*model.MetaData{isSlave}, nil
	}

	// be slave
	isSlave.SetValue(1)
	isSlave.Tags = fmt.Sprintf("%s,role=%s", isSlave.Tags, "slaver")

	data := make([]*model.MetaData, len(SlaveStatusToSend))
	for i, s := range SlaveStatusToSend {
		data[i] = model.NewMetric(s, cfg)
		switch s {
		case "Slave_SQL_Running", "Slave_IO_Running":
			data[i].SetValue(0)
			v := row.Str(res.Map(s))
			if v == "Yes" {
				data[i].SetValue(1)
			}
		default:
			v, err := row.Int64Err(res.Map(s))
			if err != nil {
				data[i].SetValue(-1)
			} else {
				data[i].SetValue(v)
			}
		}
	}
	return append(data, isSlave), nil
}
