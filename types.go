package vice

var CertProductType = struct {
	HAServer             _CertProductType
	HAGlobalServer       _CertProductType
	Server               _CertProductType
	GlobalServer         _CertProductType
	IntranetServer       _CertProductType
	IntranetGlobalServer _CertProductType
	PrivateServer        _CertProductType
	GeotrustServer       _CertProductType
	CodeSigning          _CertProductType
	JavaCodeSigning      _CertProductType
	EVCodeSigning        _CertProductType
	OFXServer            _CertProductType
}{
	"HAServer",
	"HAGlobalServer",
	"Server",
	"GlobalServer",
	"IntranetServer",
	"IntranetGlobalServer",
	"PrivateServer",
	"GeotrustServer",
	"CodeSigning",
	"JavaCodeSigning",
	"EVCodeSigning",
	"OFXServer",
}

var CertRevokementType = struct {
	KeyCompromise        _CertRevokementType
	CACompromise         _CertRevokementType
	AffiliationChanged   _CertRevokementType
	Superseded           _CertRevokementType
	CessationOfOperation _CertRevokementType
	CertificateHold      _CertRevokementType
	RemoveFromCRL        _CertRevokementType
	PrivilegeWithdrawn   _CertRevokementType
	AACompromise         _CertRevokementType
	Unspecified          _CertRevokementType
}{
	KeyCompromise:        "Key compromise",
	CACompromise:         "CA compromise",
	AffiliationChanged:   "Affiliation changed",
	Superseded:           "Superseded",
	CessationOfOperation: "Cessation of operation",
	CertificateHold:      "Certificate hold",
	RemoveFromCRL:        "Remove from CRL",
	PrivilegeWithdrawn:   "Privilege withdrawn",
	AACompromise:         "AA compromise",
	Unspecified:          "Unspecified",
}
