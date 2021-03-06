// Copyright 2015 The go-ethereum Authors
// This file is part of Webchain.
//
// Webchain is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// Webchain is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with Webchain. If not, see <http://www.gnu.org/licenses/>.

package tests

import (
	"path/filepath"
	"testing"
)

func TestRLP(t *testing.T) {
	err := RunRLPTest(filepath.Join(rlpTestDir, "rlptest.json"), nil)
	if err != nil {
		t.Fatal(err)
	}
}

func TestRLP_invalid(t *testing.T) {
	err := RunRLPTest(filepath.Join(rlpTestDir, "invalidRLPTest.json"), nil)
	if err != nil {
		t.Fatal(err)
	}
}
