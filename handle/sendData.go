package handle

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	log "github.com/Sirupsen/logrus"

	"github.com/mysql-monitor/base"
	"github.com/mysql-monitor/model"
)

func SendData(data []*model.MetaData, cfg *base.Cfg) ([]byte, error) {
	js, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	log.Debugf("Send to %s, size: %d", cfg.FalconClient, len(data))
	for _, m := range data {
		fmt.Println(m)
		log.Debugf("%s", m)
	}

	res, err := http.Post(cfg.FalconClient, "Content-Type: application/json", bytes.NewBuffer(js))
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	return ioutil.ReadAll(res.Body)
}
