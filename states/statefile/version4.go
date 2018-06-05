package statefile

import (
	"encoding/json"
	"fmt"

	version "github.com/hashicorp/go-version"
	"github.com/hashicorp/terraform/tfdiags"
)

func readStateV4(src []byte) (*File, tfdiags.Diagnostics) {
	var diags tfdiags.Diagnostics
	sV4 := &stateV4{}
	err := json.Unmarshal(src, sV4)
	if err != nil {
		diags = diags.Append(jsonUnmarshalDiags(err))
		return nil, diags
	}

	file, prepDiags := prepareStateV4(sV4)
	diags = diags.Append(prepDiags)
	return file, diags
}

func prepareStateV4(sV4 *stateV4) (*File, tfdiags.Diagnostics) {
	var diags tfdiags.Diagnostics

	var tfVersion *version.Version
	if sV4.TerraformVersion != "" {
		var err error
		tfVersion, err = version.NewVersion(sV4.TerraformVersion)
		if err != nil {
			diags = diags.Append(tfdiags.Sourceless(
				tfdiags.Error,
				"Invalid Terraform version string",
				fmt.Sprintf("State file claims to have been written by Terraform version %q, which is not a valid version string.", sV4.TerraformVersion),
			))
		}
	}

	file := &File{
		TerraformVersion: tfVersion,
		Serial:           sV4.Serial,
		Lineage:          sV4.Lineage,
	}

	// TODO: Populate the State field too

	return file, diags
}

type stateV4 struct {
	TerraformVersion string                    `json:"terraform_version"`
	Serial           uint64                    `json:"serial"`
	Lineage          string                    `json:"lineage"`
	RootOutputs      map[string]*outputStateV4 `json:"outputs"`
	Resources        []*instanceObjectStateV4  `json:"resources"`
}

type outputStateV4 struct {
	ValueRaw     json.RawMessage `json:"value"`
	ValueTypeRaw json.RawMessage `json:"type"`

	Sensitive bool `json:"sensitive"`
}

type resourceStateV4 struct {
	Module         string `json:"module,omitempty"`
	Mode           string `json:"mode"`
	Type           string `json:"type"`
	Name           string `json:"name"`
	EachMode       string `json:"each,omitempty"`
	ProviderConfig string `json:"provider"`
	Instances      string `json:"instances"`
}

type instanceObjectStateV4 struct {
	IndexKey interface{} `json:"index_key,omitempty"`
	Status   string      `json:"status,omitempty"`
	Deposed  bool        `json:"deposed,omitempty"`

	SchemaVersion  uint64            `json:"schema_version"`
	AttributesRaw  json.RawMessage   `json:"attributes"`
	AttributesFlat map[string]string `json:"attributes_flat"`

	PrivateRaw json.RawMessage `json:"private"`
}
