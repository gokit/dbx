package schema

import (
	"errors"
	"fmt"
	"strings"
)

// column type convert func
type ColumnTypeConverter func(column Column) (string, error)

// get dialect ColumnType func by dialect and column type
func ColumnTypeMapper(d Dialect, dataType DataType) (ColumnTypeConverter, error) {

	var typeFunc ColumnTypeConverter
	var err error

	switch dataType {
	case TypeChar:
		typeFunc = d.TypeChar
	case TypeString:
		typeFunc = d.TypeString
	case TypeText:
		typeFunc = d.TypeText
	case TypeMediumText:
		typeFunc = d.TypeMediumText
	case TypeLongText:
		typeFunc = d.TypeLongText
	case TypeInt:
		typeFunc = d.TypeInteger
	case TypeTinyInt:
		typeFunc = d.TypeTinyInteger
	case TypeSmallInt:
		typeFunc = d.TypeSmallInteger
	case TypeMediumInt:
		typeFunc = d.TypeMediumInteger
	case TypeBigInt:
		typeFunc = d.TypeBigInteger
	case TypeTinyBlob:
		typeFunc = d.TypeTinyBlob
	case TypeBlob:
		typeFunc = d.TypeBlob
	case TypeMediumBlob:
		typeFunc = d.TypeMediumBlob
	case TypeLongBlob:
		typeFunc = d.TypeLongBlob
	case TypeFloat:
		typeFunc = d.TypeFloat
	case TypeDouble:
		typeFunc = d.TypeDouble
	case TypeDecimal:
		typeFunc = d.TypeDecimal
	case TypeBoolean:
		typeFunc = d.TypeBoolean
	case TypeEnum:
		typeFunc = d.TypeEnum
	case TypeSet:
		typeFunc = d.TypeSet
	case TypeJson:
		typeFunc = d.TypeJson
	case TypeJsonb:
		typeFunc = d.TypeJsonb
	case TypeDate:
		typeFunc = d.TypeDate
	case TypeDateTime:
		typeFunc = d.TypeDateTime
	case TypeDateTimeTz:
		typeFunc = d.TypeDateTimeTz
	case TypeTime:
		typeFunc = d.TypeTime
	case TypeTimeTz:
		typeFunc = d.TypeTimeTz
	case TypeTimestamp:
		typeFunc = d.TypeTimestamp
	case TypeTimestampTz:
		typeFunc = d.TypeTimestampTz
	case TypeYear:
		typeFunc = d.TypeYear
	case TypeBinary:
		typeFunc = d.TypeBinary
	case TypeUUID:
		typeFunc = d.TypeUuid
	case TypeIpAddress:
		typeFunc = d.TypeIpAddress
	case TypeMacAddress:
		typeFunc = d.TypeMacAddress
	case TypeGeometry:
		typeFunc = d.TypeGeometry
	case TypePoint:
		typeFunc = d.TypePoint
	case TypeLineString:
		typeFunc = d.TypeLineString
	case TypePolygon:
		typeFunc = d.TypePolygon
	case TypeGeometryCollection:
		typeFunc = d.TypeGeometryCollection
	case TypeMultiPoint:
		typeFunc = d.TypeMultiPoint
	case TypeMultiLineString:
		typeFunc = d.TypeMultiLineString
	case TypeMultiPolygon:
		typeFunc = d.TypeMultiPolygon
	default:
		err = errors.New(fmt.Sprintf("unsupported data type %s", dataType))
	}

	return typeFunc, err
}

func ColumnType(d Dialect, column Column) (string, error) {
	typeFunc, err := ColumnTypeMapper(d, column.DataType())
	if err != nil {
		return "", err
	}
	return typeFunc(column)
}

// Convert the column type to abstract data type
func DataTypeMapper(columnType string) (DataType, error) {

	var dataType DataType
	var err error

	switch strings.ToLower(columnType) {
	case "char":
		dataType = TypeChar
	case "string", "varchar":
		dataType = TypeString
	case "text":
		dataType = TypeText
	case "mediumtext":
		dataType = TypeMediumText
	case "longtext":
		dataType = TypeLongText
	case "int", "integer", "int4", "int2":
		dataType = TypeInt
	case "tinyint", "tinyinteger":
		dataType = TypeTinyInt
	case "smallint", "smallinteger":
		dataType = TypeSmallInt
	case "mediumint", "mediuminteger":
		dataType = TypeMediumInt
	case "bigint", "biginteger", "int8":
		dataType = TypeBigInt
	case "blob":
		dataType = TypeBlob
	case "tinyblob":
		dataType = TypeTinyBlob
	case "mediumblob":
		dataType = TypeMediumBlob
	case "longblob":
		dataType = TypeLongBlob
	case "float":
		dataType = TypeFloat
	case "double":
		dataType = TypeDouble
	case "decimal":
		dataType = TypeDecimal
	case "boolean":
		dataType = TypeBoolean
	case "enum":
		dataType = TypeEnum
	case "set":
		dataType = TypeSet
	case "json":
		dataType = TypeJson
	case "jsonb":
		dataType = TypeJsonb
	case "date":
		dataType = TypeDate
	case "datetime":
		dataType = TypeDateTime
	case "datetimetz":
		dataType = TypeDateTimeTz
	case "time":
		dataType = TypeTime
	case "timetz":
		dataType = TypeTimeTz
	case "timestamp":
		dataType = TypeTimestamp
	case "timestamptz":
		dataType = TypeTimestampTz
	case "year":
		dataType = TypeYear
	case "binary":
		dataType = TypeBinary
	case "uuid":
		dataType = TypeUUID
	case "ipaddress":
		dataType = TypeIpAddress
	case "macaddress":
		dataType = TypeMacAddress
	case "geometry":
		dataType = TypeGeometry
	case "point":
		dataType = TypePoint
	case "linestring":
		dataType = TypeLineString
	case "polygon":
		dataType = TypePolygon
	case "geometrycollection":
		dataType = TypeGeometryCollection
	case "multipoint":
		dataType = TypeMultiPoint
	case "multiLinestring":
		dataType = TypeMultiLineString
	case "multiPolygon":
		dataType = TypeMultiPolygon
	default:
		err = errors.New(fmt.Sprintf("dbx: unsupported column type '%s'", columnType))
	}
	return dataType, err
}
