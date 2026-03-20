package profile

import "github.com/mlinder10/wcj/types"

func BuildResponse[T types.ProfileMapable](
	profiles []T, page, totalProfiles int,
) (types.ProfileResponse, types.APIError) {
	parsedProfiles := make([]types.Profile, len(profiles))

	if len(profiles) == 0 {
		return types.ProfileResponse{
			Users:       []types.Profile{},
			CurrentPage: page + 1,
			TotalPages:  page + 1,
		}, types.Errors.None
	}

	for i, profile := range profiles {
		parsedProfiles[i] = profile.ToProfile()
	}

	return types.ProfileResponse{
		Users:       parsedProfiles,
		CurrentPage: page + 1,
		TotalPages:  totalProfiles/userLimit + 1,
	}, types.Errors.None
}
