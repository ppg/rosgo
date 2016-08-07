// Code generated by ros-gen-go.
// source: /opt/ros/jade/share/ros/../std_msgs/msg//UInt64MultiArray.msg
// DO NOT EDIT!
package std_msgs

type _MsgUInt64MultiArray struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgUInt64MultiArray) Text() string {
	return t.text
}

func (t *_MsgUInt64MultiArray) Name() string {
	return t.name
}

func (t *_MsgUInt64MultiArray) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgUInt64MultiArray) NewMessage() ros.Message {
	m := new(UInt64MultiArray)

	return m
}

var (
	MsgUInt64MultiArray = &_MsgUInt64MultiArray{
		`# Please look at the MultiArrayLayout message definition for
# documentation on all multiarrays.

MultiArrayLayout  layout        # specification of data layout
uint64[]          data          # array of data

`,
		"std_msgs/UInt64MultiArray",
		"61643890115c64ea151267f0dcdc4f3e",
	}
)

type UInt64MultiArray struct {
	Layout MultiArrayLayout
	Data   []uint64
}

func (m *UInt64MultiArray) Type() ros.MessageType {
	return MsgUInt64MultiArray
}

func (m *UInt64MultiArray) Serialize(buf *bytes.Buffer) (err error) {
	// Layout
	if err = m.Layout.Serialize(buf); err != nil {
		return err
	}

	// Data
	binary.Write(buf, binary.LittleEndian, uint32(len(m.Data)))
	for _, e := range m.Data {
		binary.Write(buf, binary.LittleEndian, e)
	}

	return
}

func (m *String) Deserialize(buf *bytes.Reader) (err error) {
	// Layout
	if err = m.Layout.Deserialize(buf); err != nil {
		return err
	}

	// Data
	var size uint32
	if err = binary.Read(buf, binary.LittleEndian, &size); err != nil {
		return err
	}
	m.Data = make([]uint64, int(size))
	for i := 0; i < int(size); i++ {
		if err = binary.Read(buf, binary.LittleEndian, &m.Data[i]); err != nil {
			return err
		}
	}

	return
}
