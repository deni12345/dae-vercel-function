package cloud

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"os"

	"github.com/dae-vercel-function/model"
	"google.golang.org/api/option"

	"cloud.google.com/go/firestore"
	"github.com/joho/godotenv"
)

// CloudStore defines the interface for Firestore operations.
type CloudStore interface {
	// GetCollections returns all collections in the Firestore database.
	GetCollections(ctx context.Context, path string) ([]*firestore.CollectionRef, error)
	// CreateSheet create an item in sheet collection if false it return error
	CreateSheet(ctx context.Context, sheet model.Sheet) error

	ObservceCollection(ctx context.Context) ([]*model.DocumentChange, error)
}

type FireStore struct {
	credentials []byte
	// Export FireStore client
	client *firestore.Client
}

func (fs *FireStore) loadCredentials() {
	if os.Getenv("ENVIRONMENT") == "" {
		if err := godotenv.Load(); err != nil {
			LogError("Error loading .env file: %v", err)
		}
	}

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
	if fs.client, err = firestore.NewClient(ctx, projectID, option.WithCredentialsJSON(fs.credentials)); err != nil {
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
