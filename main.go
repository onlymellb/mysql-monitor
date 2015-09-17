//Lain MySQL Performance Monitor(For open-falcon)
package main

import (
	"flag"
	"fmt"
	"os"

	log "github.com/Sirupsen/logrus"

	"github.com/mysql-monitor/base"
	"github.com/mysql-monitor/handle"
	"github.com/mysql-monitor/model"
)

func main() {
	var sdir, endpoint string
	var user, passwd string
	var version bool

	flag.StringVar(&sdir, "s", "/var/lib/mysql", "mysql socket file in this dir.")
	flag.StringVar(&user, "u", "monitor", "monitor user for mysql")
	flag.StringVar(&passwd, "p", "", "user password")
	flag.StringVar(&endpoint, "ep", "", "endpoint name")
	flag.BoolVar(&version, "v", false, "show version")
	flag.Parse()

	if version {
		fmt.Println(model.VERSION)
		os.Exit(0)
	}

	go base.Timeout()
	base.CheckIsDir(sdir)
	cfg := base.NewConfig(sdir, user, passwd)
	cfg.Endpoint = model.HostName(endpoint, cfg)
	base.InitLog(cfg)

	nsmap, err := base.GetAppSockMap(cfg.SocketDir)
	if err != nil {
		log.Fatalf("get socket file list failed: %v", err)
	}
	if len(nsmap) == 0 {
		log.Fatalf("not found socket file in assign dir: %v", cfg.SocketDir)
	}

	for appname, socketpath := range nsmap {
		cfg.Appname = appname
		relay := cfg.Endpoint
		cfg.Endpoint = fmt.Sprintf("%s_%s_%s", "mysql", appname, cfg.Endpoint)
		err := handle.FetchData(&model.MysqlIns{
			Host: cfg.Host,
			Tag:  fmt.Sprintf("appname=%d", cfg.Appname),
		}, socketpath, cfg)
		if err != nil {
			log.Error(err)
		}
		cfg.Endpoint = relay
	}
}
