package steamapi

import (
	"github.com/246859/steamapi/types/publishedfile"
)

const (
	QueryFilesURL = "/IPublishedFileService/QueryFiles/v1/"

	SetDeveloperMetadataURL = "/IPublishedFileService/SetDeveloperMetadata/v1/"

	UpdateAppUGCBanURL = "/IPublishedFileService/UpdateAppUGCBan/v1/"

	UpdateBanStatusURL = "/IPublishedFileService/UpdateBanStatus/v1/"

	UpdateIncompatibleStatusURL = "/IPublishedFileService/UpdateIncompatibleStatus/v1/"

	UpdateTagsURL = "/IPublishedFileService/UpdateTags/v1/"
)

func (c *Client) PublishedFiles() *ISteamPublishedFilesService {
	return &ISteamPublishedFilesService{c: c}
}

type ISteamPublishedFilesService struct {
	c *Client
}

// QueryFiles see https://partner.steamgames.com/doc/webapi/IPublishedFileService#QueryFiles
func (i ISteamPublishedFilesService) QueryFiles(queryFileRequest publishedfile.QueryFileRequest, options ...RequestOptions) (publishedfile.QueryFileDetails, error) {
	var response publishedfile.QueryFileDetails
	queryform, err := structToMap(queryFileRequest)
	if err != nil {
		return response, err
	}
	options = joinRequestOptions(options, WithRequestQuery(queryform))
	request := i.c.Get(PublicHost, QueryFilesURL, options...)
	request.SetResult(&response)
	if _, err = request.Send(); err != nil {
		return response, err
	}
	return response, nil
}
