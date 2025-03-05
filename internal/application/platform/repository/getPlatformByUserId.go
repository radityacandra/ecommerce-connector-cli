package repository

import "context"

func (r *Repository) GetPlatformByUserId(ctx context.Context, userId string) ([]FindByUserIdAndTypeOutput, error) {
	docs := r.Db.Client.Collection("platforms").Where("user_id", "==", userId).Documents(ctx)

	defer docs.Stop()

	raw, err := docs.GetAll()
	if err != nil {
		return nil, err
	}

	var output []FindByUserIdAndTypeOutput
	for _, doc := range raw {
		var data FindByUserIdAndTypeOutput
		doc.DataTo(&data)

		output = append(output, data)
	}

	return output, nil
}
