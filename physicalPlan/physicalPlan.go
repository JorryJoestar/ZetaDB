package execution

import "ZetaDB/container"

type PhysicalPlan struct {
	entranceIterator Iterator
}

func NewPhysicalPlan(entranceIterator Iterator) *PhysicalPlan {
	return &PhysicalPlan{
		entranceIterator: entranceIterator,
	}
}

func (pp *PhysicalPlan) HasNext() bool {
	return pp.entranceIterator.HasNext()
}

func (pp *PhysicalPlan) GetNext() (*container.Tuple, error) {
	return pp.entranceIterator.GetNext()
}

func (pp *PhysicalPlan) GetSchema() *container.Schema {
	return pp.entranceIterator.GetSchema()
}
