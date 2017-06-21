//line warc.y:3
package warc

import __yyfmt__ "fmt"

//line warc.y:3
import (
	"bytes"
	"fmt"
)

func setRecord(yylex interface{}, r *record) {
	yylex.(*Tokenizer).Record = r
}

func addRecord(yylex interface{}, r *record) {
	yylex.(*Tokenizer).Records = append(yylex.(*Tokenizer).Records, r)
	yylex.(*Tokenizer).Record = NewRecord()
}

func getLatestHeader(yylex interface{}) *Header {
	return yylex.(*Tokenizer).Record.Header
}

func forceEOF(yylex interface{}) {
	yylex.(*Tokenizer).ForceEOF = true
}

//line warc.y:29
type yySymType struct {
	yys        int
	empty      struct{}
	records    []*record
	record     *record
	header     *Header
	token      string
	bytes      []byte
	recordType RecordType
}

const LEX_ERROR = 57346
const WARCINFO = 57347
const RESPONSE = 57348
const RESOURCE = 57349
const REQUEST = 57350
const METADATA = 57351
const REVISIT = 57352
const CONVERSION = 57353
const CONTINUATION = 57354
const WARC_RECORD_ID = 57355
const WARC_DATE = 57356
const CONTENT_LENGTH = 57357
const CONTENT_TYPE = 57358
const WARC_CONCURRENT_TO = 57359
const WARC_BLOCK_DIGEST = 57360
const WARC_PAYLOAD_DIGEST = 57361
const WARC_IP_ADDRESS = 57362
const WARC_REFERS_TO = 57363
const WARC_TARGET_URI = 57364
const WARC_TRUNCATED = 57365
const WARC_WARCINFO_ID = 57366
const WARC_FILENAME = 57367
const WARC_PROFILE = 57368
const WARC_IDENTIFIED_PAYLOAD_TYPE = 57369
const WARC_SEGMENT_ORIGIN_ID = 57370
const WARC_SEGMENT_NUMBER = 57371
const WARC_SEGMENT_TOTAL_LENGTH = 57372
const WARC_TYPE = 57373
const WARC_VERSION = 57374
const FIELD_KEY = 57375
const FIELD_VALUE = 57376
const BLOCK = 57377

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"LEX_ERROR",
	"WARCINFO",
	"RESPONSE",
	"RESOURCE",
	"REQUEST",
	"METADATA",
	"REVISIT",
	"CONVERSION",
	"CONTINUATION",
	"WARC_RECORD_ID",
	"WARC_DATE",
	"CONTENT_LENGTH",
	"CONTENT_TYPE",
	"WARC_CONCURRENT_TO",
	"WARC_BLOCK_DIGEST",
	"WARC_PAYLOAD_DIGEST",
	"WARC_IP_ADDRESS",
	"WARC_REFERS_TO",
	"WARC_TARGET_URI",
	"WARC_TRUNCATED",
	"WARC_WARCINFO_ID",
	"WARC_FILENAME",
	"WARC_PROFILE",
	"WARC_IDENTIFIED_PAYLOAD_TYPE",
	"WARC_SEGMENT_ORIGIN_ID",
	"WARC_SEGMENT_NUMBER",
	"WARC_SEGMENT_TOTAL_LENGTH",
	"WARC_TYPE",
	"WARC_VERSION",
	"FIELD_KEY",
	"FIELD_VALUE",
	"BLOCK",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 77

