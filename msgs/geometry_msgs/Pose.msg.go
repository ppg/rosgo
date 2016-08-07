// Code generated by ros-gen-go.
// source: /opt/ros/jade/share/ros/../geometry_msgs/msg//Pose.msg
// DO NOT EDIT!
package geometry_msgs

type _MsgPose struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgPose) Text() string {
	return t.text
}

func (t *_MsgPose) Name() string {
	return t.name
}

func (t *_MsgPose) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgPose) NewMessage() ros.Message {
	m := new(Pose)

	return m
}

var (
	MsgPose = &_MsgPose{
		`# A representation of pose in free space, composed of postion and orientation. 
Point position
Quaternion orientation
`,
		"geometry_msgs/Pose",
		"dc72bb5c46de72bac142dec5d13c6f57",
	}
)

type Pose struct {
	Position    Point
	Orientation Quaternion
}

func (m *Pose) Type() ros.MessageType {
	return MsgPose
}

func (m *Pose) Serialize(buf *bytes.Buffer) (err error) {
	// Position
	if err = m.Position.Serialize(buf); err != nil {
		return err
	}

	// Orientation
	if err = m.Orientation.Serialize(buf); err != nil {
		return err
	}

	return
}

func (m *String) Deserialize(buf *bytes.Reader) (err error) {
	// Position
	if err = m.Position.Deserialize(buf); err != nil {
		return err
	}

	// Orientation
	if err = m.Orientation.Deserialize(buf); err != nil {
		return err
	}

	return
}
