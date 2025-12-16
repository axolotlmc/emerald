package block

import (
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"time"
)

type Hash [32]byte

type BlockHeader struct {
	Version      uint32
	PreviousHash Hash
	MerkleRoot   Hash
	Timestamp    int64
	Difficulty   uint32
	Nonce        uint64
}

func NewBlockHeader(version uint32, prevHash Hash, merkleRoot Hash, difficulty uint32) *BlockHeader {
	return &BlockHeader{
		Version:      version,
		PreviousHash: prevHash,
		MerkleRoot:   merkleRoot,
		Timestamp:    time.Now().Unix(),
		Difficulty:   difficulty,
		Nonce:        0,
	}
}

func (bh *BlockHeader) Serialize() []byte {
	buf := make([]byte, 144)
	offset := 0

	binary.LittleEndian.PutUint32(buf[offset:], bh.Version)
	offset += 4

	copy(buf[offset:], bh.PreviousHash[:])
	offset += 32

	copy(buf[offset:], bh.MerkleRoot[:])
	offset += 32

	binary.LittleEndian.PutUint64(buf[offset:], uint64(bh.Timestamp))
	offset += 8

	binary.LittleEndian.PutUint32(buf[offset:], bh.Difficulty)
	offset += 4

	binary.LittleEndian.PutUint64(buf[offset:], bh.Nonce)

	return buf[:88]
}

func (bh *BlockHeader) Hash() Hash {
	data := bh.Serialize()
	hashBytes := sha256.Sum256(data)
	return Hash(hashBytes)
}

func (bh *BlockHeader) String() string {
	h := bh.Hash()
	return hex.EncodeToString(h[:])
}
