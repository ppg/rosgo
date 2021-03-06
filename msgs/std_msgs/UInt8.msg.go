// Code generated by ros-gen-go.
// source: UInt8.msg
// DO NOT EDIT!
package std_msgs

import (
	"io"

	"github.com/ppg/rosgo/ros"
)

type _MsgUInt8 struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgUInt8) Text() string {
	return t.text
}

func (t *_MsgUInt8) Name() string {
	return t.name
}

func (t *_MsgUInt8) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgUInt8) NewMessage() ros.Message {
	m := new(UInt8)

	return m
}

var (
	MsgUInt8 = &_MsgUInt8{
		`uint8 data
`,
		"std_msgs/UInt8",
		"7dd4c5dca2bc8c0277734d2bba45f965",
	}
)

type UInt8 struct {
	Data uint8
}

func (m *UInt8) Serialize(w io.Writer) (err error) {
	if err = ros.SerializeMessageField(w, "uint8", &m.Data); err != nil {
		return err
	}

	return
}

func (m *UInt8) Deserialize(r io.Reader) (err error) {
	// Data
	if err = ros.DeserializeMessageField(r, "uint8", &m.Data); err != nil {
		return err
	}

	return
}
