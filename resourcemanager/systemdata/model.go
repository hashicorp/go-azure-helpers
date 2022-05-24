package systemdata

import "encoding/json"

var _ json.Unmarshaler = &SystemData{}

type SystemData struct {
	CreatedBy          string `json:"createdBy"`
	CreatedByType      string `json:"createdByType"`
	CreatedAt          string `json:"createdAt"`
	LastModifiedBy     string `json:"lastModifiedBy"`
	LastModifiedbyType string `json:"lastModifiedbyType"`
	LastModifiedAt     string `json:"lastModifiedAt"`
}

func (s *SystemData) UnmarshalJSON(bytes []byte) error {
	return json.Unmarshal(bytes, &s)
}
