package datamodels

type ProgramItem struct {
	functionExpressionNode *ExpressionListNode
	currentParameterIndex int
	variableMap map[string] Node
	functionName string
}

func NewProgramItem(
	functionExpressionNode *ExpressionListNode,
	currentParameterIndex int,
	variableMap map[string] Node,
	functionName string,
) *ProgramItem {
	return &ProgramItem{
		functionExpressionNode,
		currentParameterIndex,
		variableMap,
		functionName,
	}
}

func NewProgramItemFromExisting(existingProgramItem *ProgramItem) *ProgramItem {
	return &ProgramItem{
		existingProgramItem.GetFunctionExpressionNode(),
		existingProgramItem.GetCurrentParameterIndex() + 1,
		existingProgramItem.GetVariableMap(),
		existingProgramItem.GetFunctionName(),
	}
}

func(p *ProgramItem) GetFunctionExpressionNode() *ExpressionListNode {
	return p.functionExpressionNode
}

func (p *ProgramItem) GetCurrentParameterIndex() int {
	return p.currentParameterIndex
}

func (p *ProgramItem) GetVariableMap() map[string] Node {
	return p.variableMap
}

func (p *ProgramItem) GetFunctionName() string {
	return p.functionName
}
