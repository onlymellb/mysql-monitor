package model

import (
	"fmt"
	"os"
	"time"

	"github.com/mysql-monitor/base"
)

type MysqlIns struct {
	Host string
	Tag  string
}

type MetaData struct {
	Metric      string      `json:"metric"`      //key
	Endpoint    string      `json:"endpoint"`    //hostname
	Value       interface{} `json:"value"`       // number or string
	CounterType string      `json:"counterType"` // GAUGE  原值   COUNTER 差值(ps)
	Tags        string      `json:"tags"`        // port=3306,k=v
	Timestamp   int64       `json:"timestamp"`
	Step        int64       `json:"step"`
}

func (m *MysqlIns) String() string {
	return fmt.Sprintf("%s:%s", m.Host, m.Tag)
}

func (m *MetaData) String() string {
	s := fmt.Sprintf("MetaData Metric:%s Endpoint:%s Value:%v CounterType:%s Tags:%s Timestamp:%d Step:%d",
		m.Metric, m.Endpoint, m.Value, m.CounterType, m.Tags, m.Timestamp, m.Step)
	return s
}

func NewMetric(name string, cfg *base.Cfg) *MetaData {
	return &MetaData{
		Metric:      name,
		Endpoint:    cfg.Endpoint,
		CounterType: dataType(name),
		Tags:        fmt.Sprintf("appname=%s", cfg.Appname),
		Timestamp:   time.Now().Unix(),
		Step:        60,
	}
}

func HostName(host string, cfg *base.Cfg) string {
	if host != "" {
		return host
	}
	host, err := os.Hostname()
	if err != nil {
		host = cfg.Host
	}
	return host
}

func (m *MetaData) SetValue(v interface{}) {
	m.Value = v
}
