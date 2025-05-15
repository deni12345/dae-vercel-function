package cloud

import (
	"context"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

// GetCollections retrieves all top-level collections from Firestore.
func (f *FireStore) GetCollections(ctx context.Context) ([]*firestore.CollectionRef, error) {
	var (
		collections     = []*firestore.CollectionRef{}
		collectionsIter = f.client.Collections(ctx)
	)

	for {
		collectionRef, err := collectionsIter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		collections = append(collections, collectionRef)
	}
	return collections, nil
}
