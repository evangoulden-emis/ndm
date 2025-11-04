package device

import "context"

type DeviceManager struct {
	Devices []Device
	ctx     context.Context
	cancel  context.CancelFunc
}

func NewDeviceManager() *DeviceManager {
	return &DeviceManager{}
}

func (dm *DeviceManager) AddDevice(device Device) {
	dm.Devices = append(dm.Devices, device)
}

func (dm *DeviceManager) RemoveDevice(device Device) {
	dm.Devices = append(dm.Devices, device)
}

func (dm *DeviceManager) EditDevice(device Device) {
	dm.Devices = append(dm.Devices, device)
}
