package client

import (
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common"

	"github.com/fbsobreira/gotron-sdk/pkg/address"

	"github.com/stretchr/testify/require"
	//github.com/fbsobreirago/gotron-sdk
)

func TestAddressTypeAssert(t *testing.T) {
	require := require.New(t)
	addr, err := address.Base58ToAddress("TRcCg4v71arvfWxFvvromcuEggKV4qEzQy")
	require.NoError(err)
	a, err := addressTypeAssert("TRcCg4v71arvfWxFvvromcuEggKV4qEzQy")
	require.NoError(err)
	fmt.Println("a", a.String())
	b, err := addressTypeAssert(addr)
	fmt.Println("b", b.String())

	c, err := addressTypeAssert(common.HexToAddress("0xab890107948621AdB234E1A6b5c926D609ae8ba8"))
	require.NoError(err)
	fmt.Println("c", c.String())
}
