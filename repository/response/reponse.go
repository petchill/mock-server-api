package response

import mResponse "mock-server-api/model/response"

func (r *responseRepository) GetResponse() (mResponse.ResponseDB, error) { // TODO: change to custom error later

	return mResponse.ResponseDB{}, nil
}
