// Code generated by ros-gen-go.
// source: /opt/ros/jade/share/ros/../diagnostic_msgs/msg//DiagnosticArray.msg
// DO NOT EDIT!
package diagnostic_msgs

type _MsgDiagnosticArray struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgDiagnosticArray) Text() string {
	return t.text
}

func (t *_MsgDiagnosticArray) Name() string {
	return t.name
}

func (t *_MsgDiagnosticArray) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgDiagnosticArray) NewMessage() ros.Message {
	m := new(DiagnosticArray)

	return m
}

var (
	MsgDiagnosticArray = &_MsgDiagnosticArray{
		`# This message is used to send diagnostic information about the state of the robot
Header header #for timestamp
DiagnosticStatus[] status # an array of components being reported on`,
		"diagnostic_msgs/DiagnosticArray",
		"5a7dd9a5697a73da9c54c05194f61e38",
	}
)

type DiagnosticArray struct {
	Header Header
	Status []DiagnosticStatus
}

func (m *DiagnosticArray) Type() ros.MessageType {
	return MsgDiagnosticArray
}

func (m *DiagnosticArray) Serialize(buf *bytes.Buffer) (err error) {
	// Header
	if err = m.Header.Serialize(buf); err != nil {
		return err
	}

	// Status
	binary.Write(buf, binary.LittleEndian, uint32(len(m.Status)))
	for _, e := range m.Status {
		if err = e.Serialize(buf); err != nil {
			return err
		}
	}

	return
}

func (m *String) Deserialize(buf *bytes.Reader) (err error) {
	// Header
	if err = m.Header.Deserialize(buf); err != nil {
		return err
	}

	// Status
	var size uint32
	if err = binary.Read(buf, binary.LittleEndian, &size); err != nil {
		return err
	}
	m.Status = make([]DiagnosticStatus, int(size))
	for i := 0; i < int(size); i++ {
		if err = m.Status[i].Deserialize(buf); err != nil {
			return err
		}
	}

	return
}
