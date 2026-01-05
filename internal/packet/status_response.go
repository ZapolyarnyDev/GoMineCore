package packet

import (
	"encoding/json"

	"gominecore/internal/proto"
)

type StatusResponse struct {
	Version struct {
		Name     string `json:"name"`
		Protocol int32  `json:"protocol"`
	} `json:"version"`

	Players struct {
		Max    int `json:"max"`
		Online int `json:"online"`
	} `json:"players"`

	Description struct {
		Text string `json:"text"`
	} `json:"description"`
}

func NewStatusResponse() *StatusResponse {
	r := &StatusResponse{}

	r.Version.Name = "GoMineCore"
	r.Version.Protocol = 769

	r.Players.Max = 2026
	r.Players.Online = 0

	r.Description.Text = "§fСервер написанный §9Zapolyarny §fна §bGo"

	return r
}

func (p *StatusResponse) Encode(w *proto.ProtoIO) error {
	data, err := json.Marshal(p)
	if err != nil {
		return err
	}

	return w.WriteString(string(data))
}
