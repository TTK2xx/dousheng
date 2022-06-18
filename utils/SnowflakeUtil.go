package utils

import (
	sn "github.com/bwmarrin/snowflake"
	"time"
)

var node *sn.Node

func Init() (err error) {
	const layout = "Jan 2, 2006 at 3:04pm (MST)"
	tm, err := time.Parse(layout, "Feb 4, 2014 at 6:05pm (PST)")
	if err != nil {
		return
	}
	sn.Epoch = tm.UnixNano() / 1000000
	node, err = sn.NewNode(64)
	return
}
func GenID() int64 {
	return node.Generate().Int64()
}
