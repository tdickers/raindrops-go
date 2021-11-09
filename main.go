package main

import (
	"time"

	"github.com/elastic/gosigar/sys/linux"
	//sock "github.com/elastic/beats/v7/metricbeat/helper/socket" // not resolving, OS X issue?
	sock "github.com/elastic/beats/metricbeat/helper/socket"
	log "github.com/sirupsen/logrus"
)

func main() {
	nlSess := sock.NewNetlinkSession()
	for _ = range time.Tick(2000 * time.Millisecond) {
		active := 0
		queued := 0
		inode_0 := 0

		sockets, err := nlSess.GetSocketList()
		if err != nil {
			log.WithError(err).Error("Unable to fetch socket list")
		}

		for _, s := range sockets {

			// TODO filter by sport to identify ones for our particular socket
			if s.Inode == 0 {
			//	log.Infof("Inode 0: %+v", s)
			// time_wait is also counted so maybe need to filter by state
			inode_0 += 1
				continue // if we were using an IP filter this would indicate a connection in the syn recvd queue, but that the filter hasn't applied to
				// it's possible this go code gives us an IP to work with
			}

			state := linux.TCPState(s.State)

			if state == linux.TCP_ESTABLISHED {
				active += 1
				//log.Infof("Active: %+v", s)
			} else if s.DstPort() == 0 && state == linux.TCP_LISTEN {
				//log.Infof("Listener: %+v", s)
				queued += int(s.RQueue)
			}
		}

		log.Infof("There are %d active and %d queued conns, inode 0: %d", active, queued, inode_0)
	}
}