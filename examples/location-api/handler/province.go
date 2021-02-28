package handler

import (
	"context"
	"errors"
	"net/http"

	"github.com/golang/protobuf/ptypes/empty"
	pbProvince "github.com/herryg91/cdd/examples/location-api/grpc/pb/province"
	crud_tbl_province "github.com/herryg91/cdd/examples/location-api/usecase/crud_tbl_province"
	grst_errors "github.com/herryg91/cdd/grst/errors"
	"google.golang.org/grpc/codes"
)

type ProvinceHandler struct {
	pbProvince.UnimplementedProvinceServer
	crudProvinceUsecase crud_tbl_province.UseCase
}

func NewProvinceHandler(crudProvinceUsecase crud_tbl_province.UseCase) pbProvince.ProvinceServer {
	return &ProvinceHandler{pbProvince.UnimplementedProvinceServer{}, crudProvinceUsecase}
}

func (h *ProvinceHandler) Get(ctx context.Context, req *pbProvince.GetReq) (*pbProvince.Province, error) {
	if err := pbProvince.ValidateRequest(req); err != nil {
		return nil, err
	}

	data, err := h.crudProvinceUsecase.GetByPrimaryKey(int(req.Id))
	if err != nil {
		if errors.Is(err, crud_tbl_province.ErrRecordNotFound) {
			return nil, grst_errors.New(http.StatusNotFound, codes.NotFound, 1101, err.Error(), &grst_errors.ErrorDetail{})
		}
		return nil, grst_errors.New(http.StatusInternalServerError, codes.Internal, 1102, err.Error(), &grst_errors.ErrorDetail{})
	}

	return &pbProvince.Province{
		Id:   int32(data.Id),
		Name: data.Name,
	}, nil
}

func (h *ProvinceHandler) GetAll(ctx context.Context, req *empty.Empty) (*pbProvince.Provinces, error) {
	datas, err := h.crudProvinceUsecase.GetAll()
	if err != nil {
		return nil, grst_errors.New(http.StatusInternalServerError, codes.Internal, 1201, err.Error(), &grst_errors.ErrorDetail{})
	}
	result := &pbProvince.Provinces{
		Provinces: []*pbProvince.Province{},
	}
	for _, data := range datas {
		result.Provinces = append(result.Provinces, &pbProvince.Province{
			Id:   int32(data.Id),
			Name: data.Name,
		})
	}
	return result, nil
}
