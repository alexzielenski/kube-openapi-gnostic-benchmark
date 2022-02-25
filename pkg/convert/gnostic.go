package convert

import (
	openapi_v2 "github.com/googleapis/gnostic/openapiv2"
	openapi_v3 "github.com/googleapis/gnostic/openapiv3"
	"k8s.io/kube-openapi/pkg/spec3"
	"k8s.io/kube-openapi/pkg/validation/spec"
)

func GnosticToKubeV2(openapi_v2.Document) *spec.Swagger {
	return nil
}

func GnosticToKubeV3(openapi_v3.Document) *spec3.OpenAPI {
	return nil
}
