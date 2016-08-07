// Code generated by ros-gen-go.
// source: /opt/ros/jade/share/ros/../geometry_msgs/msg//PointStamped.msg
// DO NOT EDIT!
package geometry_msgs

type _MsgPointStamped struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgPointStamped) Text() string {
	return t.text
}

func (t *_MsgPointStamped) Name() string {
	return t.name
}

func (t *_MsgPointStamped) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgPointStamped) NewMessage() ros.Message {
	m := new(PointStamped)

	return m
}

var (
	MsgPointStamped = &_MsgPointStamped{
		`# This represents a Point with reference coordinate frame and timestamp
Header header
Point point
`,
		"geometry_msgs/PointStamped",
		"e948b3cf3f45aaeaedb063e8b966cf1f",
	}
)

type PointStamped struct {
	Header Header
	Point  Point
}

func (m *PointStamped) Type() ros.MessageType {
	return MsgPointStamped
}

func (m *PointStamped) Serialize(buf *bytes.Buffer) (err error) {
	// Header
	if err = m.Header.Serialize(buf); err != nil {
		return err
	}

	// Point
	if err = m.Point.Serialize(buf); err != nil {
		return err
	}

	return
}

func (m *String) Deserialize(buf *bytes.Reader) (err error) {
	// Header
	if err = m.Header.Deserialize(buf); err != nil {
		return err
	}

	// Point
	if err = m.Point.Deserialize(buf); err != nil {
		return err
	}

	return
}
