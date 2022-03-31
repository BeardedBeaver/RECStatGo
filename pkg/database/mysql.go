package database

type MySQLConnector struct {
	uri string
}

func (c *MySQLConnector) AddEntry(d MetaData) error {
	return nil
}

func (c *MySQLConnector) Initialize() error {
	// TODO initialize db and table if needed
	return nil
}

func (c *MySQLConnector) Close() {

}
