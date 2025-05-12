package cloud

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"os"

	"github.com/dae-vercel-function/model"

	"cloud.google.com/go/firestore"
)

// CloudStore defines the interface for Firestore operations.
type CloudStore interface {
	// GetCollections returns all collections in the Firestore database.
	GetCollections(ctx context.Context, path string) ([]*firestore.CollectionRef, error)

	CreateSheet(ctx context.Context, sheet model.Sheet) error
}

type FireStore struct {
	credentials []byte
	// Export FireStore client
	client *firestore.Client
}

func (fs *FireStore) loadCredentials() {
	googleCredentials := os.Getenv("GOOGLE_CREDENTIALS")
	if googleCredentials == "" {
		LogError("GOOGLE_CREDENTIALS environment variable is not set")
		return
	}

	decodedCredentials, err := base64.RawStdEncoding.DecodeString(googleCredentials)
	if err != nil {
		LogError("Loading credentials on err: %v", err)
		return
	}
	fs.credentials = decodedCredentials
}

func (fs *FireStore) initClient(ctx context.Context, projectID string) {
	var err error
	if len(fs.credentials) == 0 {
		LogError("Invalid firestore credentials")
		return
	}
	if fs.client, err = firestore.NewClient(ctx, projectID); err != nil {
		LogError("Init firestore client on err: %v", err)
	}
}

func NewFireStore(ctx context.Context, projectID string) *FireStore {
	inst := &FireStore{}
	inst.loadCredentials()
	inst.initClient(ctx, projectID)
	return inst
}

func (fs *FireStore) Close() {
	if fs != nil && fs.client != nil {
		fs.client.Close()
	}
}

func LogError(format string, v ...interface{}) {
	err := fmt.Errorf(format, v...)
	log.Printf("[Error] %s", err)
}
