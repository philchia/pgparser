//line sql.y:20
package sqlparser

import __yyfmt__ "fmt"

//line sql.y:20

import "bytes"

func SetParseTree(yylex interface{}, stmt Statement) {
	yylex.(*Tokenizer).ParseTree = stmt
}

func SetAllowComments(yylex interface{}, allow bool) {
	yylex.(*Tokenizer).AllowComments = allow
}

func ForceEOF(yylex interface{}) {
	yylex.(*Tokenizer).ForceEOF = true
}

var (
	SHARE        = []byte("share")
	MODE         = []byte("mode")
	IF_BYTES     = []byte("if")
	VALUES_BYTES = []byte("values")
)

//line sql.y:45
type yySymType struct {
	yys         int
	empty       struct{}
	statement   Statement
	selStmt     SelectStatement
	byt         byte
	bytes       []byte
	bytes2      [][]byte
	str         string
	selectExprs SelectExprs
	selectExpr  SelectExpr
	columns     Columns
	colName     *ColName
	tableExprs  TableExprs
	tableExpr   TableExpr
	smTableExpr SimpleTableExpr
	tableName   *TableName
	indexHints  *IndexHints
	expr        Expr
	boolExpr    BoolExpr
	valExpr     ValExpr
	tuple       Tuple
	valExprs    ValExprs
	values      Values
	subquery    *Subquery
	caseExpr    *CaseExpr
	whens       []*When
	when        *When
	orderBy     OrderBy
	order       *Order
	limit       *Limit
	insRows     InsertRows
	updateExprs UpdateExprs
	updateExpr  *UpdateExpr
}

const LEX_ERROR = 57346
const SELECT = 57347
const INSERT = 57348
const UPDATE = 57349
const DELETE = 57350
const FROM = 57351
const WHERE = 57352
const GROUP = 57353
const HAVING = 57354
const ORDER = 57355
const BY = 57356
const LIMIT = 57357
const FOR = 57358
const ALL = 57359
const DISTINCT = 57360
const AS = 57361
const EXISTS = 57362
const NULL = 57363
const ASC = 57364
const DESC = 57365
const VALUES = 57366
const INTO = 57367
const DUPLICATE = 57368
const KEY = 57369
const DEFAULT = 57370
const SET = 57371
const SHOW = 57372
const LOCK = 57373
const ID = 57374
const STRING = 57375
const NUMBER = 57376
const VALUE_ARG = 57377
const COMMENT = 57378
const TIME = 57379
const ZONE = 57380
const SESSION = 57381
const AUTHORIZATION = 57382
const UNION = 57383
const MINUS = 57384
const EXCEPT = 57385
const INTERSECT = 57386
const JOIN = 57387
const STRAIGHT_JOIN = 57388
const LEFT = 57389
const RIGHT = 57390
const INNER = 57391
const OUTER = 57392
const CROSS = 57393
const NATURAL = 57394
const USE = 57395
const FORCE = 57396
const CONCAT = 57397
const ON = 57398
const OR = 57399
const AND = 57400
const NOT = 57401
const BETWEEN = 57402
const CASE = 57403
const WHEN = 57404
const THEN = 57405
const ELSE = 57406
const LE = 57407
const GE = 57408
const NE = 57409
const NULL_SAFE_EQUAL = 57410
const IS = 57411
const LIKE = 57412
const IN = 57413
const UNARY = 57414
const END = 57415
const BEGIN = 57416
const START = 57417
const TRANSACTION = 57418
const COMMIT = 57419
const ROLLBACK = 57420
const ISOLATION = 57421
const LEVEL = 57422
const NAMES = 57423
const REPLACE = 57424
const ADMIN = 57425
const HELP = 57426
const OFFSET = 57427
const COLLATE = 57428
const CREATE = 57429
const ALTER = 57430
const DROP = 57431
const RENAME = 57432
const TABLE = 57433
const INDEX = 57434
const VIEW = 57435
const TO = 57436
const IGNORE = 57437
const IF = 57438
const UNIQUE = 57439
const USING = 57440
const TRUNCATE = 57441

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"LEX_ERROR",
	"SELECT",
	"INSERT",
	"UPDATE",
	"DELETE",
	"FROM",
	"WHERE",
	"GROUP",
	"HAVING",
	"ORDER",
	"BY",
	"LIMIT",
	"FOR",
	"ALL",
	"DISTINCT",
	"AS",
	"EXISTS",
	"NULL",
	"ASC",
	"DESC",
	"VALUES",
	"INTO",
	"DUPLICATE",
	"KEY",
	"DEFAULT",
	"SET",
	"SHOW",
	"LOCK",
	"ID",
	"STRING",
	"NUMBER",
	"VALUE_ARG",
	"COMMENT",
	"'('",
	"'~'",
	"TIME",
	"ZONE",
	"SESSION",
	"AUTHORIZATION",
	"UNION",
	"MINUS",
	"EXCEPT",
	"INTERSECT",
	"','",
	"JOIN",
	"STRAIGHT_JOIN",
	"LEFT",
	"RIGHT",
	"INNER",
	"OUTER",
	"CROSS",
	"NATURAL",
	"USE",
	"FORCE",
	"CONCAT",
	"ON",
	"OR",
	"AND",
	"NOT",
	"BETWEEN",
	"CASE",
	"WHEN",
	"THEN",
	"ELSE",
	"'='",
	"'<'",
	"'>'",
	"LE",
	"GE",
	"NE",
	"NULL_SAFE_EQUAL",
	"IS",
	"LIKE",
	"IN",
	"'|'",
	"'&'",
	"'+'",
	"'-'",
	"'*'",
	"'/'",
	"'%'",
	"'^'",
	"'.'",
	"UNARY",
	"END",
	"BEGIN",
	"START",
	"TRANSACTION",
	"COMMIT",
	"ROLLBACK",
	"ISOLATION",
	"LEVEL",
	"NAMES",
	"REPLACE",
	"ADMIN",
	"HELP",
	"OFFSET",
	"COLLATE",
	"CREATE",
	"ALTER",
	"DROP",
	"RENAME",
	"TABLE",
	"INDEX",
	"VIEW",
	"TO",
	"IGNORE",
	"IF",
	"UNIQUE",
	"USING",
	"TRUNCATE",
	"')'",
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

const yyLast = 739

