package gql

//Executor executes given query on a data source and returns results.
type Executor interface {}


//Query
type Query interface {
	//Where applies a boolean filter on data.
	Where() Where
	//GroupBy groups the data based on a feature.
	GroupBy() Group
	//SortBy sorts data.
	SortBy() Sort
	//Aggregate data with an aggregator.
	Aggregate(Aggregator)
}

type Where interface {
	Query() Query
	And() Where
	Or() Where
}

type Aggregator interface {

}

type Group interface {
	By() Group
	Query() Query
}

type Sort interface {
	By() Sort
	Query() Query
}

func main() {

}