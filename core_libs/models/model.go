package models

import "fmt"

type DeviceCreate struct {
	Token      string  `json:"token"`
	MacAddress *string `json:"mac_address"`
}
type MacAddress struct {
	MacAddress *string `json:"mac_address"`
}

type DeviceDelete struct {
	ThingID   string `json:"thing_id"`
	ChannelID string `json:"channel_id"`
}

func (r DeviceDelete) Validate() error {
	if r.ThingID == "" {
		return fmt.Errorf("ThingID empty")
	}

	if r.ChannelID == "" {
		return fmt.Errorf("ChannelID empty")
	}

	return nil
}

type ThingRequest struct {
	Key      *string   `json:"key"`
	Name     *string   `json:"name"`
	Metadata *Metadata `json:"metadata"`
}
type ThingResponse struct {
	ID       *string   `json:"id"`
	Key      *string   `json:"key"`
	Name     *string   `json:"name"`
	Metadata *Metadata `json:"metadata"`
}
type ThingMainflux struct {
	ThingID  string `json:"thing_id"`
	ThingKey string `json:"thing_key"`
}
type ResponseCreateDevice struct {
	ThingID   string `json:"thing_id"`
	ThingKey  string `json:"thing_key"`
	ChannelID string `json:"channel_id"`
}
type Metadata struct {
	Type string `json:"type"`
}
type ChannelRequest struct {
	Name     string    `json:"name"`
	Metadata *Metadata `json:"metadata"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
