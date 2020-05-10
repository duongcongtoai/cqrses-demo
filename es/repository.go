package es

type DomainRepository interface {
	Load(aggregateTypeName string, aggregateId string) (AggregateRoot, error)
	Save(aggregate AggregateRoot) error
}
