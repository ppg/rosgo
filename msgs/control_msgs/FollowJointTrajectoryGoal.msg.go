// Code generated by ros-gen-go.
// source: /opt/ros/jade/share/ros/../control_msgs/msg//FollowJointTrajectoryGoal.msg
// DO NOT EDIT!
package control_msgs

type _MsgFollowJointTrajectoryGoal struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgFollowJointTrajectoryGoal) Text() string {
	return t.text
}

func (t *_MsgFollowJointTrajectoryGoal) Name() string {
	return t.name
}

func (t *_MsgFollowJointTrajectoryGoal) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgFollowJointTrajectoryGoal) NewMessage() ros.Message {
	m := new(FollowJointTrajectoryGoal)

	return m
}

var (
	MsgFollowJointTrajectoryGoal = &_MsgFollowJointTrajectoryGoal{
		`# ====== DO NOT MODIFY! AUTOGENERATED FROM AN ACTION DEFINITION ======
# The joint trajectory to follow
trajectory_msgs/JointTrajectory trajectory

# Tolerances for the trajectory.  If the measured joint values fall
# outside the tolerances the trajectory goal is aborted.  Any
# tolerances that are not specified (by being omitted or set to 0) are
# set to the defaults for the action server (often taken from the
# parameter server).

# Tolerances applied to the joints as the trajectory is executed.  If
# violated, the goal aborts with error_code set to
# PATH_TOLERANCE_VIOLATED.
JointTolerance[] path_tolerance

# To report success, the joints must be within goal_tolerance of the
# final trajectory value.  The goal must be achieved by time the
# trajectory ends plus goal_time_tolerance.  (goal_time_tolerance
# allows some leeway in time, so that the trajectory goal can still
# succeed even if the joints reach the goal some time after the
# precise end time of the trajectory).
#
# If the joints are not within goal_tolerance after "trajectory finish
# time" + goal_time_tolerance, the goal aborts with error_code set to
# GOAL_TOLERANCE_VIOLATED
JointTolerance[] goal_tolerance
duration goal_time_tolerance

`,
		"control_msgs/FollowJointTrajectoryGoal",
		"d6e925ca987733b7452947e90e68dd68",
	}
)

type FollowJointTrajectoryGoal struct {
	Trajectory        trajectory_msgs.JointTrajectory
	PathTolerance     []JointTolerance
	GoalTolerance     []JointTolerance
	GoalTimeTolerance ros.Duration
}

func (m *FollowJointTrajectoryGoal) Type() ros.MessageType {
	return MsgFollowJointTrajectoryGoal
}

func (m *FollowJointTrajectoryGoal) Serialize(buf *bytes.Buffer) (err error) {
	// Trajectory
	if err = m.Trajectory.Serialize(buf); err != nil {
		return err
	}

	// PathTolerance
	binary.Write(buf, binary.LittleEndian, uint32(len(m.PathTolerance)))
	for _, e := range m.PathTolerance {
		if err = e.Serialize(buf); err != nil {
			return err
		}
	}

	// GoalTolerance
	binary.Write(buf, binary.LittleEndian, uint32(len(m.GoalTolerance)))
	for _, e := range m.GoalTolerance {
		if err = e.Serialize(buf); err != nil {
			return err
		}
	}

	// GoalTimeTolerance
	binary.Write(buf, binary.LittleEndian, m.GoalTimeTolerance.Sec)
	binary.Write(buf, binary.LittleEndian, m.GoalTimeTolerance.NSec)

	return
}

func (m *String) Deserialize(buf *bytes.Reader) (err error) {
	// Trajectory
	if err = m.Trajectory.Deserialize(buf); err != nil {
		return err
	}

	// PathTolerance
	var size uint32
	if err = binary.Read(buf, binary.LittleEndian, &size); err != nil {
		return err
	}
	m.PathTolerance = make([]JointTolerance, int(size))
	for i := 0; i < int(size); i++ {
		if err = m.PathTolerance[i].Deserialize(buf); err != nil {
			return err
		}
	}

	// GoalTolerance
	var size uint32
	if err = binary.Read(buf, binary.LittleEndian, &size); err != nil {
		return err
	}
	m.GoalTolerance = make([]JointTolerance, int(size))
	for i := 0; i < int(size); i++ {
		if err = m.GoalTolerance[i].Deserialize(buf); err != nil {
			return err
		}
	}

	// GoalTimeTolerance
	if err = binary.Read(buf, binary.LittleEndian, &m.GoalTimeTolerance.Sec); err != nil {
		return err
	}
	if err = binary.Read(buf, binary.LittleEndian, &m.GoalTimeTolerance.NSec); err != nil {
		return err
	}

	return
}
