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
	Num25     Num25     `json:"25"`
	Num80     Num80     `json:"80"`
	Num443    Num443    `json:"443"`
	Domain    string    `json:"domain"`
	AlexaRank int       `json:"alexa_rank"`
	Tags      []string  `json:"tags"`
	UpdatedAt time.Time `json:"updated_at"`
	Ports     []int     `json:"ports"`
	Protocols []string  `json:"protocols"`
}
type Metadata struct {
}
type Starttls struct {
	Ehlo     string   `json:"ehlo"`
	Starttls string   `json:"starttls"`
	Banner   string   `json:"banner"`
	Metadata Metadata `json:"metadata"`
}
type SMTP struct {
	Starttls Starttls `json:"starttls"`
}
type Num25 struct {
	SMTP SMTP `json:"smtp"`
}
type Unknown struct {
	Value string `json:"value"`
	Key   string `json:"key"`
}
type Headers struct {
	Via          string    `json:"via"`
	Unknown      []Unknown `json:"unknown"`
	Age          string    `json:"age"`
	Server       string    `json:"server"`
	Connection   string    `json:"connection"`
	ContentType  string    `json:"content_type"`
	AcceptRanges string    `json:"accept_ranges"`
	CacheControl string    `json:"cache_control"`
}
type Get struct {
	Body       string   `json:"body"`
	Headers    Headers  `json:"headers"`
	StatusCode int      `json:"status_code"`
	Title      string   `json:"title"`
	StatusLine string   `json:"status_line"`
	BodySha256 string   `json:"body_sha256"`
	Metadata   Metadata `json:"metadata"`
}
type HTTPWww struct {
	Get Get `json:"get"`
}
type HTTP struct {
	Get Get `json:"get"`
}
type Num80 struct {
	HTTPWww HTTPWww `json:"http_www"`
	HTTP    HTTP    `json:"http"`
}
type CurveID struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
type EcdhParams struct {
	CurveID CurveID `json:"curve_id"`
}
type ServerKeyExchange struct {
	EcdhParams EcdhParams `json:"ecdh_params"`
}
type Subject struct {
	CommonName   []string `json:"common_name"`
	Country      []string `json:"country"`
	Organization []string `json:"organization"`
	Province     []string `json:"province"`
	Locality     []string `json:"locality"`
}
type SignatureAlgorithm struct {
	Oid  string `json:"oid"`
	Name string `json:"name"`
}
type Validity struct {
	Start  time.Time `json:"start"`
	Length int       `json:"length"`
	End    time.Time `json:"end"`
}
type CertificatePolicies struct {
	Cps []string `json:"cps,omitempty"`
	ID  string   `json:"id"`
}
type AuthorityInfoAccess struct {
	OcspUrls   []string `json:"ocsp_urls"`
	IssuerUrls []string `json:"issuer_urls"`
}
type ExtendedKeyUsage struct {
	ClientAuth bool `json:"client_auth"`
	ServerAuth bool `json:"server_auth"`
}
type SubjectAltName struct {
	DNSNames []string `json:"dns_names"`
}
type BasicConstraints struct {
	IsCa bool `json:"is_ca"`
}
type KeyUsage struct {
	KeyEncipherment  bool `json:"key_encipherment"`
	Value            int  `json:"value"`
	DigitalSignature bool `json:"digital_signature"`
}
type SignedCertificateTimestamps struct {
	LogID     string `json:"log_id"`
	Timestamp int    `json:"timestamp"`
	Version   int    `json:"version"`
	Signature string `json:"signature"`
}
type Extensions struct {
	AuthorityKeyID              string                        `json:"authority_key_id"`
	CertificatePolicies         []CertificatePolicies         `json:"certificate_policies"`
	AuthorityInfoAccess         AuthorityInfoAccess           `json:"authority_info_access"`
	ExtendedKeyUsage            ExtendedKeyUsage              `json:"extended_key_usage"`
	SubjectAltName              SubjectAltName                `json:"subject_alt_name"`
	BasicConstraints            BasicConstraints              `json:"basic_constraints"`
	CrlDistributionPoints       []string                      `json:"crl_distribution_points"`
	KeyUsage                    KeyUsage                      `json:"key_usage"`
	SignedCertificateTimestamps []SignedCertificateTimestamps `json:"signed_certificate_timestamps"`
	SubjectKeyID                string                        `json:"subject_key_id"`
}
type KeyAlgorithm struct {
	Name string `json:"name"`
}
type RsaPublicKey struct {
	Length   int    `json:"length"`
	Modulus  string `json:"modulus"`
	Exponent int    `json:"exponent"`
}
type SubjectKeyInfo struct {
	FingerprintSha256 string       `json:"fingerprint_sha256"`
	KeyAlgorithm      KeyAlgorithm `json:"key_algorithm"`
	RsaPublicKey      RsaPublicKey `json:"rsa_public_key"`
}
type Signature struct {
	SelfSigned         bool               `json:"self_signed"`
	Valid              bool               `json:"valid"`
	Value              string             `json:"value"`
	SignatureAlgorithm SignatureAlgorithm `json:"signature_algorithm"`
}
type Issuer struct {
	CommonName   []string `json:"common_name"`
	Country      []string `json:"country"`
	Organization []string `json:"organization"`
}
type Parsed struct {
	TbsNoctFingerprint     string             `json:"tbs_noct_fingerprint"`
	SubjectDn              string             `json:"subject_dn"`
	Subject                Subject            `json:"subject"`
	SignatureAlgorithm     SignatureAlgorithm `json:"signature_algorithm"`
	Redacted               bool               `json:"redacted"`
	SerialNumber           string             `json:"serial_number"`
	ValidationLevel        string             `json:"validation_level"`
	IssuerDn               string             `json:"issuer_dn"`
	FingerprintSha1        string             `json:"fingerprint_sha1"`
	Version                int                `json:"version"`
	FingerprintSha256      string             `json:"fingerprint_sha256"`
	Names                  []string           `json:"names"`
	TbsFingerprint         string             `json:"tbs_fingerprint"`
	Validity               Validity           `json:"validity"`
	Extensions             Extensions         `json:"extensions"`
	FingerprintMd5         string             `json:"fingerprint_md5"`
	SubjectKeyInfo         SubjectKeyInfo     `json:"subject_key_info"`
	Signature              Signature          `json:"signature"`
	SpkiSubjectFingerprint string             `json:"spki_subject_fingerprint"`
	Issuer                 Issuer             `json:"issuer"`
}
type Certificate struct {
	Parsed Parsed `json:"parsed"`
}
type Chain struct {
	Parsed Parsed `json:"parsed"`
}
type CipherSuite struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
type SessionTicket struct {
	Length       int `json:"length"`
	LifetimeHint int `json:"lifetime_hint"`
}

