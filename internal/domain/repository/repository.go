package repository

import (
	"context"
	"time"

	"github.com/null0routed/mariner/internal/vault"
)

type VaultRepository interface {
	Exists(ctx context.Context, userID string) (bool, error)
	Load(ctx context.Context, userID, password string) (data vault.Data, exists bool, err error)
	Save(ctx context.Context, userID, password string, data vault.Data) (error)
	Delete(ctx context.Context, userID string) error
}

type OrganizationRepository interface {
	List(ctx context.Context) (map[string]vault.Organization, error)
	Save(ctx context.Context, org vault.Organization) error
	Delete(ctx context.Context, orgID string) error
}

type PreferencesRepository interface {
	Get(ctx context.Context, userID string) (map[string]bool, error)
	Save(ctx context.Context, userID string, preferences map[string]bool) error
}

type AuditAppender interface {
	Append(ctx context.Context, even vault.AuditEvent) error
}

type AuditReader interface {
	List(ctx context.Context, filter vault.AuditFilter) (vault.AuditPage, error)
	ListActions(ctx context.Context) ([]string, error)
	Cursor(ctx context.Context) (occurredAt, eventID string, err error)
	ListAfter(ctx context.Context, occurredAt, eventID string, limit int) ([]vault.AuditRecord, error)
}

type UploadRepository interface {
	Create(ctx context.Context, upload vault.MultipartUpload) error
	Load(ctx context.Context, uploadID string) (vault.MultipartUpload, bool, error)
	SavePart(ctx context.Context, uploadID string, expectedPart int32, partsJSON, hashState string) (bool, error)
	ClaimComplete(ctx context.Context, uploadID string) (vault.MultipartUpload, bool, error)
	Finish(ctx context.Context, uploadID string, success bool) error
	Delete(ctx context.Context, uploadID string) error
	ListStale(ctx context.Context, before time.Time) ([]vault.MultipartUpload, error)
	ClaimCleanup(ctx context.Context, uploadID string, before time.Time) (bool, error)
}

type Closer interface {
	Close() error
}