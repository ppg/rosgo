// Code generated by ros-gen-go.
// source: SetBool.srv
// DO NOT EDIT!
package std_srvs

import (
	"io"

	"github.com/ppg/rosgo/ros"
)

// Service type metadata
type _SrvSetBool struct {
	name    string
	md5sum  string
	text    string
	reqType ros.MessageType
	resType ros.MessageType
}

func (t *_SrvSetBool) Name() string                  { return t.name }
func (t *_SrvSetBool) MD5Sum() string                { return t.md5sum }
func (t *_SrvSetBool) Text() string                  { return t.text }
func (t *_SrvSetBool) RequestType() ros.MessageType  { return t.reqType }
func (t *_SrvSetBool) ResponseType() ros.MessageType { return t.resType }
func (t *_SrvSetBool) NewService() ros.Service {
	return new(SetBool)
}

var (
	SrvSetBool = &_SrvSetBool{
		"std_srvs/SetBool",
		"e3669e693e08eccee0f27d72171188a9",
		`bool data # e.g. for hardware enabling / disabling
---
bool success   # indicate successful run of triggered service
string message # informational, e.g. for error messages
`,
		MsgSetBoolRequest,
		MsgSetBoolResponse,
	}
)

type SetBool struct {
	Request  SetBoolRequest
	Response SetBoolResponse
}

func (s *SetBool) ReqMessage() ros.Message { return &s.Request }
func (s *SetBool) ResMessage() ros.Message { return &s.Response }

// SetBoolRequest

type _MsgSetBoolRequest struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgSetBoolRequest) Text() string {
	return t.text
}

func (t *_MsgSetBoolRequest) Name() string {
	return t.name
}

func (t *_MsgSetBoolRequest) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgSetBoolRequest) NewMessage() ros.Message {
	m := new(SetBoolRequest)

	return m
}

var (
	MsgSetBoolRequest = &_MsgSetBoolRequest{
		`bool data # e.g. for hardware enabling / disabling
`,
		"std_srvs/SetBoolRequest",
		"e3669e693e08eccee0f27d72171188a9",
	}
)

type SetBoolRequest struct {
	Data bool
}

func (m *SetBoolRequest) Serialize(w io.Writer) (err error) {
	if err = ros.SerializeMessageField(w, "bool", &m.Data); err != nil {
		return err
	}

	return
}

func (m *SetBoolRequest) Deserialize(r io.Reader) (err error) {
	// Data
	if err = ros.DeserializeMessageField(r, "bool", &m.Data); err != nil {
		return err
	}

	return
}

// SetBoolResponse

type _MsgSetBoolResponse struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgSetBoolResponse) Text() string {
	return t.text
}

func (t *_MsgSetBoolResponse) Name() string {
	return t.name
}

func (t *_MsgSetBoolResponse) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgSetBoolResponse) NewMessage() ros.Message {
	m := new(SetBoolResponse)

	return m
}

var (
	MsgSetBoolResponse = &_MsgSetBoolResponse{
		`
bool success   # indicate successful run of triggered service
string message # informational, e.g. for error messages
`,
		"std_srvs/SetBoolResponse",
		"7bec9fbe319affc86fe89a5114580b78",
	}
)

type SetBoolResponse struct {
	Success bool
	Message string
}

func (m *SetBoolResponse) Serialize(w io.Writer) (err error) {
	if err = ros.SerializeMessageField(w, "bool", &m.Success); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "string", &m.Message); err != nil {
		return err
	}

	return
}

func (m *SetBoolResponse) Deserialize(r io.Reader) (err error) {
	// Success
	if err = ros.DeserializeMessageField(r, "bool", &m.Success); err != nil {
		return err
	}

	// Message
	if err = ros.DeserializeMessageField(r, "string", &m.Message); err != nil {
		return err
	}

	return
}
