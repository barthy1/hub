// Code generated by goa v3.3.1, DO NOT EDIT.
//
// resource HTTP server encoders and decoders
//
// Command:
// $ goa gen github.com/tektoncd/hub/api/v1/design

package server

import (
	"context"
	"net/http"
	"strconv"

	resourceviews "github.com/tektoncd/hub/api/v1/gen/resource/views"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// EncodeQueryResponse returns an encoder for responses returned by the
// resource Query endpoint.
func EncodeQueryResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*resourceviews.Resources)
		enc := encoder(ctx, w)
		body := NewQueryResponseBody(res.Projected)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeQueryRequest returns a decoder for requests sent to the resource Query
// endpoint.
func DecodeQueryRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			name       string
			catalogs   []string
			categories []string
			kinds      []string
			tags       []string
			limit      uint
			match      string
			err        error
		)
		nameRaw := r.URL.Query().Get("name")
		if nameRaw != "" {
			name = nameRaw
		}
		catalogs = r.URL.Query()["catalogs"]
		categories = r.URL.Query()["categories"]
		kinds = r.URL.Query()["kinds"]
		tags = r.URL.Query()["tags"]
		{
			limitRaw := r.URL.Query().Get("limit")
			if limitRaw == "" {
				limit = 1000
			} else {
				v, err2 := strconv.ParseUint(limitRaw, 10, strconv.IntSize)
				if err2 != nil {
					err = goa.MergeErrors(err, goa.InvalidFieldTypeError("limit", limitRaw, "unsigned integer"))
				}
				limit = uint(v)
			}
		}
		matchRaw := r.URL.Query().Get("match")
		if matchRaw != "" {
			match = matchRaw
		} else {
			match = "contains"
		}
		if !(match == "exact" || match == "contains") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("match", match, []interface{}{"exact", "contains"}))
		}
		if err != nil {
			return nil, err
		}
		payload := NewQueryPayload(name, catalogs, categories, kinds, tags, limit, match)

		return payload, nil
	}
}

// EncodeQueryError returns an encoder for errors returned by the Query
// resource endpoint.
func EncodeQueryError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "internal-error":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewQueryInternalErrorResponseBody(res)
			}
			w.Header().Set("goa-error", "internal-error")
			w.WriteHeader(http.StatusInternalServerError)
			return enc.Encode(body)
		case "invalid-kind":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewQueryInvalidKindResponseBody(res)
			}
			w.Header().Set("goa-error", "invalid-kind")
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		case "not-found":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewQueryNotFoundResponseBody(res)
			}
			w.Header().Set("goa-error", "not-found")
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeListResponse returns an encoder for responses returned by the resource
// List endpoint.
func EncodeListResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*resourceviews.Resources)
		enc := encoder(ctx, w)
		body := NewListResponseBody(res.Projected)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeListRequest returns a decoder for requests sent to the resource List
// endpoint.
func DecodeListRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			limit uint
			err   error
		)
		{
			limitRaw := r.URL.Query().Get("limit")
			if limitRaw == "" {
				limit = 1000
			} else {
				v, err2 := strconv.ParseUint(limitRaw, 10, strconv.IntSize)
				if err2 != nil {
					err = goa.MergeErrors(err, goa.InvalidFieldTypeError("limit", limitRaw, "unsigned integer"))
				}
				limit = uint(v)
			}
		}
		if err != nil {
			return nil, err
		}
		payload := NewListPayload(limit)

		return payload, nil
	}
}

// EncodeListError returns an encoder for errors returned by the List resource
// endpoint.
func EncodeListError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "internal-error":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewListInternalErrorResponseBody(res)
			}
			w.Header().Set("goa-error", "internal-error")
			w.WriteHeader(http.StatusInternalServerError)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeVersionsByIDResponse returns an encoder for responses returned by the
// resource VersionsByID endpoint.
func EncodeVersionsByIDResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*resourceviews.ResourceVersions)
		enc := encoder(ctx, w)
		body := NewVersionsByIDResponseBody(res.Projected)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeVersionsByIDRequest returns a decoder for requests sent to the
// resource VersionsByID endpoint.
func DecodeVersionsByIDRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			id  uint
			err error

			params = mux.Vars(r)
		)
		{
			idRaw := params["id"]
			v, err2 := strconv.ParseUint(idRaw, 10, strconv.IntSize)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("id", idRaw, "unsigned integer"))
			}
			id = uint(v)
		}
		if err != nil {
			return nil, err
		}
		payload := NewVersionsByIDPayload(id)

		return payload, nil
	}
}

