package FileSys

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

var VirtualDisk [100][1024]byte

type SuperBlock struct {
	INodeBitmapStart     int
	FreeBlockBitmapStart int
	DataBlockStart       int
}

func InitializeFileSystem() {
	initializeSuperBlock()
	sblock := ReadSuperBlock()
	fmt.Println(sblock)
}

func initializeSuperBlock() {
	superBlock := SuperBlock{
		INodeBitmapStart:     1,
		FreeBlockBitmapStart: 2,
		DataBlockStart:       7,
	}
	superBlockBytes := EncodeToBytes(superBlock)
	copy(VirtualDisk[0][:], superBlockBytes)
}

func ReadSuperBlock() SuperBlock {
	sBlock := SuperBlock{}
	decoder := gob.NewDecoder(bytes.NewReader(VirtualDisk[0][:]))
	err := decoder.Decode(&sBlock)
	if err != nil {
		log.Fatal("Unable to Decode superblock - better blue Screen", err)
	}
	return sBlock
}

// from https://gist.github.com/SteveBate/042960baa7a4795c3565
func EncodeToBytes(p interface{}) []byte {

	buf := bytes.Buffer{}
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(p)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("uncompressed size (bytes): ", len(buf.Bytes()))
	return buf.Bytes()
}
