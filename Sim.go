package go_modemmanager

import "github.com/godbus/dbus/v5"

// The SIM interface handles communication with SIM, USIM, and RUIM (CDMA SIM) cards.

const (
	SimInterface = ModemManagerInterface + ".Sim"

	ModemManagerSimsObjectPath = modemManagerMainObjectPath + "SIMs"
	/* Methods */

	/* Property */

)

type Sim interface {
	/* METHODS */
	// Returns object path
	GetObjectPath() dbus.ObjectPath
	//MarshalJSON() ([]byte, error)
}

func NewSim(objectPath dbus.ObjectPath) (Sim, error) {
	var sm sim
	return &sm, sm.init(SimInterface, objectPath)
}

type  sim struct {
	dbusBase
}
func (sm sim) GetObjectPath() dbus.ObjectPath {
	return sm.obj.Path()
}