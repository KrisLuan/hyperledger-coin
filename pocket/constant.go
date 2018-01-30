package pocket

import "github.com/op/go-logging"

var (
	logger = logging.MustGetLogger("foam")
)

const (
	DefaultPocketKind string	= "default"
	InitAddr string				= "foam"
	InitTotalPoint int64		= 20000000*100000000
	InitPubkey string			= "fafasdfasdfas"
	//联合键注意，addr是不能包含‘_’的
	CompositeIndexName string	= "foam_point"
)

const (
	pointInfoKey	= "pointCoinInfo"
	kindKey			= "kindKey"
)