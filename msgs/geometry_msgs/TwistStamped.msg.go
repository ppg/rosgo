// Code generated by ros-gen-go.
// source: TwistStamped.msg
// DO NOT EDIT!
package geometry_msgs

import (
	"io"

	"github.com/ppg/rosgo/msgs/std_msgs"
	"github.com/ppg/rosgo/ros"
)

type _MsgTwistStamped struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgTwistStamped) Text() string {
	return t.text
}

func (t *_MsgTwistStamped) Name() string {
	return t.name
}

func (t *_MsgTwistStamped) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgTwistStamped) NewMessage() ros.Message {
	m := new(TwistStamped)

	return m
}

var (
	MsgTwistStamped = &_MsgTwistStamped{
		`# A twist with reference coordinate frame and timestamp
Header header
Twist twist
`,
		"geometry_msgs/TwistStamped",
		"08a22ddf566b82f747df9cc6e2fbbf7a",
	}
)

type TwistStamped struct {
	Header std_msgs.Header
	Twist  Twist
}

func (m *TwistStamped) Serialize(w io.Writer) (err error) {
	if err = ros.SerializeMessageField(w, "Header", &m.Header); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "Twist", &m.Twist); err != nil {
		return err
	}

	return
}

func (m *TwistStamped) Deserialize(r io.Reader) (err error) {
	// Header
	if err = ros.DeserializeMessageField(r, "Header", &m.Header); err != nil {
		return err
	}

	// Twist
	if err = ros.DeserializeMessageField(r, "Twist", &m.Twist); err != nil {
		return err
	}

	return
}
