// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pki

import (
	"context"
	"fmt"

	"github.com/openbao/openbao/sdk/logical"
)

func (b *backend) doTrackBilling(ctx context.Context, identifiers []*ACMEIdentifier) error {
	billingView, ok := b.System().(logical.ACMEBillingSystemView)
	if !ok {
		return fmt.Errorf("failed to perform cast to ACME billing system view interface")
	}

	var realized []string
	for _, identifier := range identifiers {
		realized = append(realized, fmt.Sprintf("%s/%s", identifier.Type, identifier.OriginalValue))
	}

	return billingView.CreateActivityCountEventForIdentifiers(ctx, realized)
}
