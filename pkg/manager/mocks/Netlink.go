// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	net "net"

	mock "github.com/stretchr/testify/mock"

	netlink "github.com/vishvananda/netlink"
)

// Netlink is an autogenerated mock type for the Netlink type
type Netlink struct {
	mock.Mock
}

// BridgeVlanAdd provides a mock function with given fields: _a0, _a1, _a2, _a3, _a4, _a5
func (_m *Netlink) BridgeVlanAdd(_a0 netlink.Link, _a1 uint16, _a2 bool, _a3 bool, _a4 bool, _a5 bool) error {
	ret := _m.Called(_a0, _a1, _a2, _a3, _a4, _a5)

	var r0 error
	if rf, ok := ret.Get(0).(func(netlink.Link, uint16, bool, bool, bool, bool) error); ok {
		r0 = rf(_a0, _a1, _a2, _a3, _a4, _a5)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// BridgeVlanDel provides a mock function with given fields: _a0, _a1, _a2, _a3, _a4, _a5
func (_m *Netlink) BridgeVlanDel(_a0 netlink.Link, _a1 uint16, _a2 bool, _a3 bool, _a4 bool, _a5 bool) error {
	ret := _m.Called(_a0, _a1, _a2, _a3, _a4, _a5)

	var r0 error
	if rf, ok := ret.Get(0).(func(netlink.Link, uint16, bool, bool, bool, bool) error); ok {
		r0 = rf(_a0, _a1, _a2, _a3, _a4, _a5)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LinkByName provides a mock function with given fields: _a0
func (_m *Netlink) LinkByName(_a0 string) (netlink.Link, error) {
	ret := _m.Called(_a0)

	var r0 netlink.Link
	if rf, ok := ret.Get(0).(func(string) netlink.Link); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(netlink.Link)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LinkSetDown provides a mock function with given fields: _a0
func (_m *Netlink) LinkSetDown(_a0 netlink.Link) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(netlink.Link) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LinkSetHardwareAddr provides a mock function with given fields: _a0, _a1
func (_m *Netlink) LinkSetHardwareAddr(_a0 netlink.Link, _a1 net.HardwareAddr) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(netlink.Link, net.HardwareAddr) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LinkSetMaster provides a mock function with given fields: _a0, _a1
func (_m *Netlink) LinkSetMaster(_a0 netlink.Link, _a1 netlink.Link) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(netlink.Link, netlink.Link) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LinkSetName provides a mock function with given fields: _a0, _a1
func (_m *Netlink) LinkSetName(_a0 netlink.Link, _a1 string) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(netlink.Link, string) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LinkSetNoMaster provides a mock function with given fields: _a0
func (_m *Netlink) LinkSetNoMaster(_a0 netlink.Link) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(netlink.Link) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LinkSetNsFd provides a mock function with given fields: _a0, _a1
func (_m *Netlink) LinkSetNsFd(_a0 netlink.Link, _a1 int) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(netlink.Link, int) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LinkSetUp provides a mock function with given fields: _a0
func (_m *Netlink) LinkSetUp(_a0 netlink.Link) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(netlink.Link) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LinkSetVfHardwareAddr provides a mock function with given fields: _a0, _a1, _a2
func (_m *Netlink) LinkSetVfHardwareAddr(_a0 netlink.Link, _a1 int, _a2 net.HardwareAddr) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(netlink.Link, int, net.HardwareAddr) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
