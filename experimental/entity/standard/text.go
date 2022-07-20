package standard

import (
	"github.com/grafana/grafana-plugin-sdk-go/data"
	"github.com/grafana/grafana-plugin-sdk-go/experimental/entity"
)

type PlainTextEntity struct {
	entity.Envelope

	Body string `json:"body,omitempty"`
}

var _ entity.Kind = &PlainTextKind{}

type PlainTextKind struct {
	info entity.KindInfo
}

func NewPlainTextKind(info entity.KindInfo) *PlainTextKind {
	return &PlainTextKind{info: info}
}

func (k *PlainTextKind) Info() entity.KindInfo {
	k.info.IsRaw = true
	return k.info
}

func (k *PlainTextKind) GoType() interface{} {
	return &PlainTextEntity{}
}

func (k *PlainTextKind) Read(payload []byte) (interface{}, error) {
	// ?? make sure the payload is safe string bytes?
	g := &PlainTextEntity{}
	g.Body = string(payload)
	return g, nil
}

func (k *PlainTextKind) Validate(payload []byte, details bool) entity.ValidationResponse {
	_, err := k.Read(payload)
	if err != nil {
		return entity.ValidationResponse{
			Valid: false,
			Info: []data.Notice{
				{
					Severity: data.NoticeSeverityError,
					Text:     err.Error(),
				},
			},
		}
	}
	return entity.ValidationResponse{
		Valid:  true,
		Result: payload,
	}
}

func (k *PlainTextKind) Migrate(payload []byte, targetVersion string) entity.ValidationResponse {
	return k.Validate(payload, false) // migration is a noop
}

func (k *PlainTextKind) GetSchemaVersions() []string {
	return nil
}

func (k *PlainTextKind) GetJSONSchema(schemaVersion string) []byte {
	return nil
}
