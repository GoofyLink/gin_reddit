package snowflake

import (
	"time"

	sf "github.com/bwmarrin/snowflake"
)

var node *sf.Node

func Init(startTime string, machineId int64) (err error) {
	//fmt.Println(startTime, machineID)
	var parsedStartTime time.Time
	parsedStartTime, err = time.Parse("2006-01-02", startTime)
	if err != nil {
		return err
	}
	sf.Epoch = parsedStartTime.UnixNano() / 1000000
	node, err = sf.NewNode(machineId)
	if err != nil {
		return err
	}
	return
}
func GenID() int64 {
	return node.Generate().Int64()
}
