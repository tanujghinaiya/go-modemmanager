// Code generated by "stringer -type=MMModem3gppFacility -trimprefix=MmModem3gppFacility"; DO NOT EDIT.

package modemmanager

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[MmModem3gppFacilityNone-0]
	_ = x[MmModem3gppFacilitySim-1]
	_ = x[MmModem3gppFacilityFixedDialing-2]
	_ = x[MmModem3gppFacilityPhSim-4]
	_ = x[MmModem3gppFacilityPhFsim-8]
	_ = x[MmModem3gppFacilityNetPers-16]
	_ = x[MmModem3gppFacilityNetSubPers-32]
	_ = x[MmModem3gppFacilityProviderPers-64]
	_ = x[MmModem3gppFacilityCorpPers-128]
}

const (
	_MMModem3gppFacility_name_0 = "NoneSimFixedDialing"
	_MMModem3gppFacility_name_1 = "PhSim"
	_MMModem3gppFacility_name_2 = "PhFsim"
	_MMModem3gppFacility_name_3 = "NetPers"
	_MMModem3gppFacility_name_4 = "NetSubPers"
	_MMModem3gppFacility_name_5 = "ProviderPers"
	_MMModem3gppFacility_name_6 = "CorpPers"
)

var (
	_MMModem3gppFacility_index_0 = [...]uint8{0, 4, 7, 19}
)

func (i MMModem3gppFacility) String() string {
	switch {
	case i <= 2:
		return _MMModem3gppFacility_name_0[_MMModem3gppFacility_index_0[i]:_MMModem3gppFacility_index_0[i+1]]
	case i == 4:
		return _MMModem3gppFacility_name_1
	case i == 8:
		return _MMModem3gppFacility_name_2
	case i == 16:
		return _MMModem3gppFacility_name_3
	case i == 32:
		return _MMModem3gppFacility_name_4
	case i == 64:
		return _MMModem3gppFacility_name_5
	case i == 128:
		return _MMModem3gppFacility_name_6
	default:
		return "MMModem3gppFacility(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
