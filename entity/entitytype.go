package entity

type EntityType int

const (
	// FileType represents a regular file
	Directory EntityType = 0
	// DirectoryType represents a directory
	File EntityType = 10
)

// String returns the string representation of the entity type
func (et EntityType) String() string {
	switch et {
	case File:
		return "file"
	case Directory:
		return "directory"
	default:
		return "unknown"
	}
}
