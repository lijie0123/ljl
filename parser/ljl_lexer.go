// Code generated from /Volumes/Code/Workspace/lijie0123/ljl/Ljl.g4 by ANTLR 4.9.1. DO NOT EDIT.

package parser
import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)
// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter


var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 38, 222, 
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 
	4, 34, 9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 3, 2, 3, 2, 3, 
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 
	5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 
	9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 12, 3, 12, 3, 13, 
	3, 13, 3, 13, 3, 14, 3, 14, 3, 15, 3, 15, 3, 16, 3, 16, 3, 17, 3, 17, 3, 
	18, 3, 18, 3, 19, 3, 19, 3, 20, 3, 20, 3, 21, 3, 21, 3, 21, 3, 22, 3, 22, 
	3, 22, 3, 23, 3, 23, 3, 23, 3, 24, 3, 24, 3, 24, 3, 25, 3, 25, 3, 26, 3, 
	26, 3, 26, 3, 27, 3, 27, 3, 27, 3, 28, 3, 28, 3, 28, 3, 29, 3, 29, 3, 29, 
	3, 29, 3, 29, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 31, 3, 31, 3, 32, 5, 
	32, 165, 10, 32, 3, 32, 6, 32, 168, 10, 32, 13, 32, 14, 32, 169, 3, 33, 
	5, 33, 173, 10, 33, 3, 33, 6, 33, 176, 10, 33, 13, 33, 14, 33, 177, 3, 
	33, 3, 33, 6, 33, 182, 10, 33, 13, 33, 14, 33, 183, 3, 34, 3, 34, 7, 34, 
	188, 10, 34, 12, 34, 14, 34, 191, 11, 34, 3, 34, 3, 34, 3, 35, 3, 35, 7, 
	35, 197, 10, 35, 12, 35, 14, 35, 200, 11, 35, 3, 36, 3, 36, 3, 36, 3, 36, 
	7, 36, 206, 10, 36, 12, 36, 14, 36, 209, 11, 36, 3, 36, 3, 36, 3, 36, 3, 
	36, 3, 36, 3, 37, 6, 37, 217, 10, 37, 13, 37, 14, 37, 218, 3, 37, 3, 37, 
	4, 189, 207, 2, 38, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 
	19, 11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17, 33, 18, 35, 19, 
	37, 20, 39, 21, 41, 22, 43, 23, 45, 24, 47, 25, 49, 26, 51, 27, 53, 28, 
	55, 29, 57, 30, 59, 31, 61, 32, 63, 33, 65, 34, 67, 35, 69, 36, 71, 37, 
	73, 38, 3, 2, 6, 3, 2, 50, 59, 5, 2, 67, 92, 97, 97, 99, 124, 6, 2, 50, 
	59, 67, 92, 97, 97, 99, 124, 5, 2, 11, 12, 15, 15, 34, 34, 2, 230, 2, 3, 
	3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 
	3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 
	19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 
	2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 
	2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 
	2, 2, 2, 43, 3, 2, 2, 2, 2, 45, 3, 2, 2, 2, 2, 47, 3, 2, 2, 2, 2, 49, 3, 
	2, 2, 2, 2, 51, 3, 2, 2, 2, 2, 53, 3, 2, 2, 2, 2, 55, 3, 2, 2, 2, 2, 57, 
	3, 2, 2, 2, 2, 59, 3, 2, 2, 2, 2, 61, 3, 2, 2, 2, 2, 63, 3, 2, 2, 2, 2, 
	65, 3, 2, 2, 2, 2, 67, 3, 2, 2, 2, 2, 69, 3, 2, 2, 2, 2, 71, 3, 2, 2, 2, 
	2, 73, 3, 2, 2, 2, 3, 75, 3, 2, 2, 2, 5, 77, 3, 2, 2, 2, 7, 82, 3, 2, 2, 
	2, 9, 89, 3, 2, 2, 2, 11, 91, 3, 2, 2, 2, 13, 93, 3, 2, 2, 2, 15, 95, 3, 
	2, 2, 2, 17, 97, 3, 2, 2, 2, 19, 104, 3, 2, 2, 2, 21, 107, 3, 2, 2, 2, 
	23, 109, 3, 2, 2, 2, 25, 111, 3, 2, 2, 2, 27, 114, 3, 2, 2, 2, 29, 116, 
	3, 2, 2, 2, 31, 118, 3, 2, 2, 2, 33, 120, 3, 2, 2, 2, 35, 122, 3, 2, 2, 
	2, 37, 124, 3, 2, 2, 2, 39, 126, 3, 2, 2, 2, 41, 128, 3, 2, 2, 2, 43, 131, 
	3, 2, 2, 2, 45, 134, 3, 2, 2, 2, 47, 137, 3, 2, 2, 2, 49, 140, 3, 2, 2, 
	2, 51, 142, 3, 2, 2, 2, 53, 145, 3, 2, 2, 2, 55, 148, 3, 2, 2, 2, 57, 151, 
	3, 2, 2, 2, 59, 156, 3, 2, 2, 2, 61, 161, 3, 2, 2, 2, 63, 164, 3, 2, 2, 
	2, 65, 172, 3, 2, 2, 2, 67, 185, 3, 2, 2, 2, 69, 194, 3, 2, 2, 2, 71, 201, 
	3, 2, 2, 2, 73, 216, 3, 2, 2, 2, 75, 76, 7, 61, 2, 2, 76, 4, 3, 2, 2, 2, 
	77, 78, 7, 118, 2, 2, 78, 79, 7, 123, 2, 2, 79, 80, 7, 114, 2, 2, 80, 81, 
	7, 103, 2, 2, 81, 6, 3, 2, 2, 2, 82, 83, 7, 117, 2, 2, 83, 84, 7, 118, 
	2, 2, 84, 85, 7, 116, 2, 2, 85, 86, 7, 119, 2, 2, 86, 87, 7, 101, 2, 2, 
	87, 88, 7, 118, 2, 2, 88, 8, 3, 2, 2, 2, 89, 90, 7, 125, 2, 2, 90, 10, 
	3, 2, 2, 2, 91, 92, 7, 127, 2, 2, 92, 12, 3, 2, 2, 2, 93, 94, 7, 63, 2, 
	2, 94, 14, 3, 2, 2, 2, 95, 96, 7, 60, 2, 2, 96, 16, 3, 2, 2, 2, 97, 98, 
	7, 103, 2, 2, 98, 99, 7, 122, 2, 2, 99, 100, 7, 118, 2, 2, 100, 101, 7, 
	103, 2, 2, 101, 102, 7, 116, 2, 2, 102, 103, 7, 112, 2, 2, 103, 18, 3, 
	2, 2, 2, 104, 105, 7, 104, 2, 2, 105, 106, 7, 112, 2, 2, 106, 20, 3, 2, 
	2, 2, 107, 108, 7, 42, 2, 2, 108, 22, 3, 2, 2, 2, 109, 110, 7, 43, 2, 2, 
	110, 24, 3, 2, 2, 2, 111, 112, 7, 47, 2, 2, 112, 113, 7, 64, 2, 2, 113, 
	26, 3, 2, 2, 2, 114, 115, 7, 46, 2, 2, 115, 28, 3, 2, 2, 2, 116, 117, 7, 
	44, 2, 2, 117, 30, 3, 2, 2, 2, 118, 119, 7, 49, 2, 2, 119, 32, 3, 2, 2, 
	2, 120, 121, 7, 45, 2, 2, 121, 34, 3, 2, 2, 2, 122, 123, 7, 47, 2, 2, 123, 
	36, 3, 2, 2, 2, 124, 125, 7, 64, 2, 2, 125, 38, 3, 2, 2, 2, 126, 127, 7, 
	62, 2, 2, 127, 40, 3, 2, 2, 2, 128, 129, 7, 64, 2, 2, 129, 130, 7, 63, 
	2, 2, 130, 42, 3, 2, 2, 2, 131, 132, 7, 62, 2, 2, 132, 133, 7, 63, 2, 2, 
	133, 44, 3, 2, 2, 2, 134, 135, 7, 63, 2, 2, 135, 136, 7, 63, 2, 2, 136, 
	46, 3, 2, 2, 2, 137, 138, 7, 35, 2, 2, 138, 139, 7, 63, 2, 2, 139, 48, 
	3, 2, 2, 2, 140, 141, 7, 35, 2, 2, 141, 50, 3, 2, 2, 2, 142, 143, 7, 40, 
	2, 2, 143, 144, 7, 40, 2, 2, 144, 52, 3, 2, 2, 2, 145, 146, 7, 126, 2, 
	2, 146, 147, 7, 126, 2, 2, 147, 54, 3, 2, 2, 2, 148, 149, 7, 107, 2, 2, 
	149, 150, 7, 104, 2, 2, 150, 56, 3, 2, 2, 2, 151, 152, 7, 118, 2, 2, 152, 
	153, 7, 106, 2, 2, 153, 154, 7, 103, 2, 2, 154, 155, 7, 112, 2, 2, 155, 
	58, 3, 2, 2, 2, 156, 157, 7, 103, 2, 2, 157, 158, 7, 110, 2, 2, 158, 159, 
	7, 117, 2, 2, 159, 160, 7, 103, 2, 2, 160, 60, 3, 2, 2, 2, 161, 162, 7, 
	48, 2, 2, 162, 62, 3, 2, 2, 2, 163, 165, 7, 47, 2, 2, 164, 163, 3, 2, 2, 
	2, 164, 165, 3, 2, 2, 2, 165, 167, 3, 2, 2, 2, 166, 168, 9, 2, 2, 2, 167, 
	166, 3, 2, 2, 2, 168, 169, 3, 2, 2, 2, 169, 167, 3, 2, 2, 2, 169, 170, 
	3, 2, 2, 2, 170, 64, 3, 2, 2, 2, 171, 173, 7, 47, 2, 2, 172, 171, 3, 2, 
	2, 2, 172, 173, 3, 2, 2, 2, 173, 175, 3, 2, 2, 2, 174, 176, 9, 2, 2, 2, 
	175, 174, 3, 2, 2, 2, 176, 177, 3, 2, 2, 2, 177, 175, 3, 2, 2, 2, 177, 
	178, 3, 2, 2, 2, 178, 179, 3, 2, 2, 2, 179, 181, 7, 48, 2, 2, 180, 182, 
	9, 2, 2, 2, 181, 180, 3, 2, 2, 2, 182, 183, 3, 2, 2, 2, 183, 181, 3, 2, 
	2, 2, 183, 184, 3, 2, 2, 2, 184, 66, 3, 2, 2, 2, 185, 189, 7, 36, 2, 2, 
	186, 188, 11, 2, 2, 2, 187, 186, 3, 2, 2, 2, 188, 191, 3, 2, 2, 2, 189, 
	190, 3, 2, 2, 2, 189, 187, 3, 2, 2, 2, 190, 192, 3, 2, 2, 2, 191, 189, 
	3, 2, 2, 2, 192, 193, 7, 36, 2, 2, 193, 68, 3, 2, 2, 2, 194, 198, 9, 3, 
	2, 2, 195, 197, 9, 4, 2, 2, 196, 195, 3, 2, 2, 2, 197, 200, 3, 2, 2, 2, 
	198, 196, 3, 2, 2, 2, 198, 199, 3, 2, 2, 2, 199, 70, 3, 2, 2, 2, 200, 198, 
	3, 2, 2, 2, 201, 202, 7, 49, 2, 2, 202, 203, 7, 44, 2, 2, 203, 207, 3, 
	2, 2, 2, 204, 206, 11, 2, 2, 2, 205, 204, 3, 2, 2, 2, 206, 209, 3, 2, 2, 
	2, 207, 208, 3, 2, 2, 2, 207, 205, 3, 2, 2, 2, 208, 210, 3, 2, 2, 2, 209, 
	207, 3, 2, 2, 2, 210, 211, 7, 44, 2, 2, 211, 212, 7, 49, 2, 2, 212, 213, 
	3, 2, 2, 2, 213, 214, 8, 36, 2, 2, 214, 72, 3, 2, 2, 2, 215, 217, 9, 5, 
	2, 2, 216, 215, 3, 2, 2, 2, 217, 218, 3, 2, 2, 2, 218, 216, 3, 2, 2, 2, 
	218, 219, 3, 2, 2, 2, 219, 220, 3, 2, 2, 2, 220, 221, 8, 37, 2, 2, 221, 
	74, 3, 2, 2, 2, 12, 2, 164, 169, 172, 177, 183, 189, 198, 207, 218, 3, 
	8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "';'", "'type'", "'struct'", "'{'", "'}'", "'='", "':'", "'extern'", 
	"'fn'", "'('", "')'", "'->'", "','", "'*'", "'/'", "'+'", "'-'", "'>'", 
	"'<'", "'>='", "'<='", "'=='", "'!='", "'!'", "'&&'", "'||'", "'if'", "'then'", 
	"'else'", "'.'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", 
	"", "", "", "", "", "", "", "", "", "", "", "", "", "INT", "FLOAT", "STRING", 
	"ID", "COMMENT", "WS",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8", 
	"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16", 
	"T__17", "T__18", "T__19", "T__20", "T__21", "T__22", "T__23", "T__24", 
	"T__25", "T__26", "T__27", "T__28", "T__29", "INT", "FLOAT", "STRING", 
	"ID", "COMMENT", "WS",
}
type LjlLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames []string
	// TODO: EOF string
}

// NewLjlLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *LjlLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewLjlLexer(input antlr.CharStream) *LjlLexer {
	l := new(LjlLexer)
	lexerDeserializer := antlr.NewATNDeserializer(nil)
	lexerAtn := lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)
	lexerDecisionToDFA := make([]*antlr.DFA, len(lexerAtn.DecisionToState))
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Ljl.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// LjlLexer tokens.
const (
	LjlLexerT__0 = 1
	LjlLexerT__1 = 2
	LjlLexerT__2 = 3
	LjlLexerT__3 = 4
	LjlLexerT__4 = 5
	LjlLexerT__5 = 6
	LjlLexerT__6 = 7
	LjlLexerT__7 = 8
	LjlLexerT__8 = 9
	LjlLexerT__9 = 10
	LjlLexerT__10 = 11
	LjlLexerT__11 = 12
	LjlLexerT__12 = 13
	LjlLexerT__13 = 14
	LjlLexerT__14 = 15
	LjlLexerT__15 = 16
	LjlLexerT__16 = 17
	LjlLexerT__17 = 18
	LjlLexerT__18 = 19
	LjlLexerT__19 = 20
	LjlLexerT__20 = 21
	LjlLexerT__21 = 22
	LjlLexerT__22 = 23
	LjlLexerT__23 = 24
	LjlLexerT__24 = 25
	LjlLexerT__25 = 26
	LjlLexerT__26 = 27
	LjlLexerT__27 = 28
	LjlLexerT__28 = 29
	LjlLexerT__29 = 30
	LjlLexerINT = 31
	LjlLexerFLOAT = 32
	LjlLexerSTRING = 33
	LjlLexerID = 34
	LjlLexerCOMMENT = 35
	LjlLexerWS = 36
)
