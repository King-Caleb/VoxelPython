// Code generated from voxelpython.g4 by ANTLR 4.13.2. DO NOT EDIT.

package voxelpython_parser // voxelpython

import "github.com/antlr4-go/antlr/v4"


// voxelpythonListener is a complete listener for a parse tree produced by voxelpythonParser.
type voxelpythonListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)	

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterVarDecl is called when entering the varDecl production.
	EnterVarDecl(c *VarDeclContext)

	// EnterIfStatement is called when entering the ifStatement production.
	EnterIfStatement(c *IfStatementContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterCommand is called when entering the command production.
	EnterCommand(c *CommandContext)

	// EnterArgument is called when entering the argument production.
	EnterArgument(c *ArgumentContext)

	// EnterNumber is called when entering the Number production.
	EnterNumber(c *NumberContext)

	// EnterTrueLiteral is called when entering the TrueLiteral production.
	EnterTrueLiteral(c *TrueLiteralContext)

	// EnterMulDiv is called when entering the MulDiv production.
	EnterMulDiv(c *MulDivContext)

	// EnterAddSub is called when entering the AddSub production.
	EnterAddSub(c *AddSubContext)

	// EnterParens is called when entering the Parens production.
	EnterParens(c *ParensContext)

	// EnterVar is called when entering the Var production.
	EnterVar(c *VarContext)

	// EnterFalseLiteral is called when entering the FalseLiteral production.
	EnterFalseLiteral(c *FalseLiteralContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitVarDecl is called when exiting the varDecl production.
	ExitVarDecl(c *VarDeclContext)

	// ExitIfStatement is called when exiting the ifStatement production.
	ExitIfStatement(c *IfStatementContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitCommand is called when exiting the command production.
	ExitCommand(c *CommandContext)

	// ExitArgument is called when exiting the argument production.
	ExitArgument(c *ArgumentContext)

	// ExitNumber is called when exiting the Number production.
	ExitNumber(c *NumberContext)

	// ExitTrueLiteral is called when exiting the TrueLiteral production.
	ExitTrueLiteral(c *TrueLiteralContext)

	// ExitMulDiv is called when exiting the MulDiv production.
	ExitMulDiv(c *MulDivContext)

	// ExitAddSub is called when exiting the AddSub production.
	ExitAddSub(c *AddSubContext)

	// ExitParens is called when exiting the Parens production.
	ExitParens(c *ParensContext)

	// ExitVar is called when exiting the Var production.
	ExitVar(c *VarContext)

	// ExitFalseLiteral is called when exiting the FalseLiteral production.
	ExitFalseLiteral(c *FalseLiteralContext)
}
