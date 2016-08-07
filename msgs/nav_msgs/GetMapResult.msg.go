// Code generated by ros-gen-go.
// source: /opt/ros/jade/share/ros/../nav_msgs/msg//GetMapResult.msg
// DO NOT EDIT!
package nav_msgs

type _MsgGetMapResult struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgGetMapResult) Text() string {
	return t.text
}

func (t *_MsgGetMapResult) Name() string {
	return t.name
}

func (t *_MsgGetMapResult) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgGetMapResult) NewMessage() ros.Message {
	m := new(GetMapResult)

	return m
}

var (
	MsgGetMapResult = &_MsgGetMapResult{
		`# ====== DO NOT MODIFY! AUTOGENERATED FROM AN ACTION DEFINITION ======
nav_msgs/OccupancyGrid map
`,
		"nav_msgs/GetMapResult",
		"ed8b66960ff4f67a7b0669717662affe",
	}
)

type GetMapResult struct {
	Map OccupancyGrid
}

func (m *GetMapResult) Type() ros.MessageType {
	return MsgGetMapResult
}

func (m *GetMapResult) Serialize(buf *bytes.Buffer) (err error) {
	// Map
	if err = m.Map.Serialize(buf); err != nil {
		return err
	}

	return
}

func (m *String) Deserialize(buf *bytes.Reader) (err error) {
	// Map
	if err = m.Map.Deserialize(buf); err != nil {
		return err
	}

	return
}