type Validation struct {
	MatchesDomain  bool `json:"matches_domain"`
	BrowserTrusted bool `json:"browser_trusted"`
}
type TLS struct {
	ServerKeyExchange ServerKeyExchange `json:"server_key_exchange"`
	Certificate       Certificate       `json:"certificate"`
	Chain             []Chain           `json:"chain"`
	CipherSuite       CipherSuite       `json:"cipher_suite"`
	Version           string            `json:"version"`
	SessionTicket     SessionTicket     `json:"session_ticket"`
	Signature         Signature         `json:"signature"`
	Validation        Validation        `json:"validation"`
	OcspStapling      bool              `json:"ocsp_stapling"`
	Metadata          Metadata          `json:"metadata"`
}
type HTTPSWww struct {
	TLS TLS `json:"tls"`
}

type DheExport struct {
	Support  bool     `json:"support"`
	Metadata Metadata `json:"metadata"`
}
type Prime struct {
	Length int    `json:"length"`
	Value  string `json:"value"`
}
type Generator struct {
	Length int    `json:"length"`
	Value  string `json:"value"`
}
type DhParams struct {
	Prime     Prime     `json:"prime"`
	Generator Generator `json:"generator"`
}
type Dhe struct {
	Support  bool     `json:"support"`
	DhParams DhParams `json:"dh_params"`
	Metadata Metadata `json:"metadata"`
}
type Heartbleed struct {
	HeartbeatEnabled     bool     `json:"heartbeat_enabled"`
	HeartbleedVulnerable bool     `json:"heartbleed_vulnerable"`
	Metadata             Metadata `json:"metadata"`
}
type RsaExport struct {
	Support  bool     `json:"support"`
	Metadata Metadata `json:"metadata"`
}
type HTTPS struct {
	TLS        TLS        `json:"tls"`
	DheExport  DheExport  `json:"dhe_export"`
	Get        Get        `json:"get"`
	Dhe        Dhe        `json:"dhe"`
	Heartbleed Heartbleed `json:"heartbleed"`
	RsaExport  RsaExport  `json:"rsa_export"`
}
type Num443 struct {
	HTTPSWww HTTPSWww `json:"https_www"`
	HTTPS    HTTPS    `json:"https"`
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
