package schema1

import (
	"github.com/docker/distribution"
	"github.com/docker/distribution/manifest/schema1"
	"github.com/opencontainers/go-digest"
)

func schema1Func(b []byte) (distribution.Manifest, distribution.Descriptor, error) {
	sm := new(schema1.SignedManifest)
	err := sm.UnmarshalJSON(b)
	if err != nil {
		return nil, distribution.Descriptor{}, err
	}

	desc := distribution.Descriptor{
		Digest:    digest.FromBytes(sm.Canonical),
		Size:      int64(len(sm.Canonical)),
		MediaType: schema1.MediaTypeSignedManifest,
	}
	return sm, desc, err
}
