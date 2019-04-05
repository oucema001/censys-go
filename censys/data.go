package censys

import (
	"context"
	"net/http"
)

const (
	dataPath = "/data"
)

//DataView struct containg the returned values of the /data endpoint
type DataView struct {
	ID          string `json:"id"`
	Port        int    `json:"port"`
	Protocol    string `json:"protocol"`
	Subprotocol string `json:"subprotocol"`
	Destination string `json:"destination"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Results     struct {
		Historical []struct {
			Timestamp  string `json:"timestamp"`
			ID         string `json:"id"`
			DetailsURL string `json:"details_url"`
		} `json:"historical"`
		Latest struct {
			Timestamp  string `json:"timestamp"`
			ID         string `json:"id"`
			DetailsURL string `json:"details_url"`
		} `json:"latest"`
	} `json:"results"`
}

/*
//Data struct containg the returned values of the /data endpoint
type Data struct {
	PrimarySeries PrimarySeries `json:"primary_series"`
	RawSeries     RawSeries     `json:"raw_series"`
}

//LatestResult struc for
type LatestResult struct {
	Timestamp  string `json:"timestamp"`
	Name       string `json:"name"`
	DetailsURL string `json:"details_url"`
}

//Certificates struc for
type Certificates struct {
	Description  string       `json:"description"`
	DetailsURL   string       `json:"details_url"`
	LatestResult LatestResult `json:"latest_result"`
	ID           string       `json:"id"`
	Name         string       `json:"name"`
}

//Domain struc for
type Domain struct {
	Description  string       `json:"description"`
	DetailsURL   string       `json:"details_url"`
	LatestResult LatestResult `json:"latest_result"`
	ID           string       `json:"id"`
	Name         string       `json:"name"`
}

//Ipv4 struc for
type Ipv4 struct {
	Description  string       `json:"description"`
	DetailsURL   string       `json:"details_url"`
	LatestResult LatestResult `json:"latest_result"`
	ID           string       `json:"id"`
	Name         string       `json:"name"`
}

//PrimarySeries struc for
type PrimarySeries struct {
	Certificates Certificates `json:"certificates"`
	Domain       Domain       `json:"domain"`
	Ipv4         Ipv4         `json:"ipv4"`
}

//Two2SSHBannerFullIpv4 struc for
type Two2SSHBannerFullIpv4 struct {
	Subprotocol  string       `json:"subprotocol"`
	Description  string       `json:"description"`
	Protocol     string       `json:"protocol"`
	Name         string       `json:"name"`
	DetailsURL   string       `json:"details_url"`
	LatestResult LatestResult `json:"latest_result"`
	Destination  string       `json:"destination"`
	ID           string       `json:"id"`
	Port         int          `json:"port"`
}

//Four43HTTPSSsl3AlexaTop1Mil struc for
type Four43HTTPSSsl3AlexaTop1Mil struct {
	Subprotocol  string       `json:"subprotocol"`
	Description  string       `json:"description"`
	Protocol     string       `json:"protocol"`
	Name         string       `json:"name"`
	DetailsURL   string       `json:"details_url"`
	LatestResult LatestResult `json:"latest_result"`
	Destination  string       `json:"destination"`
	ID           string       `json:"id"`
	Port         int          `json:"port"`
}

//RawSeries struc for RawSeries
type RawSeries struct {
	Two2SSHBannerFullIpv4       Two2SSHBannerFullIpv4       `json:"22-ssh-banner-full_ipv4"`
	Four43HTTPSSsl3AlexaTop1Mil Four43HTTPSSsl3AlexaTop1Mil `json:"443-https-ssl3-alexa_top1mil"`
}
*/

//Data struct containg the returned values of the /data endpoint
type Data struct {
	PrimarySeries PrimarySeries `json:"primary_series"`
	RawSeries     RawSeries     `json:"raw_series"`
}

//AlexaTop1MillionSnapshotsDeprecatedFormat contains data returned for
type AlexaTop1MillionSnapshotsDeprecatedFormat struct {
	Description  string      `json:"description"`
	DetailsURL   string      `json:"details_url"`
	LatestResult interface{} `json:"latest_result"`
	ID           string      `json:"id"`
	Name         string      `json:"name"`
}

//AlexaTop1MillionSnapshots contains data returned for
type AlexaTop1MillionSnapshots struct {
	Description  string      `json:"description"`
	DetailsURL   string      `json:"details_url"`
	LatestResult interface{} `json:"latest_result"`
	ID           string      `json:"id"`
	Name         string      `json:"name"`
}

//AllX509CertificatesDeprecatedFormat contains data returned for
type AllX509CertificatesDeprecatedFormat struct {
	Description  string      `json:"description"`
	DetailsURL   string      `json:"details_url"`
	LatestResult interface{} `json:"latest_result"`
	ID           string      `json:"id"`
	Name         string      `json:"name"`
}

//IPv4SnapshotsDeprecatedFormat contains data returned for
type IPv4SnapshotsDeprecatedFormat struct {
	Description  string      `json:"description"`
	DetailsURL   string      `json:"details_url"`
	LatestResult interface{} `json:"latest_result"`
	ID           string      `json:"id"`
	Name         string      `json:"name"`
}

//IPv4Snapshots contains data returned for
type IPv4Snapshots struct {
	Description  string      `json:"description"`
	DetailsURL   string      `json:"details_url"`
	LatestResult interface{} `json:"latest_result"`
	ID           string      `json:"id"`
	Name         string      `json:"name"`
}

//AllX509Certificates contains data returned for
type AllX509Certificates struct {
	Description  string      `json:"description"`
	DetailsURL   string      `json:"details_url"`
	LatestResult interface{} `json:"latest_result"`
	ID           string      `json:"id"`
	Name         string      `json:"name"`
}

//IPv4BannersSnapshots contains data returned for
type IPv4BannersSnapshots struct {
	Description  string      `json:"description"`
	DetailsURL   string      `json:"details_url"`
	LatestResult interface{} `json:"latest_result"`
	ID           string      `json:"id"`
	Name         string      `json:"name"`
}

//CertificatesByDateAdded contains data returned for
type CertificatesByDateAdded struct {
	Description  string      `json:"description"`
	DetailsURL   string      `json:"details_url"`
	LatestResult interface{} `json:"latest_result"`
	ID           string      `json:"id"`
	Name         string      `json:"name"`
}

//PrimarySeries contains data returned for
type PrimarySeries struct {
	AlexaTop1MillionSnapshotsDeprecatedFormat AlexaTop1MillionSnapshotsDeprecatedFormat `json:"Alexa Top 1 Million Snapshots (Deprecated Format)"`
	AlexaTop1MillionSnapshots                 AlexaTop1MillionSnapshots                 `json:"Alexa Top 1 Million Snapshots"`
	AllX509CertificatesDeprecatedFormat       AllX509CertificatesDeprecatedFormat       `json:"All X.509 Certificates (Deprecated Format)"`
	IPv4SnapshotsDeprecatedFormat             IPv4SnapshotsDeprecatedFormat             `json:"IPv4 Snapshots (Deprecated Format)"`
	IPv4Snapshots                             IPv4Snapshots                             `json:"IPv4 Snapshots"`
	AllX509Certificates                       AllX509Certificates                       `json:"All X.509 Certificates"`
	IPv4BannersSnapshots                      IPv4BannersSnapshots                      `json:"IPv4 Banners Snapshots"`
	CertificatesByDateAdded                   CertificatesByDateAdded                   `json:"Certificates by Date Added"`
}

//RawSeries contains data returned for
type RawSeries struct {
}

//GetData calls the /data api and returns a data struct or an error
func (c *Client) GetData(ctx context.Context) (*Data, error) {
	var data Data
	req, err := c.NewRequest(http.MethodGet, dataPath, nil, nil)
	if err != nil {
		return nil, err
	}

	if err := c.Do(ctx, req, &data); err != nil {
		return nil, err
	}
	return &data, nil
}
