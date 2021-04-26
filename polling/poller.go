package polling

import (
	"context"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

type LongRunningPoller struct {
	future *azure.Future
	ctx    context.Context
	client autorest.Client
}

func (fw *LongRunningPoller) PollUntilDone() error {
	return fw.future.WaitForCompletionRef(fw.ctx, fw.client)
}
