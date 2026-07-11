// Copyright 2026 BitWise Media Group Ltd
// SPDX-License-Identifier: MIT

package report

import "testing"

func TestClassify(t *testing.T) {
	tests := []struct {
		version string
		want    Maturity
	}{
		// The <1.0.0 vs >=1.0.0 boundary (AC7/AC8).
		{"0.9.9", MaturityUnstable},
		{"1.0.0", MaturityStable},
		{"1.2.3", MaturityStable},
		{"0.0.1", MaturityUnstable},
		{"10.4.2", MaturityStable},
		// A prerelease tag wins regardless of the core version.
		{"1.0.0-alpha.1", MaturityPrerelease},
		{"2.0.0-rc.1", MaturityPrerelease},
		{"0.9.0-beta", MaturityPrerelease},
		// Missing or unparseable versions are never gated (AC9).
		{"", MaturityUnknown},
		{"not-a-version", MaturityUnknown},
	}
	for _, tt := range tests {
		t.Run(tt.version, func(t *testing.T) {
			if got := Classify(tt.version); got != tt.want {
				t.Errorf("Classify(%q) = %q, want %q", tt.version, got, tt.want)
			}
		})
	}
}

func TestParseMaturity(t *testing.T) {
	tests := []struct {
		in     string
		want   Maturity
		wantOK bool
	}{
		{"stable", MaturityStable, true},
		{"unstable", MaturityUnstable, true},
		{"prerelease", MaturityPrerelease, true},
		// "unknown" is an internal classification, never a selectable gate level.
		{"unknown", "", false},
		// The closed set is case-sensitive and rejects anything else.
		{"STABLE", "", false},
		{"", "", false},
		{"garbage", "", false},
	}
	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			got, ok := ParseMaturity(tt.in)
			if got != tt.want || ok != tt.wantOK {
				t.Errorf("ParseMaturity(%q) = (%q, %t), want (%q, %t)", tt.in, got, ok, tt.want, tt.wantOK)
			}
		})
	}
}
