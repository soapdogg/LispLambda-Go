package datamodels

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

//counterfeiter:generate -o fakes/fake_program_item.go . ProgramItem
type ProgramItem interface {
	GetFunctionExpressionNode() ExpressionListNode
	GetCurrentParameterIndex() int
	GetVariableMap() map[string] Node
	GetFunctionName() string
}

type programItem struct {
	functionExpressionNode ExpressionListNode
	currentParameterIndex int
	variableMap map[string] Node
	functionName string
}

func NewProgramItem(
	functionExpressionNode ExpressionListNode,
	currentParameterIndex int,
	variableMap map[string] Node,
	functionName string,
) *programItem {
	return &programItem{
		functionExpressionNode,
		currentParameterIndex,
		variableMap,
		functionName,
	}
}

func NewProgramItemFromScratch(
	functionExpressionNode ExpressionListNode,
	variableMap map[string] Node,
) *programItem {
	firstChild := functionExpressionNode.GetChildren()[0].(AtomNode)
	return &programItem{
		functionExpressionNode,
		0,
		variableMap,
		firstChild.GetValue(),
	}
}

func NewProgramItemFromExisting(existingProgramItem ProgramItem) *programItem {
	return &programItem{
		existingProgramItem.GetFunctionExpressionNode(),
		existingProgramItem.GetCurrentParameterIndex() + 1,
		existingProgramItem.GetVariableMap(),
		existingProgramItem.GetFunctionName(),
	}
}

func(p *programItem) GetFunctionExpressionNode() ExpressionListNode {
	return p.functionExpressionNode
}

func (p *programItem) GetCurrentParameterIndex() int {
	return p.currentParameterIndex
}

func (p *programItem) GetVariableMap() map[string] Node {
	return p.variableMap
}

func (p *programItem) GetFunctionName() string {
	return p.functionName
}

var _ ProgramItem = &programItem{}