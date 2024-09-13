package structs

type Device struct {
	Id          string `json:"id"`
	DeviceId    string `json:"device_id"`
	CreatedAt   string `json:"created_at"`
	Temperature int    `json:"temperature"`
	Warning     string `json:"warning"`
}
