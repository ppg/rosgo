// Code generated by ros-gen-go.
// source: /opt/ros/jade/share/ros/../sensor_msgs/msg//MultiEchoLaserScan.msg
// DO NOT EDIT!
package sensor_msgs

type _MsgMultiEchoLaserScan struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgMultiEchoLaserScan) Text() string {
	return t.text
}

func (t *_MsgMultiEchoLaserScan) Name() string {
	return t.name
}

func (t *_MsgMultiEchoLaserScan) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgMultiEchoLaserScan) NewMessage() ros.Message {
	m := new(MultiEchoLaserScan)

	return m
}

var (
	MsgMultiEchoLaserScan = &_MsgMultiEchoLaserScan{
		`# Single scan from a multi-echo planar laser range-finder
#
# If you have another ranging device with different behavior (e.g. a sonar
# array), please find or create a different message, since applications
# will make fairly laser-specific assumptions about this data

Header header            # timestamp in the header is the acquisition time of 
                         # the first ray in the scan.
                         #
                         # in frame frame_id, angles are measured around 
                         # the positive Z axis (counterclockwise, if Z is up)
                         # with zero angle being forward along the x axis
                         
float32 angle_min        # start angle of the scan [rad]
float32 angle_max        # end angle of the scan [rad]
float32 angle_increment  # angular distance between measurements [rad]

float32 time_increment   # time between measurements [seconds] - if your scanner
                         # is moving, this will be used in interpolating position
                         # of 3d points
float32 scan_time        # time between scans [seconds]

float32 range_min        # minimum range value [m]
float32 range_max        # maximum range value [m]

LaserEcho[] ranges       # range data [m] (Note: NaNs, values < range_min or > range_max should be discarded)
                         # +Inf measurements are out of range
                         # -Inf measurements are too close to determine exact distance.
LaserEcho[] intensities  # intensity data [device-specific units].  If your
                         # device does not provide intensities, please leave
                         # the array empty.`,
		"sensor_msgs/MultiEchoLaserScan",
		"b80691870c47673fb3c9bf3c06e0a0ab",
	}
)

type MultiEchoLaserScan struct {
	Header         Header
	AngleMin       float32
	AngleMax       float32
	AngleIncrement float32
	TimeIncrement  float32
	ScanTime       float32
	RangeMin       float32
	RangeMax       float32
	Ranges         []LaserEcho
	Intensities    []LaserEcho
}

func (m *MultiEchoLaserScan) Type() ros.MessageType {
	return MsgMultiEchoLaserScan
}

func (m *MultiEchoLaserScan) Serialize(buf *bytes.Buffer) (err error) {
	// Header
	if err = m.Header.Serialize(buf); err != nil {
		return err
	}

	// AngleMin
	binary.Write(buf, binary.LittleEndian, m.AngleMin)

	// AngleMax
	binary.Write(buf, binary.LittleEndian, m.AngleMax)

	// AngleIncrement
	binary.Write(buf, binary.LittleEndian, m.AngleIncrement)

	// TimeIncrement
	binary.Write(buf, binary.LittleEndian, m.TimeIncrement)

	// ScanTime
	binary.Write(buf, binary.LittleEndian, m.ScanTime)

	// RangeMin
	binary.Write(buf, binary.LittleEndian, m.RangeMin)

	// RangeMax
	binary.Write(buf, binary.LittleEndian, m.RangeMax)

	// Ranges
	binary.Write(buf, binary.LittleEndian, uint32(len(m.Ranges)))
	for _, e := range m.Ranges {
		if err = e.Serialize(buf); err != nil {
			return err
		}
	}

	// Intensities
	binary.Write(buf, binary.LittleEndian, uint32(len(m.Intensities)))
	for _, e := range m.Intensities {
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

	// AngleMin
	if err = binary.Read(buf, binary.LittleEndian, &m.AngleMin); err != nil {
		return err
	}

	// AngleMax
	if err = binary.Read(buf, binary.LittleEndian, &m.AngleMax); err != nil {
		return err
	}

	// AngleIncrement
	if err = binary.Read(buf, binary.LittleEndian, &m.AngleIncrement); err != nil {
		return err
	}

	// TimeIncrement
	if err = binary.Read(buf, binary.LittleEndian, &m.TimeIncrement); err != nil {
		return err
	}

	// ScanTime
	if err = binary.Read(buf, binary.LittleEndian, &m.ScanTime); err != nil {
		return err
	}

	// RangeMin
	if err = binary.Read(buf, binary.LittleEndian, &m.RangeMin); err != nil {
		return err
	}

	// RangeMax
	if err = binary.Read(buf, binary.LittleEndian, &m.RangeMax); err != nil {
		return err
	}

	// Ranges
	var size uint32
	if err = binary.Read(buf, binary.LittleEndian, &size); err != nil {
		return err
	}
	m.Ranges = make([]LaserEcho, int(size))
	for i := 0; i < int(size); i++ {
		if err = m.Ranges[i].Deserialize(buf); err != nil {
			return err
		}
	}

	// Intensities
	var size uint32
	if err = binary.Read(buf, binary.LittleEndian, &size); err != nil {
		return err
	}
	m.Intensities = make([]LaserEcho, int(size))
	for i := 0; i < int(size); i++ {
		if err = m.Intensities[i].Deserialize(buf); err != nil {
			return err
		}
	}

	return
}
