// Package actionlib_msgs contains message definitions (autogenerated).
package actionlib_msgs

import (
	"github.com/aler9/goroslib/msg"
	"github.com/aler9/goroslib/msgs/std_msgs"
	"time"
)

const (
	PENDING    uint8 = 0
	ACTIVE     uint8 = 1
	PREEMPTED  uint8 = 2
	SUCCEEDED  uint8 = 3
	ABORTED    uint8 = 4
	REJECTED   uint8 = 5
	PREEMPTING uint8 = 6
	RECALLING  uint8 = 7
	RECALLED   uint8 = 8
	LOST       uint8 = 9
)

type GoalID struct {
	msg.Package `ros:"actionlib_msgs"`
	Stamp       time.Time
	Id          string
}

type GoalStatus struct {
	msg.Package     `ros:"actionlib_msgs"`
	msg.Definitions `ros:"uint8 PENDING=0,uint8 ACTIVE=1,uint8 PREEMPTED=2,uint8 SUCCEEDED=3,uint8 ABORTED=4,uint8 REJECTED=5,uint8 PREEMPTING=6,uint8 RECALLING=7,uint8 RECALLED=8,uint8 LOST=9"`
	GoalId          GoalID
	Status          uint8
	Text            string
}

type GoalStatusArray struct {
	msg.Package `ros:"actionlib_msgs"`
	Header      std_msgs.Header
	StatusList  []GoalStatus
}