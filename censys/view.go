package censys

import (
	"context"
	"net/http"
	"net/url"
	"time"
)

type viewType string

const (
	viewPath = "/view"
	//CERTIFICATESVIEW serach for certificates
	CERTIFICATESVIEW viewType = viewPath + "/certificates/"
	//IPV4VIEW search for ipv4
	IPV4VIEW viewType = viewPath + "/ipv4/"
	//WEBSITESVIEW search for websites
	WEBSITESVIEW viewType = viewPath + "/websites/"
)

//View class containg the view result
type View struct {
	ValidationTimestamp time.Time `json:"validation_timestamp"`
	UpdatedAt           string    `json:"updated_at"`
	Raw                 string    `json:"raw"`
	ValidationNssValid  bool      `json:"validation.nss.valid"`
	Parsed              struct {
		FingerprintSha1 string `json:"fingerprint_sha1"`
		SubjectDn       string `json:"subject_dn"`
		Issuer          struct {
			CommonName []string `json:"common_name"`
		} `json:"issuer"`
		Signature struct {
			SelfSigned         bool   `json:"self_signed"`
			Valid              bool   `json:"valid"`
			Value              string `json:"value"`
			SignatureAlgorithm struct {
				Oid  string `json:"oid"`
				Name string `json:"name"`
			} `json:"signature_algorithm"`
		} `json:"signature"`
		Validity struct {
			Start time.Time `json:"start"`
			End   time.Time `json:"end"`
		} `json:"validity"`
		IssuerDn          string `json:"issuer_dn"`
		FingerprintSha256 string `json:"fingerprint_sha256"`
		Version           int    `json:"version"`
		Extensions        struct {
			AuthorityKeyID   string `json:"authority_key_id"`
			BasicConstraints struct {
				IsCa bool `json:"is_ca"`
			} `json:"basic_constraints"`
			SubjectKeyID        string        `json:"subject_key_id"`
			CertificatePolicies []interface{} `json:"certificate_policies"`
		} `json:"extensions"`
		SignatureAlgorithm struct {
			Oid  string `json:"oid"`
			Name string `json:"name"`
		} `json:"signature_algorithm"`
		SerialNumber   string `json:"serial_number"`
		FingerprintMd5 string `json:"fingerprint_md5"`
		SubjectKeyInfo struct {
			KeyAlgorithm struct {
				Oid  string `json:"oid"`
				Name string `json:"name"`
			} `json:"key_algorithm"`
			RsaPublicKey struct {
				Length   int    `json:"length"`
				Modulus  string `json:"modulus"`
				Exponent int    `json:"exponent"`
			} `json:"rsa_public_key"`
		} `json:"subject_key_info"`
		Subject struct {
			CommonName []string `json:"common_name"`
		} `json:"subject"`
	} `json:"parsed"`
}

//GetView returns a view
func (client *Client) GetView(ctx context.Context, viewty viewType, query string) (*View, error) {
	var view View
	req, err := client.NewRequest(http.MethodGet, string(viewty)+url.QueryEscape(query), nil, nil)
	if err != nil {
		return nil, err
	}
	if err = client.Do(ctx, req, &view); err != nil {
		return nil, err
	}
	return &view, nil
}
