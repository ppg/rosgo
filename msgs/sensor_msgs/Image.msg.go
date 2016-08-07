// Code generated by ros-gen-go.
// source: /opt/ros/jade/share/ros/../sensor_msgs/msg//Image.msg
// DO NOT EDIT!
package sensor_msgs

type _MsgImage struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgImage) Text() string {
	return t.text
}

func (t *_MsgImage) Name() string {
	return t.name
}

func (t *_MsgImage) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgImage) NewMessage() ros.Message {
	m := new(Image)

	return m
}

var (
	MsgImage = &_MsgImage{
		`# This message contains an uncompressed image
# (0, 0) is at top-left corner of image
#

Header header        # Header timestamp should be acquisition time of image
                     # Header frame_id should be optical frame of camera
                     # origin of frame should be optical center of cameara
                     # +x should point to the right in the image
                     # +y should point down in the image
                     # +z should point into to plane of the image
                     # If the frame_id here and the frame_id of the CameraInfo
                     # message associated with the image conflict
                     # the behavior is undefined

uint32 height         # image height, that is, number of rows
uint32 width          # image width, that is, number of columns

# The legal values for encoding are in file src/image_encodings.cpp
# If you want to standardize a new string format, join
# ros-users@lists.sourceforge.net and send an email proposing a new encoding.

string encoding       # Encoding of pixels -- channel meaning, ordering, size
                      # taken from the list of strings in include/sensor_msgs/image_encodings.h

uint8 is_bigendian    # is this data bigendian?
uint32 step           # Full row length in bytes
uint8[] data          # actual matrix data, size is (step * rows)
`,
		"sensor_msgs/Image",
		"4233ee3710fbed0dbbbbde74c8e1a251",
	}
)

type Image struct {
	Header      Header
	Height      uint32
	Width       uint32
	Encoding    string
	IsBigendian uint8
	Step        uint32
	Data        []uint8
}

func (m *Image) Type() ros.MessageType {
	return MsgImage
}

func (m *Image) Serialize(buf *bytes.Buffer) (err error) {
	// Header
	if err = m.Header.Serialize(buf); err != nil {
		return err
	}

	// Height
	binary.Write(buf, binary.LittleEndian, m.Height)

	// Width
	binary.Write(buf, binary.LittleEndian, m.Width)

	// Encoding
	binary.Write(buf, binary.LittleEndian, uint32(len([]byte(m.Encoding))))
	buf.Write([]byte(m.Encoding))

	// IsBigendian
	binary.Write(buf, binary.LittleEndian, m.IsBigendian)

	// Step
	binary.Write(buf, binary.LittleEndian, m.Step)

	// Data
	binary.Write(buf, binary.LittleEndian, uint32(len(m.Data)))
	for _, e := range m.Data {
		binary.Write(buf, binary.LittleEndian, e)
	}

	return
}

func (m *String) Deserialize(buf *bytes.Reader) (err error) {
	// Header
	if err = m.Header.Deserialize(buf); err != nil {
		return err
	}

	// Height
	if err = binary.Read(buf, binary.LittleEndian, &m.Height); err != nil {
		return err
	}

	// Width
	if err = binary.Read(buf, binary.LittleEndian, &m.Width); err != nil {
		return err
	}

	// Encoding
	var size uint32
	if err = binary.Read(buf, binary.LittleEndian, &size); err != nil {
		return err
	}
	data := make([]byte, int(size))
	if err = binary.Read(buf, binary.LittleEndian, data); err != nil {
		return err
	}
	m.Encoding = string(data)

	// IsBigendian
	if err = binary.Read(buf, binary.LittleEndian, &m.IsBigendian); err != nil {
		return err
	}

	// Step
	if err = binary.Read(buf, binary.LittleEndian, &m.Step); err != nil {
		return err
	}

	// Data
	var size uint32
	if err = binary.Read(buf, binary.LittleEndian, &size); err != nil {
		return err
	}
	m.Data = make([]uint8, int(size))
	for i := 0; i < int(size); i++ {
		if err = binary.Read(buf, binary.LittleEndian, &m.Data[i]); err != nil {
			return err
		}
	}

	return
}
