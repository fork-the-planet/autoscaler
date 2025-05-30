// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// Use the Core Services API to manage resources such as virtual cloud networks (VCNs),
// compute instances, and block storage volumes. For more information, see the console
// documentation for the Networking (https://docs.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.oracle.com/iaas/Content/Block/Concepts/overview.htm) services.
// The required permissions are documented in the
// Details for the Core Services (https://docs.oracle.com/iaas/Content/Identity/Reference/corepolicyreference.htm) article.
//

package core

import (
	"fmt"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/oci/vendor-internal/github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AppCatalogListingResourceVersionAgreements Agreements for a listing resource version.
type AppCatalogListingResourceVersionAgreements struct {

	// The OCID of the listing associated with these agreements.
	ListingId *string `mandatory:"false" json:"listingId"`

	// Listing resource version associated with these agreements.
	ListingResourceVersion *string `mandatory:"false" json:"listingResourceVersion"`

	// Oracle TOU link
	OracleTermsOfUseLink *string `mandatory:"false" json:"oracleTermsOfUseLink"`

	// EULA link
	EulaLink *string `mandatory:"false" json:"eulaLink"`

	// Date and time the agreements were retrieved, in RFC3339 (https://tools.ietf.org/html/rfc3339) format.
	// Example: `2018-03-20T12:32:53.532Z`
	TimeRetrieved *common.SDKTime `mandatory:"false" json:"timeRetrieved"`

	// A generated signature for this agreement retrieval operation which should be used in the create subscription call.
	Signature *string `mandatory:"false" json:"signature"`
}

func (m AppCatalogListingResourceVersionAgreements) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AppCatalogListingResourceVersionAgreements) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
