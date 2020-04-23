// Code generated by "stringer -type=MMSerialError -trimprefix=MMSerialError"; DO NOT EDIT.

package modemmanager

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[MmSerialErrorUnknown-0]
	_ = x[MmSerialErrorOpenFailed-1]
	_ = x[MmSerialErrorSendFailed-2]
	_ = x[MmSerialErrorResponseTimeout-3]
	_ = x[MmSerialErrorOpenFailedNoDevice-4]
	_ = x[MmSerialErrorFlashFailed-5]
	_ = x[MmSerialErrorNotOpen-6]
	_ = x[MmSerialErrorParseFailed-7]
	_ = x[MmSerialErrorFrameNotFound-8]
}

const _MMSerialError_name = "MmSerialErrorUnknownMmSerialErrorOpenFailedMmSerialErrorSendFailedMmSerialErrorResponseTimeoutMmSerialErrorOpenFailedNoDeviceMmSerialErrorFlashFailedMmSerialErrorNotOpenMmSerialErrorParseFailedMmSerialErrorFrameNotFound"

var _MMSerialError_index = [...]uint8{0, 20, 43, 66, 94, 125, 149, 169, 193, 219}

func (i MMSerialError) String() string {
	if i >= MMSerialError(len(_MMSerialError_index)-1) {
		return "MMSerialError(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _MMSerialError_name[_MMSerialError_index[i]:_MMSerialError_index[i+1]]
}