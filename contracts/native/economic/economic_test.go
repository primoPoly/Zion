/*
 * Copyright (C) 2021 The Zion Authors
 * This file is part of The Zion library.
 *
 * The Zion is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The Zion is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU Lesser General Public License
 * along with The Zion.  If not, see <http://www.gnu.org/licenses/>.
 */

package economic

import (
	"math/big"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/contracts/native"
	. "github.com/ethereum/go-ethereum/contracts/native/go_abi/economic_abi"
	nm "github.com/ethereum/go-ethereum/contracts/native/governance/node_manager"
	"github.com/ethereum/go-ethereum/contracts/native/utils"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	InitABI()
	nm.InitNodeManager()
	os.Exit(m.Run())
}

func TestTotalSupply(t *testing.T) {
	sdb := utils.NewTestStateDB()
	blockHeight := big.NewInt(40)

	contractRef := native.NewContractRef(sdb, common.EmptyAddress, common.EmptyAddress, blockHeight, common.Hash{}, 0, nil)
	ctx := native.NewNativeContract(sdb, contractRef)

	raw, err := TotalSupply(ctx)
	assert.NoError(t, err)

	output := new(big.Int)
	assert.NoError(t, utils.UnpackOutputs(ABI, MethodTotalSupply, output, raw))

	t.Log(output)
}