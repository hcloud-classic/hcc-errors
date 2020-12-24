package hcc_errors

var middleWareList = [...]string{
	"",
	// string for MiddleWare
	"-Module Name-",
	"Cello",
	"Clarinet",
	"Flute",
	"Harp",
	"Oboe",
	"Piano",
	"Piccolo",
	"Viola",
	"Violin",
	"NoVNC",
	"Scheduler",
}

var functionList = [...]string{
	"",
	// string for Category
	"-Category-",
	"Internal",
	"Driver",
	"GraphQL",
	"Grpc",
	"SQL",
	"InfluxDB",
	"RabbitMQ",
}

var actionList = [...]string{
	"",
	// string for Describe Error
	// Use Generally
	"-Operation-",
	"Initialize fail -> ",
	"Connection fail -> ",
	"Undefined error -> ",
	"Argumnet error -> ",
	"JSON marshal fail -> ",
	"JSON unmarshal fail -> ",
	"Request error -> ",
	"Response error -> ",
	"Send error -> ",
	"Receive error -> ",
	"Parsing error -> ",
	"Invalid Token Error -> ",
	"Token Expired -> ",
	"Operationfail -> ",
	"No Result -> ",
	"Timestamp Convert Failed -> ",
	"UUID Generate Failed -> ",

	// clarinet specific

	// piccolo specific
	"Prepare error -> ",
	"Execute error -> ",
	"Token Generation Error -> ",
	"Login failed -> ",
	"User exist -> ",

	// cello specific
	"volumehandlerError -> ",
	"storagepoolError -> ",
	"createvolumeError -> ",
	"preparepxeError -> ",
	"writeiscsiError -> ",

	// violin-scheduler specific
	"schedulerhandlerError -> ",

	// flute specific
	"IPMI Error->",

	// viola specific

	// piano specific
	"Read Metric Error -> ",

	// harp specific
	"interface address lookup error -> ",
	"PF error -> ",
	"DHCPD error -> ",
	"file error -> ",
	"ifconfig error -> ",
	"IP address error -> ",
	"Subnet In Use error -> ",
	"Subnet not allocated error -> ",
	"AdaptiveIP allocated error -> ",

	// violin-novnc specific

	// violin specific
	"Create Server failed -> ",
	"Create Server routine error ->",
	"Get available nodes error ->",
	"Get nodes error ->",
	"ServerNode present error ->",
}
