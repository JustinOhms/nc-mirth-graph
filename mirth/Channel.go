package mirth

type Channel struct {
	Id             string `xml:"id" json:"nodeId"`
	Name           string `xml:"name" json:"nodeName"`
	Enabled        string `xml:"enabled"`
	Description    string `xml:"description"`
	FilePath       string
	SourceType     string   `xml:"sourceConnector>properties>scheme"`
	DestinationIds []string `xml:"destinationConnectors>connector>properties>channelId"`
}
