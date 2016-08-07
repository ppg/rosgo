// Code generated by ros-gen-go.
// source: /opt/ros/jade/share/ros/../geometry_msgs/msg//WrenchStamped.msg
// DO NOT EDIT!
package geometry_msgs

type _MsgWrenchStamped struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgWrenchStamped) Text() string {
	return t.text
}

func (t *_MsgWrenchStamped) Name() string {
	return t.name
}

func (t *_MsgWrenchStamped) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgWrenchStamped) NewMessage() ros.Message {
	m := new(WrenchStamped)

	return m
}

var (
	MsgWrenchStamped = &_MsgWrenchStamped{
		`# A wrench with reference coordinate frame and timestamp
Header header
Wrench wrench
`,
		"geometry_msgs/WrenchStamped",
		"150502b356390fb151385ef7647f633e",
	}
)

type WrenchStamped struct {
	Header Header
	Wrench Wrench
}

func (m *WrenchStamped) Type() ros.MessageType {
	return MsgWrenchStamped
}

func (m *WrenchStamped) Serialize(buf *bytes.Buffer) (err error) {
	// Header
	if err = m.Header.Serialize(buf); err != nil {
		return err
	}

	// Wrench
	if err = m.Wrench.Serialize(buf); err != nil {
		return err
	}

	return
}

func (m *String) Deserialize(buf *bytes.Reader) (err error) {
	// Header
	if err = m.Header.Deserialize(buf); err != nil {
		return err
	}

	// Wrench
	if err = m.Wrench.Deserialize(buf); err != nil {
		return err
	}

	return
}