// EncodeVersionsByIDError returns an encoder for errors returned by the
// VersionsByID resource endpoint.
func EncodeVersionsByIDError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "internal-error":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewVersionsByIDInternalErrorResponseBody(res)
			}
			w.Header().Set("goa-error", "internal-error")
			w.WriteHeader(http.StatusInternalServerError)
			return enc.Encode(body)
		case "not-found":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewVersionsByIDNotFoundResponseBody(res)
			}
			w.Header().Set("goa-error", "not-found")
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeByCatalogKindNameVersionResponse returns an encoder for responses
// returned by the resource ByCatalogKindNameVersion endpoint.
func EncodeByCatalogKindNameVersionResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*resourceviews.ResourceVersion)
		enc := encoder(ctx, w)
		body := NewByCatalogKindNameVersionResponseBody(res.Projected)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeByCatalogKindNameVersionRequest returns a decoder for requests sent to
// the resource ByCatalogKindNameVersion endpoint.
func DecodeByCatalogKindNameVersionRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			catalog string
			kind    string
			name    string
			version string
			err     error

			params = mux.Vars(r)
		)
		catalog = params["catalog"]
		kind = params["kind"]
		if !(kind == "task" || kind == "pipeline") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("kind", kind, []interface{}{"task", "pipeline"}))
		}
		name = params["name"]
		version = params["version"]
		if err != nil {
			return nil, err
		}
		payload := NewByCatalogKindNameVersionPayload(catalog, kind, name, version)

		return payload, nil
	}
}

// EncodeByCatalogKindNameVersionError returns an encoder for errors returned
// by the ByCatalogKindNameVersion resource endpoint.
func EncodeByCatalogKindNameVersionError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "internal-error":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewByCatalogKindNameVersionInternalErrorResponseBody(res)
			}
			w.Header().Set("goa-error", "internal-error")
			w.WriteHeader(http.StatusInternalServerError)
			return enc.Encode(body)
		case "not-found":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewByCatalogKindNameVersionNotFoundResponseBody(res)
			}
			w.Header().Set("goa-error", "not-found")
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeByVersionIDResponse returns an encoder for responses returned by the
// resource ByVersionId endpoint.
func EncodeByVersionIDResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*resourceviews.ResourceVersion)
		enc := encoder(ctx, w)
		body := NewByVersionIDResponseBody(res.Projected)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeByVersionIDRequest returns a decoder for requests sent to the resource
// ByVersionId endpoint.
func DecodeByVersionIDRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			versionID uint
			err       error

			params = mux.Vars(r)
		)
		{
			versionIDRaw := params["versionID"]
			v, err2 := strconv.ParseUint(versionIDRaw, 10, strconv.IntSize)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("versionID", versionIDRaw, "unsigned integer"))
			}
			versionID = uint(v)
		}
		if err != nil {
			return nil, err
		}
		payload := NewByVersionIDPayload(versionID)

		return payload, nil
	}
}

// EncodeByVersionIDError returns an encoder for errors returned by the
// ByVersionId resource endpoint.
func EncodeByVersionIDError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "internal-error":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewByVersionIDInternalErrorResponseBody(res)
			}
			w.Header().Set("goa-error", "internal-error")
			w.WriteHeader(http.StatusInternalServerError)
			return enc.Encode(body)
		case "not-found":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewByVersionIDNotFoundResponseBody(res)
			}
			w.Header().Set("goa-error", "not-found")
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeByCatalogKindNameResponse returns an encoder for responses returned by
// the resource ByCatalogKindName endpoint.
func EncodeByCatalogKindNameResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*resourceviews.Resource)
		enc := encoder(ctx, w)
		body := NewByCatalogKindNameResponseBody(res.Projected)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeByCatalogKindNameRequest returns a decoder for requests sent to the
// resource ByCatalogKindName endpoint.
func DecodeByCatalogKindNameRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			catalog          string
			kind             string
			name             string
			pipelinesversion *string
			err              error

			params = mux.Vars(r)
		)
		catalog = params["catalog"]
		kind = params["kind"]
		if !(kind == "task" || kind == "pipeline") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("kind", kind, []interface{}{"task", "pipeline"}))
		}
		name = params["name"]
		pipelinesversionRaw := r.URL.Query().Get("pipelinesversion")
		if pipelinesversionRaw != "" {
			pipelinesversion = &pipelinesversionRaw
		}
		if pipelinesversion != nil {
			err = goa.MergeErrors(err, goa.ValidatePattern("pipelinesversion", *pipelinesversion, "^\\d+(?:\\.\\d+){0,2}$"))
		}
		if err != nil {
			return nil, err
		}
		payload := NewByCatalogKindNamePayload(catalog, kind, name, pipelinesversion)

		return payload, nil
	}
}

