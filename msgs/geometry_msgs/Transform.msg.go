// Code generated by ros-gen-go.
// source: /opt/ros/jade/share/ros/../geometry_msgs/msg//Transform.msg
// DO NOT EDIT!
package geometry_msgs

type _MsgTransform struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgTransform) Text() string {
	return t.text
}

func (t *_MsgTransform) Name() string {
	return t.name
}

func (t *_MsgTransform) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgTransform) NewMessage() ros.Message {
	m := new(Transform)

	return m
}

var (
	MsgTransform = &_MsgTransform{
		`# This represents the transform between two coordinate frames in free space.

Vector3 translation
Quaternion rotation
`,
		"geometry_msgs/Transform",
		"756be060b1c8cf0e64a10ba16909d887",
	}
)

type Transform struct {
	Translation Vector3
	Rotation    Quaternion
}

func (m *Transform) Type() ros.MessageType {
	return MsgTransform
}

func (m *Transform) Serialize(buf *bytes.Buffer) (err error) {
	// Translation
	if err = m.Translation.Serialize(buf); err != nil {
		return err
	}

	// Rotation
	if err = m.Rotation.Serialize(buf); err != nil {
		return err
	}

	return
}

func (m *String) Deserialize(buf *bytes.Reader) (err error) {
	// Translation
	if err = m.Translation.Deserialize(buf); err != nil {
		return err
	}

	// Rotation
	if err = m.Rotation.Deserialize(buf); err != nil {
		return err
	}

	return
}
