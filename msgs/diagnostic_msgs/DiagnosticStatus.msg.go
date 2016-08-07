// Code generated by ros-gen-go.
// source: /opt/ros/jade/share/ros/../diagnostic_msgs/msg//DiagnosticStatus.msg
// DO NOT EDIT!
package diagnostic_msgs

type _MsgDiagnosticStatus struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgDiagnosticStatus) Text() string {
	return t.text
}

func (t *_MsgDiagnosticStatus) Name() string {
	return t.name
}

func (t *_MsgDiagnosticStatus) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgDiagnosticStatus) NewMessage() ros.Message {
	m := new(DiagnosticStatus)

	return m
}

var (
	MsgDiagnosticStatus = &_MsgDiagnosticStatus{
		`# This message holds the status of an individual component of the robot.
# 

# Possible levels of operations
byte OK=0
byte WARN=1
byte ERROR=2
byte STALE=3

byte level # level of operation enumerated above 
string name # a description of the test/component reporting
string message # a description of the status
string hardware_id # a hardware unique string
KeyValue[] values # an array of values associated with the status

`,
		"diagnostic_msgs/DiagnosticStatus",
		"bfaff5a20cf26a7d7255cbaa01dae095",
	}
)

type DiagnosticStatus struct {
	Level      int8
	Name       string
	Message    string
	HardwareID string
	Values     []KeyValue
}

func (m *DiagnosticStatus) Type() ros.MessageType {
	return MsgDiagnosticStatus
}

func (m *DiagnosticStatus) Serialize(buf *bytes.Buffer) (err error) {
	// Level
	binary.Write(buf, binary.LittleEndian, m.Level)

	// Name
	binary.Write(buf, binary.LittleEndian, uint32(len([]byte(m.Name))))
	buf.Write([]byte(m.Name))

	// Message
	binary.Write(buf, binary.LittleEndian, uint32(len([]byte(m.Message))))
	buf.Write([]byte(m.Message))

	// HardwareID
	binary.Write(buf, binary.LittleEndian, uint32(len([]byte(m.HardwareID))))
	buf.Write([]byte(m.HardwareID))

	// Values
	binary.Write(buf, binary.LittleEndian, uint32(len(m.Values)))
	for _, e := range m.Values {
		if err = e.Serialize(buf); err != nil {
			return err
		}
	}

	return
}

func (m *String) Deserialize(buf *bytes.Reader) (err error) {
	// Level
	if err = binary.Read(buf, binary.LittleEndian, &m.Level); err != nil {
		return err
	}

	// Name
	var size uint32
	if err = binary.Read(buf, binary.LittleEndian, &size); err != nil {
		return err
	}
	data := make([]byte, int(size))
	if err = binary.Read(buf, binary.LittleEndian, data); err != nil {
		return err
	}
	m.Name = string(data)

	// Message
	var size uint32
	if err = binary.Read(buf, binary.LittleEndian, &size); err != nil {
		return err
	}
	data := make([]byte, int(size))
	if err = binary.Read(buf, binary.LittleEndian, data); err != nil {
		return err
	}
	m.Message = string(data)

	// HardwareID
	var size uint32
	if err = binary.Read(buf, binary.LittleEndian, &size); err != nil {
		return err
	}
	data := make([]byte, int(size))
	if err = binary.Read(buf, binary.LittleEndian, data); err != nil {
		return err
	}
	m.HardwareID = string(data)

	// Values
	var size uint32
	if err = binary.Read(buf, binary.LittleEndian, &size); err != nil {
		return err
	}
	m.Values = make([]KeyValue, int(size))
	for i := 0; i < int(size); i++ {
		if err = m.Values[i].Deserialize(buf); err != nil {
			return err
		}
	}

	return
}
