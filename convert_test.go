package parsebench

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/go-openapi/jsonreference"
	openapi_v2 "github.com/googleapis/gnostic/openapiv2"
	openapi_v3 "github.com/googleapis/gnostic/openapiv3"
	"google.golang.org/protobuf/proto"
	"gopkg.in/yaml.v3"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/kube-openapi/pkg/spec3"
	"k8s.io/kube-openapi/pkg/validation/spec"
)

func ConnectToCluster() (*kubernetes.Clientset, error) {
	// uses the current context in kubeconfig
	configOverrides := &clientcmd.ConfigOverrides{CurrentContext: ""}

	config2 := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(clientcmd.NewDefaultClientConfigLoadingRules(), configOverrides)
	config3, err := config2.ClientConfig()

	if err != nil {
		return nil, err
	}

	// config, _ := clientcmd.BuildConfigFromFlags("", "<path-to-kubeconfig>")

	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config3)
	if err != nil {
		return nil, err
	}
	return clientset, nil
}

type CaseV3 struct {
	path         string
	originalJSON []byte
	newJSON      []byte
	expect       *spec3.OpenAPI
	gnostic      *openapi_v3.Document
	yaml         *yaml.Node
	pb           []byte
}

func NewCaseV3(b *testing.B, path string, src []byte) *CaseV3 {
	res := &CaseV3{}
	res.path = path
	res.originalJSON = src
	return res
}

func CommonConversionV3(b *testing.B) []*CaseV3 {
	res := []*CaseV3{}
	err := filepath.Walk("spec",
		func(path string, _ os.FileInfo, err error) error {
			if err != nil {
				return err
			} else if strings.HasSuffix(path, ".json") {
				file, err := os.Open(path)
				if err != nil {
					return err
				}
				defer file.Close()
				bytes, err := io.ReadAll(file)
				if err != nil {
					return err
				}
				res = append(res, NewCaseV3(b, path, bytes))
			}
			return nil
		})
	if err != nil {
		b.Fatal(err)
	}

	b.Run("json->kube", func(b2 *testing.B) {
		for i := 0; i < b2.N; i++ {
			for _, c := range res {
				err = json.Unmarshal(c.originalJSON, &c.expect)
				if err != nil {
					b.Fatal(err)
				}
			}
		}
	})

	b.Run("json->gnostic", func(b2 *testing.B) {
		for i := 0; i < b2.N; i++ {
			for _, c := range res {
				c.gnostic, err = openapi_v3.ParseDocument(c.originalJSON)
				if err != nil {
					b.Fatal(err)
				}
			}
		}
	})

	b.Run("gnostic->pb", func(b2 *testing.B) {
		for i := 0; i < b2.N; i++ {
			for _, c := range res {
				c.pb, err = proto.Marshal(c.gnostic)
				if err != nil {
					b.Fatal(err)
				}
			}
		}
	})

	b.Run("pb->gnostic", func(b2 *testing.B) {
		for i := 0; i < b2.N; i++ {
			for _, c := range res {
				err = proto.Unmarshal(c.pb, c.gnostic)
				if err != nil {
					b.Fatal(err)
				}
			}
		}
	})

	b.Run("gnostic->yaml", func(b2 *testing.B) {
		for i := 0; i < b2.N; i++ {
			for _, c := range res {
				c.yaml = c.gnostic.ToRawInfo()
			}
		}
	})

	return res
}

