package encoding

import (
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"gopkg.in/go-playground/assert.v1"
)

func Test_EncodeSignalProof(t *testing.T) {
	// nolint: lll
	storageProof := "0x00edb9377a94f02e691ea208a80a7bd13dcb51dfae3a389d3001a55c5fdf3f3ae70807df1d8ffcaeefa5dc56875a03dae3494d93bbf3074018fea52137a098433d9795f455fcfaf89b6456b970605bd0f70bec98418ef322fa5bba9874ab2277225a38c19f99275cd6d25d9cd4259c366c02b750056a6b71810034f95d34ae12c1298a95e6ba03de90dbbfacc5c8edbeb6ed93a0e01c08f3435b1ca944e9229a176804445432795c14dc315925fe2480084d6b0729c30fb1c2e50e481df5d7daa25769a0c74baf362ff02f2a934538cd11bf939f64377489255aa60a12628b130f4bddd30664d6b5f91d17cc110bee38a7ce5c629585e57d777bfbb2eaf09e15"

	s := SignalProof{
		Height:       uint64(1),
		StorageProof: hexutil.MustDecode(storageProof),
		Hops:         []Hop{},
	}

	// nolint: lll
	want := "0x000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000008000000000000000000000000000000000000000000000000000000000000001a0000000000000000000000000000000000000000000000000000000000000010000edb9377a94f02e691ea208a80a7bd13dcb51dfae3a389d3001a55c5fdf3f3ae70807df1d8ffcaeefa5dc56875a03dae3494d93bbf3074018fea52137a098433d9795f455fcfaf89b6456b970605bd0f70bec98418ef322fa5bba9874ab2277225a38c19f99275cd6d25d9cd4259c366c02b750056a6b71810034f95d34ae12c1298a95e6ba03de90dbbfacc5c8edbeb6ed93a0e01c08f3435b1ca944e9229a176804445432795c14dc315925fe2480084d6b0729c30fb1c2e50e481df5d7daa25769a0c74baf362ff02f2a934538cd11bf939f64377489255aa60a12628b130f4bddd30664d6b5f91d17cc110bee38a7ce5c629585e57d777bfbb2eaf09e150000000000000000000000000000000000000000000000000000000000000000"
	proof, err := EncodeSignalProof(s)
	assert.Equal(t, nil, err)
	assert.Equal(t, hexutil.Encode(proof), want)
}
