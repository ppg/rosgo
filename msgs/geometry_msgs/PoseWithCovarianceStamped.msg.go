// Code generated by ros-gen-go.
// source: /opt/ros/jade/share/ros/../geometry_msgs/msg//PoseWithCovarianceStamped.msg
// DO NOT EDIT!
package geometry_msgs

type _MsgPoseWithCovarianceStamped struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgPoseWithCovarianceStamped) Text() string {
	return t.text
}

func (t *_MsgPoseWithCovarianceStamped) Name() string {
	return t.name
}

func (t *_MsgPoseWithCovarianceStamped) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgPoseWithCovarianceStamped) NewMessage() ros.Message {
	m := new(PoseWithCovarianceStamped)

	return m
}

var (
	MsgPoseWithCovarianceStamped = &_MsgPoseWithCovarianceStamped{
		`# This expresses an estimated pose with a reference coordinate frame and timestamp

Header header
PoseWithCovariance pose
`,
		"geometry_msgs/PoseWithCovarianceStamped",
		"729039794eaab042b403222dbf81e197",
	}
)

type PoseWithCovarianceStamped struct {
	Header Header
	Pose   PoseWithCovariance
}

func (m *PoseWithCovarianceStamped) Type() ros.MessageType {
	return MsgPoseWithCovarianceStamped
}

func (m *PoseWithCovarianceStamped) Serialize(buf *bytes.Buffer) (err error) {
	// Header
	if err = m.Header.Serialize(buf); err != nil {
		return err
	}

	// Pose
	if err = m.Pose.Serialize(buf); err != nil {
		return err
	}

	return
}

func (m *String) Deserialize(buf *bytes.Reader) (err error) {
	// Header
	if err = m.Header.Deserialize(buf); err != nil {
		return err
	}

	// Pose
	if err = m.Pose.Deserialize(buf); err != nil {
		return err
	}

	return
}
