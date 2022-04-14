package execution

type Index struct {
	indexName        string
	onTableName      string
	attriNameList    []string
	indexFirstPageId uint32
}

func NewIndex() {}

func NewIndexFromBytes() {}

func IndexToBytes() {}
