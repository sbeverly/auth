package jwt

import (
	"context"
	"crypto/ecdsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/asn1"
	b64 "encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"log"
	"math/big"
	"strings"
	"time"

	cloudkms "cloud.google.com/go/kms/apiv1"
	kmspb "google.golang.org/genproto/googleapis/cloud/kms/v1"
)

var header = []byte(`{"alg": "HS256", "typ": "JWT"}`)
var keyName = "projects/siyan-io/locations/global/keyRings/siyan-io/cryptoKeys/jwt/cryptoKeyVersions/1"

var secretKey *ecdsa.PublicKey

func init() {
	secretKey = getSecretKey(keyName)
}

type Claims struct {
	UserID int `json:"userId"`
	Iat time.Time `json:"iat"`
}

func Generate(claims *Claims) (string, error) {
	jsonClaims, _ := json.Marshal(claims)

	headerB64 := b64Encode(header)
	claimsB64 := b64Encode(jsonClaims)

	message := headerB64 + "." + claimsB64
	signature, err := signAsymmetric([]byte(message))
	sigB64 := b64Encode(signature)

	token := message + "." + sigB64
	return token, err
}

// Verify : Verify token signature
func Verify(token string) error {
	header, payload, signature, err := parseToken(token)

	if err != nil {
		return err
	}

	message := []byte(string(header) + "." + string(payload))

	return verifySignatureEC(signature, message)
}

// Claims : Extract Claims from token
func GetClaims(token string) (*Claims, error) {
	_, claimsb64, _, err := parseToken(token)

	if err != nil {
		return nil, err
	}

	claimsJSON := b64Decode(string(claimsb64))

	var claims Claims
	json.Unmarshal(claimsJSON, &claims)

	return &claims, nil
}

func parseToken(token string) ([]byte, []byte, []byte, error) {
	parts := strings.Split(token, ".")

	if len(parts) != 3 {
		return nil, nil, nil, errors.New("Malformed JWT: Does not contain 3 parts.")
	}

	signature := b64Decode(parts[2])
	header := []byte(parts[0])
	claims := []byte(parts[1])

	return header, claims, signature, nil
}

func b64Encode(data []byte) string {
	encoded := b64.URLEncoding.EncodeToString(data)
	return encoded
}

func b64Decode(str string) []byte {
	decoded, _ := b64.URLEncoding.DecodeString(str)
	return decoded
}

// signAsymmetric will sign a plaintext message using a saved asymmetric private key.
func signAsymmetric(message []byte) ([]byte, error) {
	ctx := context.Background()
	client, err := cloudkms.NewKeyManagementClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("cloudkms.NewKeyManagementClient: %v", err)
	}

	digest := sha256.New()
	digest.Write(message)

	req := &kmspb.AsymmetricSignRequest{
		Name: keyName,
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

func verifySignatureEC(signature []byte, message []byte) error {
	var parsedSig struct{ R, S *big.Int }
	if _, err := asn1.Unmarshal(signature, &parsedSig); err != nil {
		return fmt.Errorf("asn1.Unmarshal: %v", err)
	}

	hash := sha256.New()
	hash.Write(message)
	digest := hash.Sum(nil)
	if !ecdsa.Verify(secretKey, digest, parsedSig.R, parsedSig.S) {
		return fmt.Errorf("ecdsa.Verify failed on key: %s", keyName)
	}

	return nil
}

func getSecretKey(string) *ecdsa.PublicKey {
	ctx := context.Background()
	client, err := cloudkms.NewKeyManagementClient(ctx)
	if err != nil {
		log.Fatal("cloudkms.NewKeyManagementClient: %v", err)
	}

	response, err := client.GetPublicKey(ctx, &kmspb.GetPublicKeyRequest{Name: keyName})
	if err != nil {
		log.Fatal("GetPublicKey: %v", err)
	}

	block, _ := pem.Decode([]byte(response.Pem))
	abstractKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		log.Fatal("x509.ParsePKIXPublicKey: %v", err)
	}

	ecKey, ok := abstractKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("key '%s' is not EC", keyName)
	}

	return ecKey

}
