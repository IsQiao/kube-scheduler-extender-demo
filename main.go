package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"math/rand"
	"net/http"
	"time"

	v1 "k8s.io/api/core/v1"
	extenderapi "k8s.io/kube-scheduler/extender/v1"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func Filter(w http.ResponseWriter, r *http.Request) {
	var buf bytes.Buffer
	body := io.TeeReader(r.Body, &buf)
	var extenderArgs extenderapi.ExtenderArgs
	var extenderFilterResult *extenderapi.ExtenderFilterResult
	if err := json.NewDecoder(body).Decode(&extenderArgs); err != nil {
		extenderFilterResult = &extenderapi.ExtenderFilterResult{
			Error: err.Error(),
		}
	} else {
		extenderFilterResult = filter(extenderArgs)
	}

	if response, err := json.Marshal(extenderFilterResult); err != nil {
		log.Fatalln(err)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}

// The function filter() filters nodes according to predicates defined in this extender
func filter(args extenderapi.ExtenderArgs) *extenderapi.ExtenderFilterResult {
	var filteredNodes []v1.Node
	failedNodes := make(extenderapi.FailedNodesMap)
	// pod := args.Pod
	// nodes := args.Nodes.Items

	// TODO

	return &extenderapi.ExtenderFilterResult{
		Nodes: &v1.NodeList{
			Items: filteredNodes,
		},
		FailedNodes: failedNodes,
		Error:       "",
	}
}

func main() {
	http.HandleFunc("/filter", Filter)
	// TODO: add more handlers

	log.Fatal(http.ListenAndServe(":8080", nil))
}
