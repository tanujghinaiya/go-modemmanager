// Code generated by "stringer -type=MMModemContactsStorage -trimprefix=MmModemContactsStorage"; DO NOT EDIT.

package go_modemmanager

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[MmModemContactsStorageUnknown-0]
	_ = x[MmModemContactsStorageMe-1]
	_ = x[MmModemContactsStorageSm-2]
	_ = x[MmModemContactsStorageMt-3]
}

const _MMModemContactsStorage_name = "UnknownMeSmMt"

var _MMModemContactsStorage_index = [...]uint8{0, 7, 9, 11, 13}

func (i MMModemContactsStorage) String() string {
	if i >= MMModemContactsStorage(len(_MMModemContactsStorage_index)-1) {
		return "MMModemContactsStorage(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _MMModemContactsStorage_name[_MMModemContactsStorage_index[i]:_MMModemContactsStorage_index[i+1]]
}
