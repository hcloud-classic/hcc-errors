package hcc_errors

const (

	// code for MiddleWare
	modulename uint64 = (1 + iota) * 10000
	cello
	clarinet
	flute
	harp
	oboe
	piano
	piccolo
	viola
	violin
	violinNoVNC
	violinScheduler
)

const (

	// code for Category
	category uint64 = (1 + iota) * 1000
	internal        // lib
	driver          // driver
	graphql         // action
	grpc
	sql
	influxDB
	rabbitmq
)

const (

	// code for Describe Error
	// Use Generally
	operation uint64 = 1 + iota
	initFail
	connectionFail
	undefinedError
	argumentError
	jsonMarshalError
	jsonUnmarshalError
	requestError  // send Request fail
	responseError // get Response fail or has error
	sendError     // send error to client
	receiveError  // get error as result from server
	parsingError
	invalidToken
	tokenExpired
	operationFail
	noResult
	timestampConversionError
	UUIDGenerationError

	// clarinet specific

	// piccolo specific
	prepareError
	executeError
	tokenGenerationError
	loginFailed
	userExist

	// cello specific
	volumehandlerError
	poolhandlerError
	storagepoolError
	preparedeployError
	reloadobjectError
	actionError

	// violin-scheduler specific
	schedulehandlerError

	// flute specific
	ipmiError

	// viola specific

	// piano specific
	readMetricError

	// harp specific
	interfaceAddrLookupError
	pfError
	dhcpdError
	fileError
	ifconfigError
	IPAddressError
	subnetInUseError
	subnetNotAllocatedError
	adaptiveIPAllocatedError

	// violin-novnc specific

	// violin specific
	createServerFailed
	createServerRoutineError
	getAvailableNodesError
	getNodesError
	serverNodePresentError
)
