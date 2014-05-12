package buffer

// START OMIT
type Token struct {
	Data  []byte
	owner chan<- *Token
}

// done with the byte slice; returning it for re-use
func (t *Token) Return() {
	t.owner <- t
}

type BufferManager struct {
	buffer chan *Token
}

func NewBufferManager(size int) *BufferManager {
	buffer := make(chan *Token, size)
	ret := &BufferManager{buffer}
	for i := 0; i < size; i++ {
		buffer <- &Token{Data: []byte{}, owner: buffer}
	}
	return ret
}

// get a re-usable byte slice
func (b *BufferManager) GetToken() *Token {
	return <-b.buffer
}
// END OMIT
