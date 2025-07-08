// Code generated from voxelpython.g4 by ANTLR 4.13.2. DO NOT EDIT.

package voxelpython_parser

import (
	"fmt"
  	"sync"
	"unicode"
	"github.com/antlr4-go/antlr/v4"
)
// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter


type voxelpythonLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames []string
	// TODO: EOF string
}

var VoxelpythonLexerLexerStaticData struct {
  once                   sync.Once
  serializedATN          []int32
  ChannelNames           []string
  ModeNames              []string
  LiteralNames           []string
  SymbolicNames          []string
  RuleNames              []string
  PredictionContextCache *antlr.PredictionContextCache
  atn                    *antlr.ATN
  decisionToDFA          []*antlr.DFA
}

func voxelpythonlexerLexerInit() {
  staticData := &VoxelpythonLexerLexerStaticData
  staticData.ChannelNames = []string{
    "DEFAULT_TOKEN_CHANNEL", "HIDDEN",
  }
  staticData.ModeNames = []string{
    "DEFAULT_MODE",
  }
  staticData.LiteralNames = []string{
    "", "'var'", "'='", "'if'", "':'", "'elif'", "'else'", "'cmd'", "'.'", 
    "'*'", "'/'", "'+'", "'-'", "'('", "')'", "'true'", "'false'",
  }
  staticData.SymbolicNames = []string{
    "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", 
    "IDENTIFIER", "NUMBER", "STRING", "NEWLINE", "INDENT", "DEDENT", "WS", 
    "COMMENT",
  }
  staticData.RuleNames = []string{
    "T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8", 
    "T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "IDENTIFIER", 
    "NUMBER", "STRING", "NEWLINE", "INDENT", "DEDENT", "WS", "COMMENT",
  }
  staticData.PredictionContextCache = antlr.NewPredictionContextCache()
  staticData.serializedATN = []int32{
	4, 0, 24, 158, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 
	4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 
	10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 
	7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 
	20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 1, 0, 1, 0, 1, 0, 1, 0, 1, 
	1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 
	5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 
	8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13, 
	1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 
	15, 1, 16, 1, 16, 5, 16, 102, 8, 16, 10, 16, 12, 16, 105, 9, 16, 1, 17, 
	4, 17, 108, 8, 17, 11, 17, 12, 17, 109, 1, 17, 1, 17, 4, 17, 114, 8, 17, 
	11, 17, 12, 17, 115, 3, 17, 118, 8, 17, 1, 18, 1, 18, 1, 18, 1, 18, 5, 
	18, 124, 8, 18, 10, 18, 12, 18, 127, 9, 18, 1, 18, 1, 18, 1, 19, 3, 19, 
	132, 8, 19, 1, 19, 4, 19, 135, 8, 19, 11, 19, 12, 19, 136, 1, 20, 1, 20, 
	1, 21, 1, 21, 1, 22, 4, 22, 144, 8, 22, 11, 22, 12, 22, 145, 1, 22, 1, 
	22, 1, 23, 1, 23, 5, 23, 152, 8, 23, 10, 23, 12, 23, 155, 9, 23, 1, 23, 
	1, 23, 0, 0, 24, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 
	9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 
	18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 1, 0, 6, 3, 0, 65, 
	90, 95, 95, 97, 122, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 1, 0, 48, 57, 
	2, 0, 34, 34, 92, 92, 2, 0, 9, 9, 32, 32, 2, 0, 10, 10, 13, 13, 167, 0, 
	1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 
	9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 
	0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 
	0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 
	0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 
	0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 
	1, 0, 0, 0, 1, 49, 1, 0, 0, 0, 3, 53, 1, 0, 0, 0, 5, 55, 1, 0, 0, 0, 7, 
	58, 1, 0, 0, 0, 9, 60, 1, 0, 0, 0, 11, 65, 1, 0, 0, 0, 13, 70, 1, 0, 0, 
	0, 15, 74, 1, 0, 0, 0, 17, 76, 1, 0, 0, 0, 19, 78, 1, 0, 0, 0, 21, 80, 
	1, 0, 0, 0, 23, 82, 1, 0, 0, 0, 25, 84, 1, 0, 0, 0, 27, 86, 1, 0, 0, 0, 
	29, 88, 1, 0, 0, 0, 31, 93, 1, 0, 0, 0, 33, 99, 1, 0, 0, 0, 35, 107, 1, 
	0, 0, 0, 37, 119, 1, 0, 0, 0, 39, 134, 1, 0, 0, 0, 41, 138, 1, 0, 0, 0, 
	43, 140, 1, 0, 0, 0, 45, 143, 1, 0, 0, 0, 47, 149, 1, 0, 0, 0, 49, 50, 
	5, 118, 0, 0, 50, 51, 5, 97, 0, 0, 51, 52, 5, 114, 0, 0, 52, 2, 1, 0, 0, 
	0, 53, 54, 5, 61, 0, 0, 54, 4, 1, 0, 0, 0, 55, 56, 5, 105, 0, 0, 56, 57, 
	5, 102, 0, 0, 57, 6, 1, 0, 0, 0, 58, 59, 5, 58, 0, 0, 59, 8, 1, 0, 0, 0, 
	60, 61, 5, 101, 0, 0, 61, 62, 5, 108, 0, 0, 62, 63, 5, 105, 0, 0, 63, 64, 
	5, 102, 0, 0, 64, 10, 1, 0, 0, 0, 65, 66, 5, 101, 0, 0, 66, 67, 5, 108, 
	0, 0, 67, 68, 5, 115, 0, 0, 68, 69, 5, 101, 0, 0, 69, 12, 1, 0, 0, 0, 70, 
	71, 5, 99, 0, 0, 71, 72, 5, 109, 0, 0, 72, 73, 5, 100, 0, 0, 73, 14, 1, 
	0, 0, 0, 74, 75, 5, 46, 0, 0, 75, 16, 1, 0, 0, 0, 76, 77, 5, 42, 0, 0, 
	77, 18, 1, 0, 0, 0, 78, 79, 5, 47, 0, 0, 79, 20, 1, 0, 0, 0, 80, 81, 5, 
	43, 0, 0, 81, 22, 1, 0, 0, 0, 82, 83, 5, 45, 0, 0, 83, 24, 1, 0, 0, 0, 
	84, 85, 5, 40, 0, 0, 85, 26, 1, 0, 0, 0, 86, 87, 5, 41, 0, 0, 87, 28, 1, 
	0, 0, 0, 88, 89, 5, 116, 0, 0, 89, 90, 5, 114, 0, 0, 90, 91, 5, 117, 0, 
	0, 91, 92, 5, 101, 0, 0, 92, 30, 1, 0, 0, 0, 93, 94, 5, 102, 0, 0, 94, 
	95, 5, 97, 0, 0, 95, 96, 5, 108, 0, 0, 96, 97, 5, 115, 0, 0, 97, 98, 5, 
	101, 0, 0, 98, 32, 1, 0, 0, 0, 99, 103, 7, 0, 0, 0, 100, 102, 7, 1, 0, 
	0, 101, 100, 1, 0, 0, 0, 102, 105, 1, 0, 0, 0, 103, 101, 1, 0, 0, 0, 103, 
	104, 1, 0, 0, 0, 104, 34, 1, 0, 0, 0, 105, 103, 1, 0, 0, 0, 106, 108, 7, 
	2, 0, 0, 107, 106, 1, 0, 0, 0, 108, 109, 1, 0, 0, 0, 109, 107, 1, 0, 0, 
	0, 109, 110, 1, 0, 0, 0, 110, 117, 1, 0, 0, 0, 111, 113, 5, 46, 0, 0, 112, 
	114, 7, 2, 0, 0, 113, 112, 1, 0, 0, 0, 114, 115, 1, 0, 0, 0, 115, 113, 
	1, 0, 0, 0, 115, 116, 1, 0, 0, 0, 116, 118, 1, 0, 0, 0, 117, 111, 1, 0, 
	0, 0, 117, 118, 1, 0, 0, 0, 118, 36, 1, 0, 0, 0, 119, 125, 5, 34, 0, 0, 
	120, 124, 8, 3, 0, 0, 121, 122, 5, 92, 0, 0, 122, 124, 9, 0, 0, 0, 123, 
	120, 1, 0, 0, 0, 123, 121, 1, 0, 0, 0, 124, 127, 1, 0, 0, 0, 125, 123, 
	1, 0, 0, 0, 125, 126, 1, 0, 0, 0, 126, 128, 1, 0, 0, 0, 127, 125, 1, 0, 
	0, 0, 128, 129, 5, 34, 0, 0, 129, 38, 1, 0, 0, 0, 130, 132, 5, 13, 0, 0, 
	131, 130, 1, 0, 0, 0, 131, 132, 1, 0, 0, 0, 132, 133, 1, 0, 0, 0, 133, 
	135, 5, 10, 0, 0, 134, 131, 1, 0, 0, 0, 135, 136, 1, 0, 0, 0, 136, 134, 
	1, 0, 0, 0, 136, 137, 1, 0, 0, 0, 137, 40, 1, 0, 0, 0, 138, 139, 1, 0, 
	0, 0, 139, 42, 1, 0, 0, 0, 140, 141, 1, 0, 0, 0, 141, 44, 1, 0, 0, 0, 142, 
	144, 7, 4, 0, 0, 143, 142, 1, 0, 0, 0, 144, 145, 1, 0, 0, 0, 145, 143, 
	1, 0, 0, 0, 145, 146, 1, 0, 0, 0, 146, 147, 1, 0, 0, 0, 147, 148, 6, 22, 
	0, 0, 148, 46, 1, 0, 0, 0, 149, 153, 5, 35, 0, 0, 150, 152, 8, 5, 0, 0, 
	151, 150, 1, 0, 0, 0, 152, 155, 1, 0, 0, 0, 153, 151, 1, 0, 0, 0, 153, 
	154, 1, 0, 0, 0, 154, 156, 1, 0, 0, 0, 155, 153, 1, 0, 0, 0, 156, 157, 
	6, 23, 0, 0, 157, 48, 1, 0, 0, 0, 11, 0, 103, 109, 115, 117, 123, 125, 
	131, 136, 145, 153, 1, 6, 0, 0,
}
  deserializer := antlr.NewATNDeserializer(nil)
  staticData.atn = deserializer.Deserialize(staticData.serializedATN)
  atn := staticData.atn
  staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
  decisionToDFA := staticData.decisionToDFA
  for index, state := range atn.DecisionToState {
    decisionToDFA[index] = antlr.NewDFA(state, index)
  }
}

