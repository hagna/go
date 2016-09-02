// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mips32

import (
	"cmd/compile/internal/gc"
	"cmd/internal/obj"
	"cmd/internal/obj/mips"
)

const (
	LeftRdwr  uint32 = gc.LeftRead | gc.LeftWrite
	RightRdwr uint32 = gc.RightRead | gc.RightWrite
)

// This table gives the basic information about instruction
// generated by the compiler and processed in the optimizer.
// See opt.h for bit definitions.
//
// Instructions not generated need not be listed.
// As an exception to that rule, we typically write down all the
// size variants of an operation even if we just use a subset.
//
// The table is formatted for 8-space tabs.
var progtable = [mips.ALAST & obj.AMask]obj.ProgInfo{
	obj.ATYPE:     {Flags: gc.Pseudo | gc.Skip},
	obj.ATEXT:     {Flags: gc.Pseudo},
	obj.AFUNCDATA: {Flags: gc.Pseudo},
	obj.APCDATA:   {Flags: gc.Pseudo},
	obj.AUNDEF:    {Flags: gc.Break},
	obj.AUSEFIELD: {Flags: gc.OK},
	obj.ACHECKNIL: {Flags: gc.LeftRead},
	obj.AVARDEF:   {Flags: gc.Pseudo | gc.RightWrite},
	obj.AVARKILL:  {Flags: gc.Pseudo | gc.RightWrite},
	obj.AVARLIVE:  {Flags: gc.Pseudo | gc.LeftRead},

	// NOP is an internal no-op that also stands
	// for USED and SET annotations, not the MIPS opcode.
	obj.ANOP: {Flags: gc.LeftRead | gc.RightWrite},

	// Integer
	mips.AADD & obj.AMask:   {Flags: gc.SizeL | gc.LeftRead | gc.RegRead | gc.RightWrite},
	mips.AADDU & obj.AMask:  {Flags: gc.SizeL | gc.LeftRead | gc.RegRead | gc.RightWrite},
	mips.AADDV & obj.AMask:  {Flags: gc.SizeQ | gc.LeftRead | gc.RegRead | gc.RightWrite},
	mips.AADDVU & obj.AMask: {Flags: gc.SizeQ | gc.LeftRead | gc.RegRead | gc.RightWrite},
	mips.ASUB & obj.AMask:   {Flags: gc.SizeL | gc.LeftRead | gc.RegRead | gc.RightWrite},
	mips.ASUBU & obj.AMask:  {Flags: gc.SizeL | gc.LeftRead | gc.RegRead | gc.RightWrite},
	mips.ASUBV & obj.AMask:  {Flags: gc.SizeQ | gc.LeftRead | gc.RegRead | gc.RightWrite},
	mips.ASUBVU & obj.AMask: {Flags: gc.SizeQ | gc.LeftRead | gc.RegRead | gc.RightWrite},
	mips.AAND & obj.AMask:   {Flags: gc.SizeQ | gc.LeftRead | gc.RegRead | gc.RightWrite},
	mips.AOR & obj.AMask:    {Flags: gc.SizeQ | gc.LeftRead | gc.RegRead | gc.RightWrite},
	mips.AXOR & obj.AMask:   {Flags: gc.SizeQ | gc.LeftRead | gc.RegRead | gc.RightWrite},
	mips.ANOR & obj.AMask:   {Flags: gc.SizeQ | gc.LeftRead | gc.RegRead | gc.RightWrite},
	mips.AMUL & obj.AMask:   {Flags: gc.SizeL | gc.LeftRead | gc.RegRead},
	mips.AMULU & obj.AMask:  {Flags: gc.SizeL | gc.LeftRead | gc.RegRead},
	mips.AMULV & obj.AMask:  {Flags: gc.SizeQ | gc.LeftRead | gc.RegRead},
	mips.AMULVU & obj.AMask: {Flags: gc.SizeQ | gc.LeftRead | gc.RegRead},
	mips.ADIV & obj.AMask:   {Flags: gc.SizeL | gc.LeftRead | gc.RegRead},
	mips.ADIVU & obj.AMask:  {Flags: gc.SizeL | gc.LeftRead | gc.RegRead},
	mips.ADIVV & obj.AMask:  {Flags: gc.SizeQ | gc.LeftRead | gc.RegRead},
	mips.ADIVVU & obj.AMask: {Flags: gc.SizeQ | gc.LeftRead | gc.RegRead},
	mips.AREM & obj.AMask:   {Flags: gc.SizeL | gc.LeftRead | gc.RegRead},
	mips.AREMU & obj.AMask:  {Flags: gc.SizeL | gc.LeftRead | gc.RegRead},
	mips.AREMV & obj.AMask:  {Flags: gc.SizeQ | gc.LeftRead | gc.RegRead},
	mips.AREMVU & obj.AMask: {Flags: gc.SizeQ | gc.LeftRead | gc.RegRead},
	mips.ASLL & obj.AMask:   {Flags: gc.SizeL | gc.LeftRead | gc.RegRead | gc.RightWrite},
	mips.ASLLV & obj.AMask:  {Flags: gc.SizeQ | gc.LeftRead | gc.RegRead | gc.RightWrite},
	mips.ASRA & obj.AMask:   {Flags: gc.SizeL | gc.LeftRead | gc.RegRead | gc.RightWrite},
	mips.ASRAV & obj.AMask:  {Flags: gc.SizeQ | gc.LeftRead | gc.RegRead | gc.RightWrite},
	mips.ASRL & obj.AMask:   {Flags: gc.SizeL | gc.LeftRead | gc.RegRead | gc.RightWrite},
	mips.ASRLV & obj.AMask:  {Flags: gc.SizeQ | gc.LeftRead | gc.RegRead | gc.RightWrite},
	mips.ASGT & obj.AMask:   {Flags: gc.SizeQ | gc.LeftRead | gc.RegRead | gc.RightWrite},
	mips.ASGTU & obj.AMask:  {Flags: gc.SizeQ | gc.LeftRead | gc.RegRead | gc.RightWrite},

	// Floating point.
	mips.AADDF & obj.AMask:    {Flags: gc.SizeF | gc.LeftRead | gc.RegRead | gc.RightWrite},
	mips.AADDD & obj.AMask:    {Flags: gc.SizeD | gc.LeftRead | gc.RegRead | gc.RightWrite},
	mips.ASUBF & obj.AMask:    {Flags: gc.SizeF | gc.LeftRead | gc.RegRead | gc.RightWrite},
	mips.ASUBD & obj.AMask:    {Flags: gc.SizeD | gc.LeftRead | gc.RegRead | gc.RightWrite},
	mips.AMULF & obj.AMask:    {Flags: gc.SizeF | gc.LeftRead | gc.RegRead | gc.RightWrite},
	mips.AMULD & obj.AMask:    {Flags: gc.SizeD | gc.LeftRead | gc.RegRead | gc.RightWrite},
	mips.ADIVF & obj.AMask:    {Flags: gc.SizeF | gc.LeftRead | gc.RegRead | gc.RightWrite},
	mips.ADIVD & obj.AMask:    {Flags: gc.SizeD | gc.LeftRead | gc.RegRead | gc.RightWrite},
	mips.AABSF & obj.AMask:    {Flags: gc.SizeF | gc.LeftRead | gc.RightWrite},
	mips.AABSD & obj.AMask:    {Flags: gc.SizeD | gc.LeftRead | gc.RightWrite},
	mips.ANEGF & obj.AMask:    {Flags: gc.SizeF | gc.LeftRead | gc.RightWrite},
	mips.ANEGD & obj.AMask:    {Flags: gc.SizeD | gc.LeftRead | gc.RightWrite},
	mips.ACMPEQF & obj.AMask:  {Flags: gc.SizeF | gc.LeftRead | gc.RegRead},
	mips.ACMPEQD & obj.AMask:  {Flags: gc.SizeD | gc.LeftRead | gc.RegRead},
	mips.ACMPGTF & obj.AMask:  {Flags: gc.SizeF | gc.LeftRead | gc.RegRead},
	mips.ACMPGTD & obj.AMask:  {Flags: gc.SizeD | gc.LeftRead | gc.RegRead},
	mips.ACMPGEF & obj.AMask:  {Flags: gc.SizeF | gc.LeftRead | gc.RegRead},
	mips.ACMPGED & obj.AMask:  {Flags: gc.SizeD | gc.LeftRead | gc.RegRead},
	mips.AMOVFD & obj.AMask:   {Flags: gc.SizeD | gc.LeftRead | gc.RightWrite | gc.Conv},
	mips.AMOVDF & obj.AMask:   {Flags: gc.SizeF | gc.LeftRead | gc.RightWrite | gc.Conv},
	mips.AMOVFW & obj.AMask:   {Flags: gc.SizeL | gc.LeftRead | gc.RightWrite | gc.Conv},
	mips.AMOVWF & obj.AMask:   {Flags: gc.SizeF | gc.LeftRead | gc.RightWrite | gc.Conv},
	mips.AMOVDW & obj.AMask:   {Flags: gc.SizeL | gc.LeftRead | gc.RightWrite | gc.Conv},
	mips.AMOVWD & obj.AMask:   {Flags: gc.SizeD | gc.LeftRead | gc.RightWrite | gc.Conv},
	mips.AMOVFV & obj.AMask:   {Flags: gc.SizeQ | gc.LeftRead | gc.RightWrite | gc.Conv},
	mips.AMOVVF & obj.AMask:   {Flags: gc.SizeF | gc.LeftRead | gc.RightWrite | gc.Conv},
	mips.AMOVDV & obj.AMask:   {Flags: gc.SizeQ | gc.LeftRead | gc.RightWrite | gc.Conv},
	mips.AMOVVD & obj.AMask:   {Flags: gc.SizeD | gc.LeftRead | gc.RightWrite | gc.Conv},
	mips.ATRUNCFW & obj.AMask: {Flags: gc.SizeL | gc.LeftRead | gc.RightWrite | gc.Conv},
	mips.ATRUNCDW & obj.AMask: {Flags: gc.SizeL | gc.LeftRead | gc.RightWrite | gc.Conv},
	mips.ATRUNCFV & obj.AMask: {Flags: gc.SizeQ | gc.LeftRead | gc.RightWrite | gc.Conv},
	mips.ATRUNCDV & obj.AMask: {Flags: gc.SizeQ | gc.LeftRead | gc.RightWrite | gc.Conv},

	// Moves
	mips.AMOVB & obj.AMask:  {Flags: gc.SizeB | gc.LeftRead | gc.RightWrite | gc.Move | gc.Conv},
	mips.AMOVBU & obj.AMask: {Flags: gc.SizeB | gc.LeftRead | gc.RightWrite | gc.Move | gc.Conv},
	mips.AMOVH & obj.AMask:  {Flags: gc.SizeW | gc.LeftRead | gc.RightWrite | gc.Move | gc.Conv},
	mips.AMOVHU & obj.AMask: {Flags: gc.SizeW | gc.LeftRead | gc.RightWrite | gc.Move | gc.Conv},
	mips.AMOVW & obj.AMask:  {Flags: gc.SizeL | gc.LeftRead | gc.RightWrite | gc.Move | gc.Conv},
	mips.AMOVWU & obj.AMask: {Flags: gc.SizeL | gc.LeftRead | gc.RightWrite | gc.Move | gc.Conv},
	mips.AMOVV & obj.AMask:  {Flags: gc.SizeQ | gc.LeftRead | gc.RightWrite | gc.Move},
	mips.AMOVF & obj.AMask:  {Flags: gc.SizeF | gc.LeftRead | gc.RightWrite | gc.Move | gc.Conv},
	mips.AMOVD & obj.AMask:  {Flags: gc.SizeD | gc.LeftRead | gc.RightWrite | gc.Move},

	// Jumps
	mips.AJMP & obj.AMask:  {Flags: gc.Jump | gc.Break},
	mips.AJAL & obj.AMask:  {Flags: gc.Call},
	mips.ABEQ & obj.AMask:  {Flags: gc.Cjmp},
	mips.ABNE & obj.AMask:  {Flags: gc.Cjmp},
	mips.ABGEZ & obj.AMask: {Flags: gc.Cjmp},
	mips.ABLTZ & obj.AMask: {Flags: gc.Cjmp},
	mips.ABGTZ & obj.AMask: {Flags: gc.Cjmp},
	mips.ABLEZ & obj.AMask: {Flags: gc.Cjmp},
	mips.ABFPF & obj.AMask: {Flags: gc.Cjmp},
	mips.ABFPT & obj.AMask: {Flags: gc.Cjmp},
	mips.ARET & obj.AMask:  {Flags: gc.Break},
	obj.ADUFFZERO:          {Flags: gc.Call},
	obj.ADUFFCOPY:          {Flags: gc.Call},
}

