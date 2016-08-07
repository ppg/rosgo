// Code generated by ros-gen-go.
// source: /opt/ros/jade/share/ros/../tf2_msgs/msg//LookupTransformGoal.msg
// DO NOT EDIT!
package tf2_msgs

type _MsgLookupTransformGoal struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgLookupTransformGoal) Text() string {
	return t.text
}

func (t *_MsgLookupTransformGoal) Name() string {
	return t.name
}

func (t *_MsgLookupTransformGoal) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgLookupTransformGoal) NewMessage() ros.Message {
	m := new(LookupTransformGoal)

	return m
}

var (
	MsgLookupTransformGoal = &_MsgLookupTransformGoal{
		`# ====== DO NOT MODIFY! AUTOGENERATED FROM AN ACTION DEFINITION ======
#Simple API
string target_frame
string source_frame
time source_time
duration timeout

#Advanced API
time target_time
string fixed_frame

#Whether or not to use the advanced API
bool advanced

`,
		"tf2_msgs/LookupTransformGoal",
		"8cc5f822256e11b7311c2088bd622d9e",
	}
)

type LookupTransformGoal struct {
	TargetFrame string
	SourceFrame string
	SourceTime  ros.Time
	Timeout     ros.Duration
	TargetTime  ros.Time
	FixedFrame  string
	Advanced    bool
}

func (m *LookupTransformGoal) Type() ros.MessageType {
	return MsgLookupTransformGoal
}

func (m *LookupTransformGoal) Serialize(buf *bytes.Buffer) (err error) {
	// TargetFrame
	binary.Write(buf, binary.LittleEndian, uint32(len([]byte(m.TargetFrame))))
	buf.Write([]byte(m.TargetFrame))

	// SourceFrame
	binary.Write(buf, binary.LittleEndian, uint32(len([]byte(m.SourceFrame))))
	buf.Write([]byte(m.SourceFrame))

	// SourceTime
	binary.Write(buf, binary.LittleEndian, m.SourceTime.Sec)
	binary.Write(buf, binary.LittleEndian, m.SourceTime.NSec)

	// Timeout
	binary.Write(buf, binary.LittleEndian, m.Timeout.Sec)
	binary.Write(buf, binary.LittleEndian, m.Timeout.NSec)

	// TargetTime
	binary.Write(buf, binary.LittleEndian, m.TargetTime.Sec)
	binary.Write(buf, binary.LittleEndian, m.TargetTime.NSec)

	// FixedFrame
	binary.Write(buf, binary.LittleEndian, uint32(len([]byte(m.FixedFrame))))
	buf.Write([]byte(m.FixedFrame))

	// Advanced
	binary.Write(buf, binary.LittleEndian, m.Advanced)

	return
}

func (m *String) Deserialize(buf *bytes.Reader) (err error) {
	// TargetFrame
	var size uint32
	if err = binary.Read(buf, binary.LittleEndian, &size); err != nil {
		return err
	}
	data := make([]byte, int(size))
	if err = binary.Read(buf, binary.LittleEndian, data); err != nil {
		return err
	}
	m.TargetFrame = string(data)

	// SourceFrame
	var size uint32
	if err = binary.Read(buf, binary.LittleEndian, &size); err != nil {
		return err
	}
	data := make([]byte, int(size))
	if err = binary.Read(buf, binary.LittleEndian, data); err != nil {
		return err
	}
	m.SourceFrame = string(data)

	// SourceTime
	if err = binary.Read(buf, binary.LittleEndian, &m.SourceTime.Sec); err != nil {
		return err
	}
	if err = binary.Read(buf, binary.LittleEndian, &m.SourceTime.NSec); err != nil {
		return err
	}

	// Timeout
	if err = binary.Read(buf, binary.LittleEndian, &m.Timeout.Sec); err != nil {
		return err
	}
	if err = binary.Read(buf, binary.LittleEndian, &m.Timeout.NSec); err != nil {
		return err
	}

	// TargetTime
	if err = binary.Read(buf, binary.LittleEndian, &m.TargetTime.Sec); err != nil {
		return err
	}
	if err = binary.Read(buf, binary.LittleEndian, &m.TargetTime.NSec); err != nil {
		return err
	}

	// FixedFrame
	var size uint32
	if err = binary.Read(buf, binary.LittleEndian, &size); err != nil {
		return err
	}
	data := make([]byte, int(size))
	if err = binary.Read(buf, binary.LittleEndian, data); err != nil {
		return err
	}
	m.FixedFrame = string(data)

	// Advanced
	if err = binary.Read(buf, binary.LittleEndian, &m.Advanced); err != nil {
		return err
	}

	return
}
