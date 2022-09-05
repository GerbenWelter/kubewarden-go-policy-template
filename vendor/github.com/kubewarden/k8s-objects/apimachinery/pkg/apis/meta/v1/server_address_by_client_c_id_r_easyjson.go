// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package v1

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonE7272ff5DecodeGithubComKubewardenK8sObjectsApimachineryPkgApisMetaV1(in *jlexer.Lexer, out *ServerAddressByClientCIDR) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "clientCIDR":
			if in.IsNull() {
				in.Skip()
				out.ClientCIDR = nil
			} else {
				if out.ClientCIDR == nil {
					out.ClientCIDR = new(string)
				}
				*out.ClientCIDR = string(in.String())
			}
		case "serverAddress":
			if in.IsNull() {
				in.Skip()
				out.ServerAddress = nil
			} else {
				if out.ServerAddress == nil {
					out.ServerAddress = new(string)
				}
				*out.ServerAddress = string(in.String())
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonE7272ff5EncodeGithubComKubewardenK8sObjectsApimachineryPkgApisMetaV1(out *jwriter.Writer, in ServerAddressByClientCIDR) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"clientCIDR\":"
		out.RawString(prefix[1:])
		if in.ClientCIDR == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.ClientCIDR))
		}
	}
	{
		const prefix string = ",\"serverAddress\":"
		out.RawString(prefix)
		if in.ServerAddress == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.ServerAddress))
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ServerAddressByClientCIDR) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonE7272ff5EncodeGithubComKubewardenK8sObjectsApimachineryPkgApisMetaV1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ServerAddressByClientCIDR) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonE7272ff5EncodeGithubComKubewardenK8sObjectsApimachineryPkgApisMetaV1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ServerAddressByClientCIDR) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonE7272ff5DecodeGithubComKubewardenK8sObjectsApimachineryPkgApisMetaV1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ServerAddressByClientCIDR) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonE7272ff5DecodeGithubComKubewardenK8sObjectsApimachineryPkgApisMetaV1(l, v)
}