func CommonConversion(b *testing.B) (*spec.Swagger, *openapi_v2.Document, *yaml.Node, []byte) {
	// Download kube-openapi swagger json
	// clientset, err := ConnectToCluster()
	// if err != nil {
	// 	b.Fatal(err)
	// }

	// restClient := clientset.RESTClient()
	// originalJSON, err := restClient.
	// 	Get().
	// 	AbsPath("/openapi/v2").
	// 	SetHeader("Accept", "application/json").
	// 	Do(context.TODO()).
	// 	Raw()
	// if err != nil {
	// 	b.Fatal(err)
	// }
	swagFile, err := os.Open("swagger.json")
	if err != nil {
		b.Fatal(err)
	}
	originalJSON, err := io.ReadAll(swagFile)
	if err != nil {
		b.Fatal(err)
	}

	f, err := os.Create("swagger.json")
	if err != nil {
		b.Fatal(err)
	}
	defer f.Close()
	_, err2 := f.Write(originalJSON)
	if err2 != nil {
		b.Fatal(err2)
	}

	// Parse into kube-openapi types
	var result *spec.Swagger
	b.Run("json->swagger", func(b2 *testing.B) {
		for i := 0; i < b2.N; i++ {
			if err := json.Unmarshal(originalJSON, &result); err != nil {
				b2.Fatal(err)
			}
		}
	})

	// Convert to JSON
	var encodedJSON []byte
	b.Run("swagger->json", func(b2 *testing.B) {
		for i := 0; i < b2.N; i++ {
			encodedJSON, err = json.Marshal(result)
			if err != nil {
				b2.Fatal(err)
			}
		}
	})

	// Convert to gnostic
	var originalGnostic *openapi_v2.Document
	b.Run("json->gnostic", func(b2 *testing.B) {
		for i := 0; i < b2.N; i++ {
			originalGnostic, err = openapi_v2.ParseDocument(encodedJSON)
			if err != nil {
				b2.Fatal(err)
			}
		}
	})

	// Convert to PB
	var encodedProto []byte
	b.Run("gnostic->pb", func(b2 *testing.B) {
		for i := 0; i < b2.N; i++ {
			encodedProto, err = proto.Marshal(originalGnostic)
			if err != nil {
				b2.Fatal(err)
			}
		}
	})

	// Convert to gnostic
	var backToGnostic openapi_v2.Document
	b.Run("pb->gnostic", func(b2 *testing.B) {
		for i := 0; i < b2.N; i++ {
			if err := proto.Unmarshal(encodedProto, &backToGnostic); err != nil {
				b2.Fatal(err)
			}
		}
	})

	var rawInfo *yaml.Node
	b.Run("gnostic->yaml", func(b2 *testing.B) {
		for i := 0; i < b2.N; i++ {
			rawInfo = backToGnostic.ToRawInfo()
		}
	})

	// if !reflect.DeepEqual(originalGnostic, backToGnostic) {
	// 	// b.Log("gnostic roundtrip loses information")
	// }
	return result, &backToGnostic, rawInfo, originalJSON
}

func BenchmarkSlowConversion(b *testing.B) {
	originalSwagger, _, yamlNode, _ := CommonConversion(b)

	for i := 0; i < b.N; i++ {
		// Convert to JSON
		var encodedJSON []byte
		var err error
		b.Run("yaml->json", func(b2 *testing.B) {
			for i := 0; i < b2.N; i++ {
				var generic interface{}
				if err := yamlNode.Decode(&generic); err != nil {
					b.Fatal(err)
				}

				encodedJSON, err = json.Marshal(generic)
				if err != nil {
					b.Fatal(err)
				}
			}
		})

		// Convert to kube-openapi
		b.Run("json->swagger", func(b2 *testing.B) {
			for i := 0; i < b2.N; i++ {
				var decodedSwagger spec.Swagger
				if err := json.Unmarshal(encodedJSON, &decodedSwagger); err != nil {
					b.Fatal(err)
				}

				if len(originalSwagger.Definitions) != len(decodedSwagger.Definitions) {
					b.Fatalf("lost definitions %v %v", len(originalSwagger.Definitions), len(decodedSwagger.Definitions))
				}
			}
		})
	}
}

func CmpRef(a, b jsonreference.Ref) bool {
	return a.String() == b.String()
}