// EncodeByCatalogKindNameError returns an encoder for errors returned by the
// ByCatalogKindName resource endpoint.
func EncodeByCatalogKindNameError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "internal-error":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewByCatalogKindNameInternalErrorResponseBody(res)
			}
			w.Header().Set("goa-error", "internal-error")
			w.WriteHeader(http.StatusInternalServerError)
			return enc.Encode(body)
		case "not-found":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewByCatalogKindNameNotFoundResponseBody(res)
			}
			w.Header().Set("goa-error", "not-found")
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeByIDResponse returns an encoder for responses returned by the resource
// ById endpoint.
func EncodeByIDResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*resourceviews.Resource)
		enc := encoder(ctx, w)
		body := NewByIDResponseBody(res.Projected)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeByIDRequest returns a decoder for requests sent to the resource ById
// endpoint.
func DecodeByIDRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			id  uint
			err error

			params = mux.Vars(r)
		)
		{
			idRaw := params["id"]
			v, err2 := strconv.ParseUint(idRaw, 10, strconv.IntSize)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("id", idRaw, "unsigned integer"))
			}
			id = uint(v)
		}
		if err != nil {
			return nil, err
		}
		payload := NewByIDPayload(id)

		return payload, nil
	}
}

// EncodeByIDError returns an encoder for errors returned by the ById resource
// endpoint.
func EncodeByIDError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "internal-error":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewByIDInternalErrorResponseBody(res)
			}
			w.Header().Set("goa-error", "internal-error")
			w.WriteHeader(http.StatusInternalServerError)
			return enc.Encode(body)
		case "not-found":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewByIDNotFoundResponseBody(res)
			}
			w.Header().Set("goa-error", "not-found")
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// marshalResourceviewsResourceDataViewToResourceDataResponseBodyWithoutVersion
// builds a value of type *ResourceDataResponseBodyWithoutVersion from a value
// of type *resourceviews.ResourceDataView.
func marshalResourceviewsResourceDataViewToResourceDataResponseBodyWithoutVersion(v *resourceviews.ResourceDataView) *ResourceDataResponseBodyWithoutVersion {
	res := &ResourceDataResponseBodyWithoutVersion{
		ID:     *v.ID,
		Name:   *v.Name,
		Kind:   *v.Kind,
		Rating: *v.Rating,
	}
	if v.Catalog != nil {
		res.Catalog = marshalResourceviewsCatalogViewToCatalogResponseBodyMin(v.Catalog)
	}
	if v.Categories != nil {
		res.Categories = make([]*CategoryResponseBody, len(v.Categories))
		for i, val := range v.Categories {
			res.Categories[i] = marshalResourceviewsCategoryViewToCategoryResponseBody(val)
		}
	}
	if v.LatestVersion != nil {
		res.LatestVersion = marshalResourceviewsResourceVersionDataViewToResourceVersionDataResponseBodyWithoutResource(v.LatestVersion)
	}
	if v.Tags != nil {
		res.Tags = make([]*TagResponseBody, len(v.Tags))
		for i, val := range v.Tags {
			res.Tags[i] = marshalResourceviewsTagViewToTagResponseBody(val)
		}
	}
	if v.Platforms != nil {
		res.Platforms = make([]*PlatformResponseBody, len(v.Platforms))
		for i, val := range v.Platforms {
			res.Platforms[i] = marshalResourceviewsPlatformViewToPlatformResponseBody(val)
		}
	}

	return res
}

// marshalResourceviewsCatalogViewToCatalogResponseBodyMin builds a value of
// type *CatalogResponseBodyMin from a value of type *resourceviews.CatalogView.
func marshalResourceviewsCatalogViewToCatalogResponseBodyMin(v *resourceviews.CatalogView) *CatalogResponseBodyMin {
	res := &CatalogResponseBodyMin{
		ID:   *v.ID,
		Name: *v.Name,
		Type: *v.Type,
	}

	return res
}

// marshalResourceviewsCategoryViewToCategoryResponseBody builds a value of
// type *CategoryResponseBody from a value of type *resourceviews.CategoryView.
func marshalResourceviewsCategoryViewToCategoryResponseBody(v *resourceviews.CategoryView) *CategoryResponseBody {
	res := &CategoryResponseBody{
		ID:   *v.ID,
		Name: *v.Name,
	}

	return res
}

