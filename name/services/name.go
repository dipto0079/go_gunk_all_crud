package name

import (
	//"blog/category/storage"
	tpm "free_go/gunk/v1/allData"
	//"context"
)

type nameCoreStore interface {
	
}

type Svc struct {
	tpm.UnimplementedNameServiceServer
	core nameCoreStore
}


func NewNameServer(n nameCoreStore) *Svc {
	return &Svc{
		core: n,
	}
}
