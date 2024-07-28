package main

import (
	"testing"

	"github.com/miekg/dns"
)

//////////

type MX struct {
	dns.MX
}

type RecordConfig struct {
	Rdata any
	Meta  string
}

//////////

func (ty *MX) SetUp(p uint16, t string) {
	ty.Preference = p
	ty.Mx = t
}

func (rc *RecordConfig) BlessMX(p uint16, t string) {
	rc.Rdata = &MX{
		dns.MX{Preference: p,
			Mx: t,
		},
	}
}

func (rc *RecordConfig) BlessMXp(p uint16, t string) {
	rc.Rdata = &MX{}
	rc.Rdata.(*MX).Preference = p
	rc.Rdata.(*MX).Mx = t
}

//////////

//////////

func Style1() *RecordConfig {
	x := RecordConfig{}
	x.Rdata = &MX{}
	x.Rdata.(*MX).Mx = "foo"
	x.Rdata.(*MX).Preference = 99
	return &x
}

func Style1p() *RecordConfig {
	x := &RecordConfig{}

	m := &MX{}
	m.Mx = "foo"
	m.Preference = 99

	// m := &MX{
	// 	dns.MX{Mx: "foo",
	// 		Preference: 99,
	// 	} }

	x.Rdata = m
	return x
}

func Style2() *RecordConfig {
	y := RecordConfig{}
	y.Rdata = &MX{}
	y.Rdata.(*MX).SetUp(1, "bar")
	return &y
}

func Style2p() *RecordConfig {
	y := &RecordConfig{}
	y.Rdata = &MX{}
	y.Rdata.(*MX).SetUp(1, "bar")
	return y
}

//////////

func Style3() *RecordConfig {
	m := &MX{}
	z := RecordConfig{Rdata: m}
	z.Rdata.(*MX).SetUp(1, "bar")
	return &z
}

func Style3p() *RecordConfig {
	m := &MX{}
	m.SetUp(1, "bar")
	z := &RecordConfig{Rdata: m}
	return z
}

func Style4(z *RecordConfig) *RecordConfig {
	z.BlessMX(1, "bar")
	return z
}

func Style4p(z *RecordConfig) *RecordConfig {
	z.BlessMXp(1, "bar")
	return z
}

//////////

func BenchmarkStyle1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Style1()
	}
}

func BenchmarkStyle1p(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Style1p()
	}
}

func BenchmarkStyle2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Style2()
	}
}

func BenchmarkStyle2p(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Style2p()
	}
}

func BenchmarkStyle3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Style3()
	}
}

func BenchmarkStyle3p(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Style3p()
	}
}

func BenchmarkStyle4(b *testing.B) {
	z := &RecordConfig{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Style4(z)
	}
}

func BenchmarkStyle4p(b *testing.B) {
	z := &RecordConfig{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Style4p(z)
	}
}
