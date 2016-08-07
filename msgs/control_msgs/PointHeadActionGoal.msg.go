// Code generated by ros-gen-go.
// source: /opt/ros/jade/share/ros/../control_msgs/msg//PointHeadActionGoal.msg
// DO NOT EDIT!
package control_msgs

type _MsgPointHeadActionGoal struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgPointHeadActionGoal) Text() string {
	return t.text
}

func (t *_MsgPointHeadActionGoal) Name() string {
	return t.name
}

func (t *_MsgPointHeadActionGoal) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgPointHeadActionGoal) NewMessage() ros.Message {
	m := new(PointHeadActionGoal)

	return m
}

var (
	MsgPointHeadActionGoal = &_MsgPointHeadActionGoal{
		`# ====== DO NOT MODIFY! AUTOGENERATED FROM AN ACTION DEFINITION ======

Header header
actionlib_msgs/GoalID goal_id
PointHeadGoal goal
`,
		"control_msgs/PointHeadActionGoal",
		"d2151cff361301782a6ae486c4db45ab",
	}
)

type PointHeadActionGoal struct {
	Header Header
	GoalID actionlib_msgs.GoalID
	Goal   PointHeadGoal
}

func (m *PointHeadActionGoal) Type() ros.MessageType {
	return MsgPointHeadActionGoal
}

func (m *PointHeadActionGoal) Serialize(buf *bytes.Buffer) (err error) {
	// Header
	if err = m.Header.Serialize(buf); err != nil {
		return err
	}

	// GoalID
	if err = m.GoalID.Serialize(buf); err != nil {
		return err
	}

	// Goal
	if err = m.Goal.Serialize(buf); err != nil {
		return err
	}

	return
}

func (m *String) Deserialize(buf *bytes.Reader) (err error) {
	// Header
	if err = m.Header.Deserialize(buf); err != nil {
		return err
	}

	// GoalID
	if err = m.GoalID.Deserialize(buf); err != nil {
		return err
	}

	// Goal
	if err = m.Goal.Deserialize(buf); err != nil {
		return err
	}

	return
}
