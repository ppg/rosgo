// Code generated by ros-gen-go.
// source: /opt/ros/jade/share/ros/../geometry_msgs/msg//TwistWithCovariance.msg
// DO NOT EDIT!
package geometry_msgs

type _MsgTwistWithCovariance struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgTwistWithCovariance) Text() string {
	return t.text
}

func (t *_MsgTwistWithCovariance) Name() string {
	return t.name
}

func (t *_MsgTwistWithCovariance) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgTwistWithCovariance) NewMessage() ros.Message {
	m := new(TwistWithCovariance)

	return m
}

var (
	MsgTwistWithCovariance = &_MsgTwistWithCovariance{
		`# This expresses velocity in free space with uncertainty.

Twist twist

# Row-major representation of the 6x6 covariance matrix
# The orientation parameters use a fixed-axis representation.
# In order, the parameters are:
# (x, y, z, rotation about X axis, rotation about Y axis, rotation about Z axis)
float64[36] covariance
`,
		"geometry_msgs/TwistWithCovariance",
		"408e7ef4f4ec295f4663586922faacdd",
	}
)

type TwistWithCovariance struct {
	Twist      Twist
	Covariance [36]float64
}

func (m *TwistWithCovariance) Type() ros.MessageType {
	return MsgTwistWithCovariance
}

func (m *TwistWithCovariance) Serialize(buf *bytes.Buffer) (err error) {
	// Twist
	if err = m.Twist.Serialize(buf); err != nil {
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
	// Twist
	if err = m.Twist.Deserialize(buf); err != nil {
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
