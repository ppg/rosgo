// Code generated by ros-gen-go.
// source: /opt/ros/jade/share/ros/../control_msgs/msg//PointHeadResult.msg
// DO NOT EDIT!
package control_msgs

type _MsgPointHeadResult struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgPointHeadResult) Text() string {
	return t.text
}

func (t *_MsgPointHeadResult) Name() string {
	return t.name
}

func (t *_MsgPointHeadResult) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgPointHeadResult) NewMessage() ros.Message {
	m := new(PointHeadResult)

	return m
}

var (
	MsgPointHeadResult = &_MsgPointHeadResult{
		`# ====== DO NOT MODIFY! AUTOGENERATED FROM AN ACTION DEFINITION ======
`,
		"control_msgs/PointHeadResult",
		"7ac3b32c97133caf1b14edc99a50c37d",
	}
)

type PointHeadResult struct {
}

func (m *PointHeadResult) Type() ros.MessageType {
	return MsgPointHeadResult
}

func (m *PointHeadResult) Serialize(buf *bytes.Buffer) (err error) {
	return
}

func (m *String) Deserialize(buf *bytes.Reader) (err error) {
	return
}
