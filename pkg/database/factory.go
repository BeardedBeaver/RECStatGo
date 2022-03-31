package database

import "fmt"

func NewConnector(name, uri string) (Connector, error) {
	var c Connector = nil
	if name == "MySQL" {
		c = &MySQLConnector{uri}
	} else if name == "CSV" {
		c = &CSVConnector{uri}
	}
	if c == nil {
		return nil, fmt.Errorf("unknown connector type requested: %s", name)
	}
	err := c.Initialize()
	if err != nil {
		return nil, err
	}

	return c, nil
}
