//autogenerated:yes
//nolint:revive,misspell,govet,lll
package ardupilotmega

// Terrain data sent from GCS. The lat/lon and grid_spacing must be the same as a lat/lon from a TERRAIN_REQUEST. See terrain protocol docs: https://mavlink.io/en/services/terrain.html
type MessageTerrainData struct {
	// Latitude of SW corner of first grid
	Lat int32
	// Longitude of SW corner of first grid
	Lon int32
	// Grid spacing
	GridSpacing uint16
	// bit within the terrain request mask
	Gridbit uint8
	// Terrain data MSL
	Data [16]int16
}

// GetID implements the msg.Message interface.
func (*MessageTerrainData) GetID() uint32 {
	return 134
}
