package factory

import "golangpet/internal/database/reader"

func (f *DependencyFactory) createUserReader() reader.UserReaderInterface {
	return reader.NewUserReader(f.db)
}
