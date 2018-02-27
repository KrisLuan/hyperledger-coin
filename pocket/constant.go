package pocket

import "github.com/hyperledger-coin/go-logging"

var (
	logger = logging.MustGetLogger("foam")
)

const (
	DefaultPocketKind string	= "default"
	InitAddr string				= "Z5MoPT8TG24e5ncGAzNWSfFkTeH2Lrw3X"
	InitTotalPoint int64		= 20000000*100000000
	InitPubkey string			= "BEtHWP3/wgq8iPKV48ddbRwhB6E5jKX3zUS9lW70vxN+KM0UJUXBzZjFRRYgIKY2KWdtfcU5WEZp+uk0PQu8yhM="
	InitPrikey string			= "53DAFF852B6207DD92936541115BCC703C2A6062F7B7C0B497D30475156D9140"
	//联合键注意，addr是不能包含‘_’的
	CompositeIndexName string	= "foamPoint"
)

const (
	pointInfoKey	= "pointCoinInfo_key"
	kindKey			= "kind_key"
	txFeeKey		= "txFee_key"
	version 		= 1
)