var yyAct = [...]int{

	17, 11, 7, 8, 10, 9, 15, 14, 18, 22,
	23, 24, 12, 16, 13, 19, 20, 21, 25, 56,
	26, 3, 27, 17, 11, 7, 8, 10, 9, 15,
	14, 18, 22, 23, 24, 12, 16, 13, 19, 20,
	21, 25, 46, 26, 47, 45, 44, 43, 42, 41,
	40, 39, 38, 37, 36, 35, 34, 33, 32, 31,
	30, 29, 55, 52, 51, 53, 50, 54, 48, 49,
	6, 2, 5, 4, 1, 0, 28,
}
var yyPact = [...]int{

	-11, -11, -1000, 10, -1000, -13, -1000, 27, 26, 25,
	24, 23, 22, 21, 20, 19, 18, 17, 16, 15,
	14, 13, 12, 11, 8, 57, -15, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000,
}
var yyPgo = [...]int{

	0, 74, 71, 72, 70, 44,
}
var yyR1 = [...]int{

	0, 1, 1, 2, 3, 3, 4, 4, 4, 4,
	4, 4, 4, 4, 4, 4, 4, 4, 4, 4,
	4, 4, 4, 4, 4, 4, 5, 5, 5, 5,
	5, 5, 5, 5,
}
var yyR2 = [...]int{

	0, 1, 2, 3, 1, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 1, 1, 1, 1,
	1, 1, 1, 1,
}
var yyChk = [...]int{

	-1000, -1, -2, 32, -2, -3, -4, 15, 16, 18,
	17, 14, 25, 27, 20, 19, 26, 13, 21, 28,
	29, 30, 22, 23, 24, 31, 33, 35, -4, 34,
	34, 34, 34, 34, 34, 34, 34, 34, 34, 34,
	34, 34, 34, 34, 34, 34, 34, -5, 11, 12,
	9, 7, 6, 8, 10, 5, 34,
}
var yyDef = [...]int{

	0, -2, 1, 0, 2, 0, 4, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 3, 5, 6,
	7, 8, 9, 10, 11, 12, 13, 14, 15, 16,
	17, 18, 19, 20, 21, 22, 23, 24, 26, 27,
	28, 29, 30, 31, 32, 33, 25,
}
var yyTok1 = [...]int{

	1,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35,
}
var yyTok3 = [...]int{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line warc.y:57
		{
			addRecord(yylex, yyDollar[1].record)
		}
	case 2:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line warc.y:61
		{
			fmt.Println("huh?")
			addRecord(yylex, yyDollar[2].record)
			// setRecord(yylex, $2)
			//$$ = append($1, $2)
		}
	case 3:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line warc.y:70
		{
			yyVAL.record = &record{Version: string(yyDollar[1].bytes), Header: yyDollar[2].header, Content: bytes.NewReader(yyDollar[3].bytes)}
		}
	case 4:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line warc.y:76
		{
			yyVAL.header = yyDollar[1].header
		}
	case 5:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line warc.y:80
		{
			yyVAL.header = yyDollar[2].header
		}
	case 6:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line warc.y:86
		{
			// TODO - convert to number
			h := getLatestHeader(yylex)
			h.ContentLength = string(yyDollar[2].bytes)
			yyVAL.header = h
		}
	case 7:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line warc.y:93
		{
			h := getLatestHeader(yylex)
			h.ContentType = string(yyDollar[2].bytes)
			yyVAL.header = h
		}
	case 8:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line warc.y:99
		{
			h := getLatestHeader(yylex)
			h.WARCBlockDigest = string(yyDollar[2].bytes)
			yyVAL.header = h
		}
	case 9:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line warc.y:105
		{
			h := getLatestHeader(yylex)
			h.WARCConcurrentTo = string(yyDollar[2].bytes)
			yyVAL.header = h
		}
	case 10:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line warc.y:111
		{
			h := getLatestHeader(yylex)
			h.WARCDate = string(yyDollar[2].bytes)
			yyVAL.header = h
		}
	case 11:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line warc.y:117
		{
			h := getLatestHeader(yylex)
			h.WARCFilename = string(yyDollar[2].bytes)
			yyVAL.header = h
		}
	case 12:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line warc.y:123
		{
			h := getLatestHeader(yylex)
			h.WARCIdentifiedPayloadType = string(yyDollar[2].bytes)
			yyVAL.header = h
		}
	case 13:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line warc.y:129
		{
			h := getLatestHeader(yylex)
			h.WARCIPAddress = string(yyDollar[2].bytes)
			yyVAL.header = h
		}
	case 14:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line warc.y:135
		{
			h := getLatestHeader(yylex)
			h.WARCPayloadDigest = string(yyDollar[2].bytes)
			yyVAL.header = h
		}
	case 15:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line warc.y:141
		{
			h := getLatestHeader(yylex)
			h.WARCProfile = string(yyDollar[2].bytes)
			yyVAL.header = h
		}
	case 16:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line warc.y:147
		{
			h := getLatestHeader(yylex)
			h.WARCRecordId = string(yyDollar[2].bytes)
			yyVAL.header = h
		}
	case 17:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line warc.y:153
		{
			h := getLatestHeader(yylex)
			h.WARCRefersTo = string(yyDollar[2].bytes)
			yyVAL.header = h
		}
	case 18:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line warc.y:159
		{
			h := getLatestHeader(yylex)
			h.WARCSegmentOriginID = string(yyDollar[2].bytes)
			yyVAL.header = h
		}
	case 19:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line warc.y:165
		{
			h := getLatestHeader(yylex)
			h.WARCSegmentNumber = string(yyDollar[2].bytes)
			yyVAL.header = h
		}
	case 20:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line warc.y:171
		{
			h := getLatestHeader(yylex)
			h.WARCSegmentTotalLength = string(yyDollar[2].bytes)
			yyVAL.header = h
		}
	case 21:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line warc.y:177
		{
			h := getLatestHeader(yylex)
			h.WARCTargetUri = string(yyDollar[2].bytes)
			yyVAL.header = h
		}
	case 22:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line warc.y:183
		{
			h := getLatestHeader(yylex)
			h.WARCTruncated = string(yyDollar[2].bytes)
			yyVAL.header = h
		}
	case 23:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line warc.y:189
		{
			h := getLatestHeader(yylex)
			h.WARCInfoId = string(yyDollar[2].bytes)
			yyVAL.header = h
		}
	case 24:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line warc.y:195
		{
			h := getLatestHeader(yylex)
			h.WARCType = yyDollar[2].recordType
			yyVAL.header = h
		}
	case 25:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line warc.y:201
		{
			h := getLatestHeader(yylex)
			h.CustomFields[string(yyDollar[1].bytes)] = string(yyDollar[2].bytes)
			yyVAL.header = h
		}
	case 26:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line warc.y:209
		{
			yyVAL.recordType = RecordTypeConversion
		}
	case 27:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line warc.y:213
		{
			yyVAL.recordType = RecordTypeContinuation
		}
	case 28:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line warc.y:217
		{
			yyVAL.recordType = RecordTypeMetadata
		}
	case 29:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line warc.y:221
		{
			yyVAL.recordType = RecordTypeResource
		}
	case 30:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line warc.y:225
		{
			yyVAL.recordType = RecordTypeResponse
		}
	case 31:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line warc.y:229
		{
			yyVAL.recordType = RecordTypeRequest
		}
	case 32:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line warc.y:233
		{
			yyVAL.recordType = RecordTypeRevisit
		}
	case 33:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line warc.y:237
		{
			yyVAL.recordType = RecordTypeWarcInfo
		}
	}
	goto yystack /* stack new state and value */
}