// voxelpythonLexerInit initializes any static state used to implement voxelpythonLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewvoxelpythonLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func VoxelpythonLexerInit() {
  staticData := &VoxelpythonLexerLexerStaticData
  staticData.once.Do(voxelpythonlexerLexerInit)
}

// NewvoxelpythonLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewvoxelpythonLexer(input antlr.CharStream) *voxelpythonLexer {
  VoxelpythonLexerInit()
	l := new(voxelpythonLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
  staticData := &VoxelpythonLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "voxelpython.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// voxelpythonLexer tokens.
const (
	voxelpythonLexerT__0 = 1
	voxelpythonLexerT__1 = 2
	voxelpythonLexerT__2 = 3
	voxelpythonLexerT__3 = 4
	voxelpythonLexerT__4 = 5
	voxelpythonLexerT__5 = 6
	voxelpythonLexerT__6 = 7
	voxelpythonLexerT__7 = 8
	voxelpythonLexerT__8 = 9
	voxelpythonLexerT__9 = 10
	voxelpythonLexerT__10 = 11
	voxelpythonLexerT__11 = 12
	voxelpythonLexerT__12 = 13
	voxelpythonLexerT__13 = 14
	voxelpythonLexerT__14 = 15
	voxelpythonLexerT__15 = 16
	voxelpythonLexerIDENTIFIER = 17
	voxelpythonLexerNUMBER = 18
	voxelpythonLexerSTRING = 19
	voxelpythonLexerNEWLINE = 20
	voxelpythonLexerINDENT = 21
	voxelpythonLexerDEDENT = 22
	voxelpythonLexerWS = 23
	voxelpythonLexerCOMMENT = 24
)

