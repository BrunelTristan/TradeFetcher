package fetcher

type FakeFetcher struct {
}

func NewFakeFetcher() IFetcher {
	return &FakeFetcher{}
}
