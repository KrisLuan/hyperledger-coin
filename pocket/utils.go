package pocket

import "github.com/golang/protobuf/proto"

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