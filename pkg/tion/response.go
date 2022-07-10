package tion

import (
	bytes2 "bytes"

	"fmt"

	"github.com/go-errors/errors"
)

// Status of the breezer
type Status struct {
	Enabled         bool
	HeaterEnabled   bool
	SoundEnabled    bool
	TimerEnabled    bool
	Speed           int8
	Gate            int8 // 0 - indoor, 1 - mixed, 2 - outdoor; when 0 than heater off; when 1 speed 1,2 unavailiable
	TempTarget      int8
	TempOut         int8 // Outcoming from device - inside
	TempIn          int8 // Incoming to device - outside
	FiltersRemains  int16
	Hours           int8
	Minutes         int8
	ErrorCode       int8
	Productivity    int8 //m3pH
	RunDays         int16
	FirmwareVersion int16
	Todo            int8
	ResetFilters    bool
}

// GateStatus from int
func GateStatus(v int8) string {
	switch v {
	case 0:
		return "indoor"
	case 1:
		return "mixed"
	case 2:
		return "outdoor"
	default:
		return "unknown"
	}
}

// GateStatus string presenation
func (s Status) GateStatus() string {
	return GateStatus(s.Gate)
}

// SetGateStatus by given string
func (s *Status) SetGateStatus(str string) {
	switch str {
	case "indoor":
		s.Gate = 0
	case "mixed":
		s.Gate = 1
	case "outdoor":
		s.Gate = 2
	}
}

// BeautyString presetation of the status
func (s Status) BeautyString() string {
	return fmt.Sprintf("\nStatus: %s, Heater: %s, Sound: %s\nTarget: %d C, In: %d C, Out: %d C\nSpeed %d, Gate: %s, Error: %d, FW: %x\nFilters remain: %d days, Uptime %d days %02d:%02d\n",
		sts(s.Enabled), sts(s.HeaterEnabled), sts(s.SoundEnabled),
		s.TempTarget, s.TempIn, s.TempOut,
		s.Speed, s.GateStatus(), s.ErrorCode, s.FirmwareVersion,
		s.FiltersRemains, s.RunDays, s.Hours, s.Minutes)
}

func sts(b bool) string {
	if b {
		return "on"
	}
	return "off"
}

// FromBytes to Status
func FromBytes(bytes []byte) (*Status, error) {
	if len(bytes) < 20 {
		return nil, errors.New(fmt.Sprintf("Expecting 20 bytes array. Got %d", len(bytes)))
	}
	buffer := bytes2.NewBuffer(bytes[2:])

	bt := rb(buffer)
	tr := Status{}
	tr.Speed = int8(int(bt) & 0xF)
	tr.Gate = bt >> 4
	tmp, _ := buffer.ReadByte()
	tr.TempTarget = int8(tmp)

	bt = rb(buffer)
	if bt&1 != 0 {
		tr.HeaterEnabled = true
	}
	if bt&2 != 0 {
		tr.Enabled = true
	}
	if bt&4 != 0 {
		tr.TimerEnabled = true
	}
	if bt&8 != 0 {
		tr.SoundEnabled = true
	}
	tr.Todo = rb(buffer)
	tr.TempOut = (rb(buffer) + rb(buffer)) / 2
	tr.TempIn = rb(buffer)
	tr.FiltersRemains = ri(buffer)
	tr.Hours = rb(buffer)
	tr.Minutes = rb(buffer)
	tr.ErrorCode = rb(buffer)
	tr.Productivity = rb(buffer)
	tr.RunDays = ri(buffer)
	tr.FirmwareVersion = ri(buffer)
	return &tr, nil
}

func rb(b *bytes2.Buffer) int8 {
	bt, _ := b.ReadByte()
	return int8(bt & 0xFF)
}

func rub(b *bytes2.Buffer) int16 {
	bt, _ := b.ReadByte()
	return int16(bt)
}

func ri(b *bytes2.Buffer) int16 {
	iv := rub(b) + rub(b)<<8
	return iv
}
