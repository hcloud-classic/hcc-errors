package hcc_errors

const (
	PianoInternalInitFail       uint64 = piano + internal + initFail
	PianoInternalConnectionFail uint64 = piano + internal + connectionFail
	PianoInternalParsingError   uint64 = piano + internal + parsingError
	PianoInternalOperationFail  uint64 = piano + internal + operationFail

	PianoGrpcArgumentError uint64 = piano + grpc + argumentError

	PianoInfluxDBReadMetricError uint64 = piano + influxDB + readMetricError
)
