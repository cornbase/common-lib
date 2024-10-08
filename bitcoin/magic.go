package bitcoin

import (
	"bytes"
	"crypto/sha256"

	"github.com/btcsuite/btcd/wire"
)

// License/Source https://github.com/BitonicNL/verify-signed-message/blob/main/LICENSE

// varIntProtoVer is the protocol version to use for serializing N as a VarInt
// Copied from https://github.com/btcsuite/btcd/blob/v0.23.3/btcutil/gcs/gcs.go#L37
const varIntProtoVer uint32 = 0

// Signed message are prepended with this magicMessage
// Taken from https://bitcoin.stackexchange.com/a/77325
const MagicMessage = "\x18Bitcoin Signed Message:\n"

// Signed message via BIP-322 are prepended with this bip322Tag
// Taken from https://github.com/bitcoin/bips/blob/master/bip-0322.mediawiki#full
const BIP322Tag = "BIP0322-signed-message"

// SHA256 returns the SHA256 hash of the input data
func SHA256(data []byte) (hash []byte) {
	sha := sha256.New()
	sha.Write(data[:])
	hash = sha.Sum(nil)
	return
}

// CreateMagicMessage builds a properly signed message.
func CreateMagicMessage(message string) string {
	buffer := bytes.Buffer{}
	buffer.Grow(wire.VarIntSerializeSize(uint64(len(message))))

	// If we cannot write the VarInt, just panic since that should never happen
	if err := wire.WriteVarInt(&buffer, varIntProtoVer, uint64(len(message))); err != nil {
		panic(err)
	}

	return MagicMessage + buffer.String() + message
}

// CreateMagicMessageBIP322 builds a properly signed message (in BIP-322 format).
func CreateMagicMessageBIP322(message []byte) [32]byte {
	tagHash := sha256.Sum256([]byte(BIP322Tag))
	sum := append(tagHash[:], tagHash[:]...) // Append tagHash twice
	sum = append(sum, message...)

	return sha256.Sum256(sum)
}
