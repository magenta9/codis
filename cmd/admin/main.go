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
	config --proxy=ADDR [--auth=AUTH] [config|model|stats|slots]
	config --proxy=ADDR [--auth=AUTH]  --start
	config --proxy=ADDR [--auth=AUTH]  --shutdown
	config --proxy=ADDR [--auth=AUTH]  --log-level=LEVEL
	config --proxy=ADDR [--auth=AUTH]  --fillslots=FILE [--locked]
	config --proxy=ADDR [--auth=AUTH]  --reset-stats
	config --proxy=ADDR [--auth=AUTH]  --forcegc
	config --coor=ip:port [config|model|stats|slots|group|proxy]
	config --coor=ip:port  --slots-assign   --beg=ID --end=ID (--gid=ID|--offline) [--confirm]
	config --coor=ip:port  --slots-status
	config --coor=ip:port  --list-proxy
	config --coor=ip:port  --create-proxy   --addr=ADDR
	config --coor=ip:port  --online-proxy   --addr=ADDR
	config --coor=ip:port  --remove-proxy  (--addr=ADDR|--token=TOKEN|--pid=ID)       [--force]
	config --coor=ip:port  --reinit-proxy  (--addr=ADDR|--token=TOKEN|--pid=ID|--all) [--force]
	config --coor=ip:port  --proxy-status
	config --coor=ip:port  --list-group
	config --coor=ip:port  --create-group   --gid=ID
	config --coor=ip:port  --remove-group   --gid=ID
	config --coor=ip:port  --resync-group  [--gid=ID | --all]
	config --coor=ip:port  --group-add      --gid=ID --addr=ADDR [--datacenter=DATACENTER]
	config --coor=ip:port  --group-del      --gid=ID --addr=ADDR
	config --coor=ip:port  --group-status
	config --coor=ip:port  --replica-groups --gid=ID --addr=ADDR (--enable|--disable)
	config --coor=ip:port  --promote-server --gid=ID --addr=ADDR
	config --coor=ip:port  --sync-action    --create --addr=ADDR
	config --coor=ip:port  --sync-action    --remove --addr=ADDR
	config --coor=ip:port  --slot-action    --create --sid=ID --gid=ID
	config --coor=ip:port  --slot-action    --remove --sid=ID
	config --coor=ip:port  --slot-action    --create-some  --gid-from=ID --gid-to=ID --num-slots=N
	config --coor=ip:port  --slot-action    --create-range --beg=ID --end=ID --gid=ID
	config --coor=ip:port  --slot-action    --interval=VALUE
	config --coor=ip:port  --slot-action    --disabled=VALUE
	config --coor=ip:port  --rebalance     [--confirm]
	config --coor=ip:port  --sentinel-add   --addr=ADDR
	config --coor=ip:port  --sentinel-del   --addr=ADDR [--force]
	config --coor=ip:port  --sentinel-resync
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
