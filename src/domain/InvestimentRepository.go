package domain

type InvestimentRepository interface {
	Create(ativo *Ativo) (*Ativo, error)
	FindByCode(code string) (*Ativo, error)
	Count() int64
}
