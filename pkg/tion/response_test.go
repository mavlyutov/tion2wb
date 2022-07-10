package tion

import (
	"fmt"
	"testing"
)

func Test01(t *testing.T) {
	// 2 - b3 10 22 12 0b 00 12 12 08 60 01 13 14 00 2d 08 00 32 00 5a
	// 1 - b3 10 21 12 0b 00 10 10 08 60 01 13 16 00 1e 08 00 32 00 5a
	bts := []byte{0xb3, 0x10, 0x21, 0x12, 0x0b, 0x00, 0x10, 0x10, 0x08, 0x60, 0x01, 0x13, 0x16, 0x00, 0x1e, 0x08, 0x00, 0x32, 0x00, 0x5a}
	fmt.Println(bts)
	resp, err := FromBytes(bts)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(resp.BeautyString())
	if resp.FiltersRemains != 352 {
		t.Fatal(resp.FiltersRemains)
	}
}

func Test02_Minus(t *testing.T) {
	//179 16 33 18 11 0 19 19 253 73 1 10 20 0 30 31 0 50 0 90 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
	bts := []byte{179, 16, 33, 18, 11, 0, 19, 19, 253, 73, 1, 10, 20, 0, 30, 31, 0, 50, 0, 90, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	resp, err := FromBytes(bts)
	if err != nil {
		t.Fatal(err)
	}

	if resp.TempIn != -3 {
		t.Fatal(resp.TempIn)
	}
}

func TestInt(t *testing.T) {
	bts := []byte{179, 16, 33, 0, 10, 0, 9, 8, 5, 107, 0, 11, 9, 0, 30, 215, 0, 50, 0, 90}
	fmt.Println(bts)
	resp, err := FromBytes(bts)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(resp.BeautyString())
	fmt.Println(resp.FirmwareVersion)
	fmt.Println(resp.FiltersRemains)
	fmt.Println(resp.RunDays)
}

func TestDays(t *testing.T) {
	bts := []byte{179, 16, 34, 15, 10, 0, 13, 12, 10, 249, 255, 18, 16, 0, 45, 73, 1, 50, 0, 90}
	resp, err := FromBytes(bts)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(resp.BeautyString())
	fmt.Println(resp.FiltersRemains)
	fmt.Println(resp.RunDays)
}
