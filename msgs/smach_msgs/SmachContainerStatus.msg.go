// Code generated by ros-gen-go.
// source: /opt/ros/jade/share/ros/../smach_msgs/msg//SmachContainerStatus.msg
// DO NOT EDIT!
package smach_msgs

type _MsgSmachContainerStatus struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgSmachContainerStatus) Text() string {
	return t.text
}

func (t *_MsgSmachContainerStatus) Name() string {
	return t.name
}

func (t *_MsgSmachContainerStatus) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgSmachContainerStatus) NewMessage() ros.Message {
	m := new(SmachContainerStatus)

	return m
}

var (
	MsgSmachContainerStatus = &_MsgSmachContainerStatus{
		`Header header

# The path to this node in the server
string path

# The initial state description
# Effects an arc from the top state to each one
string[] initial_states

# The current state description
string[] active_states

# A pickled user data structure
# i.e. the UserData's internal dictionary
string local_data

# Debugging info string
string info
`,
		"smach_msgs/SmachContainerStatus",
		"2a07c0f13a55f9c0e61b861363c73741",
	}
)

type SmachContainerStatus struct {
	Header        Header
	Path          string
	InitialStates []string
	ActiveStates  []string
	LocalData     string
	Info          string
}

func (m *SmachContainerStatus) Type() ros.MessageType {
	return MsgSmachContainerStatus
}

func (m *SmachContainerStatus) Serialize(buf *bytes.Buffer) (err error) {
	// Header
	if err = m.Header.Serialize(buf); err != nil {
		return err
	}

	// Path
	binary.Write(buf, binary.LittleEndian, uint32(len([]byte(m.Path))))
	buf.Write([]byte(m.Path))

	// InitialStates
	binary.Write(buf, binary.LittleEndian, uint32(len(m.InitialStates)))
	for _, e := range m.InitialStates {
		binary.Write(buf, binary.LittleEndian, uint32(len([]byte(e))))
		buf.Write([]byte(e))
	}

	// ActiveStates
	binary.Write(buf, binary.LittleEndian, uint32(len(m.ActiveStates)))
	for _, e := range m.ActiveStates {
		binary.Write(buf, binary.LittleEndian, uint32(len([]byte(e))))
		buf.Write([]byte(e))
	}

	// LocalData
	binary.Write(buf, binary.LittleEndian, uint32(len([]byte(m.LocalData))))
	buf.Write([]byte(m.LocalData))

	// Info
	binary.Write(buf, binary.LittleEndian, uint32(len([]byte(m.Info))))
	buf.Write([]byte(m.Info))

	return
}

func (m *String) Deserialize(buf *bytes.Reader) (err error) {
	// Header
	if err = m.Header.Deserialize(buf); err != nil {
		return err
	}

	// Path
	var size uint32
	if err = binary.Read(buf, binary.LittleEndian, &size); err != nil {
		return err
	}
	data := make([]byte, int(size))
	if err = binary.Read(buf, binary.LittleEndian, data); err != nil {
		return err
	}
	m.Path = string(data)

	// InitialStates
	var size uint32
	if err = binary.Read(buf, binary.LittleEndian, &size); err != nil {
		return err
	}
	m.InitialStates = make([]string, int(size))
	for i := 0; i < int(size); i++ {
		var size uint32
		if err = binary.Read(buf, binary.LittleEndian, &size); err != nil {
			return err
		}
		data := make([]byte, int(size))
		if err = binary.Read(buf, binary.LittleEndian, data); err != nil {
			return err
		}
		m.InitialStates[i] = string(data)
	}

	// ActiveStates
	var size uint32
	if err = binary.Read(buf, binary.LittleEndian, &size); err != nil {
		return err
	}
	m.ActiveStates = make([]string, int(size))
	for i := 0; i < int(size); i++ {
		var size uint32
		if err = binary.Read(buf, binary.LittleEndian, &size); err != nil {
			return err
		}
		data := make([]byte, int(size))
		if err = binary.Read(buf, binary.LittleEndian, data); err != nil {
			return err
		}
		m.ActiveStates[i] = string(data)
	}

	// LocalData
	var size uint32
	if err = binary.Read(buf, binary.LittleEndian, &size); err != nil {
		return err
	}
	data := make([]byte, int(size))
	if err = binary.Read(buf, binary.LittleEndian, data); err != nil {
		return err
	}
	m.LocalData = string(data)

	// Info
	var size uint32
	if err = binary.Read(buf, binary.LittleEndian, &size); err != nil {
		return err
	}
	data := make([]byte, int(size))
	if err = binary.Read(buf, binary.LittleEndian, data); err != nil {
		return err
	}
	m.Info = string(data)

	return
}
