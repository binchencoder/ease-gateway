package runtime

import (
	"net/http"
	"testing"

	"golang.org/x/net/context"
	"google.golang.org/protobuf/proto"
)

type GatewayServiceHookFake struct {
}

func (g *GatewayServiceHookFake) Bootstrap(sgs map[string]*ServiceGroup) error {
	return nil
}
func (g *GatewayServiceHookFake) RequestReceived(w http.ResponseWriter, r *http.Request) (ctx context.Context, err error) {
	return nil, nil
}

func (g *GatewayServiceHookFake) RequestAccepted(ctx context.Context, svc *Service, m *Method, w http.ResponseWriter, r *http.Request) (ctxret context.Context, err error) {
	return nil, nil
}

func (g *GatewayServiceHookFake) RequestParsed(ctx context.Context, svc *Service, m *Method, in proto.Message, meta *ServerMetadata) error {
	return nil
}

func (g *GatewayServiceHookFake) RequestHandled(ctx context.Context, svc *Service, m *Method, out proto.Message, meta *ServerMetadata, err error) {
}

func TestHook(t *testing.T) {
	fakeHook := GatewayServiceHookFake{}
	err := SetGatewayServiceHook(&fakeHook)
	if err != nil {
		t.Errorf("Failed to set hook err=%+v", err)
	}
}