func proginfo(p *obj.Prog) {
	info := &p.Info
	*info = progtable[p.As&obj.AMask]
	if info.Flags == 0 {
		gc.Fatalf("proginfo: unknown instruction %v", p)
	}

	if (info.Flags&gc.RegRead != 0) && p.Reg == 0 {
		info.Flags &^= gc.RegRead
		info.Flags |= gc.RightRead /*CanRegRead |*/
	}

	if (p.From.Type == obj.TYPE_MEM || p.From.Type == obj.TYPE_ADDR) && p.From.Reg != 0 {
		info.Regindex |= RtoB(int(p.From.Reg))
	}

	if (p.To.Type == obj.TYPE_MEM || p.To.Type == obj.TYPE_ADDR) && p.To.Reg != 0 {
		info.Regindex |= RtoB(int(p.To.Reg))
	}

	if p.From.Type == obj.TYPE_ADDR && p.From.Sym != nil && (info.Flags&gc.LeftRead != 0) {
		info.Flags &^= gc.LeftRead
		info.Flags |= gc.LeftAddr
	}

	if p.As == obj.ADUFFZERO {
		info.Reguse |= 1<<0 | RtoB(mips.REGRT1)
		info.Regset |= RtoB(mips.REGRT1)
	}

	if p.As == obj.ADUFFCOPY {
		// TODO(austin) Revisit when duffcopy is implemented
		info.Reguse |= RtoB(mips.REGRT1) | RtoB(mips.REGRT2) | RtoB(mips.REG_R3)

		info.Regset |= RtoB(mips.REGRT1) | RtoB(mips.REGRT2)
	}
}
