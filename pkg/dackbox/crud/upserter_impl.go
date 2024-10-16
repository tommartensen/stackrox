package crud

import (
	"github.com/stackrox/rox/pkg/dackbox"
	"github.com/stackrox/rox/pkg/protocompat"
)

type upserterImpl struct {
	keyFunc ProtoKeyFunction

	addToIndex bool
}

// UpsertIn saves the input object and adds a reference to it from the input parentKey if one is passed in.
func (uc *upserterImpl) UpsertIn(parentKey []byte, msg protocompat.Message, dackTxn *dackbox.Transaction) error {
	// Generate key.
	key := uc.keyFunc(msg)

	// If a parent key is set, add the generated key to the parent's child list.
	if len(parentKey) != 0 {
		dackTxn.Graph().AddRefs(parentKey, key)
	}
	if uc.addToIndex {
		dackTxn.MarkDirty(key, msg)
	}

	// Marshal an upsert the base object.
	toWrite, err := protocompat.Marshal(msg)
	if err != nil {
		return err
	}
	dackTxn.Set(key, toWrite)
	return nil
}
