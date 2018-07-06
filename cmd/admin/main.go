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
	config [-v] --proxy=ADDR [--auth=AUTH] [config|model|stats|slots]
	config [-v] --proxy=ADDR [--auth=AUTH]  --start
	config [-v] --proxy=ADDR [--auth=AUTH]  --shutdown
	config [-v] --proxy=ADDR [--auth=AUTH]  --log-level=LEVEL
	config [-v] --proxy=ADDR [--auth=AUTH]  --fillslots=FILE [--locked]
	config [-v] --proxy=ADDR [--auth=AUTH]  --reset-stats
	config [-v] --proxy=ADDR [--auth=AUTH]  --forcegc
	config [-v] --dashboard=ADDR           [config|model|stats|slots|group|proxy]
	config [-v] --dashboard=ADDR            --shutdown
	config [-v] --dashboard=ADDR            --reload
	config [-v] --dashboard=ADDR            --log-level=LEVEL
	config [-v] --dashboard=ADDR            --slots-assign   --beg=ID --end=ID (--gid=ID|--offline) [--confirm]
	config [-v] --dashboard=ADDR            --slots-status
	config [-v] --dashboard=ADDR            --list-proxy
	config [-v] --dashboard=ADDR            --create-proxy   --addr=ADDR
	config [-v] --dashboard=ADDR            --online-proxy   --addr=ADDR
	config [-v] --dashboard=ADDR            --remove-proxy  (--addr=ADDR|--token=TOKEN|--pid=ID)       [--force]
	config [-v] --dashboard=ADDR            --reinit-proxy  (--addr=ADDR|--token=TOKEN|--pid=ID|--all) [--force]
	config [-v] --dashboard=ADDR            --proxy-status
	config [-v] --dashboard=ADDR            --list-group
	config [-v] --dashboard=ADDR            --create-group   --gid=ID
	config [-v] --dashboard=ADDR            --remove-group   --gid=ID
	config [-v] --dashboard=ADDR            --resync-group  [--gid=ID | --all]
	config [-v] --dashboard=ADDR            --group-add      --gid=ID --addr=ADDR [--datacenter=DATACENTER]
	config [-v] --dashboard=ADDR            --group-del      --gid=ID --addr=ADDR
	config [-v] --dashboard=ADDR            --group-status
	config [-v] --dashboard=ADDR            --replica-groups --gid=ID --addr=ADDR (--enable|--disable)
	config [-v] --dashboard=ADDR            --promote-server --gid=ID --addr=ADDR
	config [-v] --dashboard=ADDR            --sync-action    --create --addr=ADDR
	config [-v] --dashboard=ADDR            --sync-action    --remove --addr=ADDR
	config [-v] --dashboard=ADDR            --slot-action    --create --sid=ID --gid=ID
	config [-v] --dashboard=ADDR            --slot-action    --remove --sid=ID
	config [-v] --dashboard=ADDR            --slot-action    --create-some  --gid-from=ID --gid-to=ID --num-slots=N
	config [-v] --dashboard=ADDR            --slot-action    --create-range --beg=ID --end=ID --gid=ID
	config [-v] --dashboard=ADDR            --slot-action    --interval=VALUE
	config [-v] --dashboard=ADDR            --slot-action    --disabled=VALUE
	config [-v] --dashboard=ADDR            --rebalance     [--confirm]
	config [-v] --dashboard=ADDR            --sentinel-add   --addr=ADDR
	config [-v] --dashboard=ADDR            --sentinel-del   --addr=ADDR [--force]
	config [-v] --dashboard=ADDR            --sentinel-resync
	config [-v] --remove-lock               --product=NAME (--zookeeper=ADDR [--zookeeper-auth=USR:PWD]|--etcd=ADDR [--etcd-auth=USR:PWD]|--filesystem=ROOT)
	config [-v] --config-dump               --product=NAME (--zookeeper=ADDR [--zookeeper-auth=USR:PWD]|--etcd=ADDR [--etcd-auth=USR:PWD]|--filesystem=ROOT) [-1]
	config [-v] --config-convert=FILE
	config [-v] --config-restore=FILE       --product=NAME (--zookeeper=ADDR [--zookeeper-auth=USR:PWD]|--etcd=ADDR [--etcd-auth=USR:PWD]|--filesystem=ROOT) [--confirm]
	config [-v] --dashboard-list                           (--zookeeper=ADDR [--zookeeper-auth=USR:PWD]|--etcd=ADDR [--etcd-auth=USR:PWD]|--filesystem=ROOT)

Options:
	-a AUTH, --auth=AUTH
	-x ADDR, --addr=ADDR
	-t TOKEN, --token=TOKEN
	-g ID, --gid=ID
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
	case d["--dashboard"] != nil:
		new(cmdDashboard).Main(d)
	default:
		new(cmdAdmin).Main(d)
	}
}
