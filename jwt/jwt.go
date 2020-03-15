package jwt

import (
	"context"
	"crypto/sha256"
	b64 "encoding/base64"
	"fmt"

	cloudkms "cloud.google.com/go/kms/apiv1"
	kmspb "google.golang.org/genproto/googleapis/cloud/kms/v1"
)

var header = []byte(`{"alg": "HS256", "typ": "JWT"}`)

func Generate(payload []byte) (string, error) {
	headerB64 := b64.URLEncoding.EncodeToString(header)
	payloadB64 := b64.URLEncoding.EncodeToString(payload)

	message := headerB64 + "." + payloadB64
	signature, err := signAsymmetric([]byte(message))
	sigB64 := b64.URLEncoding.EncodeToString(signature)

	token := message + "." + sigB64
	return token, err
}

// signAsymmetric will sign a plaintext message using a saved asymmetric private key.
func signAsymmetric(message []byte) ([]byte, error) {
	name := "projects/siyan-io/locations/global/keyRings/siyan-io/cryptoKeys/jwt/cryptoKeyVersions/1"

	ctx := context.Background()
	client, err := cloudkms.NewKeyManagementClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("cloudkms.NewKeyManagementClient: %v", err)
	}

	digest := sha256.New()
	digest.Write(message)

	req := &kmspb.AsymmetricSignRequest{
		Name: name,
		Digest: &kmspb.Digest{
			Digest: &kmspb.Digest_Sha256{
				Sha256: digest.Sum(nil),
			},
		},
	}

	response, err := client.AsymmetricSign(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("AsymmetricSign: %v", err)
	}

	return response.Signature, nil
}

func parse(tkn string) string {
	return tkn
}

func IsValid(tkn string) string {
	return parse(tkn)
}