var yyAct = [...]int{

	58, 366, 333, 55, 398, 122, 101, 188, 292, 360,
	67, 197, 269, 183, 123, 3, 233, 56, 145, 251,
	135, 139, 144, 103, 211, 210, 175, 407, 154, 64,
	128, 407, 72, 407, 149, 80, 305, 90, 205, 267,
	148, 61, 62, 63, 64, 153, 70, 72, 205, 105,
	205, 87, 104, 93, 191, 65, 61, 62, 63, 226,
	66, 70, 313, 314, 315, 316, 317, 180, 318, 319,
	152, 143, 73, 39, 40, 41, 42, 130, 377, 288,
	132, 79, 238, 80, 136, 49, 376, 73, 68, 69,
	146, 74, 89, 76, 257, 409, 375, 77, 99, 408,
	129, 406, 131, 68, 69, 81, 296, 266, 172, 255,
	82, 83, 84, 258, 298, 116, 241, 118, 239, 71,
	85, 324, 177, 173, 352, 354, 124, 65, 242, 356,
	187, 209, 141, 126, 71, 179, 121, 109, 195, 178,
	190, 105, 174, 176, 201, 181, 412, 134, 207, 184,
	211, 210, 158, 162, 199, 150, 300, 210, 105, 88,
	105, 104, 161, 104, 232, 229, 163, 164, 165, 166,
	167, 168, 169, 170, 171, 184, 286, 246, 353, 150,
	150, 211, 210, 235, 254, 256, 253, 225, 227, 276,
	196, 102, 283, 248, 203, 186, 245, 372, 240, 112,
	113, 114, 109, 261, 374, 282, 281, 361, 231, 275,
	177, 159, 301, 228, 373, 273, 265, 287, 262, 247,
	194, 350, 237, 361, 274, 110, 111, 112, 113, 114,
	109, 280, 137, 346, 344, 150, 349, 348, 347, 345,
	243, 159, 234, 297, 234, 277, 278, 180, 204, 291,
	382, 308, 302, 289, 117, 313, 314, 315, 316, 317,
	303, 318, 319, 115, 105, 150, 20, 104, 105, 294,
	307, 309, 279, 392, 306, 284, 285, 199, 380, 311,
	290, 159, 273, 78, 323, 310, 205, 272, 140, 106,
	391, 105, 271, 272, 104, 263, 295, 331, 271, 325,
	332, 330, 390, 140, 199, 66, 299, 230, 140, 108,
	107, 110, 111, 112, 113, 114, 109, 39, 40, 41,
	42, 120, 273, 273, 342, 343, 339, 119, 155, 98,
	91, 381, 65, 322, 208, 363, 357, 355, 338, 337,
	404, 362, 326, 327, 106, 368, 321, 91, 260, 259,
	89, 20, 21, 22, 23, 405, 329, 202, 192, 189,
	336, 185, 178, 133, 108, 107, 110, 111, 112, 113,
	114, 109, 393, 157, 379, 24, 25, 388, 386, 20,
	156, 138, 249, 193, 396, 96, 94, 236, 359, 397,
	334, 399, 399, 399, 364, 367, 400, 401, 198, 371,
	335, 106, 36, 105, 358, 387, 104, 389, 413, 410,
	293, 370, 341, 414, 234, 415, 100, 411, 378, 402,
	20, 108, 107, 110, 111, 112, 113, 114, 109, 44,
	19, 124, 18, 17, 16, 30, 31, 15, 32, 33,
	394, 395, 367, 34, 35, 14, 384, 385, 26, 27,
	29, 28, 53, 13, 142, 250, 64, 75, 304, 72,
	37, 252, 43, 127, 200, 403, 383, 65, 61, 62,
	63, 365, 66, 70, 50, 369, 52, 154, 64, 340,
	244, 72, 106, 125, 45, 46, 47, 48, 182, 148,
	61, 62, 63, 60, 153, 70, 57, 86, 59, 73,
	92, 264, 108, 107, 110, 111, 112, 113, 114, 109,
	54, 212, 151, 351, 270, 68, 69, 312, 268, 152,
	147, 73, 320, 206, 95, 20, 51, 108, 107, 110,
	111, 112, 113, 114, 109, 38, 97, 68, 69, 146,
	154, 64, 12, 11, 72, 10, 71, 154, 64, 9,
	8, 72, 65, 61, 62, 63, 7, 153, 70, 65,
	61, 62, 63, 64, 153, 70, 72, 6, 71, 5,
	160, 4, 2, 1, 65, 61, 62, 63, 0, 66,
	70, 0, 152, 20, 73, 0, 0, 0, 0, 152,
	0, 73, 107, 110, 111, 112, 113, 114, 109, 64,
	68, 69, 72, 0, 0, 0, 73, 68, 69, 0,
	65, 61, 62, 63, 0, 66, 70, 0, 0, 0,
	0, 0, 68, 69, 0, 0, 0, 0, 0, 0,
	0, 71, 64, 0, 0, 72, 0, 0, 71, 0,
	0, 0, 73, 65, 61, 62, 63, 106, 66, 70,
	328, 0, 0, 71, 0, 0, 0, 0, 68, 69,
	0, 0, 0, 0, 0, 0, 0, 108, 107, 110,
	111, 112, 113, 114, 109, 73, 0, 0, 0, 0,
	0, 0, 0, 106, 0, 0, 0, 214, 216, 71,
	0, 68, 69, 218, 219, 220, 221, 222, 223, 224,
	217, 215, 213, 108, 107, 110, 111, 112, 113, 114,
	109, 106, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 71, 0, 0, 0, 0, 0, 0, 0,
	0, 108, 107, 110, 111, 112, 113, 114, 109,
}
var yyPact = [...]int{

	346, -1000, -1000, 274, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 435, -15, -27, -1, 4,
	-1000, 29, -1000, -1000, -1000, 60, 298, -1000, 415, 369,
	-1000, -1000, -1000, 367, -1000, -75, 318, 407, 95, 653,
	223, 21, 212, -1000, -1000, -1000, -1000, 611, 290, 284,
	-1000, -1000, -1000, -1000, -1000, 50, 578, -1000, -1000, -1000,
	-1000, -1000, -1000, 611, -81, -7, 298, -1000, -4, 298,
	-1000, 331, -91, 298, -91, -1000, 356, 271, -1000, 46,
	-1000, -1000, -35, -1000, -1000, 457, -1000, 292, 355, 344,
	318, 194, 542, -1000, 85, -1000, 611, 611, 611, 611,
	611, 611, 611, 611, 611, -1000, 13, -1000, -1000, 8,
	457, 330, 20, 30, 653, 84, 653, 329, 133, 298,
	-1000, 327, -1000, -55, 326, 363, 161, 298, 318, 374,
	300, 325, 318, -1000, 239, -1000, -1000, 315, 45, 121,
	625, -1000, 527, 520, 270, -1000, 318, 300, 404, 300,
	-1000, 286, 23, 449, 145, 513, -1000, 117, 117, 52,
	52, 52, -1000, -1000, 3, 457, 1, -1000, 42, -1000,
	611, -1000, 110, -1000, 527, -1000, 362, -94, -1000, 81,
	-1000, 317, -1000, -1000, 316, -1000, 266, -1000, 268, 274,
	-8, -1000, -1000, -1000, 255, 457, -1000, -1000, 298, 107,
	527, 527, 611, 268, 129, 611, 611, 155, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 625, -36, 625, -1000,
	415, 251, 234, 397, 527, -1000, 611, 653, -1000, -1000,
	-9, -1000, 298, 653, 26, -1000, 611, 90, -1000, -1000,
	153, 298, -1000, -73, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 374, 300, 204, -1000, -1000, 300, 232, 207,
	314, 261, 35, -1000, -1000, -1000, -1000, -1000, 96, 653,
	-1000, 268, 611, 611, 653, 589, -1000, 335, -1000, 374,
	300, 397, 375, 386, 121, 653, -1000, -1000, -1000, 653,
	611, 307, -1000, -1000, 306, -1000, -1000, 194, 268, -1000,
	401, 255, 255, -1000, -1000, 186, 185, 189, 188, 173,
	68, -1000, 305, 14, 304, -1000, 653, 343, 611, -1000,
	148, 164, 375, -1000, 611, 611, 653, -1000, -1000, -1000,
	399, 385, 207, 138, -1000, 166, -1000, 156, -1000, -1000,
	-1000, -1000, -11, -21, -29, -1000, -1000, -1000, 611, 653,
	-1000, 348, -1000, -1000, 231, 203, -1000, 424, -1000, 397,
	527, 611, 527, -1000, -1000, 265, 253, 236, 653, 345,
	611, 611, 611, -1000, -1000, -1000, 375, 121, 200, 121,
	298, 298, 298, 412, 653, 653, -1000, 324, -14, -1000,
	-16, -20, 300, -1000, 410, 69, -1000, 298, -1000, -1000,
	194, -1000, 298, -1000, 298, -1000,
}
var yyPgo = [...]int{

	0, 573, 572, 14, 571, 569, 567, 556, 550, 549,
	545, 543, 542, 462, 536, 535, 524, 283, 22, 18,
	523, 522, 520, 518, 12, 517, 514, 51, 513, 4,
	16, 34, 512, 511, 11, 510, 59, 17, 5, 501,
	498, 10, 496, 3, 493, 488, 13, 483, 480, 479,
	475, 8, 471, 1, 466, 2, 465, 21, 464, 9,
	6, 23, 147, 463, 461, 458, 457, 455, 0, 7,
	454, 453, 445, 437, 434, 433, 432, 430, 429,
}
var yyR1 = [...]int{

	0, 1, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 2, 2, 3,
	3, 3, 4, 4, 74, 74, 5, 6, 7, 7,
	7, 7, 8, 8, 8, 8, 8, 71, 71, 72,
	73, 75, 75, 76, 77, 9, 9, 9, 10, 10,
	10, 11, 12, 12, 12, 78, 13, 14, 14, 15,
	15, 15, 15, 15, 16, 16, 18, 18, 19, 19,
	19, 22, 22, 20, 20, 20, 23, 23, 24, 24,
	24, 24, 21, 21, 21, 25, 25, 25, 25, 25,
	25, 25, 25, 25, 26, 26, 26, 27, 27, 28,
	28, 28, 28, 29, 29, 30, 30, 31, 31, 31,
	31, 31, 32, 32, 32, 32, 32, 32, 32, 32,
	32, 32, 33, 33, 33, 33, 33, 33, 33, 34,
	34, 39, 39, 37, 37, 41, 38, 38, 36, 36,
	36, 36, 36, 36, 36, 36, 36, 36, 36, 36,
	36, 36, 36, 36, 36, 36, 40, 40, 42, 42,
	42, 44, 47, 47, 45, 45, 46, 48, 48, 43,
	43, 43, 35, 35, 35, 35, 49, 49, 50, 50,
	51, 51, 52, 52, 53, 54, 54, 54, 55, 55,
	55, 55, 56, 56, 56, 57, 57, 58, 58, 59,
	59, 60, 60, 61, 61, 62, 62, 63, 63, 17,
	17, 64, 64, 64, 64, 64, 65, 65, 66, 66,
	67, 67, 68, 69, 70, 70,
}
var yyR2 = [...]int{

	0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 4,
	12, 3, 8, 8, 6, 6, 8, 7, 3, 4,
	4, 6, 2, 3, 4, 3, 2, 1, 2, 1,
	1, 4, 2, 2, 4, 5, 8, 4, 6, 7,
	4, 5, 4, 5, 5, 0, 2, 0, 2, 1,
	2, 1, 1, 1, 0, 1, 1, 3, 1, 2,
	3, 1, 1, 0, 1, 2, 1, 3, 3, 3,
	3, 5, 0, 1, 2, 1, 1, 2, 3, 2,
	3, 2, 2, 2, 1, 3, 1, 1, 3, 0,
	5, 5, 5, 1, 3, 0, 2, 1, 3, 3,
	2, 3, 3, 3, 4, 3, 4, 5, 6, 3,
	4, 2, 1, 1, 1, 1, 1, 1, 1, 2,
	1, 1, 3, 3, 1, 3, 1, 3, 1, 1,
	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	2, 3, 4, 5, 4, 1, 1, 1, 1, 1,
	1, 5, 0, 1, 1, 2, 4, 0, 2, 1,
	3, 5, 1, 1, 1, 1, 0, 3, 0, 2,
	0, 3, 1, 3, 2, 0, 1, 1, 0, 2,
	4, 4, 0, 2, 4, 0, 3, 1, 3, 0,
	5, 1, 3, 3, 3, 0, 2, 0, 3, 0,
	1, 1, 1, 1, 1, 1, 0, 1, 0, 1,
	0, 2, 1, 0, 0, 1,
}
var yyChk = [...]int{

	-1000, -1, -2, -3, -4, -5, -6, -7, -8, -9,
	-10, -11, -12, -71, -72, -73, -74, -75, -76, -77,
	5, 6, 7, 8, 29, 30, 102, 103, 105, 104,
	89, 90, 92, 93, 97, 98, 56, 114, -15, 43,
	44, 45, 46, -13, -78, -13, -13, -13, -13, -36,
	39, 91, 41, 17, -35, -43, -37, -42, -68, -40,
	-44, 33, 34, 35, 21, 32, 37, -41, 80, 81,
	38, 111, 24, 64, 106, -66, 108, 112, -17, 108,
	110, 106, 106, 107, 108, 91, -13, -27, 99, 32,
	-68, 32, -13, -3, 17, -16, 18, -14, -17, -27,
	9, -60, 96, -61, -43, -68, 58, 79, 78, 85,
	80, 81, 82, 83, 84, 40, 94, 42, -36, 37,
	37, 86, -38, -3, -36, -47, -36, -63, 111, 107,
	-68, 106, -68, 32, -62, 111, -68, -62, 25, -57,
	37, 86, -70, 106, -18, -19, 82, -22, 32, -31,
	-36, -32, 62, 37, 20, 36, 25, 29, -27, 47,
	28, -36, 68, -36, -36, -36, -36, -36, -36, -36,
	-36, -36, 95, 115, -18, 18, -18, -68, 32, 115,
	47, 115, -45, -46, 65, 32, 62, -68, -69, 32,
	-69, 109, 32, 20, 59, -68, -27, -34, 24, -3,
	-58, -43, 32, -27, 9, 47, -20, -68, 19, 86,
	61, 60, -33, 77, 62, 76, 63, 75, 68, 69,
	70, 71, 72, 73, 74, -31, -36, -31, -36, -41,
	37, -27, -60, -30, 10, -61, 101, -36, 59, 115,
	-18, 115, 86, -36, -48, -46, 67, -31, -69, 20,
	-67, 113, -64, 105, 103, 28, 104, 13, 32, 32,
	32, -69, -57, 29, -39, -37, 115, 47, -23, -24,
	-26, 37, 32, -41, -19, -68, 82, -31, -31, -36,
	-37, 77, 76, 63, -36, -36, 21, 62, 115, -57,
	29, -30, -51, 13, -31, -36, 115, -68, 88, -36,
	66, 59, -68, -69, -65, 109, -34, -60, 47, -43,
	-30, 47, -25, 48, 49, 50, 51, 52, 54, 55,
	-21, 32, 19, -24, 86, -37, -36, -36, 61, 21,
	-34, -60, -51, -55, 15, 14, -36, 32, 32, -37,
	-49, 11, -24, -24, 48, 53, 48, 53, 48, 48,
	48, -28, 56, 110, 57, 32, 115, 32, 61, -36,
	-59, 59, -59, -55, -36, -52, -53, -36, -69, -50,
	12, 14, 59, 48, 48, 107, 107, 107, -36, 26,
	47, 100, 47, -54, 22, 23, -51, -31, -38, -31,
	37, 37, 37, 27, -36, -36, -53, -55, -29, -68,
	-29, -29, 7, -56, 16, 31, 115, 47, 115, 115,
	-60, 7, 77, -68, -68, -68,
}
var yyDef = [...]int{

	0, -2, 1, 2, 3, 4, 5, 6, 7, 8,
	9, 10, 11, 12, 13, 14, 15, 16, 17, 18,
	55, 55, 55, 55, 55, 0, 218, 209, 0, 0,
	37, 0, 39, 40, 55, 0, 0, 55, 0, 59,
	61, 62, 63, 64, 57, 209, 0, 0, 0, 32,
	0, 0, 0, 36, 138, 139, 140, 0, 169, 0,
	155, 172, 173, 174, 175, 222, 0, 134, 158, 159,
	160, 156, 157, 162, 207, 0, 0, 219, 0, 0,
	210, 0, 205, 0, 205, 38, 0, 195, 42, 97,
	43, 222, 224, 21, 60, 0, 65, 56, 0, 0,
	0, 28, 0, 201, 0, 169, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 33, 0, 35, 150, 0,
	0, 0, 0, 0, 136, 0, 163, 0, 0, 0,
	223, 0, 223, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 225, 19, 66, 68, 73, 222, 71,
	72, 107, 0, 0, 0, 58, 0, 0, 105, 0,
	29, 30, 0, 141, 142, 143, 144, 145, 146, 147,
	148, 149, 34, 151, 0, 0, 0, 170, 222, 133,
	0, 135, 167, 164, 0, 223, 0, 220, 47, 0,
	50, 0, 52, 206, 0, 223, 195, 41, 0, 130,
	0, 197, 98, 44, 0, 0, 69, 74, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 122, 123,
	124, 125, 126, 127, 128, 110, 0, 0, 136, 121,
	0, 195, 105, 180, 0, 202, 0, 203, 204, 152,
	0, 154, 0, 137, 0, 165, 0, 0, 45, 208,
	0, 0, 223, 216, 211, 212, 213, 214, 215, 51,
	53, 54, 0, 0, 129, 131, 196, 0, 105, 76,
	82, 0, 94, 96, 67, 75, 70, 108, 109, 112,
	113, 0, 0, 0, 115, 0, 119, 0, 111, 0,
	0, 180, 188, 0, 106, 31, 153, 171, 161, 168,
	0, 0, 221, 48, 0, 217, 24, 25, 0, 198,
	176, 0, 0, 85, 86, 0, 0, 0, 0, 0,
	99, 83, 0, 0, 0, 114, 116, 0, 0, 120,
	199, 199, 188, 27, 0, 0, 166, 223, 49, 132,
	178, 0, 77, 80, 87, 0, 89, 0, 91, 92,
	93, 78, 0, 0, 0, 84, 79, 95, 0, 117,
	22, 0, 23, 26, 189, 181, 182, 185, 46, 180,
	0, 0, 0, 88, 90, 0, 0, 0, 118, 0,
	0, 0, 0, 184, 186, 187, 188, 179, 177, 81,
	0, 0, 0, 0, 190, 191, 183, 192, 0, 103,
	0, 0, 0, 20, 0, 0, 100, 0, 101, 102,
	200, 193, 0, 104, 0, 194,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 84, 79, 3,
	37, 115, 82, 80, 47, 81, 86, 83, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	69, 68, 70, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 85, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 78, 3, 38,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 39, 40, 41, 42, 43,
	44, 45, 46, 48, 49, 50, 51, 52, 53, 54,
	55, 56, 57, 58, 59, 60, 61, 62, 63, 64,
	65, 66, 67, 71, 72, 73, 74, 75, 76, 77,
	87, 88, 89, 90, 91, 92, 93, 94, 95, 96,
	97, 98, 99, 100, 101, 102, 103, 104, 105, 106,
	107, 108, 109, 110, 111, 112, 113, 114,
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
		//line sql.y:195
		{
			SetParseTree(yylex, yyDollar[1].statement)
		}
	case 2:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:201
		{
			yyVAL.statement = yyDollar[1].selStmt
		}
	case 19:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:223
		{
			yyVAL.selStmt = &SimpleSelect{Comments: Comments(yyDollar[2].bytes2), Distinct: yyDollar[3].str, SelectExprs: yyDollar[4].selectExprs}
		}
	case 20:
		yyDollar = yyS[yypt-12 : yypt+1]
		//line sql.y:227
		{
			yyVAL.selStmt = &Select{Comments: Comments(yyDollar[2].bytes2), Distinct: yyDollar[3].str, SelectExprs: yyDollar[4].selectExprs, From: yyDollar[6].tableExprs, Where: NewWhere(AST_WHERE, yyDollar[7].boolExpr), GroupBy: GroupBy(yyDollar[8].valExprs), Having: NewWhere(AST_HAVING, yyDollar[9].boolExpr), OrderBy: yyDollar[10].orderBy, Limit: yyDollar[11].limit, Lock: yyDollar[12].str}
		}
	case 21:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:231
		{
			yyVAL.selStmt = &Union{Type: yyDollar[2].str, Left: yyDollar[1].selStmt, Right: yyDollar[3].selStmt}
		}
	case 22:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line sql.y:238
		{
			yyVAL.statement = &Insert{Comments: Comments(yyDollar[2].bytes2), Ignore: yyDollar[3].str, Table: yyDollar[5].tableName, Columns: yyDollar[6].columns, Rows: yyDollar[7].insRows, OnDup: OnDup(yyDollar[8].updateExprs)}
		}
	case 23:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line sql.y:242
		{
			cols := make(Columns, 0, len(yyDollar[7].updateExprs))
			vals := make(ValTuple, 0, len(yyDollar[7].updateExprs))
			for _, col := range yyDollar[7].updateExprs {
				cols = append(cols, &NonStarExpr{Expr: col.Name})
				vals = append(vals, col.Expr)
			}
			yyVAL.statement = &Insert{Comments: Comments(yyDollar[2].bytes2), Ignore: yyDollar[3].str, Table: yyDollar[5].tableName, Columns: cols, Rows: Values{vals}, OnDup: OnDup(yyDollar[8].updateExprs)}
		}
	case 24:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:254
		{
			yyVAL.statement = &Replace{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[4].tableName, Columns: yyDollar[5].columns, Rows: yyDollar[6].insRows}
		}
	case 25:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:258
		{
			cols := make(Columns, 0, len(yyDollar[6].updateExprs))
			vals := make(ValTuple, 0, len(yyDollar[6].updateExprs))
			for _, col := range yyDollar[6].updateExprs {
				cols = append(cols, &NonStarExpr{Expr: col.Name})
				vals = append(vals, col.Expr)
			}
			yyVAL.statement = &Replace{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[4].tableName, Columns: cols, Rows: Values{vals}}
		}
	case 26:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line sql.y:271
		{
			yyVAL.statement = &Update{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[3].tableName, Exprs: yyDollar[5].updateExprs, Where: NewWhere(AST_WHERE, yyDollar[6].boolExpr), OrderBy: yyDollar[7].orderBy, Limit: yyDollar[8].limit}
		}
	case 27:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line sql.y:277
		{
			yyVAL.statement = &Delete{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[4].tableName, Where: NewWhere(AST_WHERE, yyDollar[5].boolExpr), OrderBy: yyDollar[6].orderBy, Limit: yyDollar[7].limit}
		}
	case 28:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:283
		{
			yyVAL.statement = &Set{Comments: Comments(yyDollar[2].bytes2), Exprs: yyDollar[3].updateExprs}
		}
	case 29:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:287
		{
			yyVAL.statement = &Set{Comments: Comments(yyDollar[2].bytes2), Exprs: UpdateExprs{&UpdateExpr{Name: &ColName{Name: []byte("names")}, Expr: StrVal("default")}}}
		}
	case 30:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:291
		{
			yyVAL.statement = &Set{Comments: Comments(yyDollar[2].bytes2), Exprs: UpdateExprs{&UpdateExpr{Name: &ColName{Name: []byte("names")}, Expr: yyDollar[4].valExpr}}}
		}
	case 31:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:295
		{
			yyVAL.statement = &Set{
				Comments: Comments(yyDollar[2].bytes2),
				Exprs: UpdateExprs{
					&UpdateExpr{
						Name: &ColName{Name: []byte("names")}, Expr: yyDollar[4].valExpr,
					},
					&UpdateExpr{
						Name: &ColName{Name: []byte("collate")}, Expr: yyDollar[6].valExpr,
					},
				},
			}
		}
	case 32:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:312
		{
			yyVAL.statement = &Show{
				Name: String(yyDollar[2].valExpr),
			}
		}
	case 33:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:318
		{
			yyVAL.statement = &Show{
				Name: "TIME ZONE",
			}
		}
	case 34:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:324
		{
			yyVAL.statement = &Show{
				Name: "TRANSACTION ISOLATION LEVEL",
			}
		}
	case 35:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:330
		{
			yyVAL.statement = &Show{
				Name: "SESSION AUTHORIZATION",
			}
		}
	case 36:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:336
		{
			yyVAL.statement = &Show{
				Name: "ALL",
			}
		}
	case 37:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:345
		{
			yyVAL.statement = &Begin{}
		}
	case 38:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:349
		{
			yyVAL.statement = &Begin{}
		}
	case 39:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:356
		{
			yyVAL.statement = &Commit{}
		}
	case 40:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:362
		{
			yyVAL.statement = &Rollback{}
		}
	case 41:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:368
		{
			yyVAL.statement = &Admin{Region: yyDollar[2].tableName, Columns: yyDollar[3].columns, Rows: yyDollar[4].insRows}
		}
	case 42:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:372
		{
			yyVAL.statement = &AdminHelp{}
		}
	case 43:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:378
		{
			yyVAL.statement = &UseDB{DB: string(yyDollar[2].bytes)}
		}
	case 44:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:384
		{
			yyVAL.statement = &Truncate{Comments: Comments(yyDollar[2].bytes2), TableOpt: yyDollar[3].str, Table: yyDollar[4].tableName}
		}
	case 45:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:390
		{
			yyVAL.statement = &DDL{Action: AST_CREATE, NewName: yyDollar[4].bytes}
		}
	case 46:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line sql.y:394
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[7].bytes, NewName: yyDollar[7].bytes}
		}
	case 47:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:399
		{
			yyVAL.statement = &DDL{Action: AST_CREATE, NewName: yyDollar[3].bytes}
		}
	case 48:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:405
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Ignore: yyDollar[2].str, Table: yyDollar[4].bytes, NewName: yyDollar[4].bytes}
		}
	case 49:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line sql.y:409
		{
			// Change this to a rename statement
			yyVAL.statement = &DDL{Action: AST_RENAME, Ignore: yyDollar[2].str, Table: yyDollar[4].bytes, NewName: yyDollar[7].bytes}
		}
	case 50:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:414
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[3].bytes, NewName: yyDollar[3].bytes}
		}
	case 51:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:420
		{
			yyVAL.statement = &DDL{Action: AST_RENAME, Table: yyDollar[3].bytes, NewName: yyDollar[5].bytes}
		}
	case 52:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:426
		{
			yyVAL.statement = &DDL{Action: AST_DROP, Table: yyDollar[4].bytes}
		}
	case 53:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:430
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[5].bytes, NewName: yyDollar[5].bytes}
		}
	case 54:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:435
		{
			yyVAL.statement = &DDL{Action: AST_DROP, Table: yyDollar[4].bytes}
		}
	case 55:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:440
		{
			SetAllowComments(yylex, true)
		}
	case 56:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:444
		{
			yyVAL.bytes2 = yyDollar[2].bytes2
			SetAllowComments(yylex, false)
		}
	case 57:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:450
		{
			yyVAL.bytes2 = nil
		}
	case 58:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:454
		{
			yyVAL.bytes2 = append(yyDollar[1].bytes2, yyDollar[2].bytes)
		}
	case 59:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:460
		{
			yyVAL.str = AST_UNION
		}
	case 60:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:464
		{
			yyVAL.str = AST_UNION_ALL
		}
	case 61:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:468
		{
			yyVAL.str = AST_SET_MINUS
		}
	case 62:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:472
		{
			yyVAL.str = AST_EXCEPT
		}
	case 63:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:476
		{
			yyVAL.str = AST_INTERSECT
		}
	case 64:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:481
		{
			yyVAL.str = ""
		}
	case 65:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:485
		{
			yyVAL.str = AST_DISTINCT
		}
	case 66:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:491
		{
			yyVAL.selectExprs = SelectExprs{yyDollar[1].selectExpr}
		}
	case 67:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:495
		{
			yyVAL.selectExprs = append(yyVAL.selectExprs, yyDollar[3].selectExpr)
		}
	case 68:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:501
		{
			yyVAL.selectExpr = &StarExpr{}
		}
	case 69:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:505
		{
			yyVAL.selectExpr = &NonStarExpr{Expr: yyDollar[1].expr, As: yyDollar[2].bytes}
		}
	case 70:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:509
		{
			yyVAL.selectExpr = &StarExpr{TableName: yyDollar[1].bytes}
		}
	case 71:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:515
		{
			yyVAL.expr = yyDollar[1].boolExpr
		}
	case 72:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:519
		{
			yyVAL.expr = yyDollar[1].valExpr
		}
	case 73:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:524
		{
			yyVAL.bytes = nil
		}
	case 74:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:528
		{
			yyVAL.bytes = yyDollar[1].bytes
		}
	case 75:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:532
		{
			yyVAL.bytes = yyDollar[2].bytes
		}
	case 76:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:538
		{
			yyVAL.tableExprs = TableExprs{yyDollar[1].tableExpr}
		}
	case 77:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:542
		{
			yyVAL.tableExprs = append(yyVAL.tableExprs, yyDollar[3].tableExpr)
		}
	case 78:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:548
		{
			yyVAL.tableExpr = &AliasedTableExpr{Expr: yyDollar[1].smTableExpr, As: yyDollar[2].bytes, Hints: yyDollar[3].indexHints}
		}
	case 79:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:552
		{
			yyVAL.tableExpr = &ParenTableExpr{Expr: yyDollar[2].tableExpr}
		}
	case 80:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:556
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr}
		}
	case 81:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:560
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr, On: yyDollar[5].boolExpr}
		}
	case 82:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:565
		{
			yyVAL.bytes = nil
		}
	case 83:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:569
		{
			yyVAL.bytes = yyDollar[1].bytes
		}
	case 84:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:573
		{
			yyVAL.bytes = yyDollar[2].bytes
		}
	case 85:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:579
		{
			yyVAL.str = AST_JOIN
		}
	case 86:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:583
		{
			yyVAL.str = AST_STRAIGHT_JOIN
		}
	case 87:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:587
		{
			yyVAL.str = AST_LEFT_JOIN
		}
	case 88:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:591
		{
			yyVAL.str = AST_LEFT_JOIN
		}
	case 89:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:595
		{
			yyVAL.str = AST_RIGHT_JOIN
		}
	case 90:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:599
		{
			yyVAL.str = AST_RIGHT_JOIN
		}
	case 91:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:603
		{
			yyVAL.str = AST_JOIN
		}
	case 92:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:607
		{
			yyVAL.str = AST_CROSS_JOIN
		}
	case 93:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:611
		{
			yyVAL.str = AST_NATURAL_JOIN
		}
	case 94:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:617
		{
			yyVAL.smTableExpr = &TableName{Name: yyDollar[1].bytes}
		}
	case 95:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:621
		{
			yyVAL.smTableExpr = &TableName{Qualifier: yyDollar[1].bytes, Name: yyDollar[3].bytes}
		}
	case 96:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:625
		{
			yyVAL.smTableExpr = yyDollar[1].subquery
		}
	case 97:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:631
		{
			yyVAL.tableName = &TableName{Name: yyDollar[1].bytes}
		}
	case 98:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:635
		{
			yyVAL.tableName = &TableName{Qualifier: yyDollar[1].bytes, Name: yyDollar[3].bytes}
		}
	case 99:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:640
		{
			yyVAL.indexHints = nil
		}
	case 100:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:644
		{
			yyVAL.indexHints = &IndexHints{Type: AST_USE, Indexes: yyDollar[4].bytes2}
		}
	case 101:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:648
		{
			yyVAL.indexHints = &IndexHints{Type: AST_IGNORE, Indexes: yyDollar[4].bytes2}
		}
	case 102:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:652
		{
			yyVAL.indexHints = &IndexHints{Type: AST_FORCE, Indexes: yyDollar[4].bytes2}
		}
	case 103:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:658
		{
			yyVAL.bytes2 = [][]byte{yyDollar[1].bytes}
		}
	case 104:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:662
		{
			yyVAL.bytes2 = append(yyDollar[1].bytes2, yyDollar[3].bytes)
		}
	case 105:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:667
		{
			yyVAL.boolExpr = nil
		}
	case 106:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:671
		{
			yyVAL.boolExpr = yyDollar[2].boolExpr
		}
	case 108:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:678
		{
			yyVAL.boolExpr = &AndExpr{Left: yyDollar[1].boolExpr, Right: yyDollar[3].boolExpr}
		}
	case 109:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:682
		{
			yyVAL.boolExpr = &OrExpr{Left: yyDollar[1].boolExpr, Right: yyDollar[3].boolExpr}
		}
	case 110:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:686
		{
			yyVAL.boolExpr = &NotExpr{Expr: yyDollar[2].boolExpr}
		}
	case 111:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:690
		{
			yyVAL.boolExpr = &ParenBoolExpr{Expr: yyDollar[2].boolExpr}
		}
	case 112:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:696
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: yyDollar[2].str, Right: yyDollar[3].valExpr}
		}
	case 113:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:700
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_IN, Right: yyDollar[3].tuple}
		}
	case 114:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:704
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_NOT_IN, Right: yyDollar[4].tuple}
		}
	case 115:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:708
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_LIKE, Right: yyDollar[3].valExpr}
		}
	case 116:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:712
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_NOT_LIKE, Right: yyDollar[4].valExpr}
		}
	case 117:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:716
		{
			yyVAL.boolExpr = &RangeCond{Left: yyDollar[1].valExpr, Operator: AST_BETWEEN, From: yyDollar[3].valExpr, To: yyDollar[5].valExpr}
		}
	case 118:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:720
		{
			yyVAL.boolExpr = &RangeCond{Left: yyDollar[1].valExpr, Operator: AST_NOT_BETWEEN, From: yyDollar[4].valExpr, To: yyDollar[6].valExpr}
		}
	case 119:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:724
		{
			yyVAL.boolExpr = &NullCheck{Operator: AST_IS_NULL, Expr: yyDollar[1].valExpr}
		}
	case 120:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:728
		{
			yyVAL.boolExpr = &NullCheck{Operator: AST_IS_NOT_NULL, Expr: yyDollar[1].valExpr}
		}
	case 121:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:732
		{
			yyVAL.boolExpr = &ExistsExpr{Subquery: yyDollar[2].subquery}
		}
	case 122:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:738
		{
			yyVAL.str = AST_EQ
		}
	case 123:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:742
		{
			yyVAL.str = AST_LT
		}
	case 124:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:746
		{
			yyVAL.str = AST_GT
		}
	case 125:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:750
		{
			yyVAL.str = AST_LE
		}
	case 126:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:754
		{
			yyVAL.str = AST_GE
		}
	case 127:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:758
		{
			yyVAL.str = AST_NE
		}
	case 128:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:762
		{
			yyVAL.str = AST_NSE
		}
	case 129:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:768
		{
			yyVAL.insRows = yyDollar[2].values
		}
	case 130:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:772
		{
			yyVAL.insRows = yyDollar[1].selStmt
		}
	case 131:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:778
		{
			yyVAL.values = Values{yyDollar[1].tuple}
		}
	case 132:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:782
		{
			yyVAL.values = append(yyDollar[1].values, yyDollar[3].tuple)
		}
	case 133:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:788
		{
			yyVAL.tuple = ValTuple(yyDollar[2].valExprs)
		}
	case 134:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:792
		{
			yyVAL.tuple = yyDollar[1].subquery
		}
	case 135:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:798
		{
			yyVAL.subquery = &Subquery{yyDollar[2].selStmt}
		}
	case 136:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:804
		{
			yyVAL.valExprs = ValExprs{yyDollar[1].valExpr}
		}
	case 137:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:808
		{
			yyVAL.valExprs = append(yyDollar[1].valExprs, yyDollar[3].valExpr)
		}
	case 138:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:814
		{
			yyVAL.valExpr = yyDollar[1].valExpr
		}
	case 139:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:818
		{
			yyVAL.valExpr = yyDollar[1].colName
		}
	case 140:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:822
		{
			yyVAL.valExpr = yyDollar[1].tuple
		}
	case 141:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:826
		{
			yyVAL.valExpr = &ConcatExpr{Left: yyDollar[1].valExpr, Right: yyDollar[3].valExpr}
		}
	case 142:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:830
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_BITAND, Right: yyDollar[3].valExpr}
		}
	case 143:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:834
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_BITOR, Right: yyDollar[3].valExpr}
		}
	case 144:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:838
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_BITXOR, Right: yyDollar[3].valExpr}
		}
	case 145:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:842
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_PLUS, Right: yyDollar[3].valExpr}
		}
	case 146:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:846
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_MINUS, Right: yyDollar[3].valExpr}
		}
	case 147:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:850
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_MULT, Right: yyDollar[3].valExpr}
		}
	case 148:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:854
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_DIV, Right: yyDollar[3].valExpr}
		}
	case 149:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:858
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_MOD, Right: yyDollar[3].valExpr}
		}
	case 150:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:862
		{
			if num, ok := yyDollar[2].valExpr.(NumVal); ok {
				switch yyDollar[1].byt {
				case '-':
					yyVAL.valExpr = append(NumVal("-"), num...)
				case '+':
					yyVAL.valExpr = num
				default:
					yyVAL.valExpr = &UnaryExpr{Operator: yyDollar[1].byt, Expr: yyDollar[2].valExpr}
				}
			} else {
				yyVAL.valExpr = &UnaryExpr{Operator: yyDollar[1].byt, Expr: yyDollar[2].valExpr}
			}
		}
	case 151:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:877
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes}
		}
	case 152:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:881
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes, Exprs: yyDollar[3].selectExprs}
		}
	case 153:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:885
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes, Distinct: true, Exprs: yyDollar[4].selectExprs}
		}
	case 154:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:889
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes, Exprs: yyDollar[3].selectExprs}
		}
	case 155:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:893
		{
			yyVAL.valExpr = yyDollar[1].caseExpr
		}
	case 156:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:899
		{
			yyVAL.bytes = IF_BYTES
		}
	case 157:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:903
		{
			yyVAL.bytes = VALUES_BYTES
		}
	case 158:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:909
		{
			yyVAL.byt = AST_UPLUS
		}
	case 159:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:913
		{
			yyVAL.byt = AST_UMINUS
		}
	case 160:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:917
		{
			yyVAL.byt = AST_TILDA
		}
	case 161:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:923
		{
			yyVAL.caseExpr = &CaseExpr{Expr: yyDollar[2].valExpr, Whens: yyDollar[3].whens, Else: yyDollar[4].valExpr}
		}
	case 162:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:928
		{
			yyVAL.valExpr = nil
		}
	case 163:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:932
		{
			yyVAL.valExpr = yyDollar[1].valExpr
		}
	case 164:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:938
		{
			yyVAL.whens = []*When{yyDollar[1].when}
		}
	case 165:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:942
		{
			yyVAL.whens = append(yyDollar[1].whens, yyDollar[2].when)
		}
	case 166:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:948
		{
			yyVAL.when = &When{Cond: yyDollar[2].boolExpr, Val: yyDollar[4].valExpr}
		}
	case 167:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:953
		{
			yyVAL.valExpr = nil
		}
	case 168:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:957
		{
			yyVAL.valExpr = yyDollar[2].valExpr
		}
	case 169:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:963
		{
			yyVAL.colName = &ColName{Name: yyDollar[1].bytes}
		}
	case 170:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:967
		{
			yyVAL.colName = &ColName{Qualifier: yyDollar[1].bytes, Name: yyDollar[3].bytes}
		}
	case 171:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:971
		{
			yyVAL.colName = &ColName{Qualifier: yyDollar[3].bytes, Name: yyDollar[5].bytes}
		}
	case 172:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:977
		{
			yyVAL.valExpr = StrVal(yyDollar[1].bytes)
		}
	case 173:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:981
		{
			yyVAL.valExpr = NumVal(yyDollar[1].bytes)
		}
	case 174:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:985
		{
			yyVAL.valExpr = ValArg(yyDollar[1].bytes)
		}
	case 175:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:989
		{
			yyVAL.valExpr = &NullVal{}
		}
	case 176:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:994
		{
			yyVAL.valExprs = nil
		}
	case 177:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:998
		{
			yyVAL.valExprs = yyDollar[3].valExprs
		}
	case 178:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1003
		{
			yyVAL.boolExpr = nil
		}
	case 179:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1007
		{
			yyVAL.boolExpr = yyDollar[2].boolExpr
		}
	case 180:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1012
		{
			yyVAL.orderBy = nil
		}
	case 181:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1016
		{
			yyVAL.orderBy = yyDollar[3].orderBy
		}
	case 182:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1022
		{
			yyVAL.orderBy = OrderBy{yyDollar[1].order}
		}
	case 183:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1026
		{
			yyVAL.orderBy = append(yyDollar[1].orderBy, yyDollar[3].order)
		}
	case 184:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1032
		{
			yyVAL.order = &Order{Expr: yyDollar[1].valExpr, Direction: yyDollar[2].str}
		}
	case 185:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1037
		{
			yyVAL.str = AST_ASC
		}
	case 186:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1041
		{
			yyVAL.str = AST_ASC
		}
	case 187:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1045
		{
			yyVAL.str = AST_DESC
		}
	case 188:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1050
		{
			yyVAL.limit = nil
		}
	case 189:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1054
		{
			yyVAL.limit = &Limit{Rowcount: yyDollar[2].valExpr}
		}
	case 190:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:1058
		{
			yyVAL.limit = &Limit{Offset: yyDollar[2].valExpr, Rowcount: yyDollar[4].valExpr}
		}
	case 191:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:1062
		{
			yyVAL.limit = &Limit{Offset: yyDollar[4].valExpr, Rowcount: yyDollar[2].valExpr}
		}
	case 192:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1067
		{
			yyVAL.str = ""
		}
	case 193:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1071
		{
			yyVAL.str = AST_FOR_UPDATE
		}
	case 194:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:1075
		{
			if !bytes.Equal(yyDollar[3].bytes, SHARE) {
				yylex.Error("expecting share")
				return 1
			}
			if !bytes.Equal(yyDollar[4].bytes, MODE) {
				yylex.Error("expecting mode")
				return 1
			}
			yyVAL.str = AST_SHARE_MODE
		}
	case 195:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1088
		{
			yyVAL.columns = nil
		}
	case 196:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1092
		{
			yyVAL.columns = yyDollar[2].columns
		}
	case 197:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1098
		{
			yyVAL.columns = Columns{&NonStarExpr{Expr: yyDollar[1].colName}}
		}
	case 198:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1102
		{
			yyVAL.columns = append(yyVAL.columns, &NonStarExpr{Expr: yyDollar[3].colName})
		}
	case 199:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1107
		{
			yyVAL.updateExprs = nil
		}
	case 200:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:1111
		{
			yyVAL.updateExprs = yyDollar[5].updateExprs
		}
	case 201:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1117
		{
			yyVAL.updateExprs = UpdateExprs{yyDollar[1].updateExpr}
		}
	case 202:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1121
		{
			yyVAL.updateExprs = append(yyDollar[1].updateExprs, yyDollar[3].updateExpr)
		}
	case 203:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1127
		{
			yyVAL.updateExpr = &UpdateExpr{Name: yyDollar[1].colName, Expr: yyDollar[3].valExpr}
		}
	case 204:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1131
		{
			yyVAL.updateExpr = &UpdateExpr{Name: yyDollar[1].colName, Expr: StrVal("ON")}
		}
	case 205:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1136
		{
			yyVAL.empty = struct{}{}
		}
	case 206:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1138
		{
			yyVAL.empty = struct{}{}
		}
	case 207:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1141
		{
			yyVAL.empty = struct{}{}
		}
	case 208:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1143
		{
			yyVAL.empty = struct{}{}
		}
	case 209:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1146
		{
			yyVAL.str = ""
		}
	case 210:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1148
		{
			yyVAL.str = AST_IGNORE
		}
	case 211:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1152
		{
			yyVAL.empty = struct{}{}
		}
	case 212:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1154
		{
			yyVAL.empty = struct{}{}
		}
	case 213:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1156
		{
			yyVAL.empty = struct{}{}
		}
	case 214:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1158
		{
			yyVAL.empty = struct{}{}
		}
	case 215:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1160
		{
			yyVAL.empty = struct{}{}
		}
	case 216:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1163
		{
			yyVAL.empty = struct{}{}
		}
	case 217:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1165
		{
			yyVAL.empty = struct{}{}
		}
	case 218:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1168
		{
			yyVAL.empty = struct{}{}
		}
	case 219:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1170
		{
			yyVAL.empty = struct{}{}
		}
	case 220:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1173
		{
			yyVAL.empty = struct{}{}
		}
	case 221:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1175
		{
			yyVAL.empty = struct{}{}
		}
	case 222:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1179
		{
			yyVAL.bytes = bytes.ToLower(yyDollar[1].bytes)
		}
	case 223:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1184
		{
			ForceEOF(yylex)
		}
	case 224:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1189
		{
			yyVAL.str = ""
		}
	case 225:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1193
		{
			yyVAL.str = AST_TABLE
		}
	}
	goto yystack /* stack new state and value */
}
