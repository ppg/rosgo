// Code generated by ros-gen-go.
// source: /opt/ros/jade/share/ros/../geometry_msgs/msg//PoseWithCovariance.msg
// DO NOT EDIT!
package geometry_msgs

type _MsgPoseWithCovariance struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgPoseWithCovariance) Text() string {
	return t.text
}

func (t *_MsgPoseWithCovariance) Name() string {
	return t.name
}

func (t *_MsgPoseWithCovariance) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgPoseWithCovariance) NewMessage() ros.Message {
	m := new(PoseWithCovariance)

	return m
}

var (
	MsgPoseWithCovariance = &_MsgPoseWithCovariance{
		`# This represents a pose in free space with uncertainty.

Pose pose

# Row-major representation of the 6x6 covariance matrix
# The orientation parameters use a fixed-axis representation.
# In order, the parameters are:
# (x, y, z, rotation about X axis, rotation about Y axis, rotation about Z axis)
float64[36] covariance
`,
		"geometry_msgs/PoseWithCovariance",
		"4ec31161b30291389f54fb885685270a",
	}
)

type PoseWithCovariance struct {
	Pose       Pose
	Covariance [36]float64
}

func (m *PoseWithCovariance) Type() ros.MessageType {
	return MsgPoseWithCovariance
}

func (m *PoseWithCovariance) Serialize(buf *bytes.Buffer) (err error) {
	// Pose
	if err = m.Pose.Serialize(buf); err != nil {
		return err
	}

	// Covariance
	binary.Write(buf, binary.LittleEndian, uint32(len(m.Covariance)))
	for _, e := range m.Covariance {
		binary.Write(buf, binary.LittleEndian, e)
	}

	return
}

func (m *String) Deserialize(buf *bytes.Reader) (err error) {
	// Pose
	if err = m.Pose.Deserialize(buf); err != nil {
		return err
	}

	// Covariance
	var size uint32
	if err = binary.Read(buf, binary.LittleEndian, &size); err != nil {
		return err
	}
	if size > 36 {
		return fmt.Errorf("array size for Covariance too large: expected=36, got=%d", size)
	}
	for i := 0; i < int(size); i++ {
		if err = binary.Read(buf, binary.LittleEndian, &m.Covariance[i]); err != nil {
			return err
		}
	}

	return
}
