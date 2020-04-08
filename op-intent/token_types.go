package opintent

import (
	"github.com/HyperService-Consortium/go-uip/internal/token_types"
	"github.com/HyperService-Consortium/go-uip/op-intent/lexer"
	"github.com/HyperService-Consortium/go-uip/op-intent/parser"
)

type RawAccountI = token_types.RawAccountI
type NamespacedNameAccountI = token_types.NamespacedNameAccountI
type NameAccountI = token_types.NameAccountI
type Token = token_types.Token
type TokParam = token_types.Param
type LocalStateVariableI = token_types.LocalStateVariableI
type StateVariableI = token_types.StateVariableI
type NamespacedRawAccountI = token_types.NamespacedRawAccountI
type ConstantI = token_types.ConstantI
type UnaryExpressionI = token_types.UnaryExpressionI
type BinaryExpressionI = token_types.BinaryExpressionI

// the LeftName intent is before RightName intent

type Dependency = parser.Dependency
type DependenciesInfo = parser.DependenciesInfo
type RawIntentsI = parser.RawIntentsI
type TxIntents = parser.TxIntents
type MerkleProofProposal = parser.MerkleProofProposal

type RawDependency = lexer.RawDependency
type RawDependencies = lexer.RawDependencies
type RawDependenciesI = lexer.RawDependenciesI
type RawDependencyI = lexer.RawDependencyI

type RootIntents = lexer.RootIntents
type InvokeIntent = lexer.InvokeIntent
type PaymentIntent = lexer.PaymentIntent
type IfIntent = lexer.IfIntent
type LoopIntent = lexer.LoopIntent
type ContractInvokeMeta = lexer.ContractInvokeMeta
