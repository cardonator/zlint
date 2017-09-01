// lint_cert_contains_unique_identifier.go
/************************************************
 These fields MUST only appear if the version is 2 or 3 (Section 4.1.2.1).
 These fields MUST NOT appear if the version is 1. The subject and issuer
 unique identifiers are present in the certificate to handle the possibility
 of reuse of subject and/or issuer names over time. This profile RECOMMENDS
 that names not be reused for different entities and that Internet certificates
 not make use of unique identifiers. CAs conforming to this profile MUST NOT
 generate certificates with unique identifiers. Applications conforming to
 this profile SHOULD be capable of parsing certificates that include unique
 identifiers, but there are no processing requirements associated with the
 unique identifiers.
************************************************/

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type CertContainsUniqueIdentifier struct{}

func (l *CertContainsUniqueIdentifier) Initialize() error {
	return nil
}

func (l *CertContainsUniqueIdentifier) CheckApplies(cert *x509.Certificate) bool {
	return true
}

func (l *CertContainsUniqueIdentifier) Execute(cert *x509.Certificate) *LintResult {
	if cert.IssuerUniqueId.Bytes == nil && cert.SubjectUniqueId.Bytes == nil {
		return &LintResult{Status: Pass}
	} //else
	return &LintResult{Status: Error}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_cert_contains_unique_identifier",
		Description:   "CAs MUST NOT generate certificate with unique identifiers",
		Source:        "RFC 5280: 4.1.2.8",
		EffectiveDate: util.RFC5280Date,
		Lint:          &CertContainsUniqueIdentifier{},
	})
}
