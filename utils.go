package gortmp

type IdGen struct {
    lastId uint32
    Next chan uint32
}

func NewIdGen() (*IdGen){
    gen := &IdGen{lastId:0, Next: make(chan uint32, 10)}
    go func() {
        for {
            gen.Next <- gen.lastId
            gen.lastId++
        }
    }()
    return gen
}
