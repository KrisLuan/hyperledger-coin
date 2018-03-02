package pocket

import "github.com/hyperledger-coin/go-logging"

var (
	logger = logging.MustGetLogger("foam")
)

const (
	DefaultPocketKind string	= "default"
	InitAddr string				= "iSC2H4Ad1QqaF2szSwYL2UbruN8QaZ6V2"
	InitTotalPoint int64		= 20000000*100000000
	InitPubkey string			= "BN6GgRrFvsh1M0/SjjUaTy+VNxvCJKOV+3OmHaPNju/w5Tc7nf9spSR/irzoh3y5jPNg6A1Aig8jMlT21sp2OZU="
	InitPrikey string			= "53DAFF852B6207DD92936541115BCC703C2A6062F7B7C0B497D30475156D9140"
	//联合键注意，addr是不能包含‘_’的
	CompositeIndexName string	= "foam_point"
)

const (
	pointInfoKey	= "pointCoinInfo_key"
	kindKey			= "kind_key"
	txFeeKey		= "txFee_key"
	version 		= 1
)