package lcs

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBigNumbers(t *testing.T) {
	type MyStruct struct {
		A *big.Int
		B *big.Int
	}

	a := &MyStruct{
		A: big.NewInt(10),
		B: big.NewInt(100500),
	}

	bytes, err := Marshal(a)

	if err != nil {
		t.Fatal(err)
	}

	b := &MyStruct{}
	err = Unmarshal(bytes, &b)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Print: %d %d", a.A, b.A)

	require.True(t, a.A.String() == b.A.String())
	require.True(t, a.B.String() == b.B.String())
	require.Equal(t, a, b)
}
