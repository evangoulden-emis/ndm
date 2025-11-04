package device

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"go_ndm.github.com/setup"
	"os"
	"sync"
)

var (
	once sync.Once
	dm   *Manager
)

// GetManager returns the singleton instance of the device manager.
func GetManager() *Manager {
	once.Do(func() {
		dm = newManager()

		// Load devices from file after manager is initialized
		cfg := setup.GetConfig()
		if cfg == nil {
			dm.Devices = []Device{}
			return
		}
		err := dm.loadFromFile(cfg.DevicesFile, &dm.Devices)
		if err != nil {
			dm.Devices = []Device{} // fallback to empty list
		}
	})
	return dm
}

type Manager struct {
	Devices []Device
	ctx     context.Context
	cancel  context.CancelFunc
}

func newManager() *Manager {
	ctx, cancel := context.WithCancel(context.Background())
	return &Manager{
		ctx:    ctx,
		cancel: cancel,
	}
}

func (dm *Manager) AddDevice(device Device) {
	device.UUID = uuid.New()
	dm.Devices = append(dm.Devices, device)
	_ = dm.saveToFile("devices.json")
}

func (dm *Manager) RemoveDevice(device Device) {
	for i, d := range dm.Devices {
		if d.UUID == device.UUID {
			dm.Devices = append(dm.Devices[:i], dm.Devices[i+1:]...)
			break
		}
	}
	_ = dm.saveToFile("devices.json")
}

func (dm *Manager) EditDevice(device Device) {
	for i, d := range dm.Devices {
		if d.UUID == device.UUID {
			dm.Devices[i] = device
			break
		}
	}
	_ = dm.saveToFile("devices.json")
}

func (dm *Manager) GetDeviceList() []Device {
	return dm.Devices
}

func (dm *Manager) loadFromFile(filename string, target interface{}) error {
	jsonData, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	return json.Unmarshal(jsonData, target)
}

func (dm *Manager) saveToFile(filename string) error {
	jsonData, err := json.MarshalIndent(dm.GetDeviceList(), "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, jsonData, 0644)
}
