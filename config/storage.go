package config

type StorageConfig struct {
	root string
}

func NewStorageConfig(root string) *StorageConfig {
	return &StorageConfig{
		root: root,
	}
}

func (s *StorageConfig) GetRoot() string {
	return s.root
}
