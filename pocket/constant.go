package pocket

import "github.com/op/go-logging"

var (
	logger = logging.MustGetLogger("foam")
)

const (
	DefaultPocketKind string	= "default"
	InitAddr string				= "foam"
	InitTotalPoint int64		= 20000000*100000000
)

const (
	pointInfoKey	= "pointCoinInfo"
	kindKey			= "kindKey"
)