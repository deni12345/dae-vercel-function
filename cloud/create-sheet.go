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
	docRef, _, err := f.client.Collection("sheet").Add(ctx, req)
	if err != nil {
		return nil, err
	}

	var res *model.Sheet
	if err = f.documentToModel(ctx, docRef, res); err != nil {
		return nil, err
	}
	return res, nil
}

func (f *FireStore) documentToModel(ctx context.Context, docRef *firestore.DocumentRef, res model.ICommon) error {
	docSnap, err := docRef.Get(ctx)
	if err != nil {
		return err
	}

	if err := docSnap.DataTo(&res); err != nil {
		return err
	}

	res.SetID(docRef.ID)
	return nil
}
