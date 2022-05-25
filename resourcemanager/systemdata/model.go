package systemdata

type SystemData struct {
    CreatedBy          string `json:"createdBy"`
    CreatedByType      string `json:"createdByType"`
    CreatedAt          string `json:"createdAt"`
    LastModifiedBy     string `json:"lastModifiedBy"`
    LastModifiedbyType string `json:"lastModifiedbyType"`
    LastModifiedAt     string `json:"lastModifiedAt"`
}
