package main

import (
	"fmt"
)

type SRegister uint64

type Registers struct {
	reg [256]struct {
		val byte
		min byte
		max byte
	}
	current SRegister
}

const (
	RegAutoAnswer SRegister = iota
	RegRingCount
	RegEscapeSeqChar
	RegCarriageReturnChar
	RegLineFeedChar
	RegBackspaceChar
	RegBlindDialWait
	RegWaitForCarrierDelay
	RegCommaDelay
	RegCarrierDetectResponseTime
	RegHangupDelay
	RegMultiFreqToneDuration
	RegEscapeSeqGuardTime
	RegStatusOptions         = 14
	RegGeneralBitmapOptions  = 21
	RegSpeakerResultsOptions = 22
	RegDelayToDTR            = 25
	RegInactivityTimeout     = 30
	RegLineMode              = 90
)

/* Flags/Masks for RegSpeakerResultsOptions Bitmap */
const (
	/* 0b11111100 */
	VolumeMask = 0xfc
	/* 0b00000001 */
	VolumeLow = 0x01
	/* 0b00000010 */
	VolumeMed = 0x02
	/* 0b00000011 */
	VolumeHigh = 0x03
	/* 0b11110011 */
	SpeakerMask = 0xF3
	/* 0b00000100 */
	SpeakerOffCarrier = 0x04
	/* 0b00001000 */
	SpeakerOn = 0x08
	/* 0b00000100 */
	SpeakerOnHandshake = 0x0C
	/* 0b10001111 */
	ResultsLevelMask = 0x8F
	/* 0b01000000 */
	ResultsLevel1 = 0x40
	/* 0b01010000 */
	ResultsLevel2 = 0x50
	/* 0b01100000 */
	ResultsLevel3 = 0x60
	/* 0b01110000 */
	ResultsLevel4 = 0x70
)

/* Flags for RegStatusOptions Bitmap */
const (
	_ byte = 1 << iota
	Echo
	Quiet
	Verbose
	_
	Tone
	_
	Outgoing
)

func NewRegisters() *Registers {
	var r Registers
	r.current = 0
	r.Reset()
	return &r
}

func (r *Registers) Reset() {
	for i, _ := range r.reg {
		r.reg[i].val = 0
		r.reg[i].min = 0
		r.reg[i].max = 255
	}

	r.Write(RegAutoAnswer, 0)
	r.Write(RegRingCount, 0)
	r.Write(RegEscapeSeqChar, byte('+'))
	r.Write(RegCarriageReturnChar, byte('\r'))
	r.Write(RegLineFeedChar, byte('\n'))
	r.Write(RegBackspaceChar, byte('\b'))
	r.Write(RegBlindDialWait, 4)
	r.Write(RegWaitForCarrierDelay, 40)
	r.Write(RegCommaDelay, 2)
	r.Write(RegCarrierDetectResponseTime, 6)
	r.Write(RegHangupDelay, 14)
	r.Write(RegMultiFreqToneDuration, 95)
	r.Write(RegEscapeSeqGuardTime, 50)
	r.Write(RegStatusOptions, 138)
	r.Write(RegInactivityTimeout, 0)
	r.Write(RegGeneralBitmapOptions, 52)
	r.Write(RegSpeakerResultsOptions, 118)
	r.Write(RegDelayToDTR, 5)
	r.Write(RegLineMode, 0)
}

func (r *Registers) SetConstraint(reg SRegister, min byte, max byte) {

}

func (r *Registers) Write(reg SRegister, value byte) error {
	if reg < 0 || reg > 255 {
		return fmt.Errorf("Register %d out of range: must be between 0 and 255", reg)
	}
	if value < r.reg[reg].min || value > r.reg[reg].max {
		return fmt.Errorf("Value %d of register %d out of range: must be between %d and %d", value, reg, r.reg[reg].min, r.reg[reg].max)
	}
	r.reg[reg].val = value
	return nil
}

func (r *Registers) Read(reg SRegister) (byte, error) {
	if reg < 0 || reg > 255 {
		return 0, fmt.Errorf("Register %d out of range: must be between 0 and 255", reg)
	}
	return r.reg[reg].val, nil
}

func (r *Registers) SetCurrent(reg SRegister) error {
	if reg < 0 || reg > 255 {
		return fmt.Errorf("Register %d out of range: must be between 0 and 255", reg)
	}
	r.current = reg
	return nil
}

func (r *Registers) WriteCurrent(value byte) error {
	return r.Write(r.current, value)
}

func (r *Registers) ReadCurrent() (byte, error) {
	return r.Read(r.current)
}
