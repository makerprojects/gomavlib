//autogenerated:yes
//nolint:revive,misspell,govet,lll
package matrixpilot

// Configure WiFi AP SSID, password, and mode. This message is re-emitted as an acknowledgement by the AP. The message may also be explicitly requested using MAV_CMD_REQUEST_MESSAGE
type MessageWifiConfigAp struct {
	// Name of Wi-Fi network (SSID). Blank to leave it unchanged when setting. Current SSID when sent back as a response.
	Ssid string `mavlen:"32"`
	// Password. Blank for an open AP. MD5 hash when message is sent back as a response.
	Password string `mavlen:"64"`
	// WiFi Mode.
	Mode WIFI_CONFIG_AP_MODE `mavenum:"int8" mavext:"true"`
	// Message acceptance response (sent back to GS).
	Response WIFI_CONFIG_AP_RESPONSE `mavenum:"int8" mavext:"true"`
}

// GetID implements the msg.Message interface.
func (*MessageWifiConfigAp) GetID() uint32 {
	return 299
}
