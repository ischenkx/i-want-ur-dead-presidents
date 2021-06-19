package billing

type Client interface {
	GetBalances(ids []string) ([]int64, error)

	RegisterWallet(id string) error
}
