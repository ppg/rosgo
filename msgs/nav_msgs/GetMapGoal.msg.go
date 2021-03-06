// Code generated by ros-gen-go.
// source: GetMapGoal.msg
// DO NOT EDIT!
package nav_msgs

import (
	"io"

	"github.com/ppg/rosgo/ros"
)

type _MsgGetMapGoal struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgGetMapGoal) Text() string {
	return t.text
}

func (t *_MsgGetMapGoal) Name() string {
	return t.name
}

func (t *_MsgGetMapGoal) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgGetMapGoal) NewMessage() ros.Message {
	m := new(GetMapGoal)

	return m
}

var (
	MsgGetMapGoal = &_MsgGetMapGoal{
		`# ====== DO NOT MODIFY! AUTOGENERATED FROM AN ACTION DEFINITION ======
# Get the map as a nav_msgs/OccupancyGrid
`,
		"nav_msgs/GetMapGoal",
		"f1dc1c7cf7d5ef6d4fbd2db76ba527ec",
	}
)

type GetMapGoal struct {
}

func (m *GetMapGoal) Serialize(w io.Writer) (err error) {
	return
}

func (m *GetMapGoal) Deserialize(r io.Reader) (err error) {
	return
}
