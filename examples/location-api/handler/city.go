package handler

import (
	"context"
	"errors"
	"net/http"

	crud_city_usecase "github.com/herryg91/cdd/examples/location-api/app/usecase/crud_city"
	search_usecase "github.com/herryg91/cdd/examples/location-api/app/usecase/search"
	pbCity "github.com/herryg91/cdd/examples/location-api/handler/grst/city"
	grst_errors "github.com/herryg91/cdd/grst/errors"
	"google.golang.org/grpc/codes"
)

type CityHandler struct {
	pbCity.UnimplementedCityServer
	crud_city_uc crud_city_usecase.UseCase
	search_uc    search_usecase.UseCase
}

func NewCityHandler(crud_city_uc crud_city_usecase.UseCase, search_uc search_usecase.UseCase) pbCity.CityServer {
	return &CityHandler{crud_city_uc: crud_city_uc, search_uc: search_uc}
}

func (h *CityHandler) Get(ctx context.Context, req *pbCity.GetReq) (*pbCity.City, error) {
	if err := pbCity.ValidateRequest(req); err != nil {
		return nil, err
	}
	data, err := h.crud_city_uc.GetByPrimaryKey(int(req.Id))
	if err != nil {
		if errors.Is(err, crud_city_usecase.ErrRecordNotFound) {
			return nil, grst_errors.New(http.StatusNotFound, codes.NotFound, 2101, err.Error(), &grst_errors.ErrorDetail{})
		}
		return nil, grst_errors.New(http.StatusInternalServerError, codes.Internal, 2102, err.Error(), &grst_errors.ErrorDetail{})
	}

	return &pbCity.City{
		Id:         int32(data.Id),
		ProvinceId: int32(data.ProvinceId),
		Name:       data.Name,
	}, nil
}
func (h *CityHandler) Search(ctx context.Context, req *pbCity.SearchReq) (*pbCity.CityProfiles, error) {
	if err := pbCity.ValidateRequest(req); err != nil {
		return nil, err
	}

	searchResult, err := h.search_uc.Search(req.Keyword)
	if err != nil {
		return nil, grst_errors.New(http.StatusInternalServerError, codes.Internal, 2202, err.Error(), &grst_errors.ErrorDetail{})
	}
	resp := &pbCity.CityProfiles{
		Cities: []*pbCity.CityProfile{},
	}
	for _, data := range searchResult {
		resp.Cities = append(resp.Cities, &pbCity.CityProfile{
			Id:           int32(data.Id),
			Name:         data.Name,
			ProvinceId:   int32(data.ProvinceId),
			ProvinceName: data.ProvinceName,
		})
	}
	return resp, nil
}
