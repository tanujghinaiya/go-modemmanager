// Code generated by "stringer -type=MMModem3gppSubscriptionState -trimprefix=MmModem3gppSubscriptionState"; DO NOT EDIT.

package go_modemmanager

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[MmModem3gppSubscriptionStateUnknown-0]
	_ = x[MmModem3gppSubscriptionStateUnprovisioned-1]
	_ = x[MmModem3gppSubscriptionStateProvisioned-2]
	_ = x[MmModem3gppSubscriptionStateOutOfData-3]
}

const _MMModem3gppSubscriptionState_name = "UnknownUnprovisionedProvisionedOutOfData"

var _MMModem3gppSubscriptionState_index = [...]uint8{0, 7, 20, 31, 40}

func (i MMModem3gppSubscriptionState) String() string {
	if i >= MMModem3gppSubscriptionState(len(_MMModem3gppSubscriptionState_index)-1) {
		return "MMModem3gppSubscriptionState(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _MMModem3gppSubscriptionState_name[_MMModem3gppSubscriptionState_index[i]:_MMModem3gppSubscriptionState_index[i+1]]
}
