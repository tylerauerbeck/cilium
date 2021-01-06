package cabf_br

/*
 * ZLint Copyright 2020 Regents of the University of Michigan
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not
 * use this file except in compliance with the License. You may obtain a copy
 * of the License at http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
 * implied. See the License for the specific language governing
 * permissions and limitations under the License.
 */

// If the Certificate asserts the policy identifier of 2.23.140.1.2.1, then it MUST NOT include
// organizationName, streetAddress, localityName, stateOrProvinceName, or postalCode in the Subject field.

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v2/lint"
	"github.com/zmap/zlint/v2/util"
)

type certPolicyConflictsWithStreet struct{}

func (l *certPolicyConflictsWithStreet) Initialize() error {
	return nil
}

func (l *certPolicyConflictsWithStreet) CheckApplies(cert *x509.Certificate) bool {
	return util.SliceContainsOID(cert.PolicyIdentifiers, util.BRDomainValidatedOID) && !util.IsCACert(cert)
}

func (l *certPolicyConflictsWithStreet) Execute(cert *x509.Certificate) *lint.LintResult {
	var out lint.LintResult
	if util.TypeInName(&cert.Subject, util.StreetAddressOID) {
		out.Status = lint.Error
	} else {
		out.Status = lint.Pass
	}
	return &out
}

func init() {
	lint.RegisterLint(&lint.Lint{
		Name:          "e_cab_dv_conflicts_with_street",
		Description:   "If certificate policy 2.23.140.1.2.1 (CA/B BR domain validated) is included, streetAddress MUST NOT be included in subject",
		Citation:      "BRs: 7.1.6.1",
		Source:        lint.CABFBaselineRequirements,
		EffectiveDate: util.CABEffectiveDate,
		Lint:          &certPolicyConflictsWithStreet{},
	})
}
