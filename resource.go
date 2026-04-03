package kubeslim

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"k8s.io/apimachinery/pkg/runtime/schema"
)

// List retrieves the resource list for the given GVR and JSON-decodes it into T.
func List[T any](ctx context.Context, client *Client, gvr schema.GroupVersionResource) (T, error) {
	var zero T

	url := client.baseURL + buildListPath(gvr)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return zero, err
	}
	req.Header.Set("Accept", "application/json")

	resp, err := client.httpClient.Do(req)
	if err != nil {
		return zero, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		body, _ := io.ReadAll(resp.Body)
		return zero, fmt.Errorf("kubernetes api error: %s %s", resp.Status, string(body))
	}

	var result T
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return zero, err
	}
	return result, nil
}

// buildListPath returns the URL path for listing a resource.
func buildListPath(gvr schema.GroupVersionResource) string {
	var prefix string
	if gvr.Group == "" || gvr.Group == "core" {
		prefix = "/api/" + gvr.Version
	} else {
		prefix = "/apis/" + gvr.Group + "/" + gvr.Version
	}
	return prefix + "/" + gvr.Resource
}

// Get retrieves a single resource by name and optionally namespace, decoding it into T.
func Get[T any](ctx context.Context, client *Client, gvr schema.GroupVersionResource, namespace, name string) (T, error) {
	var zero T

	url := client.baseURL + buildGetPath(gvr, namespace, name)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return zero, err
	}
	req.Header.Set("Accept", "application/json")

	resp, err := client.httpClient.Do(req)
	if err != nil {
		return zero, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		body, _ := io.ReadAll(resp.Body)
		return zero, fmt.Errorf("kubernetes api error: %s %s", resp.Status, string(body))
	}

	var result T
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return zero, err
	}
	return result, nil
}

// buildGetPath returns the URL path for getting a specific resource.
func buildGetPath(gvr schema.GroupVersionResource, namespace, name string) string {
	var prefix string
	if gvr.Group == "" || gvr.Group == "core" {
		prefix = "/api/" + gvr.Version
	} else {
		prefix = "/apis/" + gvr.Group + "/" + gvr.Version
	}

	if namespace != "" {
		return prefix + "/namespaces/" + namespace + "/" + gvr.Resource + "/" + name
	}
	return prefix + "/" + gvr.Resource + "/" + name
}
