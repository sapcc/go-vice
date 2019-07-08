package vice

import (
	"context"
)

const revokeBasePath = "/vswebservices/rest/services/revoke"

type _CertRevokementType string

type RevokeRequest struct {
	CertSerial string              `url:"certSerial"`
	Reason     _CertRevokementType `url:"reason"`
	Challenge  string              `url:"challenge,omitempty"`
}

type Revokement struct {
	ViceResponse ViceResponse
}

func (c *CertificatesServiceOp) Revoke(ctx context.Context, er *RevokeRequest) (*Revokement, error) {
	req, err := c.client.newRequest(ctx, "POST", revokeBasePath, er)
	if err != nil {
		return nil, err
	}

	revokement := new(Revokement)
	err = c.client.Do(req, revokement)
	if err != nil {
		return nil, err
	}

	return revokement, nil
}
