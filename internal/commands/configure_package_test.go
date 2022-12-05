package commands_test

import (
	"reflect"
	"testing"

	"github.com/massdriver-cloud/mass/internal/api"
	"github.com/massdriver-cloud/mass/internal/commands"
)

func TestConfigurePackage(t *testing.T) {
	params := map[string]interface{}{
		"cidr": "10.0.0.0/16",
	}

	client := mockClientWithJSONResponseMap(map[string]interface{}{
		"getPackageByNamingConvention": mockQueryResponse("getPackageByNamingConvention", api.Package{
			Manifest: api.Manifest{ID: "manifest-id"},
			Target:   api.Target{ID: "target-id"},
		}),
		"configurePackage": map[string]interface{}{
			"data": map[string]interface{}{
				"configurePackage": map[string]interface{}{
					"result": map[string]interface{}{
						"id":     "pkg-id",
						"params": string(mustMarshalJSON(params)),
					},
					"successful": true,
				},
			},
		},
	})

	pkg, err := commands.ConfigurePackage(client, "faux-org-id", "ecomm-prod-cache", params)
	if err != nil {
		t.Fatal(err)
	}

	got := pkg.Params
	want := params

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}