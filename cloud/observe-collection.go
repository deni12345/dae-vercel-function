package cloud

import (
	"context"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/dae-vercel-function/model"
	"google.golang.org/api/iterator"
)

func (f *FireStore) ObservceCollection(ctx context.Context, sheetID string) ([]*model.DocumentChange, error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	sheetColRef := f.client.Collection("sheet")
	iter := sheetColRef.Snapshots(ctx)
	defer func() {
		cancel()
		iter.Stop()
	}()

	var result []*model.DocumentChange
	for {
		snap, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return result, err
		}

		for _, change := range snap.Changes {
			sheet := model.Sheet{}
			switch change.Kind {
			case firestore.DocumentAdded, firestore.DocumentModified:
				if err := f.documentToModel(ctx, change.Doc.Ref, &sheet); err != nil {
					result = append(result, &model.DocumentChange{Sheet: sheet, Action: model.ActionDict[int(change.Kind)]})
					LogError("Fail to observe on sheet %v", change.Doc.Ref.ID)
				}
			case firestore.DocumentRemoved:
				sheet.ID = change.Doc.Ref.ID
				result = append(result, &model.DocumentChange{Sheet: sheet, Action: model.ActionDict[int(change.Kind)]})
			}
		}

	}
	return result, nil
}
