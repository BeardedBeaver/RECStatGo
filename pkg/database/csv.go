package database

import (
	"fmt"
	"os"
)

// CSVConnector struct implements a Connector interface.
// Interprets uri as a file name and dumps passed metadata
// to a file in a form of CSV
type CSVConnector struct {
	uri string
}

func (c *CSVConnector) AddEntry(d MetaData) error {
	file, err := os.OpenFile(c.uri, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	s := fmt.Sprintf("%d, %s, %s, %s, %s, %s, %s\n",
		d.Timestamp,
		d.ProgramName,
		d.Version,
		d.Platform,
		d.InstallationId,
		d.Mac,
		d.UserName)
	_, err = file.WriteString(s)
	if err != nil {
		return err
	}
	return nil
}

func (c *CSVConnector) Initialize() error {
	return nil
}

func (c *CSVConnector) Close() {

}
