package lumberjack

type RotateType string

var (
	RotateDateNotNeed RotateType = ""
	RotateDaily       RotateType = "daily"
	RotateHourly      RotateType = "hourly"
	RotateSize        RotateType = "size"
)

func IsLegalRotateType(t RotateType) bool {
	return t == RotateDateNotNeed || t == RotateDaily || t == RotateHourly || t == RotateSize
}
