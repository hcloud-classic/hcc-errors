package hcc_errors

const (
	// On
	CelloInternalInitFail                 = cello + internal + initFail
	CelloInternalOperationFail            = cello + internal + operationFail
	CelloInternalUUIDGenerationError      = cello + internal + UUIDGenerationError
	CelloInternalTimeStampConversionError = cello + internal + timestampConversionError

	CelloInternalVolumeHandleError = cello + internal + volumehandlerError
	CelloInternalPoolHandleError   = cello + internal + poolhandlerError
	CelloInternalStoragePoolError  = cello + internal + storagepoolError
	CelloInternalPrepareDeploy     = cello + internal + preparedeployError
	CelloInternalReloadObject      = cello + internal + reloadobjectError
	CelloInternalAction            = cello + internal + actionError

	CelloGrpcArgumentError = cello + grpc + argumentError
	CelloGrpcRequestError  = cello + grpc + requestError

	CelloSQLOperationFail = cello + sql + operationFail
	CelloSQLNoResult      = cello + sql + noResult
	CelloSQLArgumentError = cello + sql + argumentError
)
