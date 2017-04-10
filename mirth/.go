package mirth

type Channel struct {
	Id             string `xml:"id"`
	Name           string `xml:"name"`
	Enabled        string `xml:"enabled"`
	Description    string `xml:"description"`
	FilePath       string
	SourceType     string   `xml:"sourceConnector>properties>scheme"`
	DestinationIds []string `xml:"destinationConnectors>connector>properties>channelId"`
}
