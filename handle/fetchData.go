package handle

import (
	"time"

	"github.com/mysql-monitor/base"
	"github.com/mysql-monitor/collect"
	"github.com/mysql-monitor/model"

	log "github.com/Sirupsen/logrus"
	"github.com/ziutek/mymysql/mysql"
	_ "github.com/ziutek/mymysql/native"
)

func MysqlAlive(cfg *base.Cfg, m *model.MysqlIns, ok bool) {
	data := model.NewMetric("mysql_alive_local", cfg)
	data.SetValue(1)
	if !ok {
		data.SetValue(0)
	}
	msg, err := SendData([]*model.MetaData{data}, cfg)
	if err != nil {
		log.Errorf("Send alive data failed: %v", err)
		return
	}
	log.Infof("Alive data response %v: %s", m, string(msg))
}

func FetchData(m *model.MysqlIns, socketfile string, cfg *base.Cfg) (err error) {
	defer func() {
		MysqlAlive(cfg, m, err == nil)
	}()

	//db := mysql.New("tcp", "", fmt.Sprintf("%s:%d", m.Host, m.Port), cfg.User, cfg.Pass)
	db := mysql.New("unix", "", socketfile, cfg.User, cfg.Pass)
	db.SetTimeout(500 * time.Millisecond)
	if err = db.Connect(); err != nil {
		return
	}
	defer db.Close()

	data := make([]*model.MetaData, 0)
	threadNums, err := collect.ProcessNum(db, cfg)
	if err != nil {
		return
	}
	data = append(data, threadNums)
	slaveState, err := collect.SlaveStatus(db, cfg)
	if err != nil {
		return
	}
	data = append(data, slaveState...)

	msg, err := SendData(data, cfg)
	if err != nil {
		return
	}
	log.Infof("Send response %v: %s", m, string(msg))
	return
}
