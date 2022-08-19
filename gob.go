package safe_gob

import (
	"encoding/gob"
	"io"
)

func Write(pack Pack, writer io.Writer) {
	gobEncoder := gob.NewEncoder(writer)
	err := gobEncoder.Encode(pack)
	if err != nil {
		panic(err)
	}
}

func Read(reader io.Reader) Pack {
	pack := new(Pack)
	gobDecoder := gob.NewDecoder(reader)
	err := gobDecoder.Decode(pack)
	if err != nil {
		panic(err)
	}
	return *pack
}
