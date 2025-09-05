package device

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// GetDrives returns a list of connected USB mass storage devices.
func GetDrives() ([]string, error) {
	const blockClassPath = "/sys/class/block"
	var usbDevices []string

	devices, err := os.ReadDir(blockClassPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read %s: %w", blockClassPath, err)
	}

	for _, device := range devices {
		deviceName := device.Name()
		devicePath := filepath.Join(blockClassPath, deviceName)

		realPath, err := os.Readlink(devicePath)
		if err != nil {
			continue
		}

		if strings.Contains(realPath, "/usb") {
			devPath := filepath.Join("/dev", deviceName)
			usbDevices = append(usbDevices, devPath)
		}
	}

	return usbDevices, nil
}