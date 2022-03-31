package database

import "fmt"

//MetaData struct describes a data passed by an app to the server
type MetaData struct {
	Timestamp      int64
	ProgramName    string
	Version        string
	Platform       string
	InstallationId string
	Mac            string
	UserName       string
}

//IsValid function checks if a passed metadata is valid,
// meaning that each member has a valid length
func (m *MetaData) IsValid() error {
	// program name, version and platform should not be empty
	if len(m.ProgramName) == 0 {
		return fmt.Errorf("program name is empty")
	}
	if len(m.Version) == 0 {
		return fmt.Errorf("program version is empty")
	}
	if len(m.Platform) == 0 {
		return fmt.Errorf("platform name is empty")
	}

	// TODO check if length is too long?

	// mac and username are hashes
	if len(m.Mac) != 32 {
		return fmt.Errorf("mac hash is invalid")
	}
	if len(m.UserName) != 32 {
		return fmt.Errorf("username hash is invalid")
	}

	// uuid is a valid dash-containing uuid
	if len(m.InstallationId) != 36 {
		return fmt.Errorf("installation id hash is invalid")
	}
	return nil
}

// Connector is an interface for an object that writes
// metadata to a certain storage
type Connector interface {
	AddEntry(d MetaData) error
	Initialize() error
	Close()
}
