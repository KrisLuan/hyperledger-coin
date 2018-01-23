package pocket

import (
	"github.com/golang/protobuf/proto"
	"strings"
)

func ParsePocket(data []byte) (*Pocket, error) {
	if data == nil || len(data) == 0 {
		return nil, ErrNoAccount
	}

	pocket := new(Pocket)
	if err := proto.Unmarshal(data, pocket); err != nil {
		return nil, err
	}

	return pocket, nil
}

func ParsePointInfo(data []byte) (*PointInfo, error) {
	if data == nil || len(data) == 0 {
		return nil, ErrKeyNoData
	}
	pointInfo := new(PointInfo)
	if err := proto.Unmarshal(data, pointInfo); err != nil {
		return nil, err
	}

	return pointInfo, nil
}

func ParsePointKind(data []byte) (*PointKind, error) {
	if data == nil || len(data) == 0 {
		return nil, ErrKeyNoData
	}

	pointKind := new(PointKind)
	if err := proto.Unmarshal(data, pointKind); err != nil {
		return nil, err
	}

	return pointKind, nil
}

func IsValidAddr(addr string) bool {
	return strings.Contains(addr, "_")
}

func VerifyAddr(addr string, pubkey string) bool {
	return true
}