package database

type DataBase interface {
	Start() error
	Exit() error
}

type MetqlDataBase struct {
	masterbase string
}

func NewDatabase(masterbase string) DataBase {
	db := &MetqlDataBase{}
	db.masterbase = masterbase
	return db
}

func (db *MetqlDataBase) Start() error {
	// init
	return nil
}

func (db *MetqlDataBase) Exit() error {
	// cleanup
	return nil
}
