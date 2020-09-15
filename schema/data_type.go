package schema

// the abstract DB type of this column
type DataType string

const (
	TypeChar               DataType = "char"
	TypeString             DataType = "string"
	TypeText               DataType = "text"
	TypeMediumText         DataType = "mediumText"
	TypeLongText           DataType = "longText"
	TypeTinyInt            DataType = "tinyInt"
	TypeSmallInt           DataType = "smallInt"
	TypeMediumInt          DataType = "mediumInt"
	TypeInt                DataType = "int"
	TypeBigInt             DataType = "bigInt"
	TypeTinyBlob           DataType = "tinyBlob"
	TypeBlob               DataType = "blob"
	TypeMediumBlob         DataType = "mediumBlob"
	TypeLongBlob           DataType = "longBlob"
	TypeFloat              DataType = "float"
	TypeDouble             DataType = "double"
	TypeDecimal            DataType = "decimal"
	TypeDateTime           DataType = "datetime"
	TypeDateTimeTz         DataType = "datetimeTz"
	TypeTimestamp          DataType = "timestamp"
	TypeTimestampTz        DataType = "timestampTz"
	TypeTime               DataType = "time"
	TypeTimeTz             DataType = "timeTz"
	TypeDate               DataType = "date"
	TypeYear               DataType = "year"
	TypeBinary             DataType = "binary"
	TypeBoolean            DataType = "boolean"
	TypeJson               DataType = "json"
	TypeJsonb              DataType = "jsonb"
	TypeEnum               DataType = "enum"
	TypeSet                DataType = "set"
	TypeUUID               DataType = "uuid"
	TypeIpAddress          DataType = "ipAddress"
	TypeMacAddress         DataType = "macAddress"
	TypeGeometry           DataType = "geometry"
	TypePoint              DataType = "point"
	TypeLineString         DataType = "lineString"
	TypePolygon            DataType = "polygon"
	TypeGeometryCollection DataType = "geometryCollection"
	TypeMultiPoint         DataType = "multiPoint"
	TypeMultiLineString    DataType = "multiLineString"
	TypeMultiPolygon       DataType = "multiPolygon"
	TypeMultiPolygonZ      DataType = "MultiPolygonZ"
)