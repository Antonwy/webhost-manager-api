package zones

type Zone struct {
	ID                   string `db:"id" json:"id"`
	Name                 string `db:"name" json:"name"`
	SyncedWithCloudflare bool   `db:"synced_with_cloudflare" json:"syncedWithCloudflare"`
}
