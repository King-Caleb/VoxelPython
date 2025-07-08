// Code generated from voxelpython.g4 by ANTLR 4.13.2. DO NOT EDIT.

package voxelpython_parser // voxelpython

import "github.com/antlr4-go/antlr/v4"

// BasevoxelpythonListener is a complete listener for a parse tree produced by voxelpythonParser.
type BasevoxelpythonListener struct{}

var _ voxelpythonListener = &BasevoxelpythonListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasevoxelpythonListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasevoxelpythonListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasevoxelpythonListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasevoxelpythonListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BasevoxelpythonListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BasevoxelpythonListener) ExitProgram(ctx *ProgramContext) {}

// EnterStatement is called when production statement is entered.
func (s *BasevoxelpythonListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BasevoxelpythonListener) ExitStatement(ctx *StatementContext) {}

// EnterVarDecl is called when production varDecl is entered.
func (s *BasevoxelpythonListener) EnterVarDecl(ctx *VarDeclContext) {}

// ExitVarDecl is called when production varDecl is exited.
func (s *BasevoxelpythonListener) ExitVarDecl(ctx *VarDeclContext) {}

// EnterIfStatement is called when production ifStatement is entered.
func (s *BasevoxelpythonListener) EnterIfStatement(ctx *IfStatementContext) {}

// ExitIfStatement is called when production ifStatement is exited.
func (s *BasevoxelpythonListener) ExitIfStatement(ctx *IfStatementContext) {}

// EnterBlock is called when production block is entered.
func (s *BasevoxelpythonListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BasevoxelpythonListener) ExitBlock(ctx *BlockContext) {}

// EnterCommand is called when production command is entered.
func (s *BasevoxelpythonListener) EnterCommand(ctx *CommandContext) {}

// ExitCommand is called when production command is exited.
func (s *BasevoxelpythonListener) ExitCommand(ctx *CommandContext) {}

// EnterArgument is called when production argument is entered.
func (s *BasevoxelpythonListener) EnterArgument(ctx *ArgumentContext) {}

// ExitArgument is called when production argument is exited.
func (s *BasevoxelpythonListener) ExitArgument(ctx *ArgumentContext) {}

// EnterNumber is called when production Number is entered.
func (s *BasevoxelpythonListener) EnterNumber(ctx *NumberContext) {}

// ExitNumber is called when production Number is exited.
func (s *BasevoxelpythonListener) ExitNumber(ctx *NumberContext) {}

// EnterTrueLiteral is called when production TrueLiteral is entered.
func (s *BasevoxelpythonListener) EnterTrueLiteral(ctx *TrueLiteralContext) {}

// ExitTrueLiteral is called when production TrueLiteral is exited.
func (s *BasevoxelpythonListener) ExitTrueLiteral(ctx *TrueLiteralContext) {}

// EnterMulDiv is called when production MulDiv is entered.
func (s *BasevoxelpythonListener) EnterMulDiv(ctx *MulDivContext) {}

// ExitMulDiv is called when production MulDiv is exited.
func (s *BasevoxelpythonListener) ExitMulDiv(ctx *MulDivContext) {}

// EnterAddSub is called when production AddSub is entered.
func (s *BasevoxelpythonListener) EnterAddSub(ctx *AddSubContext) {}

// ExitAddSub is called when production AddSub is exited.
func (s *BasevoxelpythonListener) ExitAddSub(ctx *AddSubContext) {}

// EnterParens is called when production Parens is entered.
func (s *BasevoxelpythonListener) EnterParens(ctx *ParensContext) {}

// ExitParens is called when production Parens is exited.
func (s *BasevoxelpythonListener) ExitParens(ctx *ParensContext) {}

// EnterVar is called when production Var is entered.
func (s *BasevoxelpythonListener) EnterVar(ctx *VarContext) {}

// ExitVar is called when production Var is exited.
func (s *BasevoxelpythonListener) ExitVar(ctx *VarContext) {}

// EnterFalseLiteral is called when production FalseLiteral is entered.
func (s *BasevoxelpythonListener) EnterFalseLiteral(ctx *FalseLiteralContext) {}

// ExitFalseLiteral is called when production FalseLiteral is exited.
func (s *BasevoxelpythonListener) ExitFalseLiteral(ctx *FalseLiteralContext) {}
