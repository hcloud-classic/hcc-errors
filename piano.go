package hcc_errors

const (
	PianoInternalInitFail           uint64 = piano + internal + initFail
	PianoInternalConnectionFail     uint64 = piano + internal + connectionFail
	PianoInternalParsingError       uint64 = piano + internal + parsingError
	PianoInternalOperationFail      uint64 = piano + internal + operationFail
	PianoInternalJsonMarshalError   uint64 = piano + internal + jsonMarshalError
	PianoInternalJsonUnmarshalError uint64 = piano + internal + jsonUnmarshalError

	PianoGrpcArgumentError  uint64 = piano + grpc + argumentError
	PianoGrpcConnectionFail uint64 = piano + grpc + connectionFail

	PianoSQLOperationFail uint64 = piano + sql + operationFail

	PianoInfluxDBReadMetricError uint64 = piano + influxDB + readMetricError
)
