package processor

import "github.com/2lovecode/powerjob-go/worker/extension/processor"

type Loader interface {
	Load(definition *processor.Definition)
}
