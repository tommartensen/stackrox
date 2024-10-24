// Code generated by cp generator. DO NOT EDIT.
// This file was originally generated with
// //go:generate cp ../../../central/nodecomponentedge/dackbox/crud.go  nodecomponentedge/crud.go
package dackbox

import (
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/dackbox/crud"
	"github.com/stackrox/rox/pkg/dbhelper"
	"github.com/stackrox/rox/pkg/protocompat"
)

var (
	// Bucket stores the node to component edges.
	Bucket = []byte("node_to_comp")

	// BucketHandler is the bucket's handler.
	BucketHandler = &dbhelper.BucketHandler{BucketPrefix: Bucket}

	// Reader reads storage.NodeComponentEdges directly from the store.
	Reader = crud.NewReader(
		crud.WithAllocFunction(alloc),
	)

	// Upserter writes storage.NodeComponentEdges directly to the store.
	Upserter = crud.NewUpserter(
		crud.WithKeyFunction(crud.PrefixKey(Bucket, keyFunc)),
		crud.AddToIndex(),
	)

	// Deleter deletes the edges from the store.
	Deleter = crud.NewDeleter(
		crud.Shared(),
		crud.RemoveFromIndex(),
	)
)

func keyFunc(msg protocompat.Message) []byte {
	return []byte(msg.(*storage.NodeComponentEdge).GetId())
}

func alloc() protocompat.Message {
	return &storage.NodeComponentEdge{}
}
