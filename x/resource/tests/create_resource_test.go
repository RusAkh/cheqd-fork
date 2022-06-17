package tests

import (
	"crypto/ed25519"
	"crypto/sha256"
	// "crypto/sha256"
	"fmt"
	"testing"

	"github.com/cheqd/cheqd-node/x/resource/types"

	"github.com/stretchr/testify/require"
)

func TestCreateResource(t *testing.T) {
	keys := GenerateTestKeys()
	cases := []struct {
		valid             bool
		name              string
		signerKeys        map[string]ed25519.PrivateKey
		msg               *types.MsgCreateResourcePayload
		previousVersionId string
		errMsg            string
	}{
		{
			valid: true,
			name:  "Valid: Works",
			signerKeys: map[string]ed25519.PrivateKey{
				ExistingDIDKey: keys[ExistingDIDKey].PrivateKey,
			},
			msg: &types.MsgCreateResourcePayload{
				CollectionId: ExistingDIDIdentifier,
				Id:           ResourceId,
				Name:         "Test Resource Name",
				ResourceType: CLSchemaType,
				MimeType:     JsonResourceType,
				Data:         []byte(SchemaData),
			},
			previousVersionId: "",
			errMsg: "",
		},
		{
			valid: true,
			name:  "Valid: Add new resource version",
			signerKeys: map[string]ed25519.PrivateKey{
				ExistingDIDKey: keys[ExistingDIDKey].PrivateKey,
			},
			msg: &types.MsgCreateResourcePayload{
				CollectionId: ExistingResource().CollectionId,
				Id:           ResourceId,
				Name:         ExistingResource().Name,
				ResourceType: ExistingResource().ResourceType,
				MimeType:     ExistingResource().MimeType,
				Data:         ExistingResource().Data,
			},
			previousVersionId: ExistingResource().Id,
			errMsg: "",
		},
		{
			valid:      false,
			name:       "Not Valid: No signature",
			signerKeys: map[string]ed25519.PrivateKey{},
			msg: &types.MsgCreateResourcePayload{
				CollectionId: ExistingDIDIdentifier,
				Id:           ResourceId,
				Name:         "Test Resource Name",
				ResourceType: CLSchemaType,
				MimeType:     JsonResourceType,
				Data:         []byte(SchemaData),
			},
			errMsg: fmt.Sprintf("signer: %s: signature is required but not found", ExistingDID),
		},
		{
			valid:      false,
			name:       "Not Valid: Invalid Resource Id",
			signerKeys: map[string]ed25519.PrivateKey{},
			msg: &types.MsgCreateResourcePayload{
				CollectionId: ExistingDIDIdentifier,
				Id:           IncorrectResourceId,
				Name:         "Test Resource Name",
				ResourceType: CLSchemaType,
				MimeType:     JsonResourceType,
				Data:         []byte(SchemaData),
			},
			errMsg: fmt.Sprintf("signer: %s: signature is required but not found", ExistingDID),
		},
		{
			valid:      false,
			name:       "Not Valid: DID Doc is not found",
			signerKeys: map[string]ed25519.PrivateKey{},
			msg: &types.MsgCreateResourcePayload{
				CollectionId: NotFoundDIDIdentifier,
				Id:           IncorrectResourceId,
				Name:         "Test Resource Name",
				ResourceType: CLSchemaType,
				MimeType:     JsonResourceType,
				Data:         []byte(SchemaData),
			},
			errMsg: fmt.Sprintf("signer: %s: signature is required but not found", ExistingDID),
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			msg := tc.msg
			resourceSetup := InitEnv(t, keys[ExistingDIDKey].PublicKey, keys[ExistingDIDKey].PrivateKey)

			resource, err := resourceSetup.SendCreateResource(msg, tc.signerKeys)
			if tc.valid {
				require.Nil(t, err)

				didStateValue, err := resourceSetup.Keeper.GetDid(&resourceSetup.Ctx, resource.CollectionId)
				require.Nil(t, err)
				require.Contains(t, didStateValue.Metadata.Resources, resource.Id)

				require.Equal(t, tc.msg.CollectionId, resource.CollectionId)
				require.Equal(t, tc.msg.Id, resource.Id)
				require.Equal(t, tc.msg.MimeType, resource.MimeType)
				require.Equal(t, tc.msg.ResourceType, resource.ResourceType)
				require.Equal(t, tc.msg.Data, resource.Data)
				require.Equal(t, tc.msg.Name, resource.Name)
				require.Equal(t, string(sha256.New().Sum(resource.Data)), resource.Checksum)
				require.Equal(t, tc.previousVersionId, resource.PreviousVersionId)
				// if tc.previousVersionId != "" {
				// 	require.Equal(t, resource.Id, resource.NextVersionId)
				// }
			} else {
				require.Error(t, err)
				require.Equal(t, tc.errMsg, err.Error())
			}
		})
	}
}

// func TestHandler_ResourceDocAlreadyExists(t *testing.T) {
// 	setup := Setup()

// 	_, _, _ = setup.InitDid(AliceDID)
// 	_, _, err := setup.InitDid(AliceDID)

// 	require.Error(t, err)
// 	require.Equal(t, fmt.Sprintf("%s: DID Doc exists", AliceDID), err.Error())
// }
