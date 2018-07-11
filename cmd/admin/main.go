// Copyright 2016 CodisLabs. All Rights Reserved.
// Licensed under the MIT (MIT-LICENSE.txt) license.

package main

import (
	"github.com/docopt/docopt-go"

	"github.com/magenta9/codis/pkg/utils/log"
)

func main() {
	const usage = `
Usage:
	config --coor=ip:port [model|stats|slots|dbserver|proxy]
	config --coor=ip:port  --proxy-list
	config --coor=ip:port  --proxy-add      --addr=ADDR
	config --coor=ip:port  --proxy-remove   --addr=ADDR
	config --coor=ip:port  --proxy-status
	config --coor=ip:port  --db-list
	config --coor=ip:port  --db-add      --dbid=ID --addr=ADDR
	config --coor=ip:port  --db-del      --dbid=ID --addr=ADDR
	config --coor=ip:port  --db-status
	config --coor=ip:port  --slots-status
	config --coor=ip:port  --slot-remove    --sid=ID
	config --coor=ip:port  --slot-add       --sid=ID --dbid=ID
	config --coor=ip:port  --slot-migrate   --dbid=ID --beginid=ID --endid=ID
	config --coor=ip:port  --rebalance
`

	d, err := docopt.Parse(usage, nil, true, "", false)
	if err != nil {
		log.PanicError(err, "parse arguments failed")
	}
	log.SetLevel(log.LevelInfo)

	if d["-v"].(bool) {
		log.SetLevel(log.LevelDebug)
	}

	switch {
	case d["--proxy"] != nil:
		new(cmdProxy).Main(d)
	case d["--coor"] != nil:
		new(cmdDashboard).Main(d)
	default:
		new(cmdAdmin).Main(d)
	}
}