// marshalResourceviewsResourceVersionDataViewToResourceVersionDataResponseBodyWithoutResource
// builds a value of type *ResourceVersionDataResponseBodyWithoutResource from
// a value of type *resourceviews.ResourceVersionDataView.
func marshalResourceviewsResourceVersionDataViewToResourceVersionDataResponseBodyWithoutResource(v *resourceviews.ResourceVersionDataView) *ResourceVersionDataResponseBodyWithoutResource {
	res := &ResourceVersionDataResponseBodyWithoutResource{
		ID:                  *v.ID,
		Version:             *v.Version,
		DisplayName:         *v.DisplayName,
		Description:         *v.Description,
		MinPipelinesVersion: *v.MinPipelinesVersion,
		RawURL:              *v.RawURL,
		WebURL:              *v.WebURL,
		UpdatedAt:           *v.UpdatedAt,
	}
	if v.Platforms != nil {
		res.Platforms = make([]*PlatformResponseBody, len(v.Platforms))
		for i, val := range v.Platforms {
			res.Platforms[i] = marshalResourceviewsPlatformViewToPlatformResponseBody(val)
		}
	}

	return res
}

// marshalResourceviewsPlatformViewToPlatformResponseBody builds a value of
// type *PlatformResponseBody from a value of type *resourceviews.PlatformView.
func marshalResourceviewsPlatformViewToPlatformResponseBody(v *resourceviews.PlatformView) *PlatformResponseBody {
	res := &PlatformResponseBody{
		ID:   *v.ID,
		Name: *v.Name,
	}

	return res
}

// marshalResourceviewsTagViewToTagResponseBody builds a value of type
// *TagResponseBody from a value of type *resourceviews.TagView.
func marshalResourceviewsTagViewToTagResponseBody(v *resourceviews.TagView) *TagResponseBody {
	res := &TagResponseBody{
		ID:   *v.ID,
		Name: *v.Name,
	}

	return res
}

// marshalResourceviewsVersionsViewToVersionsResponseBody builds a value of
// type *VersionsResponseBody from a value of type *resourceviews.VersionsView.
func marshalResourceviewsVersionsViewToVersionsResponseBody(v *resourceviews.VersionsView) *VersionsResponseBody {
	res := &VersionsResponseBody{}
	if v.Latest != nil {
		res.Latest = marshalResourceviewsResourceVersionDataViewToResourceVersionDataResponseBodyMin(v.Latest)
	}
	if v.Versions != nil {
		res.Versions = make([]*ResourceVersionDataResponseBodyMin, len(v.Versions))
		for i, val := range v.Versions {
			res.Versions[i] = marshalResourceviewsResourceVersionDataViewToResourceVersionDataResponseBodyMin(val)
		}
	}

	return res
}

// marshalResourceviewsResourceVersionDataViewToResourceVersionDataResponseBodyMin
// builds a value of type *ResourceVersionDataResponseBodyMin from a value of
// type *resourceviews.ResourceVersionDataView.
func marshalResourceviewsResourceVersionDataViewToResourceVersionDataResponseBodyMin(v *resourceviews.ResourceVersionDataView) *ResourceVersionDataResponseBodyMin {
	res := &ResourceVersionDataResponseBodyMin{
		ID:      *v.ID,
		Version: *v.Version,
		RawURL:  *v.RawURL,
		WebURL:  *v.WebURL,
	}
	if v.Platforms != nil {
		res.Platforms = make([]*PlatformResponseBody, len(v.Platforms))
		for i, val := range v.Platforms {
			res.Platforms[i] = marshalResourceviewsPlatformViewToPlatformResponseBody(val)
		}
	}

	return res
}

// marshalResourceviewsResourceVersionDataViewToResourceVersionDataResponseBody
// builds a value of type *ResourceVersionDataResponseBody from a value of type
// *resourceviews.ResourceVersionDataView.
func marshalResourceviewsResourceVersionDataViewToResourceVersionDataResponseBody(v *resourceviews.ResourceVersionDataView) *ResourceVersionDataResponseBody {
	res := &ResourceVersionDataResponseBody{
		ID:                  *v.ID,
		Version:             *v.Version,
		DisplayName:         *v.DisplayName,
		Description:         *v.Description,
		MinPipelinesVersion: *v.MinPipelinesVersion,
		RawURL:              *v.RawURL,
		WebURL:              *v.WebURL,
		UpdatedAt:           *v.UpdatedAt,
	}
	if v.Platforms != nil {
		res.Platforms = make([]*PlatformResponseBody, len(v.Platforms))
		for i, val := range v.Platforms {
			res.Platforms[i] = marshalResourceviewsPlatformViewToPlatformResponseBody(val)
		}
	}
	if v.Resource != nil {
		res.Resource = marshalResourceviewsResourceDataViewToResourceDataResponseBodyInfo(v.Resource)
	}

	return res
}

