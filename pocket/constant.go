package pocket

import "github.com/hyperledger-coin/go-logging"

var (
	logger = logging.MustGetLogger("foam")
)

const (
	DefaultPocketKind string	= "default"
	InitAddr string				= "caRNGB62WSf3AiCToXn2iCGU1RJ9tr4Vu"
	InitTotalPoint int64		= 20000000*100000000
	InitPubkey string			= "BNFHetucyZHP1FtrtRz0UhWwSfIfsHY0NA+2ZJuRS+KOIClJEwjxnzSUKhRqyz2tt/XfrCykxkb7Qrm2gjeMJRU="
	InitPrikey string			= "53DAFF852B6207DD92936541115BCC703C2A6062F7B7C0B497D30475156D9140"
	//联合键注意，addr是不能包含‘_’的
	CompositeIndexName string	= "foam_point"
)

const (
	pointInfoKey	= "pointCoinInfo"
	kindKey			= "kindKey"
)