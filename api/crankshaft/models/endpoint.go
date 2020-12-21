package models


type APIObservationEndpoint struct {
	EndpointID     string `json:"endpoint_id"`
	InConnections  uint64 `json:"in_connections"`
	OutConnections uint64 `json:"out_connections"`
	LocalResets    uint64 `json:"local_resets"`
	RemoteResets   uint64 `json:"remote_resets"`
}