// marshalResourceviewsResourceDataViewToResourceDataResponseBodyInfo builds a
// value of type *ResourceDataResponseBodyInfo from a value of type
// *resourceviews.ResourceDataView.
func marshalResourceviewsResourceDataViewToResourceDataResponseBodyInfo(v *resourceviews.ResourceDataView) *ResourceDataResponseBodyInfo {
	res := &ResourceDataResponseBodyInfo{
		ID:     *v.ID,
		Name:   *v.Name,
		Kind:   *v.Kind,
		Rating: *v.Rating,
	}
	if v.Catalog != nil {
		res.Catalog = marshalResourceviewsCatalogViewToCatalogResponseBodyMin(v.Catalog)
	}
	if v.Categories != nil {
		res.Categories = make([]*CategoryResponseBody, len(v.Categories))
		for i, val := range v.Categories {
			res.Categories[i] = marshalResourceviewsCategoryViewToCategoryResponseBody(val)
		}
	}
	if v.Tags != nil {
		res.Tags = make([]*TagResponseBody, len(v.Tags))
		for i, val := range v.Tags {
			res.Tags[i] = marshalResourceviewsTagViewToTagResponseBody(val)
		}
	}
	if v.Platforms != nil {
		res.Platforms = make([]*PlatformResponseBody, len(v.Platforms))
		for i, val := range v.Platforms {
			res.Platforms[i] = marshalResourceviewsPlatformViewToPlatformResponseBody(val)
		}
	}

	return res
}

// marshalResourceviewsResourceDataViewToResourceDataResponseBody builds a
// value of type *ResourceDataResponseBody from a value of type
// *resourceviews.ResourceDataView.
func marshalResourceviewsResourceDataViewToResourceDataResponseBody(v *resourceviews.ResourceDataView) *ResourceDataResponseBody {
	res := &ResourceDataResponseBody{
		ID:     *v.ID,
		Name:   *v.Name,
		Kind:   *v.Kind,
		Rating: *v.Rating,
	}
	if v.Catalog != nil {
		res.Catalog = marshalResourceviewsCatalogViewToCatalogResponseBodyMin(v.Catalog)
	}
	if v.Categories != nil {
		res.Categories = make([]*CategoryResponseBody, len(v.Categories))
		for i, val := range v.Categories {
			res.Categories[i] = marshalResourceviewsCategoryViewToCategoryResponseBody(val)
		}
	}
	if v.LatestVersion != nil {
		res.LatestVersion = marshalResourceviewsResourceVersionDataViewToResourceVersionDataResponseBodyWithoutResource(v.LatestVersion)
	}
	if v.Tags != nil {
		res.Tags = make([]*TagResponseBody, len(v.Tags))
		for i, val := range v.Tags {
			res.Tags[i] = marshalResourceviewsTagViewToTagResponseBody(val)
		}
	}
	if v.Platforms != nil {
		res.Platforms = make([]*PlatformResponseBody, len(v.Platforms))
		for i, val := range v.Platforms {
			res.Platforms[i] = marshalResourceviewsPlatformViewToPlatformResponseBody(val)
		}
	}
	if v.Versions != nil {
		res.Versions = make([]*ResourceVersionDataResponseBodyTiny, len(v.Versions))
		for i, val := range v.Versions {
			res.Versions[i] = marshalResourceviewsResourceVersionDataViewToResourceVersionDataResponseBodyTiny(val)
		}
	}

	return res
}

// marshalResourceviewsResourceVersionDataViewToResourceVersionDataResponseBodyTiny
// builds a value of type *ResourceVersionDataResponseBodyTiny from a value of
// type *resourceviews.ResourceVersionDataView.
func marshalResourceviewsResourceVersionDataViewToResourceVersionDataResponseBodyTiny(v *resourceviews.ResourceVersionDataView) *ResourceVersionDataResponseBodyTiny {
	res := &ResourceVersionDataResponseBodyTiny{
		ID:      *v.ID,
		Version: *v.Version,
	}

	return res
}