func BenchmarkFastConversion(b *testing.B) {
	originalSwagger, _, rawInfo, _ := CommonConversion(b)

	for i := 0; i < b.N; i++ {
		b.Run("yaml->swagger", func(b2 *testing.B) {
			for i := 0; i < b2.N; i++ {
				// Convert to yaml.Node
				yamlNode := &yaml.Node{
					Kind:        yaml.DocumentNode,
					Content:     []*yaml.Node{rawInfo},
					HeadComment: "",
				}

				// Convert to kube-openapi
				var decodedSwagger spec.Swagger

				if err := yamlNode.Decode(&decodedSwagger); err != nil {
					b2.Fatal(err)
				}

				if len(originalSwagger.Definitions) != len(decodedSwagger.Definitions) {
					b2.Fatal("lost definitions")
				}
				// fmt.Println(cmp.Diff(originalSwagger, &decodedSwagger, cmp.Comparer(CmpRef)))
			}
			// b.Log(cmp.Equal(originalSwagger, &decodedSwagger, cmp.Comparer(CmpRef)))
		})
	}
}

// func BenchmarkEasyJSON(b *testing.B) {
// 	originalSwagger, _, _, src := CommonConversion(b)
// 	for i := 0; i < b.N; i++ {
// 		b.Run("easyjson->swagger", func(b2 *testing.B) {
// 			for i := 0; i < b2.N; i++ {
// 				var decodedSwagger spec.Swagger
// 				if err := easyjson.Unmarshal(src, &decodedSwagger); err != nil {
// 					b2.Fatal(err)
// 				}

// 				if len(originalSwagger.Definitions) != len(decodedSwagger.Definitions) {
// 					b2.Fatal("lost definitions")
// 				}
// 			}
// 		})
// 	}
// }

func BenchmarkFastConversionV3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cases := CommonConversionV3(b)
		b.Run("yaml->kube", func(b2 *testing.B) {
			for i := 0; i < b2.N; i++ {
				for _, c := range cases {
					yamlNode := &yaml.Node{
						Kind:        yaml.DocumentNode,
						Content:     []*yaml.Node{c.yaml},
						HeadComment: "",
					}

					// Convert to kube-openapi
					var decodedSwagger spec3.OpenAPI
					if err := yamlNode.Decode(&decodedSwagger); err != nil {
						b2.Fatal(err)
					}

					if len(c.expect.Components.Schemas) != len(decodedSwagger.Components.Schemas) {
						b2.Fatal("lost definitions")
					}
				}
			}
		})
	}
}

func BenchmarkSlowConversionV3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cases := CommonConversionV3(b)
		b.Run("gnostic->json", func(b2 *testing.B) {
			for i := 0; i < b2.N; i++ {
				for _, c := range cases {
					var err error
					c.newJSON, err = json.Marshal(c.gnostic)
					if err != nil {
						b2.Fatal(err)
					}
				}
			}
		})

		b.Run("json->kube", func(b2 *testing.B) {
			for i := 0; i < b2.N; i++ {
				for _, c := range cases {
					var decodedSwagger spec3.OpenAPI
					if err := json.Unmarshal(c.originalJSON, &decodedSwagger); err != nil {
						b2.Fatal(err)
					}

					if len(c.expect.Components.Schemas) != len(decodedSwagger.Components.Schemas) {
						b2.Fatal("lost definitions")
					}
				}
			}
		})
	}
}

// BenchmarkFastConversion/json->swagger-8         	       2	 573748519 ns/op	95734136 B/op	 1381621 allocs/op
// BenchmarkFastConversion/swagger->json-8         	       7	 169633842 ns/op	71493013 B/op	  274090 allocs/op
// BenchmarkFastConversion/json->gnostic-8         	       4	 254860818 ns/op	81450744 B/op	 1248399 allocs/op
// BenchmarkFastConversion/gnostic->pb-8           	     100	  24392588 ns/op	 2899969 B/op	       1 allocs/op
// BenchmarkFastConversion/pb->gnostic-8           	      63	  18125862 ns/op	 9480707 B/op	  123829 allocs/op
// BenchmarkFastConversion/gnostic->yaml-8         	      32	  34916985 ns/op	32855202 B/op	  264562 allocs/op
// BenchmarkFastConversion/yaml->swagger-8         	       6	 201319739 ns/op	66092041 B/op	 1354028 allocs/op
