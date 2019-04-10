package censys

import (
	"context"
	"fmt"
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
	Num0 struct {
		Lookup struct {
			Dmarc struct {
				Raw string `json:"raw"`
			} `json:"dmarc"`
			Axfr struct {
				Support bool `json:"support"`
				Servers []struct {
					Status string `json:"status"`
					Server string `json:"server"`
					Error  string `json:"error"`
				} `json:"servers"`
				Truncated bool `json:"truncated"`
			} `json:"axfr"`
			Spf struct {
				Raw string `json:"raw"`
			} `json:"spf"`
		} `json:"lookup"`
	} `json:"0"`
	Num25 struct {
		SMTP struct {
			Starttls struct {
				Ehlo string `json:"ehlo"`
				TLS  struct {
					ServerKeyExchange struct {
						EcdhParams struct {
							CurveID struct {
								ID   int    `json:"id"`
								Name string `json:"name"`
							} `json:"curve_id"`
						} `json:"ecdh_params"`
					} `json:"server_key_exchange"`
					Certificate struct {
						Parsed struct {
							TbsNoctFingerprint string `json:"tbs_noct_fingerprint"`
							SubjectDn          string `json:"subject_dn"`
							Subject            struct {
								CommonName   []string `json:"common_name"`
								Country      []string `json:"country"`
								Organization []string `json:"organization"`
								Province     []string `json:"province"`
								Locality     []string `json:"locality"`
							} `json:"subject"`
							SignatureAlgorithm struct {
								Oid  string `json:"oid"`
								Name string `json:"name"`
							} `json:"signature_algorithm"`
							Redacted          bool     `json:"redacted"`
							SerialNumber      string   `json:"serial_number"`
							ValidationLevel   string   `json:"validation_level"`
							IssuerDn          string   `json:"issuer_dn"`
							FingerprintSha1   string   `json:"fingerprint_sha1"`
							Version           int      `json:"version"`
							FingerprintSha256 string   `json:"fingerprint_sha256"`
							Names             []string `json:"names"`
							TbsFingerprint    string   `json:"tbs_fingerprint"`
							Validity          struct {
								Start  time.Time `json:"start"`
								Length int       `json:"length"`
								End    time.Time `json:"end"`
							} `json:"validity"`
							Extensions struct {
								AuthorityKeyID      string `json:"authority_key_id"`
								CertificatePolicies []struct {
									ID string `json:"id"`
								} `json:"certificate_policies"`
								AuthorityInfoAccess struct {
									OcspUrls   []string `json:"ocsp_urls"`
									IssuerUrls []string `json:"issuer_urls"`
								} `json:"authority_info_access"`
								ExtendedKeyUsage struct {
									ServerAuth bool `json:"server_auth"`
								} `json:"extended_key_usage"`
								SubjectAltName struct {
									DNSNames []string `json:"dns_names"`
								} `json:"subject_alt_name"`
								BasicConstraints struct {
									IsCa bool `json:"is_ca"`
								} `json:"basic_constraints"`
								CrlDistributionPoints []string `json:"crl_distribution_points"`
								SubjectKeyID          string   `json:"subject_key_id"`
							} `json:"extensions"`
							FingerprintMd5 string `json:"fingerprint_md5"`
							SubjectKeyInfo struct {
								FingerprintSha256 string `json:"fingerprint_sha256"`
								KeyAlgorithm      struct {
									Name string `json:"name"`
								} `json:"key_algorithm"`
								RsaPublicKey struct {
									Length   int    `json:"length"`
									Modulus  string `json:"modulus"`
									Exponent int    `json:"exponent"`
								} `json:"rsa_public_key"`
							} `json:"subject_key_info"`
							Signature struct {
								SelfSigned         bool   `json:"self_signed"`
								Valid              bool   `json:"valid"`
								Value              string `json:"value"`
								SignatureAlgorithm struct {
									Oid  string `json:"oid"`
									Name string `json:"name"`
								} `json:"signature_algorithm"`
							} `json:"signature"`
							SpkiSubjectFingerprint string `json:"spki_subject_fingerprint"`
							Issuer                 struct {
								CommonName   []string `json:"common_name"`
								Country      []string `json:"country"`
								Organization []string `json:"organization"`
							} `json:"issuer"`
						} `json:"parsed"`
					} `json:"certificate"`
					Chain []struct {
						Parsed struct {
							TbsNoctFingerprint string `json:"tbs_noct_fingerprint"`
							SubjectDn          string `json:"subject_dn"`
							Subject            struct {
								CommonName   []string `json:"common_name"`
								Country      []string `json:"country"`
								Organization []string `json:"organization"`
							} `json:"subject"`
							SignatureAlgorithm struct {
								Oid  string `json:"oid"`
								Name string `json:"name"`
							} `json:"signature_algorithm"`
							Redacted          bool   `json:"redacted"`
							SerialNumber      string `json:"serial_number"`
							ValidationLevel   string `json:"validation_level"`
							IssuerDn          string `json:"issuer_dn"`
							FingerprintSha1   string `json:"fingerprint_sha1"`
							Version           int    `json:"version"`
							FingerprintSha256 string `json:"fingerprint_sha256"`
							Validity          struct {
								Start  time.Time `json:"start"`
								Length int       `json:"length"`
								End    time.Time `json:"end"`
							} `json:"validity"`
							TbsFingerprint string `json:"tbs_fingerprint"`
							Extensions     struct {
								AuthorityKeyID      string `json:"authority_key_id"`
								CertificatePolicies []struct {
									Cps []string `json:"cps"`
									ID  string   `json:"id"`
								} `json:"certificate_policies"`
								ExtendedKeyUsage struct {
									ClientAuth bool `json:"client_auth"`
									ServerAuth bool `json:"server_auth"`
								} `json:"extended_key_usage"`
								AuthorityInfoAccess struct {
									OcspUrls []string `json:"ocsp_urls"`
								} `json:"authority_info_access"`
								BasicConstraints struct {
									MaxPathLen int  `json:"max_path_len"`
									IsCa       bool `json:"is_ca"`
								} `json:"basic_constraints"`
								CrlDistributionPoints []string `json:"crl_distribution_points"`
								KeyUsage              struct {
									CertificateSign  bool `json:"certificate_sign"`
									CrlSign          bool `json:"crl_sign"`
									Value            int  `json:"value"`
									DigitalSignature bool `json:"digital_signature"`
								} `json:"key_usage"`
								SubjectKeyID string `json:"subject_key_id"`
							} `json:"extensions"`
							FingerprintMd5 string `json:"fingerprint_md5"`
							SubjectKeyInfo struct {
								FingerprintSha256 string `json:"fingerprint_sha256"`
								KeyAlgorithm      struct {
									Name string `json:"name"`
								} `json:"key_algorithm"`
								RsaPublicKey struct {
									Length   int    `json:"length"`
									Modulus  string `json:"modulus"`
									Exponent int    `json:"exponent"`
								} `json:"rsa_public_key"`
							} `json:"subject_key_info"`
							Signature struct {
								SelfSigned         bool   `json:"self_signed"`
								Valid              bool   `json:"valid"`
								Value              string `json:"value"`
								SignatureAlgorithm struct {
									Oid  string `json:"oid"`
									Name string `json:"name"`
								} `json:"signature_algorithm"`
							} `json:"signature"`
							SpkiSubjectFingerprint string `json:"spki_subject_fingerprint"`
							Issuer                 struct {
								CommonName         []string `json:"common_name"`
								Organization       []string `json:"organization"`
								OrganizationalUnit []string `json:"organizational_unit"`
							} `json:"issuer"`
						} `json:"parsed"`
					} `json:"chain"`
					CipherSuite struct {
						ID   string `json:"id"`
						Name string `json:"name"`
					} `json:"cipher_suite"`
					Version      string `json:"version"`
					OcspStapling bool   `json:"ocsp_stapling"`
					Signature    struct {
						HashAlgorithm      string `json:"hash_algorithm"`
						Valid              bool   `json:"valid"`
						SignatureAlgorithm string `json:"signature_algorithm"`
					} `json:"signature"`
					Scts []struct {
						LogID     string `json:"log_id"`
						Timestamp int    `json:"timestamp"`
						Version   int    `json:"version"`
						Signature string `json:"signature"`
					} `json:"scts"`
					Validation struct {
						BrowserTrusted bool `json:"browser_trusted"`
					} `json:"validation"`
				} `json:"tls"`
				Starttls string `json:"starttls"`
				Banner   string `json:"banner"`
				Metadata struct {
					Product      string `json:"product"`
					Description  string `json:"description"`
					Manufacturer string `json:"manufacturer"`
				} `json:"metadata"`
			} `json:"starttls"`
		} `json:"smtp"`
	} `json:"25"`
	Num80 struct {
		HTTPWww struct {
			Get struct {
				Body    string `json:"body"`
				Headers struct {
					XXSSProtection string `json:"x_xss_protection"`
					Unknown        []struct {
						Value string `json:"value"`
						Key   string `json:"key"`
					} `json:"unknown"`
					Expires       string `json:"expires"`
					Server        string `json:"server"`
					ContentType   string `json:"content_type"`
					P3P           string `json:"p3p"`
					XFrameOptions string `json:"x_frame_options"`
					CacheControl  string `json:"cache_control"`
				} `json:"headers"`
				StatusCode int    `json:"status_code"`
				Title      string `json:"title"`
				StatusLine string `json:"status_line"`
				BodySha256 string `json:"body_sha256"`
				Metadata   struct {
				} `json:"metadata"`
			} `json:"get"`
		} `json:"http_www"`
		HTTP struct {
			Get struct {
				Body    string `json:"body"`
				Headers struct {
					XXSSProtection string `json:"x_xss_protection"`
					Unknown        []struct {
						Value string `json:"value"`
						Key   string `json:"key"`
					} `json:"unknown"`
					Expires       string `json:"expires"`
					Server        string `json:"server"`
					ContentType   string `json:"content_type"`
					P3P           string `json:"p3p"`
					XFrameOptions string `json:"x_frame_options"`
					CacheControl  string `json:"cache_control"`
				} `json:"headers"`
				StatusCode int    `json:"status_code"`
				Title      string `json:"title"`
				StatusLine string `json:"status_line"`
				BodySha256 string `json:"body_sha256"`
				Metadata   struct {
					Product      string `json:"product"`
					Description  string `json:"description"`
					Manufacturer string `json:"manufacturer"`
				} `json:"metadata"`
			} `json:"get"`
		} `json:"http"`
	} `json:"80"`
	Num443 struct {
		HTTPS struct {
			TLS struct {
				ServerKeyExchange struct {
					EcdhParams struct {
						CurveID struct {
							ID   int    `json:"id"`
							Name string `json:"name"`
						} `json:"curve_id"`
					} `json:"ecdh_params"`
				} `json:"server_key_exchange"`
				Certificate struct {
					Parsed struct {
						TbsNoctFingerprint string `json:"tbs_noct_fingerprint"`
						SubjectDn          string `json:"subject_dn"`
						Subject            struct {
							CommonName   []string `json:"common_name"`
							Country      []string `json:"country"`
							Organization []string `json:"organization"`
							Province     []string `json:"province"`
							Locality     []string `json:"locality"`
						} `json:"subject"`
						SignatureAlgorithm struct {
							Oid  string `json:"oid"`
							Name string `json:"name"`
						} `json:"signature_algorithm"`
						Redacted          bool     `json:"redacted"`
						SerialNumber      string   `json:"serial_number"`
						ValidationLevel   string   `json:"validation_level"`
						IssuerDn          string   `json:"issuer_dn"`
						FingerprintSha1   string   `json:"fingerprint_sha1"`
						Version           int      `json:"version"`
						FingerprintSha256 string   `json:"fingerprint_sha256"`
						Names             []string `json:"names"`
						TbsFingerprint    string   `json:"tbs_fingerprint"`
						Validity          struct {
							Start  time.Time `json:"start"`
							Length int       `json:"length"`
							End    time.Time `json:"end"`
						} `json:"validity"`
						Extensions struct {
							AuthorityKeyID      string `json:"authority_key_id"`
							CertificatePolicies []struct {
								ID string `json:"id"`
							} `json:"certificate_policies"`
							AuthorityInfoAccess struct {
								OcspUrls   []string `json:"ocsp_urls"`
								IssuerUrls []string `json:"issuer_urls"`
							} `json:"authority_info_access"`
							ExtendedKeyUsage struct {
								ServerAuth bool `json:"server_auth"`
							} `json:"extended_key_usage"`
							SubjectAltName struct {
								DNSNames []string `json:"dns_names"`
							} `json:"subject_alt_name"`
							BasicConstraints struct {
								IsCa bool `json:"is_ca"`
							} `json:"basic_constraints"`
							CrlDistributionPoints []string `json:"crl_distribution_points"`
							KeyUsage              struct {
								Value            int  `json:"value"`
								DigitalSignature bool `json:"digital_signature"`
							} `json:"key_usage"`
							SubjectKeyID string `json:"subject_key_id"`
						} `json:"extensions"`
						FingerprintMd5 string `json:"fingerprint_md5"`
						SubjectKeyInfo struct {
							FingerprintSha256 string `json:"fingerprint_sha256"`
							KeyAlgorithm      struct {
								Name string `json:"name"`
							} `json:"key_algorithm"`
							EcdsaPublicKey struct {
								B      string `json:"b"`
								Curve  string `json:"curve"`
								Gy     string `json:"gy"`
								N      string `json:"n"`
								P      string `json:"p"`
								Length int    `json:"length"`
								Pub    string `json:"pub"`
								Y      string `json:"y"`
								X      string `json:"x"`
								Gx     string `json:"gx"`
							} `json:"ecdsa_public_key"`
						} `json:"subject_key_info"`
						Signature struct {
							SelfSigned         bool   `json:"self_signed"`
							Valid              bool   `json:"valid"`
							Value              string `json:"value"`
							SignatureAlgorithm struct {
								Oid  string `json:"oid"`
								Name string `json:"name"`
							} `json:"signature_algorithm"`
						} `json:"signature"`
						SpkiSubjectFingerprint string `json:"spki_subject_fingerprint"`
						Issuer                 struct {
							CommonName   []string `json:"common_name"`
							Country      []string `json:"country"`
							Organization []string `json:"organization"`
						} `json:"issuer"`
					} `json:"parsed"`
				} `json:"certificate"`
				Chain []struct {
					Parsed struct {
						TbsNoctFingerprint string `json:"tbs_noct_fingerprint"`
						SubjectDn          string `json:"subject_dn"`
						Subject            struct {
							CommonName   []string `json:"common_name"`
							Country      []string `json:"country"`
							Organization []string `json:"organization"`
						} `json:"subject"`
						SignatureAlgorithm struct {
							Oid  string `json:"oid"`
							Name string `json:"name"`
						} `json:"signature_algorithm"`
						Redacted          bool   `json:"redacted"`
						SerialNumber      string `json:"serial_number"`
						ValidationLevel   string `json:"validation_level"`
						IssuerDn          string `json:"issuer_dn"`
						FingerprintSha1   string `json:"fingerprint_sha1"`
						Version           int    `json:"version"`
						FingerprintSha256 string `json:"fingerprint_sha256"`
						Validity          struct {
							Start  time.Time `json:"start"`
							Length int       `json:"length"`
							End    time.Time `json:"end"`
						} `json:"validity"`
						TbsFingerprint string `json:"tbs_fingerprint"`
						Extensions     struct {
							AuthorityKeyID      string `json:"authority_key_id"`
							CertificatePolicies []struct {
								Cps []string `json:"cps"`
								ID  string   `json:"id"`
							} `json:"certificate_policies"`
							ExtendedKeyUsage struct {
								ClientAuth bool `json:"client_auth"`
								ServerAuth bool `json:"server_auth"`
							} `json:"extended_key_usage"`
							AuthorityInfoAccess struct {
								OcspUrls []string `json:"ocsp_urls"`
							} `json:"authority_info_access"`
							BasicConstraints struct {
								MaxPathLen int  `json:"max_path_len"`
								IsCa       bool `json:"is_ca"`
							} `json:"basic_constraints"`
							CrlDistributionPoints []string `json:"crl_distribution_points"`
							KeyUsage              struct {
								CertificateSign  bool `json:"certificate_sign"`
								CrlSign          bool `json:"crl_sign"`
								Value            int  `json:"value"`
								DigitalSignature bool `json:"digital_signature"`
							} `json:"key_usage"`
							SubjectKeyID string `json:"subject_key_id"`
						} `json:"extensions"`
						FingerprintMd5 string `json:"fingerprint_md5"`
						SubjectKeyInfo struct {
							FingerprintSha256 string `json:"fingerprint_sha256"`
							KeyAlgorithm      struct {
								Name string `json:"name"`
							} `json:"key_algorithm"`
							RsaPublicKey struct {
								Length   int    `json:"length"`
								Modulus  string `json:"modulus"`
								Exponent int    `json:"exponent"`
							} `json:"rsa_public_key"`
						} `json:"subject_key_info"`
						Signature struct {
							SelfSigned         bool   `json:"self_signed"`
							Valid              bool   `json:"valid"`
							Value              string `json:"value"`
							SignatureAlgorithm struct {
								Oid  string `json:"oid"`
								Name string `json:"name"`
							} `json:"signature_algorithm"`
						} `json:"signature"`
						SpkiSubjectFingerprint string `json:"spki_subject_fingerprint"`
						Issuer                 struct {
							CommonName         []string `json:"common_name"`
							Organization       []string `json:"organization"`
							OrganizationalUnit []string `json:"organizational_unit"`
						} `json:"issuer"`
					} `json:"parsed"`
				} `json:"chain"`
				CipherSuite struct {
					ID   string `json:"id"`
					Name string `json:"name"`
				} `json:"cipher_suite"`
				Version       string `json:"version"`
				SessionTicket struct {
					Length       int `json:"length"`
					LifetimeHint int `json:"lifetime_hint"`
				} `json:"session_ticket"`
				Signature struct {
					HashAlgorithm      string `json:"hash_algorithm"`
					Valid              bool   `json:"valid"`
					SignatureAlgorithm string `json:"signature_algorithm"`
				} `json:"signature"`
				Scts []struct {
					LogID     string `json:"log_id"`
					Timestamp int    `json:"timestamp"`
					Version   int    `json:"version"`
					Signature string `json:"signature"`
				} `json:"scts"`
				Validation struct {
					MatchesDomain  bool `json:"matches_domain"`
					BrowserTrusted bool `json:"browser_trusted"`
				} `json:"validation"`
				OcspStapling bool `json:"ocsp_stapling"`
				Metadata     struct {
				} `json:"metadata"`
			} `json:"tls"`
			DheExport struct {
				Support  bool `json:"support"`
				Metadata struct {
				} `json:"metadata"`
			} `json:"dhe_export"`
			Get struct {
				Body    string `json:"body"`
				Headers struct {
					XXSSProtection string `json:"x_xss_protection"`
					Unknown        []struct {
						Value string `json:"value"`
						Key   string `json:"key"`
					} `json:"unknown"`
					Expires       string `json:"expires"`
					Server        string `json:"server"`
					AltSvc        string `json:"alt_svc"`
					ContentType   string `json:"content_type"`
					P3P           string `json:"p3p"`
					XFrameOptions string `json:"x_frame_options"`
					CacheControl  string `json:"cache_control"`
				} `json:"headers"`
				StatusCode int    `json:"status_code"`
				Title      string `json:"title"`
				StatusLine string `json:"status_line"`
				BodySha256 string `json:"body_sha256"`
				Metadata   struct {
					Product      string `json:"product"`
					Description  string `json:"description"`
					Manufacturer string `json:"manufacturer"`
				} `json:"metadata"`
			} `json:"get"`
			Dhe struct {
				Support  bool `json:"support"`
				Metadata struct {
				} `json:"metadata"`
			} `json:"dhe"`
			Heartbleed struct {
				HeartbeatEnabled     bool `json:"heartbeat_enabled"`
				HeartbleedVulnerable bool `json:"heartbleed_vulnerable"`
				Metadata             struct {
				} `json:"metadata"`
			} `json:"heartbleed"`
			RsaExport struct {
				Support  bool `json:"support"`
				Metadata struct {
				} `json:"metadata"`
			} `json:"rsa_export"`
		} `json:"https"`
		HTTPSWww struct {
			TLS struct {
				ServerKeyExchange struct {
					EcdhParams struct {
						CurveID struct {
							ID   int    `json:"id"`
							Name string `json:"name"`
						} `json:"curve_id"`
					} `json:"ecdh_params"`
				} `json:"server_key_exchange"`
				Certificate struct {
					Parsed struct {
						TbsNoctFingerprint string `json:"tbs_noct_fingerprint"`
						SubjectDn          string `json:"subject_dn"`
						Subject            struct {
							CommonName   []string `json:"common_name"`
							Country      []string `json:"country"`
							Organization []string `json:"organization"`
							Province     []string `json:"province"`
							Locality     []string `json:"locality"`
						} `json:"subject"`
						SignatureAlgorithm struct {
							Oid  string `json:"oid"`
							Name string `json:"name"`
						} `json:"signature_algorithm"`
						Redacted          bool     `json:"redacted"`
						SerialNumber      string   `json:"serial_number"`
						ValidationLevel   string   `json:"validation_level"`
						IssuerDn          string   `json:"issuer_dn"`
						FingerprintSha1   string   `json:"fingerprint_sha1"`
						Version           int      `json:"version"`
						FingerprintSha256 string   `json:"fingerprint_sha256"`
						Names             []string `json:"names"`
						TbsFingerprint    string   `json:"tbs_fingerprint"`
						Validity          struct {
							Start  time.Time `json:"start"`
							Length int       `json:"length"`
							End    time.Time `json:"end"`
						} `json:"validity"`
						Extensions struct {
							AuthorityKeyID      string `json:"authority_key_id"`
							CertificatePolicies []struct {
								ID string `json:"id"`
							} `json:"certificate_policies"`
							AuthorityInfoAccess struct {
								OcspUrls   []string `json:"ocsp_urls"`
								IssuerUrls []string `json:"issuer_urls"`
							} `json:"authority_info_access"`
							ExtendedKeyUsage struct {
								ServerAuth bool `json:"server_auth"`
							} `json:"extended_key_usage"`
							SubjectAltName struct {
								DNSNames []string `json:"dns_names"`
							} `json:"subject_alt_name"`
							BasicConstraints struct {
								IsCa bool `json:"is_ca"`
							} `json:"basic_constraints"`
							CrlDistributionPoints []string `json:"crl_distribution_points"`
							KeyUsage              struct {
								Value            int  `json:"value"`
								DigitalSignature bool `json:"digital_signature"`
							} `json:"key_usage"`
							SubjectKeyID string `json:"subject_key_id"`
						} `json:"extensions"`
						FingerprintMd5 string `json:"fingerprint_md5"`
						SubjectKeyInfo struct {
							FingerprintSha256 string `json:"fingerprint_sha256"`
							KeyAlgorithm      struct {
								Name string `json:"name"`
							} `json:"key_algorithm"`
							EcdsaPublicKey struct {
								B      string `json:"b"`
								Curve  string `json:"curve"`
								Gy     string `json:"gy"`
								N      string `json:"n"`
								P      string `json:"p"`
								Length int    `json:"length"`
								Pub    string `json:"pub"`
								Y      string `json:"y"`
								X      string `json:"x"`
								Gx     string `json:"gx"`
							} `json:"ecdsa_public_key"`
						} `json:"subject_key_info"`
						Signature struct {
							SelfSigned         bool   `json:"self_signed"`
							Valid              bool   `json:"valid"`
							Value              string `json:"value"`
							SignatureAlgorithm struct {
								Oid  string `json:"oid"`
								Name string `json:"name"`
							} `json:"signature_algorithm"`
						} `json:"signature"`
						SpkiSubjectFingerprint string `json:"spki_subject_fingerprint"`
						Issuer                 struct {
							CommonName   []string `json:"common_name"`
							Country      []string `json:"country"`
							Organization []string `json:"organization"`
						} `json:"issuer"`
					} `json:"parsed"`
				} `json:"certificate"`
				Chain []struct {
					Parsed struct {
						TbsNoctFingerprint string `json:"tbs_noct_fingerprint"`
						SubjectDn          string `json:"subject_dn"`
						Subject            struct {
							CommonName   []string `json:"common_name"`
							Country      []string `json:"country"`
							Organization []string `json:"organization"`
						} `json:"subject"`
						SignatureAlgorithm struct {
							Oid  string `json:"oid"`
							Name string `json:"name"`
						} `json:"signature_algorithm"`
						Redacted          bool   `json:"redacted"`
						SerialNumber      string `json:"serial_number"`
						ValidationLevel   string `json:"validation_level"`
						IssuerDn          string `json:"issuer_dn"`
						FingerprintSha1   string `json:"fingerprint_sha1"`
						Version           int    `json:"version"`
						FingerprintSha256 string `json:"fingerprint_sha256"`
						Validity          struct {
							Start  time.Time `json:"start"`
							Length int       `json:"length"`
							End    time.Time `json:"end"`
						} `json:"validity"`
						TbsFingerprint string `json:"tbs_fingerprint"`
						Extensions     struct {
							AuthorityKeyID      string `json:"authority_key_id"`
							CertificatePolicies []struct {
								Cps []string `json:"cps"`
								ID  string   `json:"id"`
							} `json:"certificate_policies"`
							ExtendedKeyUsage struct {
								ClientAuth bool `json:"client_auth"`
								ServerAuth bool `json:"server_auth"`
							} `json:"extended_key_usage"`
							AuthorityInfoAccess struct {
								OcspUrls []string `json:"ocsp_urls"`
							} `json:"authority_info_access"`
							BasicConstraints struct {
								MaxPathLen int  `json:"max_path_len"`
								IsCa       bool `json:"is_ca"`
							} `json:"basic_constraints"`
							CrlDistributionPoints []string `json:"crl_distribution_points"`
							KeyUsage              struct {
								CertificateSign  bool `json:"certificate_sign"`
								CrlSign          bool `json:"crl_sign"`
								Value            int  `json:"value"`
								DigitalSignature bool `json:"digital_signature"`
							} `json:"key_usage"`
							SubjectKeyID string `json:"subject_key_id"`
						} `json:"extensions"`
						FingerprintMd5 string `json:"fingerprint_md5"`
						SubjectKeyInfo struct {
							FingerprintSha256 string `json:"fingerprint_sha256"`
							KeyAlgorithm      struct {
								Name string `json:"name"`
							} `json:"key_algorithm"`
							RsaPublicKey struct {
								Length   int    `json:"length"`
								Modulus  string `json:"modulus"`
								Exponent int    `json:"exponent"`
							} `json:"rsa_public_key"`
						} `json:"subject_key_info"`
						Signature struct {
							SelfSigned         bool   `json:"self_signed"`
							Valid              bool   `json:"valid"`
							Value              string `json:"value"`
							SignatureAlgorithm struct {
								Oid  string `json:"oid"`
								Name string `json:"name"`
							} `json:"signature_algorithm"`
						} `json:"signature"`
						SpkiSubjectFingerprint string `json:"spki_subject_fingerprint"`
						Issuer                 struct {
							CommonName         []string `json:"common_name"`
							Organization       []string `json:"organization"`
							OrganizationalUnit []string `json:"organizational_unit"`
						} `json:"issuer"`
					} `json:"parsed"`
				} `json:"chain"`
				CipherSuite struct {
					ID   string `json:"id"`
					Name string `json:"name"`
				} `json:"cipher_suite"`
				Version       string `json:"version"`
				SessionTicket struct {
					Length       int `json:"length"`
					LifetimeHint int `json:"lifetime_hint"`
				} `json:"session_ticket"`
				Signature struct {
					HashAlgorithm      string `json:"hash_algorithm"`
					Valid              bool   `json:"valid"`
					SignatureAlgorithm string `json:"signature_algorithm"`
				} `json:"signature"`
				Scts []struct {
					LogID     string `json:"log_id"`
					Timestamp int    `json:"timestamp"`
					Version   int    `json:"version"`
					Signature string `json:"signature"`
				} `json:"scts"`
				Validation struct {
					MatchesDomain  bool `json:"matches_domain"`
					BrowserTrusted bool `json:"browser_trusted"`
				} `json:"validation"`
				OcspStapling bool `json:"ocsp_stapling"`
				Metadata     struct {
				} `json:"metadata"`
			} `json:"tls"`
		} `json:"https_www"`
	} `json:"443"`
	Domain    string    `json:"domain"`
	AlexaRank int       `json:"alexa_rank"`
	Tags      []string  `json:"tags"`
	UpdatedAt time.Time `json:"updated_at"`
	Ports     []int     `json:"ports"`
	Protocols []string  `json:"protocols"`
}

//GetView returns a view of a result obtained by search
func (client *Client) GetView(ctx context.Context, viewty viewType, query string) (*View, error) {
	var view View
	var s = string(viewty) + url.QueryEscape(query)
	fmt.Println(s)
	req, err := client.NewRequest(http.MethodGet, s, nil, nil)
	if err != nil {
		return nil, err
	}
	if err = client.Do(ctx, req, &view); err != nil {
		return nil, err
	}
	return &view, nil
}
