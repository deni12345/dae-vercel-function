package cloud

import (
	"context"

	"github.com/dae-vercel-function/model"

	"cloud.google.com/go/firestore"
)

const (
	SheetCollection = "sheet"
)

func (f *FireStore) CreateSheet(ctx context.Context, req model.Sheet) (*model.Sheet, error) {
	docRef, _, err := f.client.Collection(SheetCollection).Add(ctx, req)
	if err != nil {
		return nil, err
	}

	snapshot, err := docRef.Get(ctx)
	if err != nil {
		return nil, err
	}

	res := &model.Sheet{}
	if err := f.documentToModel(ctx, snapshot, res); err != nil {
		return nil, err
	}
	return res, nil
}

func (f *FireStore) documentToModel(ctx context.Context, docSnapshot *firestore.DocumentSnapshot, res model.ICommon) error {
	if err := docSnapshot.DataTo(&res); err != nil {
		return err
	}

	res.SetID(docSnapshot.Ref.ID)
	return nil
}
