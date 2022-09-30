package domain

type InvestimentRepository interface {
	Create(ativo *Ativo) (*Ativo, error)
}
