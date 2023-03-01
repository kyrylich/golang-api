package factory

import "golangpet/internal/database/writer"

func (f *DependencyFactory) createUserWriter() writer.UserWriterInterface {
	return writer.NewUserWriter(f.db)
}
