package storage

import "sync"

type keyTableManager struct {
	userTable    *table
	assertTable  *table
	viewTable    *table
	indexTable   *table
	triggerTable *table
	psmTable     *table
}

//use GetBufferPool to get the unique bufferPool
var ktmInstance *keyTableManager
var ktmOnce sync.Once

func GetKeyTableManager() *keyTableManager {
	ktmOnce.Do(func() {
		ktmInstance = &keyTableManager{}
	})
	return ktmInstance
}
