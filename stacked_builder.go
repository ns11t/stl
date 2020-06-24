package stl

import (
	"context"

	"github.com/gofrs/uuid"
)

// NewStacked creates an instance of stacked Builder.
func NewStacked() Builder {
	id, err := uuid.NewV4()
	if err != nil {
		panic(err)
	}

	return &stackedBuilder{id: id.String()}
}

type stackedBuilder struct {
	id     string
	prefix string
	shs    []string
	exs    []string
	v      Vault
}

func (t *stackedBuilder) Shared(name string) Builder {
	if name == "" {
		panic("resource name could not be empty")
	}

	t.prefix = t.prefix + name
	t.shs = append(t.shs, t.prefix)

	return t
}

func (t *stackedBuilder) Exclusive(name string) Builder {
	if name == "" {
		panic("resource name could not be empty")
	}

	t.prefix = t.prefix + name
	t.exs = append(t.exs, t.prefix)

	return t
}

func (t *stackedBuilder) ListShared() []string {
	return t.shs
}

func (t *stackedBuilder) ListExclusive() []string {
	return t.exs
}

func (t *stackedBuilder) ID() string {
	return t.id
}

func (t *stackedBuilder) ToTx() Tx {
	return t
}

func (t *stackedBuilder) ToLocker(v Vault) Locker {
	t.v = v
	return t
}

func (t *stackedBuilder) Lock() {
	_ = t.v.Lock(context.Background(), t)
}

func (t *stackedBuilder) Unlock() {
	t.v.Unlock(t)
}

func (t *stackedBuilder) LockWithContext(ctx context.Context) error {
	return t.v.Lock(ctx, t)
}
