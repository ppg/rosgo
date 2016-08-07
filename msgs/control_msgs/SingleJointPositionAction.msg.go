// Code generated by ros-gen-go.
// source: /opt/ros/jade/share/ros/../control_msgs/msg//SingleJointPositionAction.msg
// DO NOT EDIT!
package control_msgs

type _MsgSingleJointPositionAction struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgSingleJointPositionAction) Text() string {
	return t.text
}

func (t *_MsgSingleJointPositionAction) Name() string {
	return t.name
}

func (t *_MsgSingleJointPositionAction) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgSingleJointPositionAction) NewMessage() ros.Message {
	m := new(SingleJointPositionAction)

	return m
}

var (
	MsgSingleJointPositionAction = &_MsgSingleJointPositionAction{
		`# ====== DO NOT MODIFY! AUTOGENERATED FROM AN ACTION DEFINITION ======

SingleJointPositionActionGoal action_goal
SingleJointPositionActionResult action_result
SingleJointPositionActionFeedback action_feedback
`,
		"control_msgs/SingleJointPositionAction",
		"d31fd05f88af960673f983d86a899f4d",
	}
)

type SingleJointPositionAction struct {
	ActionGoal     SingleJointPositionActionGoal
	ActionResult   SingleJointPositionActionResult
	ActionFeedback SingleJointPositionActionFeedback
}

func (m *SingleJointPositionAction) Type() ros.MessageType {
	return MsgSingleJointPositionAction
}

func (m *SingleJointPositionAction) Serialize(buf *bytes.Buffer) (err error) {
	// ActionGoal
	if err = m.ActionGoal.Serialize(buf); err != nil {
		return err
	}

	// ActionResult
	if err = m.ActionResult.Serialize(buf); err != nil {
		return err
	}

	// ActionFeedback
	if err = m.ActionFeedback.Serialize(buf); err != nil {
		return err
	}

	return
}

func (m *String) Deserialize(buf *bytes.Reader) (err error) {
	// ActionGoal
	if err = m.ActionGoal.Deserialize(buf); err != nil {
		return err
	}

	// ActionResult
	if err = m.ActionResult.Deserialize(buf); err != nil {
		return err
	}

	// ActionFeedback
	if err = m.ActionFeedback.Deserialize(buf); err != nil {
		return err
	}

	return
}
