// Code generated by "stringer -type=SpaceType"; DO NOT EDIT.

package types

import "strconv"

const (
	_SpaceType_name_0 = "SpaceTypeUndefined"
	_SpaceType_name_1 = "SpaceTypeProtocolSpaceTypeImplementation"
)

var (
	_SpaceType_index_1 = [...]uint8{0, 17, 40}
)

func (i SpaceType) String() string {
	switch {
	case i == -1:
		return _SpaceType_name_0
	case 1 <= i && i <= 2:
		i -= 1
		return _SpaceType_name_1[_SpaceType_index_1[i]:_SpaceType_index_1[i+1]]
	default:
		return "SpaceType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
