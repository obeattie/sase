//line query/tokeniser.rl:1
package query

import (
	"fmt"
	"time"

	log "github.com/cihub/seelog"
)

//line query/tokeniser.rl:11
//line query/tokeniser.go:16
const sase_start int = 1
const sase_first_final int = 130
const sase_error int = 0

const sase_en_main int = 1

//line query/tokeniser.rl:12
func tokenize(data string) ([]*token, error) {
	var (
		cs        int             // current state
		p         int = 0         // data offset
		pe        int = len(data) // data end offset
		eof       int = pe        // eof offset
		mark      int = -1
		tokens        = make([]*token, 0, 10)
		proposals     = make([]*proposedToken, 0, 100)
		start         = time.Now()
	)
	defer log.Debugf("[Tokenizer] Took %s", time.Since(start).String())

	// Add a single token onto the committed tokens list to collect the actual tokens
	root := &token{tt: ttGeneric}
	tokens = append(tokens, root)

	markedText := func() string {
		return data[mark:p]
	}

	propose := func(typ tt) {
		log.Tracef("[Tokenizer] propose: %s", typ.String())
		proposals = append(proposals, &proposedToken{
			token: &token{
				tt: typ,
			},
			i:  len(tokens),
			pi: len(proposals),
		})
	}

	proposal := func(typ tt) *proposedToken {
		for i := len(proposals) - 1; i >= 0; i-- {
			pt := proposals[i]
			if pt.tt == typ {
				return pt
			}
		}
		return nil
	}

	setText := func(typ tt) string {
		text := markedText()
		proposal(typ).content = text
		return text
	}

	commit := func(typ tt) *token {
		log.Tracef("[Tokenizer] commit: %s", typ.String())
		pt := proposal(typ)
		t := pt.token
		// Add everything on the tokens list that has been committed since this token was proposed as children of this
		// token
		children := tokens[pt.i:]
		t.children = append(make([]*token, 0, len(children)), children...)
		tokens = append(tokens[:pt.i], t)
		// Remove any "defunct" proposals (like this one)
		log.Tracef("[Tokenizer] commit removing %d defunct proposals", len(proposals)-pt.pi)
		proposals = proposals[:pt.pi]
		return t
	}

	// Suppress "variable not used"
	var (
		_ = cs
		_ = p
		_ = pe
		_ = eof
		_ = mark
		_ = tokens
		_ = proposals
		_ = propose
		_ = proposal
		_ = setText
		_ = commit
	)

//line query/tokeniser.go:105
	{
		cs = sase_start
	}

//line query/tokeniser.go:110
	{
		if p == pe {
			goto _test_eof
		}
		switch cs {
		case 1:
			goto st_case_1
		case 0:
			goto st_case_0
		case 2:
			goto st_case_2
		case 3:
			goto st_case_3
		case 4:
			goto st_case_4
		case 5:
			goto st_case_5
		case 6:
			goto st_case_6
		case 7:
			goto st_case_7
		case 8:
			goto st_case_8
		case 9:
			goto st_case_9
		case 10:
			goto st_case_10
		case 130:
			goto st_case_130
		case 131:
			goto st_case_131
		case 132:
			goto st_case_132
		case 11:
			goto st_case_11
		case 12:
			goto st_case_12
		case 13:
			goto st_case_13
		case 14:
			goto st_case_14
		case 15:
			goto st_case_15
		case 16:
			goto st_case_16
		case 17:
			goto st_case_17
		case 18:
			goto st_case_18
		case 19:
			goto st_case_19
		case 20:
			goto st_case_20
		case 21:
			goto st_case_21
		case 133:
			goto st_case_133
		case 134:
			goto st_case_134
		case 22:
			goto st_case_22
		case 23:
			goto st_case_23
		case 24:
			goto st_case_24
		case 25:
			goto st_case_25
		case 26:
			goto st_case_26
		case 27:
			goto st_case_27
		case 28:
			goto st_case_28
		case 29:
			goto st_case_29
		case 30:
			goto st_case_30
		case 31:
			goto st_case_31
		case 32:
			goto st_case_32
		case 33:
			goto st_case_33
		case 34:
			goto st_case_34
		case 35:
			goto st_case_35
		case 36:
			goto st_case_36
		case 37:
			goto st_case_37
		case 38:
			goto st_case_38
		case 135:
			goto st_case_135
		case 136:
			goto st_case_136
		case 137:
			goto st_case_137
		case 39:
			goto st_case_39
		case 40:
			goto st_case_40
		case 41:
			goto st_case_41
		case 42:
			goto st_case_42
		case 43:
			goto st_case_43
		case 138:
			goto st_case_138
		case 44:
			goto st_case_44
		case 45:
			goto st_case_45
		case 46:
			goto st_case_46
		case 47:
			goto st_case_47
		case 48:
			goto st_case_48
		case 49:
			goto st_case_49
		case 50:
			goto st_case_50
		case 51:
			goto st_case_51
		case 52:
			goto st_case_52
		case 53:
			goto st_case_53
		case 54:
			goto st_case_54
		case 55:
			goto st_case_55
		case 56:
			goto st_case_56
		case 57:
			goto st_case_57
		case 58:
			goto st_case_58
		case 59:
			goto st_case_59
		case 60:
			goto st_case_60
		case 61:
			goto st_case_61
		case 62:
			goto st_case_62
		case 63:
			goto st_case_63
		case 64:
			goto st_case_64
		case 65:
			goto st_case_65
		case 66:
			goto st_case_66
		case 67:
			goto st_case_67
		case 68:
			goto st_case_68
		case 69:
			goto st_case_69
		case 139:
			goto st_case_139
		case 70:
			goto st_case_70
		case 71:
			goto st_case_71
		case 72:
			goto st_case_72
		case 73:
			goto st_case_73
		case 74:
			goto st_case_74
		case 140:
			goto st_case_140
		case 75:
			goto st_case_75
		case 76:
			goto st_case_76
		case 77:
			goto st_case_77
		case 78:
			goto st_case_78
		case 79:
			goto st_case_79
		case 80:
			goto st_case_80
		case 81:
			goto st_case_81
		case 82:
			goto st_case_82
		case 83:
			goto st_case_83
		case 84:
			goto st_case_84
		case 85:
			goto st_case_85
		case 86:
			goto st_case_86
		case 87:
			goto st_case_87
		case 88:
			goto st_case_88
		case 89:
			goto st_case_89
		case 141:
			goto st_case_141
		case 90:
			goto st_case_90
		case 91:
			goto st_case_91
		case 92:
			goto st_case_92
		case 93:
			goto st_case_93
		case 94:
			goto st_case_94
		case 95:
			goto st_case_95
		case 96:
			goto st_case_96
		case 97:
			goto st_case_97
		case 98:
			goto st_case_98
		case 99:
			goto st_case_99
		case 100:
			goto st_case_100
		case 101:
			goto st_case_101
		case 102:
			goto st_case_102
		case 103:
			goto st_case_103
		case 104:
			goto st_case_104
		case 105:
			goto st_case_105
		case 106:
			goto st_case_106
		case 107:
			goto st_case_107
		case 108:
			goto st_case_108
		case 109:
			goto st_case_109
		case 110:
			goto st_case_110
		case 111:
			goto st_case_111
		case 112:
			goto st_case_112
		case 113:
			goto st_case_113
		case 114:
			goto st_case_114
		case 115:
			goto st_case_115
		case 116:
			goto st_case_116
		case 117:
			goto st_case_117
		case 118:
			goto st_case_118
		case 119:
			goto st_case_119
		case 120:
			goto st_case_120
		case 121:
			goto st_case_121
		case 122:
			goto st_case_122
		case 123:
			goto st_case_123
		case 124:
			goto st_case_124
		case 125:
			goto st_case_125
		case 126:
			goto st_case_126
		case 127:
			goto st_case_127
		case 128:
			goto st_case_128
		case 129:
			goto st_case_129
		}
		goto st_out
	st1:
		if p++; p == pe {
			goto _test_eof1
		}
	st_case_1:
		switch data[p] {
		case 32:
			goto st1
		case 69:
			goto st2
		case 101:
			goto st2
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st1
		}
		goto st0
	st_case_0:
	st0:
		cs = 0
		goto _out
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
		switch data[p] {
		case 86:
			goto st3
		case 118:
			goto st3
		}
		goto st0
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
		switch data[p] {
		case 69:
			goto st4
		case 101:
			goto st4
		}
		goto st0
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
		switch data[p] {
		case 78:
			goto st5
		case 110:
			goto st5
		}
		goto st0
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
		switch data[p] {
		case 84:
			goto st6
		case 116:
			goto st6
		}
		goto st0
	st6:
		if p++; p == pe {
			goto _test_eof6
		}
	st_case_6:
		if data[p] == 32 {
			goto st7
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st7
		}
		goto st0
	st7:
		if p++; p == pe {
			goto _test_eof7
		}
	st_case_7:
		switch data[p] {
		case 32:
			goto st7
		case 33:
			goto tr8
		case 65:
			goto tr9
		case 78:
			goto tr11
		case 83:
			goto tr12
		case 95:
			goto tr10
		case 97:
			goto tr9
		case 110:
			goto tr11
		case 115:
			goto tr12
		}
		switch {
		case data[p] < 66:
			if 9 <= data[p] && data[p] <= 13 {
				goto st7
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr10
			}
		default:
			goto tr10
		}
		goto st0
	tr8:
//line query/tokeniser.rl:156
		propose(ttEventClause)
//line query/tokeniser.rl:136
		propose(ttNegatedDecl)
		goto st8
	st8:
		if p++; p == pe {
			goto _test_eof8
		}
	st_case_8:
//line query/tokeniser.go:532
		switch data[p] {
		case 32:
			goto st9
		case 40:
			goto st10
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st9
		}
		goto st0
	st9:
		if p++; p == pe {
			goto _test_eof9
		}
	st_case_9:
		if data[p] == 40 {
			goto st10
		}
		goto st0
	st10:
		if p++; p == pe {
			goto _test_eof10
		}
	st_case_10:
		switch data[p] {
		case 32:
			goto st10
		case 41:
			goto st130
		case 65:
			goto tr16
		case 95:
			goto tr17
		case 97:
			goto tr16
		}
		switch {
		case data[p] < 66:
			if 9 <= data[p] && data[p] <= 13 {
				goto st10
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr17
			}
		default:
			goto tr17
		}
		goto st0
	tr88:
//line query/tokeniser.rl:112
		setText(ttEventDeclAlias)
//line query/tokeniser.rl:113
		commit(ttEventDeclAlias)
//line query/tokeniser.rl:117
		commit(ttEventDecl)
		goto st130
	tr99:
//line query/tokeniser.rl:128
		commit(ttAnyDecl)
		goto st130
	st130:
		if p++; p == pe {
			goto _test_eof130
		}
	st_case_130:
//line query/tokeniser.go:599
		switch data[p] {
		case 32:
			goto tr213
		case 59:
			goto tr214
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr213
		}
		goto st0
	tr213:
//line query/tokeniser.rl:140
		commit(ttNegatedDecl)
//line query/tokeniser.rl:157
		commit(ttEventClause)
		goto st131
	tr236:
//line query/tokeniser.rl:112
		setText(ttEventDeclAlias)
//line query/tokeniser.rl:113
		commit(ttEventDeclAlias)
//line query/tokeniser.rl:117
		commit(ttEventDecl)
//line query/tokeniser.rl:157
		commit(ttEventClause)
		goto st131
	tr239:
//line query/tokeniser.rl:128
		commit(ttAnyDecl)
//line query/tokeniser.rl:157
		commit(ttEventClause)
		goto st131
	tr241:
//line query/tokeniser.rl:151
		commit(ttSeqDecl)
//line query/tokeniser.rl:157
		commit(ttEventClause)
		goto st131
	st131:
		if p++; p == pe {
			goto _test_eof131
		}
	st_case_131:
//line query/tokeniser.go:643
		switch data[p] {
		case 32:
			goto st131
		case 59:
			goto st132
		case 87:
			goto st11
		case 119:
			goto st11
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st131
		}
		goto st0
	tr214:
//line query/tokeniser.rl:140
		commit(ttNegatedDecl)
//line query/tokeniser.rl:157
		commit(ttEventClause)
		goto st132
	tr221:
//line query/tokeniser.rl:182
		propose(ttAttributeSelector)
//line query/tokeniser.rl:183
		setText(ttAttributeSelector)
//line query/tokeniser.rl:184
		commit(ttAttributeSelector)
//line query/tokeniser.rl:189
		commit(ttPredicate)
		goto st132
	tr231:
//line query/tokeniser.rl:199
		setText(ttDuration)
//line query/tokeniser.rl:200
		commit(ttDuration)
//line query/tokeniser.rl:204
		commit(ttWithinClause)
		goto st132
	tr235:
//line query/tokeniser.rl:183
		setText(ttAttributeSelector)
//line query/tokeniser.rl:184
		commit(ttAttributeSelector)
//line query/tokeniser.rl:189
		commit(ttPredicate)
		goto st132
	tr238:
//line query/tokeniser.rl:112
		setText(ttEventDeclAlias)
//line query/tokeniser.rl:113
		commit(ttEventDeclAlias)
//line query/tokeniser.rl:117
		commit(ttEventDecl)
//line query/tokeniser.rl:157
		commit(ttEventClause)
		goto st132
	tr240:
//line query/tokeniser.rl:128
		commit(ttAnyDecl)
//line query/tokeniser.rl:157
		commit(ttEventClause)
		goto st132
	tr242:
//line query/tokeniser.rl:151
		commit(ttSeqDecl)
//line query/tokeniser.rl:157
		commit(ttEventClause)
		goto st132
	st132:
		if p++; p == pe {
			goto _test_eof132
		}
	st_case_132:
//line query/tokeniser.go:717
		if data[p] == 32 {
			goto st132
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st132
		}
		goto st0
	st11:
		if p++; p == pe {
			goto _test_eof11
		}
	st_case_11:
		switch data[p] {
		case 72:
			goto st12
		case 73:
			goto st29
		case 104:
			goto st12
		case 105:
			goto st29
		}
		goto st0
	st12:
		if p++; p == pe {
			goto _test_eof12
		}
	st_case_12:
		switch data[p] {
		case 69:
			goto st13
		case 101:
			goto st13
		}
		goto st0
	st13:
		if p++; p == pe {
			goto _test_eof13
		}
	st_case_13:
		switch data[p] {
		case 82:
			goto st14
		case 114:
			goto st14
		}
		goto st0
	st14:
		if p++; p == pe {
			goto _test_eof14
		}
	st_case_14:
		switch data[p] {
		case 69:
			goto st15
		case 101:
			goto st15
		}
		goto st0
	st15:
		if p++; p == pe {
			goto _test_eof15
		}
	st_case_15:
		if data[p] == 32 {
			goto st16
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st16
		}
		goto st0
	tr43:
//line query/tokeniser.rl:171
		commit(ttConjunction)
		goto st16
	tr46:
//line query/tokeniser.rl:175
		commit(ttDisjunction)
		goto st16
	st16:
		if p++; p == pe {
			goto _test_eof16
		}
	st_case_16:
//line query/tokeniser.go:802
		switch data[p] {
		case 32:
			goto st16
		case 95:
			goto tr24
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto st16
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr24
			}
		default:
			goto tr24
		}
		goto st0
	tr24:
//line query/tokeniser.rl:187
		propose(ttPredicate)
//line query/tokeniser.rl:92
		mark = p
		goto st17
	st17:
		if p++; p == pe {
			goto _test_eof17
		}
	st_case_17:
//line query/tokeniser.go:833
		switch data[p] {
		case 32:
			goto tr25
		case 33:
			goto tr26
		case 46:
			goto tr27
		case 60:
			goto tr29
		case 61:
			goto tr30
		case 62:
			goto tr31
		case 95:
			goto st17
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr25
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st17
				}
			case data[p] >= 65:
				goto st17
			}
		default:
			goto st17
		}
		goto st0
	tr25:
//line query/tokeniser.rl:182
		propose(ttAttributeSelector)
//line query/tokeniser.rl:183
		setText(ttAttributeSelector)
//line query/tokeniser.rl:184
		commit(ttAttributeSelector)
		goto st18
	tr76:
//line query/tokeniser.rl:183
		setText(ttAttributeSelector)
//line query/tokeniser.rl:184
		commit(ttAttributeSelector)
		goto st18
	st18:
		if p++; p == pe {
			goto _test_eof18
		}
	st_case_18:
//line query/tokeniser.go:887
		switch data[p] {
		case 32:
			goto st18
		case 33:
			goto tr33
		case 60:
			goto tr34
		case 61:
			goto tr35
		case 62:
			goto tr36
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st18
		}
		goto st0
	tr26:
//line query/tokeniser.rl:182
		propose(ttAttributeSelector)
//line query/tokeniser.rl:183
		setText(ttAttributeSelector)
//line query/tokeniser.rl:184
		commit(ttAttributeSelector)
//line query/tokeniser.rl:161
		propose(ttEq)
		goto st19
	tr33:
//line query/tokeniser.rl:161
		propose(ttEq)
		goto st19
	tr77:
//line query/tokeniser.rl:183
		setText(ttAttributeSelector)
//line query/tokeniser.rl:184
		commit(ttAttributeSelector)
//line query/tokeniser.rl:161
		propose(ttEq)
		goto st19
	st19:
		if p++; p == pe {
			goto _test_eof19
		}
	st_case_19:
//line query/tokeniser.go:931
		if data[p] == 61 {
			goto st20
		}
		goto st0
	st20:
		if p++; p == pe {
			goto _test_eof20
		}
	st_case_20:
		switch data[p] {
		case 32:
			goto tr38
		case 95:
			goto tr39
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr38
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr39
			}
		default:
			goto tr39
		}
		goto st0
	tr38:
//line query/tokeniser.rl:161
		commit(ttEq)
		goto st21
	tr62:
//line query/tokeniser.rl:163
		commit(ttEq)
		goto st21
	tr65:
//line query/tokeniser.rl:165
		commit(ttEq)
		goto st21
	tr68:
//line query/tokeniser.rl:160
		commit(ttEq)
		goto st21
	tr70:
//line query/tokeniser.rl:162
		commit(ttEq)
		goto st21
	tr73:
//line query/tokeniser.rl:164
		commit(ttEq)
		goto st21
	st21:
		if p++; p == pe {
			goto _test_eof21
		}
	st_case_21:
//line query/tokeniser.go:989
		switch data[p] {
		case 32:
			goto st21
		case 95:
			goto tr41
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto st21
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr41
			}
		default:
			goto tr41
		}
		goto st0
	tr39:
//line query/tokeniser.rl:161
		commit(ttEq)
//line query/tokeniser.rl:92
		mark = p
		goto st133
	tr41:
//line query/tokeniser.rl:92
		mark = p
		goto st133
	tr64:
//line query/tokeniser.rl:163
		commit(ttEq)
//line query/tokeniser.rl:92
		mark = p
		goto st133
	tr66:
//line query/tokeniser.rl:165
		commit(ttEq)
//line query/tokeniser.rl:92
		mark = p
		goto st133
	tr69:
//line query/tokeniser.rl:160
		commit(ttEq)
//line query/tokeniser.rl:92
		mark = p
		goto st133
	tr72:
//line query/tokeniser.rl:162
		commit(ttEq)
//line query/tokeniser.rl:92
		mark = p
		goto st133
	tr74:
//line query/tokeniser.rl:164
		commit(ttEq)
//line query/tokeniser.rl:92
		mark = p
		goto st133
	st133:
		if p++; p == pe {
			goto _test_eof133
		}
	st_case_133:
//line query/tokeniser.go:1054
		switch data[p] {
		case 32:
			goto tr218
		case 46:
			goto tr219
		case 59:
			goto tr221
		case 95:
			goto st133
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr218
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st133
				}
			case data[p] >= 65:
				goto st133
			}
		default:
			goto st133
		}
		goto st0
	tr218:
//line query/tokeniser.rl:182
		propose(ttAttributeSelector)
//line query/tokeniser.rl:183
		setText(ttAttributeSelector)
//line query/tokeniser.rl:184
		commit(ttAttributeSelector)
//line query/tokeniser.rl:189
		commit(ttPredicate)
		goto st134
	tr233:
//line query/tokeniser.rl:183
		setText(ttAttributeSelector)
//line query/tokeniser.rl:184
		commit(ttAttributeSelector)
//line query/tokeniser.rl:189
		commit(ttPredicate)
		goto st134
	st134:
		if p++; p == pe {
			goto _test_eof134
		}
	st_case_134:
//line query/tokeniser.go:1106
		switch data[p] {
		case 32:
			goto st134
		case 38:
			goto tr223
		case 59:
			goto st132
		case 65:
			goto tr224
		case 79:
			goto tr225
		case 87:
			goto st28
		case 94:
			goto tr227
		case 97:
			goto tr224
		case 111:
			goto tr225
		case 119:
			goto st28
		case 124:
			goto tr228
		case 226:
			goto tr229
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st134
		}
		goto st0
	tr223:
//line query/tokeniser.rl:170
		propose(ttConjunction)
		goto st22
	st22:
		if p++; p == pe {
			goto _test_eof22
		}
	st_case_22:
//line query/tokeniser.go:1146
		if data[p] == 38 {
			goto st23
		}
		goto st0
	tr227:
//line query/tokeniser.rl:170
		propose(ttConjunction)
		goto st23
	st23:
		if p++; p == pe {
			goto _test_eof23
		}
	st_case_23:
//line query/tokeniser.go:1160
		if data[p] == 32 {
			goto tr43
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr43
		}
		goto st0
	tr224:
//line query/tokeniser.rl:170
		propose(ttConjunction)
		goto st24
	st24:
		if p++; p == pe {
			goto _test_eof24
		}
	st_case_24:
//line query/tokeniser.go:1177
		switch data[p] {
		case 78:
			goto st25
		case 110:
			goto st25
		}
		goto st0
	st25:
		if p++; p == pe {
			goto _test_eof25
		}
	st_case_25:
		switch data[p] {
		case 68:
			goto st23
		case 100:
			goto st23
		}
		goto st0
	tr225:
//line query/tokeniser.rl:174
		propose(ttDisjunction)
		goto st26
	st26:
		if p++; p == pe {
			goto _test_eof26
		}
	st_case_26:
//line query/tokeniser.go:1206
		switch data[p] {
		case 82:
			goto st27
		case 114:
			goto st27
		}
		goto st0
	st27:
		if p++; p == pe {
			goto _test_eof27
		}
	st_case_27:
		if data[p] == 32 {
			goto tr46
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr46
		}
		goto st0
	st28:
		if p++; p == pe {
			goto _test_eof28
		}
	st_case_28:
		switch data[p] {
		case 73:
			goto st29
		case 105:
			goto st29
		}
		goto st0
	st29:
		if p++; p == pe {
			goto _test_eof29
		}
	st_case_29:
		switch data[p] {
		case 84:
			goto st30
		case 116:
			goto st30
		}
		goto st0
	st30:
		if p++; p == pe {
			goto _test_eof30
		}
	st_case_30:
		switch data[p] {
		case 72:
			goto st31
		case 104:
			goto st31
		}
		goto st0
	st31:
		if p++; p == pe {
			goto _test_eof31
		}
	st_case_31:
		switch data[p] {
		case 73:
			goto st32
		case 105:
			goto st32
		}
		goto st0
	st32:
		if p++; p == pe {
			goto _test_eof32
		}
	st_case_32:
		switch data[p] {
		case 78:
			goto st33
		case 110:
			goto st33
		}
		goto st0
	st33:
		if p++; p == pe {
			goto _test_eof33
		}
	st_case_33:
		if data[p] == 32 {
			goto st34
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st34
		}
		goto st0
	st34:
		if p++; p == pe {
			goto _test_eof34
		}
	st_case_34:
		switch data[p] {
		case 32:
			goto st34
		case 43:
			goto tr52
		case 45:
			goto tr52
		}
		switch {
		case data[p] > 13:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr53
			}
		case data[p] >= 9:
			goto st34
		}
		goto st0
	tr52:
//line query/tokeniser.rl:203
		propose(ttWithinClause)
//line query/tokeniser.rl:92
		mark = p
//line query/tokeniser.rl:198
		propose(ttDuration)
		goto st35
	st35:
		if p++; p == pe {
			goto _test_eof35
		}
	st_case_35:
//line query/tokeniser.go:1333
		if 48 <= data[p] && data[p] <= 57 {
			goto st36
		}
		goto st0
	tr53:
//line query/tokeniser.rl:203
		propose(ttWithinClause)
//line query/tokeniser.rl:92
		mark = p
//line query/tokeniser.rl:198
		propose(ttDuration)
		goto st36
	st36:
		if p++; p == pe {
			goto _test_eof36
		}
	st_case_36:
//line query/tokeniser.go:1351
		switch data[p] {
		case 46:
			goto st37
		case 72:
			goto st135
		case 77:
			goto st137
		case 78:
			goto st39
		case 83:
			goto st135
		case 85:
			goto st39
		case 104:
			goto st135
		case 109:
			goto st137
		case 110:
			goto st39
		case 115:
			goto st135
		case 117:
			goto st39
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st36
		}
		goto st0
	st37:
		if p++; p == pe {
			goto _test_eof37
		}
	st_case_37:
		if 48 <= data[p] && data[p] <= 57 {
			goto st38
		}
		goto st0
	st38:
		if p++; p == pe {
			goto _test_eof38
		}
	st_case_38:
		switch data[p] {
		case 72:
			goto st135
		case 77:
			goto st137
		case 78:
			goto st39
		case 83:
			goto st135
		case 85:
			goto st39
		case 104:
			goto st135
		case 109:
			goto st137
		case 110:
			goto st39
		case 115:
			goto st135
		case 117:
			goto st39
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st38
		}
		goto st0
	st135:
		if p++; p == pe {
			goto _test_eof135
		}
	st_case_135:
		switch data[p] {
		case 32:
			goto tr230
		case 59:
			goto tr231
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr230
		}
		goto st0
	tr230:
//line query/tokeniser.rl:199
		setText(ttDuration)
//line query/tokeniser.rl:200
		commit(ttDuration)
//line query/tokeniser.rl:204
		commit(ttWithinClause)
		goto st136
	st136:
		if p++; p == pe {
			goto _test_eof136
		}
	st_case_136:
//line query/tokeniser.go:1448
		switch data[p] {
		case 32:
			goto st136
		case 59:
			goto st132
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st136
		}
		goto st0
	st137:
		if p++; p == pe {
			goto _test_eof137
		}
	st_case_137:
		switch data[p] {
		case 32:
			goto tr230
		case 59:
			goto tr231
		case 83:
			goto st135
		case 115:
			goto st135
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr230
		}
		goto st0
	st39:
		if p++; p == pe {
			goto _test_eof39
		}
	st_case_39:
		switch data[p] {
		case 83:
			goto st135
		case 115:
			goto st135
		}
		goto st0
	tr228:
//line query/tokeniser.rl:174
		propose(ttDisjunction)
		goto st40
	st40:
		if p++; p == pe {
			goto _test_eof40
		}
	st_case_40:
//line query/tokeniser.go:1499
		if data[p] == 124 {
			goto st27
		}
		goto st0
	tr229:
//line query/tokeniser.rl:174
		propose(ttDisjunction)
		goto st41
	st41:
		if p++; p == pe {
			goto _test_eof41
		}
	st_case_41:
//line query/tokeniser.go:1513
		if data[p] == 136 {
			goto st42
		}
		goto st0
	st42:
		if p++; p == pe {
			goto _test_eof42
		}
	st_case_42:
		if data[p] == 168 {
			goto st27
		}
		goto st0
	tr219:
//line query/tokeniser.rl:182
		propose(ttAttributeSelector)
		goto st43
	st43:
		if p++; p == pe {
			goto _test_eof43
		}
	st_case_43:
//line query/tokeniser.go:1536
		if data[p] == 95 {
			goto st138
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st138
			}
		case data[p] >= 65:
			goto st138
		}
		goto st0
	st138:
		if p++; p == pe {
			goto _test_eof138
		}
	st_case_138:
		switch data[p] {
		case 32:
			goto tr233
		case 46:
			goto st43
		case 59:
			goto tr235
		case 95:
			goto st138
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr233
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st138
				}
			case data[p] >= 65:
				goto st138
			}
		default:
			goto st138
		}
		goto st0
	tr29:
//line query/tokeniser.rl:182
		propose(ttAttributeSelector)
//line query/tokeniser.rl:183
		setText(ttAttributeSelector)
//line query/tokeniser.rl:184
		commit(ttAttributeSelector)
//line query/tokeniser.rl:163
		propose(ttEq)
//line query/tokeniser.rl:165
		propose(ttEq)
		goto st44
	tr34:
//line query/tokeniser.rl:163
		propose(ttEq)
//line query/tokeniser.rl:165
		propose(ttEq)
		goto st44
	tr79:
//line query/tokeniser.rl:183
		setText(ttAttributeSelector)
//line query/tokeniser.rl:184
		commit(ttAttributeSelector)
//line query/tokeniser.rl:163
		propose(ttEq)
//line query/tokeniser.rl:165
		propose(ttEq)
		goto st44
	st44:
		if p++; p == pe {
			goto _test_eof44
		}
	st_case_44:
//line query/tokeniser.go:1615
		switch data[p] {
		case 32:
			goto tr62
		case 61:
			goto st45
		case 95:
			goto tr64
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr62
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr64
			}
		default:
			goto tr64
		}
		goto st0
	st45:
		if p++; p == pe {
			goto _test_eof45
		}
	st_case_45:
		switch data[p] {
		case 32:
			goto tr65
		case 95:
			goto tr66
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr65
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr66
			}
		default:
			goto tr66
		}
		goto st0
	tr30:
//line query/tokeniser.rl:182
		propose(ttAttributeSelector)
//line query/tokeniser.rl:183
		setText(ttAttributeSelector)
//line query/tokeniser.rl:184
		commit(ttAttributeSelector)
//line query/tokeniser.rl:160
		propose(ttEq)
		goto st46
	tr35:
//line query/tokeniser.rl:160
		propose(ttEq)
		goto st46
	tr80:
//line query/tokeniser.rl:183
		setText(ttAttributeSelector)
//line query/tokeniser.rl:184
		commit(ttAttributeSelector)
//line query/tokeniser.rl:160
		propose(ttEq)
		goto st46
	st46:
		if p++; p == pe {
			goto _test_eof46
		}
	st_case_46:
//line query/tokeniser.go:1688
		if data[p] == 61 {
			goto st47
		}
		goto st0
	st47:
		if p++; p == pe {
			goto _test_eof47
		}
	st_case_47:
		switch data[p] {
		case 32:
			goto tr68
		case 95:
			goto tr69
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr68
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr69
			}
		default:
			goto tr69
		}
		goto st0
	tr31:
//line query/tokeniser.rl:182
		propose(ttAttributeSelector)
//line query/tokeniser.rl:183
		setText(ttAttributeSelector)
//line query/tokeniser.rl:184
		commit(ttAttributeSelector)
//line query/tokeniser.rl:162
		propose(ttEq)
//line query/tokeniser.rl:164
		propose(ttEq)
		goto st48
	tr36:
//line query/tokeniser.rl:162
		propose(ttEq)
//line query/tokeniser.rl:164
		propose(ttEq)
		goto st48
	tr81:
//line query/tokeniser.rl:183
		setText(ttAttributeSelector)
//line query/tokeniser.rl:184
		commit(ttAttributeSelector)
//line query/tokeniser.rl:162
		propose(ttEq)
//line query/tokeniser.rl:164
		propose(ttEq)
		goto st48
	st48:
		if p++; p == pe {
			goto _test_eof48
		}
	st_case_48:
//line query/tokeniser.go:1750
		switch data[p] {
		case 32:
			goto tr70
		case 61:
			goto st49
		case 95:
			goto tr72
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr70
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr72
			}
		default:
			goto tr72
		}
		goto st0
	st49:
		if p++; p == pe {
			goto _test_eof49
		}
	st_case_49:
		switch data[p] {
		case 32:
			goto tr73
		case 95:
			goto tr74
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr73
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr74
			}
		default:
			goto tr74
		}
		goto st0
	tr27:
//line query/tokeniser.rl:182
		propose(ttAttributeSelector)
		goto st50
	st50:
		if p++; p == pe {
			goto _test_eof50
		}
	st_case_50:
//line query/tokeniser.go:1805
		if data[p] == 95 {
			goto st51
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st51
			}
		case data[p] >= 65:
			goto st51
		}
		goto st0
	st51:
		if p++; p == pe {
			goto _test_eof51
		}
	st_case_51:
		switch data[p] {
		case 32:
			goto tr76
		case 33:
			goto tr77
		case 46:
			goto st50
		case 60:
			goto tr79
		case 61:
			goto tr80
		case 62:
			goto tr81
		case 95:
			goto st51
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr76
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st51
				}
			case data[p] >= 65:
				goto st51
			}
		default:
			goto st51
		}
		goto st0
	tr16:
//line query/tokeniser.rl:116
		propose(ttEventDecl)
//line query/tokeniser.rl:92
		mark = p
//line query/tokeniser.rl:105
		propose(ttEventDeclType)
//line query/tokeniser.rl:124
		propose(ttAnyDecl)
		goto st52
	st52:
		if p++; p == pe {
			goto _test_eof52
		}
	st_case_52:
//line query/tokeniser.go:1872
		switch data[p] {
		case 32:
			goto tr82
		case 78:
			goto st58
		case 95:
			goto st57
		case 110:
			goto st58
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr82
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st57
				}
			case data[p] >= 65:
				goto st57
			}
		default:
			goto st57
		}
		goto st0
	tr82:
//line query/tokeniser.rl:106
		setText(ttEventDeclType)
//line query/tokeniser.rl:107
		commit(ttEventDeclType)
		goto st53
	st53:
		if p++; p == pe {
			goto _test_eof53
		}
	st_case_53:
//line query/tokeniser.go:1912
		switch data[p] {
		case 32:
			goto st53
		case 95:
			goto tr86
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto st53
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr86
			}
		default:
			goto tr86
		}
		goto st0
	tr86:
//line query/tokeniser.rl:92
		mark = p
//line query/tokeniser.rl:111
		propose(ttEventDeclAlias)
		goto st54
	st54:
		if p++; p == pe {
			goto _test_eof54
		}
	st_case_54:
//line query/tokeniser.go:1943
		switch data[p] {
		case 32:
			goto tr87
		case 41:
			goto tr88
		case 44:
			goto tr89
		case 95:
			goto st54
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr87
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st54
				}
			case data[p] >= 65:
				goto st54
			}
		default:
			goto st54
		}
		goto st0
	tr87:
//line query/tokeniser.rl:112
		setText(ttEventDeclAlias)
//line query/tokeniser.rl:113
		commit(ttEventDeclAlias)
//line query/tokeniser.rl:117
		commit(ttEventDecl)
		goto st55
	tr98:
//line query/tokeniser.rl:128
		commit(ttAnyDecl)
		goto st55
	st55:
		if p++; p == pe {
			goto _test_eof55
		}
	st_case_55:
//line query/tokeniser.go:1989
		switch data[p] {
		case 32:
			goto st55
		case 41:
			goto st130
		case 44:
			goto st56
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st55
		}
		goto st0
	tr89:
//line query/tokeniser.rl:112
		setText(ttEventDeclAlias)
//line query/tokeniser.rl:113
		commit(ttEventDeclAlias)
//line query/tokeniser.rl:117
		commit(ttEventDecl)
		goto st56
	tr100:
//line query/tokeniser.rl:128
		commit(ttAnyDecl)
		goto st56
	st56:
		if p++; p == pe {
			goto _test_eof56
		}
	st_case_56:
//line query/tokeniser.go:2019
		switch data[p] {
		case 32:
			goto st56
		case 65:
			goto tr16
		case 95:
			goto tr17
		case 97:
			goto tr16
		}
		switch {
		case data[p] < 66:
			if 9 <= data[p] && data[p] <= 13 {
				goto st56
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr17
			}
		default:
			goto tr17
		}
		goto st0
	tr17:
//line query/tokeniser.rl:116
		propose(ttEventDecl)
//line query/tokeniser.rl:92
		mark = p
//line query/tokeniser.rl:105
		propose(ttEventDeclType)
		goto st57
	st57:
		if p++; p == pe {
			goto _test_eof57
		}
	st_case_57:
//line query/tokeniser.go:2056
		switch data[p] {
		case 32:
			goto tr82
		case 95:
			goto st57
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr82
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st57
				}
			case data[p] >= 65:
				goto st57
			}
		default:
			goto st57
		}
		goto st0
	st58:
		if p++; p == pe {
			goto _test_eof58
		}
	st_case_58:
		switch data[p] {
		case 32:
			goto tr82
		case 89:
			goto st59
		case 95:
			goto st57
		case 121:
			goto st59
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr82
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st57
				}
			case data[p] >= 65:
				goto st57
			}
		default:
			goto st57
		}
		goto st0
	st59:
		if p++; p == pe {
			goto _test_eof59
		}
	st_case_59:
		switch data[p] {
		case 32:
			goto tr94
		case 40:
			goto st61
		case 95:
			goto st57
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr94
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st57
				}
			case data[p] >= 65:
				goto st57
			}
		default:
			goto st57
		}
		goto st0
	tr94:
//line query/tokeniser.rl:106
		setText(ttEventDeclType)
//line query/tokeniser.rl:107
		commit(ttEventDeclType)
		goto st60
	st60:
		if p++; p == pe {
			goto _test_eof60
		}
	st_case_60:
//line query/tokeniser.go:2156
		switch data[p] {
		case 32:
			goto st53
		case 40:
			goto st61
		case 95:
			goto tr86
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto st53
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr86
			}
		default:
			goto tr86
		}
		goto st0
	st61:
		if p++; p == pe {
			goto _test_eof61
		}
	st_case_61:
		switch data[p] {
		case 32:
			goto st61
		case 41:
			goto st62
		case 95:
			goto tr97
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto st61
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr97
			}
		default:
			goto tr97
		}
		goto st0
	tr106:
//line query/tokeniser.rl:112
		setText(ttEventDeclAlias)
//line query/tokeniser.rl:113
		commit(ttEventDeclAlias)
//line query/tokeniser.rl:117
		commit(ttEventDecl)
		goto st62
	st62:
		if p++; p == pe {
			goto _test_eof62
		}
	st_case_62:
//line query/tokeniser.go:2217
		switch data[p] {
		case 32:
			goto tr98
		case 41:
			goto tr99
		case 44:
			goto tr100
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr98
		}
		goto st0
	tr97:
//line query/tokeniser.rl:116
		propose(ttEventDecl)
//line query/tokeniser.rl:92
		mark = p
//line query/tokeniser.rl:105
		propose(ttEventDeclType)
		goto st63
	st63:
		if p++; p == pe {
			goto _test_eof63
		}
	st_case_63:
//line query/tokeniser.go:2243
		switch data[p] {
		case 32:
			goto tr101
		case 95:
			goto st63
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr101
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st63
				}
			case data[p] >= 65:
				goto st63
			}
		default:
			goto st63
		}
		goto st0
	tr101:
//line query/tokeniser.rl:106
		setText(ttEventDeclType)
//line query/tokeniser.rl:107
		commit(ttEventDeclType)
		goto st64
	st64:
		if p++; p == pe {
			goto _test_eof64
		}
	st_case_64:
//line query/tokeniser.go:2279
		switch data[p] {
		case 32:
			goto st64
		case 95:
			goto tr104
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto st64
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr104
			}
		default:
			goto tr104
		}
		goto st0
	tr104:
//line query/tokeniser.rl:92
		mark = p
//line query/tokeniser.rl:111
		propose(ttEventDeclAlias)
		goto st65
	st65:
		if p++; p == pe {
			goto _test_eof65
		}
	st_case_65:
//line query/tokeniser.go:2310
		switch data[p] {
		case 32:
			goto tr105
		case 41:
			goto tr106
		case 44:
			goto tr107
		case 95:
			goto st65
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr105
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st65
				}
			case data[p] >= 65:
				goto st65
			}
		default:
			goto st65
		}
		goto st0
	tr105:
//line query/tokeniser.rl:112
		setText(ttEventDeclAlias)
//line query/tokeniser.rl:113
		commit(ttEventDeclAlias)
//line query/tokeniser.rl:117
		commit(ttEventDecl)
		goto st66
	st66:
		if p++; p == pe {
			goto _test_eof66
		}
	st_case_66:
//line query/tokeniser.go:2352
		switch data[p] {
		case 32:
			goto st66
		case 41:
			goto st62
		case 44:
			goto st67
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st66
		}
		goto st0
	tr107:
//line query/tokeniser.rl:112
		setText(ttEventDeclAlias)
//line query/tokeniser.rl:113
		commit(ttEventDeclAlias)
//line query/tokeniser.rl:117
		commit(ttEventDecl)
		goto st67
	st67:
		if p++; p == pe {
			goto _test_eof67
		}
	st_case_67:
//line query/tokeniser.go:2378
		switch data[p] {
		case 32:
			goto st67
		case 95:
			goto tr97
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto st67
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr97
			}
		default:
			goto tr97
		}
		goto st0
	tr9:
//line query/tokeniser.rl:156
		propose(ttEventClause)
//line query/tokeniser.rl:124
		propose(ttAnyDecl)
//line query/tokeniser.rl:116
		propose(ttEventDecl)
//line query/tokeniser.rl:92
		mark = p
//line query/tokeniser.rl:105
		propose(ttEventDeclType)
		goto st68
	st68:
		if p++; p == pe {
			goto _test_eof68
		}
	st_case_68:
//line query/tokeniser.go:2415
		switch data[p] {
		case 32:
			goto tr111
		case 78:
			goto st71
		case 95:
			goto st70
		case 110:
			goto st71
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr111
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st70
				}
			case data[p] >= 65:
				goto st70
			}
		default:
			goto st70
		}
		goto st0
	tr111:
//line query/tokeniser.rl:106
		setText(ttEventDeclType)
//line query/tokeniser.rl:107
		commit(ttEventDeclType)
		goto st69
	st69:
		if p++; p == pe {
			goto _test_eof69
		}
	st_case_69:
//line query/tokeniser.go:2455
		switch data[p] {
		case 32:
			goto st69
		case 95:
			goto tr115
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto st69
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr115
			}
		default:
			goto tr115
		}
		goto st0
	tr115:
//line query/tokeniser.rl:92
		mark = p
//line query/tokeniser.rl:111
		propose(ttEventDeclAlias)
		goto st139
	st139:
		if p++; p == pe {
			goto _test_eof139
		}
	st_case_139:
//line query/tokeniser.go:2486
		switch data[p] {
		case 32:
			goto tr236
		case 59:
			goto tr238
		case 95:
			goto st139
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr236
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st139
				}
			case data[p] >= 65:
				goto st139
			}
		default:
			goto st139
		}
		goto st0
	tr10:
//line query/tokeniser.rl:156
		propose(ttEventClause)
//line query/tokeniser.rl:116
		propose(ttEventDecl)
//line query/tokeniser.rl:92
		mark = p
//line query/tokeniser.rl:105
		propose(ttEventDeclType)
		goto st70
	st70:
		if p++; p == pe {
			goto _test_eof70
		}
	st_case_70:
//line query/tokeniser.go:2528
		switch data[p] {
		case 32:
			goto tr111
		case 95:
			goto st70
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr111
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st70
				}
			case data[p] >= 65:
				goto st70
			}
		default:
			goto st70
		}
		goto st0
	st71:
		if p++; p == pe {
			goto _test_eof71
		}
	st_case_71:
		switch data[p] {
		case 32:
			goto tr111
		case 89:
			goto st72
		case 95:
			goto st70
		case 121:
			goto st72
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr111
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st70
				}
			case data[p] >= 65:
				goto st70
			}
		default:
			goto st70
		}
		goto st0
	st72:
		if p++; p == pe {
			goto _test_eof72
		}
	st_case_72:
		switch data[p] {
		case 32:
			goto tr117
		case 40:
			goto st74
		case 95:
			goto st70
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr117
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st70
				}
			case data[p] >= 65:
				goto st70
			}
		default:
			goto st70
		}
		goto st0
	tr117:
//line query/tokeniser.rl:106
		setText(ttEventDeclType)
//line query/tokeniser.rl:107
		commit(ttEventDeclType)
		goto st73
	st73:
		if p++; p == pe {
			goto _test_eof73
		}
	st_case_73:
//line query/tokeniser.go:2628
		switch data[p] {
		case 32:
			goto st69
		case 40:
			goto st74
		case 95:
			goto tr115
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto st69
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr115
			}
		default:
			goto tr115
		}
		goto st0
	st74:
		if p++; p == pe {
			goto _test_eof74
		}
	st_case_74:
		switch data[p] {
		case 32:
			goto st74
		case 41:
			goto st140
		case 95:
			goto tr120
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto st74
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr120
			}
		default:
			goto tr120
		}
		goto st0
	tr126:
//line query/tokeniser.rl:112
		setText(ttEventDeclAlias)
//line query/tokeniser.rl:113
		commit(ttEventDeclAlias)
//line query/tokeniser.rl:117
		commit(ttEventDecl)
		goto st140
	st140:
		if p++; p == pe {
			goto _test_eof140
		}
	st_case_140:
//line query/tokeniser.go:2689
		switch data[p] {
		case 32:
			goto tr239
		case 59:
			goto tr240
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr239
		}
		goto st0
	tr120:
//line query/tokeniser.rl:116
		propose(ttEventDecl)
//line query/tokeniser.rl:92
		mark = p
//line query/tokeniser.rl:105
		propose(ttEventDeclType)
		goto st75
	st75:
		if p++; p == pe {
			goto _test_eof75
		}
	st_case_75:
//line query/tokeniser.go:2713
		switch data[p] {
		case 32:
			goto tr121
		case 95:
			goto st75
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr121
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st75
				}
			case data[p] >= 65:
				goto st75
			}
		default:
			goto st75
		}
		goto st0
	tr121:
//line query/tokeniser.rl:106
		setText(ttEventDeclType)
//line query/tokeniser.rl:107
		commit(ttEventDeclType)
		goto st76
	st76:
		if p++; p == pe {
			goto _test_eof76
		}
	st_case_76:
//line query/tokeniser.go:2749
		switch data[p] {
		case 32:
			goto st76
		case 95:
			goto tr124
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto st76
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr124
			}
		default:
			goto tr124
		}
		goto st0
	tr124:
//line query/tokeniser.rl:92
		mark = p
//line query/tokeniser.rl:111
		propose(ttEventDeclAlias)
		goto st77
	st77:
		if p++; p == pe {
			goto _test_eof77
		}
	st_case_77:
//line query/tokeniser.go:2780
		switch data[p] {
		case 32:
			goto tr125
		case 41:
			goto tr126
		case 44:
			goto tr127
		case 95:
			goto st77
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr125
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st77
				}
			case data[p] >= 65:
				goto st77
			}
		default:
			goto st77
		}
		goto st0
	tr125:
//line query/tokeniser.rl:112
		setText(ttEventDeclAlias)
//line query/tokeniser.rl:113
		commit(ttEventDeclAlias)
//line query/tokeniser.rl:117
		commit(ttEventDecl)
		goto st78
	st78:
		if p++; p == pe {
			goto _test_eof78
		}
	st_case_78:
//line query/tokeniser.go:2822
		switch data[p] {
		case 32:
			goto st78
		case 41:
			goto st140
		case 44:
			goto st79
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st78
		}
		goto st0
	tr127:
//line query/tokeniser.rl:112
		setText(ttEventDeclAlias)
//line query/tokeniser.rl:113
		commit(ttEventDeclAlias)
//line query/tokeniser.rl:117
		commit(ttEventDecl)
		goto st79
	st79:
		if p++; p == pe {
			goto _test_eof79
		}
	st_case_79:
//line query/tokeniser.go:2848
		switch data[p] {
		case 32:
			goto st79
		case 95:
			goto tr120
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto st79
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr120
			}
		default:
			goto tr120
		}
		goto st0
	tr11:
//line query/tokeniser.rl:156
		propose(ttEventClause)
//line query/tokeniser.rl:136
		propose(ttNegatedDecl)
//line query/tokeniser.rl:116
		propose(ttEventDecl)
//line query/tokeniser.rl:92
		mark = p
//line query/tokeniser.rl:105
		propose(ttEventDeclType)
		goto st80
	st80:
		if p++; p == pe {
			goto _test_eof80
		}
	st_case_80:
//line query/tokeniser.go:2885
		switch data[p] {
		case 32:
			goto tr111
		case 79:
			goto st81
		case 95:
			goto st70
		case 111:
			goto st81
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr111
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st70
				}
			case data[p] >= 65:
				goto st70
			}
		default:
			goto st70
		}
		goto st0
	st81:
		if p++; p == pe {
			goto _test_eof81
		}
	st_case_81:
		switch data[p] {
		case 32:
			goto tr111
		case 84:
			goto st82
		case 95:
			goto st70
		case 116:
			goto st82
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr111
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st70
				}
			case data[p] >= 65:
				goto st70
			}
		default:
			goto st70
		}
		goto st0
	st82:
		if p++; p == pe {
			goto _test_eof82
		}
	st_case_82:
		switch data[p] {
		case 32:
			goto tr133
		case 40:
			goto st10
		case 95:
			goto st70
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr133
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st70
				}
			case data[p] >= 65:
				goto st70
			}
		default:
			goto st70
		}
		goto st0
	tr133:
//line query/tokeniser.rl:106
		setText(ttEventDeclType)
//line query/tokeniser.rl:107
		commit(ttEventDeclType)
		goto st83
	st83:
		if p++; p == pe {
			goto _test_eof83
		}
	st_case_83:
//line query/tokeniser.go:2989
		switch data[p] {
		case 32:
			goto st69
		case 40:
			goto st10
		case 95:
			goto tr115
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto st69
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr115
			}
		default:
			goto tr115
		}
		goto st0
	tr12:
//line query/tokeniser.rl:156
		propose(ttEventClause)
//line query/tokeniser.rl:147
		propose(ttSeqDecl)
//line query/tokeniser.rl:116
		propose(ttEventDecl)
//line query/tokeniser.rl:92
		mark = p
//line query/tokeniser.rl:105
		propose(ttEventDeclType)
		goto st84
	st84:
		if p++; p == pe {
			goto _test_eof84
		}
	st_case_84:
//line query/tokeniser.go:3028
		switch data[p] {
		case 32:
			goto tr111
		case 69:
			goto st85
		case 95:
			goto st70
		case 101:
			goto st85
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr111
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st70
				}
			case data[p] >= 65:
				goto st70
			}
		default:
			goto st70
		}
		goto st0
	st85:
		if p++; p == pe {
			goto _test_eof85
		}
	st_case_85:
		switch data[p] {
		case 32:
			goto tr111
		case 81:
			goto st86
		case 95:
			goto st70
		case 113:
			goto st86
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr111
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st70
				}
			case data[p] >= 65:
				goto st70
			}
		default:
			goto st70
		}
		goto st0
	st86:
		if p++; p == pe {
			goto _test_eof86
		}
	st_case_86:
		switch data[p] {
		case 32:
			goto tr136
		case 40:
			goto st88
		case 95:
			goto st70
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr136
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st70
				}
			case data[p] >= 65:
				goto st70
			}
		default:
			goto st70
		}
		goto st0
	tr136:
//line query/tokeniser.rl:106
		setText(ttEventDeclType)
//line query/tokeniser.rl:107
		commit(ttEventDeclType)
		goto st87
	st87:
		if p++; p == pe {
			goto _test_eof87
		}
	st_case_87:
//line query/tokeniser.go:3132
		switch data[p] {
		case 32:
			goto st69
		case 40:
			goto st88
		case 95:
			goto tr115
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto st69
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr115
			}
		default:
			goto tr115
		}
		goto st0
	st88:
		if p++; p == pe {
			goto _test_eof88
		}
	st_case_88:
		switch data[p] {
		case 32:
			goto st89
		case 33:
			goto tr139
		case 41:
			goto st141
		case 65:
			goto tr141
		case 78:
			goto tr143
		case 95:
			goto tr142
		case 97:
			goto tr141
		case 110:
			goto tr143
		}
		switch {
		case data[p] < 66:
			if 9 <= data[p] && data[p] <= 13 {
				goto st89
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr142
			}
		default:
			goto tr142
		}
		goto st0
	st89:
		if p++; p == pe {
			goto _test_eof89
		}
	st_case_89:
		switch data[p] {
		case 32:
			goto st89
		case 41:
			goto st141
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st89
		}
		goto st0
	tr160:
//line query/tokeniser.rl:112
		setText(ttEventDeclAlias)
//line query/tokeniser.rl:113
		commit(ttEventDeclAlias)
//line query/tokeniser.rl:117
		commit(ttEventDecl)
		goto st141
	tr169:
//line query/tokeniser.rl:128
		commit(ttAnyDecl)
		goto st141
	tr150:
//line query/tokeniser.rl:140
		commit(ttNegatedDecl)
		goto st141
	st141:
		if p++; p == pe {
			goto _test_eof141
		}
	st_case_141:
//line query/tokeniser.go:3226
		switch data[p] {
		case 32:
			goto tr241
		case 59:
			goto tr242
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr241
		}
		goto st0
	tr139:
//line query/tokeniser.rl:136
		propose(ttNegatedDecl)
		goto st90
	st90:
		if p++; p == pe {
			goto _test_eof90
		}
	st_case_90:
//line query/tokeniser.go:3246
		switch data[p] {
		case 32:
			goto st91
		case 40:
			goto st92
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st91
		}
		goto st0
	st91:
		if p++; p == pe {
			goto _test_eof91
		}
	st_case_91:
		if data[p] == 40 {
			goto st92
		}
		goto st0
	st92:
		if p++; p == pe {
			goto _test_eof92
		}
	st_case_92:
		switch data[p] {
		case 32:
			goto st92
		case 41:
			goto st93
		case 65:
			goto tr147
		case 95:
			goto tr148
		case 97:
			goto tr147
		}
		switch {
		case data[p] < 66:
			if 9 <= data[p] && data[p] <= 13 {
				goto st92
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr148
			}
		default:
			goto tr148
		}
		goto st0
	tr190:
//line query/tokeniser.rl:112
		setText(ttEventDeclAlias)
//line query/tokeniser.rl:113
		commit(ttEventDeclAlias)
//line query/tokeniser.rl:117
		commit(ttEventDecl)
		goto st93
	tr201:
//line query/tokeniser.rl:128
		commit(ttAnyDecl)
		goto st93
	st93:
		if p++; p == pe {
			goto _test_eof93
		}
	st_case_93:
//line query/tokeniser.go:3313
		switch data[p] {
		case 32:
			goto tr149
		case 41:
			goto tr150
		case 44:
			goto tr151
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr149
		}
		goto st0
	tr159:
//line query/tokeniser.rl:112
		setText(ttEventDeclAlias)
//line query/tokeniser.rl:113
		commit(ttEventDeclAlias)
//line query/tokeniser.rl:117
		commit(ttEventDecl)
		goto st94
	tr168:
//line query/tokeniser.rl:128
		commit(ttAnyDecl)
		goto st94
	tr149:
//line query/tokeniser.rl:140
		commit(ttNegatedDecl)
		goto st94
	st94:
		if p++; p == pe {
			goto _test_eof94
		}
	st_case_94:
//line query/tokeniser.go:3347
		switch data[p] {
		case 32:
			goto st94
		case 41:
			goto st141
		case 44:
			goto st95
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st94
		}
		goto st0
	tr161:
//line query/tokeniser.rl:112
		setText(ttEventDeclAlias)
//line query/tokeniser.rl:113
		commit(ttEventDeclAlias)
//line query/tokeniser.rl:117
		commit(ttEventDecl)
		goto st95
	tr170:
//line query/tokeniser.rl:128
		commit(ttAnyDecl)
		goto st95
	tr151:
//line query/tokeniser.rl:140
		commit(ttNegatedDecl)
		goto st95
	st95:
		if p++; p == pe {
			goto _test_eof95
		}
	st_case_95:
//line query/tokeniser.go:3381
		switch data[p] {
		case 32:
			goto st95
		case 33:
			goto tr139
		case 65:
			goto tr141
		case 78:
			goto tr143
		case 95:
			goto tr142
		case 97:
			goto tr141
		case 110:
			goto tr143
		}
		switch {
		case data[p] < 66:
			if 9 <= data[p] && data[p] <= 13 {
				goto st95
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr142
			}
		default:
			goto tr142
		}
		goto st0
	tr141:
//line query/tokeniser.rl:116
		propose(ttEventDecl)
//line query/tokeniser.rl:92
		mark = p
//line query/tokeniser.rl:105
		propose(ttEventDeclType)
//line query/tokeniser.rl:124
		propose(ttAnyDecl)
		goto st96
	st96:
		if p++; p == pe {
			goto _test_eof96
		}
	st_case_96:
//line query/tokeniser.go:3426
		switch data[p] {
		case 32:
			goto tr154
		case 78:
			goto st100
		case 95:
			goto st99
		case 110:
			goto st100
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr154
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st99
				}
			case data[p] >= 65:
				goto st99
			}
		default:
			goto st99
		}
		goto st0
	tr154:
//line query/tokeniser.rl:106
		setText(ttEventDeclType)
//line query/tokeniser.rl:107
		commit(ttEventDeclType)
		goto st97
	st97:
		if p++; p == pe {
			goto _test_eof97
		}
	st_case_97:
//line query/tokeniser.go:3466
		switch data[p] {
		case 32:
			goto st97
		case 95:
			goto tr158
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto st97
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr158
			}
		default:
			goto tr158
		}
		goto st0
	tr158:
//line query/tokeniser.rl:92
		mark = p
//line query/tokeniser.rl:111
		propose(ttEventDeclAlias)
		goto st98
	st98:
		if p++; p == pe {
			goto _test_eof98
		}
	st_case_98:
//line query/tokeniser.go:3497
		switch data[p] {
		case 32:
			goto tr159
		case 41:
			goto tr160
		case 44:
			goto tr161
		case 95:
			goto st98
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr159
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st98
				}
			case data[p] >= 65:
				goto st98
			}
		default:
			goto st98
		}
		goto st0
	tr142:
//line query/tokeniser.rl:116
		propose(ttEventDecl)
//line query/tokeniser.rl:92
		mark = p
//line query/tokeniser.rl:105
		propose(ttEventDeclType)
		goto st99
	st99:
		if p++; p == pe {
			goto _test_eof99
		}
	st_case_99:
//line query/tokeniser.go:3539
		switch data[p] {
		case 32:
			goto tr154
		case 95:
			goto st99
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr154
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st99
				}
			case data[p] >= 65:
				goto st99
			}
		default:
			goto st99
		}
		goto st0
	st100:
		if p++; p == pe {
			goto _test_eof100
		}
	st_case_100:
		switch data[p] {
		case 32:
			goto tr154
		case 89:
			goto st101
		case 95:
			goto st99
		case 121:
			goto st101
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr154
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st99
				}
			case data[p] >= 65:
				goto st99
			}
		default:
			goto st99
		}
		goto st0
	st101:
		if p++; p == pe {
			goto _test_eof101
		}
	st_case_101:
		switch data[p] {
		case 32:
			goto tr164
		case 40:
			goto st103
		case 95:
			goto st99
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr164
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st99
				}
			case data[p] >= 65:
				goto st99
			}
		default:
			goto st99
		}
		goto st0
	tr164:
//line query/tokeniser.rl:106
		setText(ttEventDeclType)
//line query/tokeniser.rl:107
		commit(ttEventDeclType)
		goto st102
	st102:
		if p++; p == pe {
			goto _test_eof102
		}
	st_case_102:
//line query/tokeniser.go:3639
		switch data[p] {
		case 32:
			goto st97
		case 40:
			goto st103
		case 95:
			goto tr158
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto st97
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr158
			}
		default:
			goto tr158
		}
		goto st0
	st103:
		if p++; p == pe {
			goto _test_eof103
		}
	st_case_103:
		switch data[p] {
		case 32:
			goto st103
		case 41:
			goto st104
		case 95:
			goto tr167
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto st103
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr167
			}
		default:
			goto tr167
		}
		goto st0
	tr176:
//line query/tokeniser.rl:112
		setText(ttEventDeclAlias)
//line query/tokeniser.rl:113
		commit(ttEventDeclAlias)
//line query/tokeniser.rl:117
		commit(ttEventDecl)
		goto st104
	st104:
		if p++; p == pe {
			goto _test_eof104
		}
	st_case_104:
//line query/tokeniser.go:3700
		switch data[p] {
		case 32:
			goto tr168
		case 41:
			goto tr169
		case 44:
			goto tr170
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr168
		}
		goto st0
	tr167:
//line query/tokeniser.rl:116
		propose(ttEventDecl)
//line query/tokeniser.rl:92
		mark = p
//line query/tokeniser.rl:105
		propose(ttEventDeclType)
		goto st105
	st105:
		if p++; p == pe {
			goto _test_eof105
		}
	st_case_105:
//line query/tokeniser.go:3726
		switch data[p] {
		case 32:
			goto tr171
		case 95:
			goto st105
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr171
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st105
				}
			case data[p] >= 65:
				goto st105
			}
		default:
			goto st105
		}
		goto st0
	tr171:
//line query/tokeniser.rl:106
		setText(ttEventDeclType)
//line query/tokeniser.rl:107
		commit(ttEventDeclType)
		goto st106
	st106:
		if p++; p == pe {
			goto _test_eof106
		}
	st_case_106:
//line query/tokeniser.go:3762
		switch data[p] {
		case 32:
			goto st106
		case 95:
			goto tr174
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto st106
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr174
			}
		default:
			goto tr174
		}
		goto st0
	tr174:
//line query/tokeniser.rl:92
		mark = p
//line query/tokeniser.rl:111
		propose(ttEventDeclAlias)
		goto st107
	st107:
		if p++; p == pe {
			goto _test_eof107
		}
	st_case_107:
//line query/tokeniser.go:3793
		switch data[p] {
		case 32:
			goto tr175
		case 41:
			goto tr176
		case 44:
			goto tr177
		case 95:
			goto st107
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr175
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st107
				}
			case data[p] >= 65:
				goto st107
			}
		default:
			goto st107
		}
		goto st0
	tr175:
//line query/tokeniser.rl:112
		setText(ttEventDeclAlias)
//line query/tokeniser.rl:113
		commit(ttEventDeclAlias)
//line query/tokeniser.rl:117
		commit(ttEventDecl)
		goto st108
	st108:
		if p++; p == pe {
			goto _test_eof108
		}
	st_case_108:
//line query/tokeniser.go:3835
		switch data[p] {
		case 32:
			goto st108
		case 41:
			goto st104
		case 44:
			goto st109
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st108
		}
		goto st0
	tr177:
//line query/tokeniser.rl:112
		setText(ttEventDeclAlias)
//line query/tokeniser.rl:113
		commit(ttEventDeclAlias)
//line query/tokeniser.rl:117
		commit(ttEventDecl)
		goto st109
	st109:
		if p++; p == pe {
			goto _test_eof109
		}
	st_case_109:
//line query/tokeniser.go:3861
		switch data[p] {
		case 32:
			goto st109
		case 95:
			goto tr167
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto st109
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr167
			}
		default:
			goto tr167
		}
		goto st0
	tr143:
//line query/tokeniser.rl:116
		propose(ttEventDecl)
//line query/tokeniser.rl:92
		mark = p
//line query/tokeniser.rl:105
		propose(ttEventDeclType)
//line query/tokeniser.rl:136
		propose(ttNegatedDecl)
		goto st110
	st110:
		if p++; p == pe {
			goto _test_eof110
		}
	st_case_110:
//line query/tokeniser.go:3896
		switch data[p] {
		case 32:
			goto tr154
		case 79:
			goto st111
		case 95:
			goto st99
		case 111:
			goto st111
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr154
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st99
				}
			case data[p] >= 65:
				goto st99
			}
		default:
			goto st99
		}
		goto st0
	st111:
		if p++; p == pe {
			goto _test_eof111
		}
	st_case_111:
		switch data[p] {
		case 32:
			goto tr154
		case 84:
			goto st112
		case 95:
			goto st99
		case 116:
			goto st112
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr154
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st99
				}
			case data[p] >= 65:
				goto st99
			}
		default:
			goto st99
		}
		goto st0
	st112:
		if p++; p == pe {
			goto _test_eof112
		}
	st_case_112:
		switch data[p] {
		case 32:
			goto tr183
		case 40:
			goto st92
		case 95:
			goto st99
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr183
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st99
				}
			case data[p] >= 65:
				goto st99
			}
		default:
			goto st99
		}
		goto st0
	tr183:
//line query/tokeniser.rl:106
		setText(ttEventDeclType)
//line query/tokeniser.rl:107
		commit(ttEventDeclType)
		goto st113
	st113:
		if p++; p == pe {
			goto _test_eof113
		}
	st_case_113:
//line query/tokeniser.go:4000
		switch data[p] {
		case 32:
			goto st97
		case 40:
			goto st92
		case 95:
			goto tr158
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto st97
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr158
			}
		default:
			goto tr158
		}
		goto st0
	tr147:
//line query/tokeniser.rl:116
		propose(ttEventDecl)
//line query/tokeniser.rl:92
		mark = p
//line query/tokeniser.rl:105
		propose(ttEventDeclType)
//line query/tokeniser.rl:124
		propose(ttAnyDecl)
		goto st114
	st114:
		if p++; p == pe {
			goto _test_eof114
		}
	st_case_114:
//line query/tokeniser.go:4037
		switch data[p] {
		case 32:
			goto tr184
		case 78:
			goto st120
		case 95:
			goto st119
		case 110:
			goto st120
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr184
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st119
				}
			case data[p] >= 65:
				goto st119
			}
		default:
			goto st119
		}
		goto st0
	tr184:
//line query/tokeniser.rl:106
		setText(ttEventDeclType)
//line query/tokeniser.rl:107
		commit(ttEventDeclType)
		goto st115
	st115:
		if p++; p == pe {
			goto _test_eof115
		}
	st_case_115:
//line query/tokeniser.go:4077
		switch data[p] {
		case 32:
			goto st115
		case 95:
			goto tr188
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto st115
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr188
			}
		default:
			goto tr188
		}
		goto st0
	tr188:
//line query/tokeniser.rl:92
		mark = p
//line query/tokeniser.rl:111
		propose(ttEventDeclAlias)
		goto st116
	st116:
		if p++; p == pe {
			goto _test_eof116
		}
	st_case_116:
//line query/tokeniser.go:4108
		switch data[p] {
		case 32:
			goto tr189
		case 41:
			goto tr190
		case 44:
			goto tr191
		case 95:
			goto st116
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr189
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st116
				}
			case data[p] >= 65:
				goto st116
			}
		default:
			goto st116
		}
		goto st0
	tr189:
//line query/tokeniser.rl:112
		setText(ttEventDeclAlias)
//line query/tokeniser.rl:113
		commit(ttEventDeclAlias)
//line query/tokeniser.rl:117
		commit(ttEventDecl)
		goto st117
	tr200:
//line query/tokeniser.rl:128
		commit(ttAnyDecl)
		goto st117
	st117:
		if p++; p == pe {
			goto _test_eof117
		}
	st_case_117:
//line query/tokeniser.go:4154
		switch data[p] {
		case 32:
			goto st117
		case 41:
			goto st93
		case 44:
			goto st118
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st117
		}
		goto st0
	tr191:
//line query/tokeniser.rl:112
		setText(ttEventDeclAlias)
//line query/tokeniser.rl:113
		commit(ttEventDeclAlias)
//line query/tokeniser.rl:117
		commit(ttEventDecl)
		goto st118
	tr202:
//line query/tokeniser.rl:128
		commit(ttAnyDecl)
		goto st118
	st118:
		if p++; p == pe {
			goto _test_eof118
		}
	st_case_118:
//line query/tokeniser.go:4184
		switch data[p] {
		case 32:
			goto st118
		case 65:
			goto tr147
		case 95:
			goto tr148
		case 97:
			goto tr147
		}
		switch {
		case data[p] < 66:
			if 9 <= data[p] && data[p] <= 13 {
				goto st118
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr148
			}
		default:
			goto tr148
		}
		goto st0
	tr148:
//line query/tokeniser.rl:116
		propose(ttEventDecl)
//line query/tokeniser.rl:92
		mark = p
//line query/tokeniser.rl:105
		propose(ttEventDeclType)
		goto st119
	st119:
		if p++; p == pe {
			goto _test_eof119
		}
	st_case_119:
//line query/tokeniser.go:4221
		switch data[p] {
		case 32:
			goto tr184
		case 95:
			goto st119
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr184
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st119
				}
			case data[p] >= 65:
				goto st119
			}
		default:
			goto st119
		}
		goto st0
	st120:
		if p++; p == pe {
			goto _test_eof120
		}
	st_case_120:
		switch data[p] {
		case 32:
			goto tr184
		case 89:
			goto st121
		case 95:
			goto st119
		case 121:
			goto st121
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr184
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st119
				}
			case data[p] >= 65:
				goto st119
			}
		default:
			goto st119
		}
		goto st0
	st121:
		if p++; p == pe {
			goto _test_eof121
		}
	st_case_121:
		switch data[p] {
		case 32:
			goto tr196
		case 40:
			goto st123
		case 95:
			goto st119
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr196
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st119
				}
			case data[p] >= 65:
				goto st119
			}
		default:
			goto st119
		}
		goto st0
	tr196:
//line query/tokeniser.rl:106
		setText(ttEventDeclType)
//line query/tokeniser.rl:107
		commit(ttEventDeclType)
		goto st122
	st122:
		if p++; p == pe {
			goto _test_eof122
		}
	st_case_122:
//line query/tokeniser.go:4321
		switch data[p] {
		case 32:
			goto st115
		case 40:
			goto st123
		case 95:
			goto tr188
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto st115
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr188
			}
		default:
			goto tr188
		}
		goto st0
	st123:
		if p++; p == pe {
			goto _test_eof123
		}
	st_case_123:
		switch data[p] {
		case 32:
			goto st123
		case 41:
			goto st124
		case 95:
			goto tr199
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto st123
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr199
			}
		default:
			goto tr199
		}
		goto st0
	tr208:
//line query/tokeniser.rl:112
		setText(ttEventDeclAlias)
//line query/tokeniser.rl:113
		commit(ttEventDeclAlias)
//line query/tokeniser.rl:117
		commit(ttEventDecl)
		goto st124
	st124:
		if p++; p == pe {
			goto _test_eof124
		}
	st_case_124:
//line query/tokeniser.go:4382
		switch data[p] {
		case 32:
			goto tr200
		case 41:
			goto tr201
		case 44:
			goto tr202
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr200
		}
		goto st0
	tr199:
//line query/tokeniser.rl:116
		propose(ttEventDecl)
//line query/tokeniser.rl:92
		mark = p
//line query/tokeniser.rl:105
		propose(ttEventDeclType)
		goto st125
	st125:
		if p++; p == pe {
			goto _test_eof125
		}
	st_case_125:
//line query/tokeniser.go:4408
		switch data[p] {
		case 32:
			goto tr203
		case 95:
			goto st125
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr203
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st125
				}
			case data[p] >= 65:
				goto st125
			}
		default:
			goto st125
		}
		goto st0
	tr203:
//line query/tokeniser.rl:106
		setText(ttEventDeclType)
//line query/tokeniser.rl:107
		commit(ttEventDeclType)
		goto st126
	st126:
		if p++; p == pe {
			goto _test_eof126
		}
	st_case_126:
//line query/tokeniser.go:4444
		switch data[p] {
		case 32:
			goto st126
		case 95:
			goto tr206
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto st126
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr206
			}
		default:
			goto tr206
		}
		goto st0
	tr206:
//line query/tokeniser.rl:92
		mark = p
//line query/tokeniser.rl:111
		propose(ttEventDeclAlias)
		goto st127
	st127:
		if p++; p == pe {
			goto _test_eof127
		}
	st_case_127:
//line query/tokeniser.go:4475
		switch data[p] {
		case 32:
			goto tr207
		case 41:
			goto tr208
		case 44:
			goto tr209
		case 95:
			goto st127
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr207
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st127
				}
			case data[p] >= 65:
				goto st127
			}
		default:
			goto st127
		}
		goto st0
	tr207:
//line query/tokeniser.rl:112
		setText(ttEventDeclAlias)
//line query/tokeniser.rl:113
		commit(ttEventDeclAlias)
//line query/tokeniser.rl:117
		commit(ttEventDecl)
		goto st128
	st128:
		if p++; p == pe {
			goto _test_eof128
		}
	st_case_128:
//line query/tokeniser.go:4517
		switch data[p] {
		case 32:
			goto st128
		case 41:
			goto st124
		case 44:
			goto st129
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st128
		}
		goto st0
	tr209:
//line query/tokeniser.rl:112
		setText(ttEventDeclAlias)
//line query/tokeniser.rl:113
		commit(ttEventDeclAlias)
//line query/tokeniser.rl:117
		commit(ttEventDecl)
		goto st129
	st129:
		if p++; p == pe {
			goto _test_eof129
		}
	st_case_129:
//line query/tokeniser.go:4543
		switch data[p] {
		case 32:
			goto st129
		case 95:
			goto tr199
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto st129
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr199
			}
		default:
			goto tr199
		}
		goto st0
	st_out:
	_test_eof1:
		cs = 1
		goto _test_eof
	_test_eof2:
		cs = 2
		goto _test_eof
	_test_eof3:
		cs = 3
		goto _test_eof
	_test_eof4:
		cs = 4
		goto _test_eof
	_test_eof5:
		cs = 5
		goto _test_eof
	_test_eof6:
		cs = 6
		goto _test_eof
	_test_eof7:
		cs = 7
		goto _test_eof
	_test_eof8:
		cs = 8
		goto _test_eof
	_test_eof9:
		cs = 9
		goto _test_eof
	_test_eof10:
		cs = 10
		goto _test_eof
	_test_eof130:
		cs = 130
		goto _test_eof
	_test_eof131:
		cs = 131
		goto _test_eof
	_test_eof132:
		cs = 132
		goto _test_eof
	_test_eof11:
		cs = 11
		goto _test_eof
	_test_eof12:
		cs = 12
		goto _test_eof
	_test_eof13:
		cs = 13
		goto _test_eof
	_test_eof14:
		cs = 14
		goto _test_eof
	_test_eof15:
		cs = 15
		goto _test_eof
	_test_eof16:
		cs = 16
		goto _test_eof
	_test_eof17:
		cs = 17
		goto _test_eof
	_test_eof18:
		cs = 18
		goto _test_eof
	_test_eof19:
		cs = 19
		goto _test_eof
	_test_eof20:
		cs = 20
		goto _test_eof
	_test_eof21:
		cs = 21
		goto _test_eof
	_test_eof133:
		cs = 133
		goto _test_eof
	_test_eof134:
		cs = 134
		goto _test_eof
	_test_eof22:
		cs = 22
		goto _test_eof
	_test_eof23:
		cs = 23
		goto _test_eof
	_test_eof24:
		cs = 24
		goto _test_eof
	_test_eof25:
		cs = 25
		goto _test_eof
	_test_eof26:
		cs = 26
		goto _test_eof
	_test_eof27:
		cs = 27
		goto _test_eof
	_test_eof28:
		cs = 28
		goto _test_eof
	_test_eof29:
		cs = 29
		goto _test_eof
	_test_eof30:
		cs = 30
		goto _test_eof
	_test_eof31:
		cs = 31
		goto _test_eof
	_test_eof32:
		cs = 32
		goto _test_eof
	_test_eof33:
		cs = 33
		goto _test_eof
	_test_eof34:
		cs = 34
		goto _test_eof
	_test_eof35:
		cs = 35
		goto _test_eof
	_test_eof36:
		cs = 36
		goto _test_eof
	_test_eof37:
		cs = 37
		goto _test_eof
	_test_eof38:
		cs = 38
		goto _test_eof
	_test_eof135:
		cs = 135
		goto _test_eof
	_test_eof136:
		cs = 136
		goto _test_eof
	_test_eof137:
		cs = 137
		goto _test_eof
	_test_eof39:
		cs = 39
		goto _test_eof
	_test_eof40:
		cs = 40
		goto _test_eof
	_test_eof41:
		cs = 41
		goto _test_eof
	_test_eof42:
		cs = 42
		goto _test_eof
	_test_eof43:
		cs = 43
		goto _test_eof
	_test_eof138:
		cs = 138
		goto _test_eof
	_test_eof44:
		cs = 44
		goto _test_eof
	_test_eof45:
		cs = 45
		goto _test_eof
	_test_eof46:
		cs = 46
		goto _test_eof
	_test_eof47:
		cs = 47
		goto _test_eof
	_test_eof48:
		cs = 48
		goto _test_eof
	_test_eof49:
		cs = 49
		goto _test_eof
	_test_eof50:
		cs = 50
		goto _test_eof
	_test_eof51:
		cs = 51
		goto _test_eof
	_test_eof52:
		cs = 52
		goto _test_eof
	_test_eof53:
		cs = 53
		goto _test_eof
	_test_eof54:
		cs = 54
		goto _test_eof
	_test_eof55:
		cs = 55
		goto _test_eof
	_test_eof56:
		cs = 56
		goto _test_eof
	_test_eof57:
		cs = 57
		goto _test_eof
	_test_eof58:
		cs = 58
		goto _test_eof
	_test_eof59:
		cs = 59
		goto _test_eof
	_test_eof60:
		cs = 60
		goto _test_eof
	_test_eof61:
		cs = 61
		goto _test_eof
	_test_eof62:
		cs = 62
		goto _test_eof
	_test_eof63:
		cs = 63
		goto _test_eof
	_test_eof64:
		cs = 64
		goto _test_eof
	_test_eof65:
		cs = 65
		goto _test_eof
	_test_eof66:
		cs = 66
		goto _test_eof
	_test_eof67:
		cs = 67
		goto _test_eof
	_test_eof68:
		cs = 68
		goto _test_eof
	_test_eof69:
		cs = 69
		goto _test_eof
	_test_eof139:
		cs = 139
		goto _test_eof
	_test_eof70:
		cs = 70
		goto _test_eof
	_test_eof71:
		cs = 71
		goto _test_eof
	_test_eof72:
		cs = 72
		goto _test_eof
	_test_eof73:
		cs = 73
		goto _test_eof
	_test_eof74:
		cs = 74
		goto _test_eof
	_test_eof140:
		cs = 140
		goto _test_eof
	_test_eof75:
		cs = 75
		goto _test_eof
	_test_eof76:
		cs = 76
		goto _test_eof
	_test_eof77:
		cs = 77
		goto _test_eof
	_test_eof78:
		cs = 78
		goto _test_eof
	_test_eof79:
		cs = 79
		goto _test_eof
	_test_eof80:
		cs = 80
		goto _test_eof
	_test_eof81:
		cs = 81
		goto _test_eof
	_test_eof82:
		cs = 82
		goto _test_eof
	_test_eof83:
		cs = 83
		goto _test_eof
	_test_eof84:
		cs = 84
		goto _test_eof
	_test_eof85:
		cs = 85
		goto _test_eof
	_test_eof86:
		cs = 86
		goto _test_eof
	_test_eof87:
		cs = 87
		goto _test_eof
	_test_eof88:
		cs = 88
		goto _test_eof
	_test_eof89:
		cs = 89
		goto _test_eof
	_test_eof141:
		cs = 141
		goto _test_eof
	_test_eof90:
		cs = 90
		goto _test_eof
	_test_eof91:
		cs = 91
		goto _test_eof
	_test_eof92:
		cs = 92
		goto _test_eof
	_test_eof93:
		cs = 93
		goto _test_eof
	_test_eof94:
		cs = 94
		goto _test_eof
	_test_eof95:
		cs = 95
		goto _test_eof
	_test_eof96:
		cs = 96
		goto _test_eof
	_test_eof97:
		cs = 97
		goto _test_eof
	_test_eof98:
		cs = 98
		goto _test_eof
	_test_eof99:
		cs = 99
		goto _test_eof
	_test_eof100:
		cs = 100
		goto _test_eof
	_test_eof101:
		cs = 101
		goto _test_eof
	_test_eof102:
		cs = 102
		goto _test_eof
	_test_eof103:
		cs = 103
		goto _test_eof
	_test_eof104:
		cs = 104
		goto _test_eof
	_test_eof105:
		cs = 105
		goto _test_eof
	_test_eof106:
		cs = 106
		goto _test_eof
	_test_eof107:
		cs = 107
		goto _test_eof
	_test_eof108:
		cs = 108
		goto _test_eof
	_test_eof109:
		cs = 109
		goto _test_eof
	_test_eof110:
		cs = 110
		goto _test_eof
	_test_eof111:
		cs = 111
		goto _test_eof
	_test_eof112:
		cs = 112
		goto _test_eof
	_test_eof113:
		cs = 113
		goto _test_eof
	_test_eof114:
		cs = 114
		goto _test_eof
	_test_eof115:
		cs = 115
		goto _test_eof
	_test_eof116:
		cs = 116
		goto _test_eof
	_test_eof117:
		cs = 117
		goto _test_eof
	_test_eof118:
		cs = 118
		goto _test_eof
	_test_eof119:
		cs = 119
		goto _test_eof
	_test_eof120:
		cs = 120
		goto _test_eof
	_test_eof121:
		cs = 121
		goto _test_eof
	_test_eof122:
		cs = 122
		goto _test_eof
	_test_eof123:
		cs = 123
		goto _test_eof
	_test_eof124:
		cs = 124
		goto _test_eof
	_test_eof125:
		cs = 125
		goto _test_eof
	_test_eof126:
		cs = 126
		goto _test_eof
	_test_eof127:
		cs = 127
		goto _test_eof
	_test_eof128:
		cs = 128
		goto _test_eof
	_test_eof129:
		cs = 129
		goto _test_eof

	_test_eof:
		{
		}
		if p == eof {
			switch cs {
			case 140:
//line query/tokeniser.rl:128
				commit(ttAnyDecl)
//line query/tokeniser.rl:157
				commit(ttEventClause)
			case 130:
//line query/tokeniser.rl:140
				commit(ttNegatedDecl)
//line query/tokeniser.rl:157
				commit(ttEventClause)
			case 141:
//line query/tokeniser.rl:151
				commit(ttSeqDecl)
//line query/tokeniser.rl:157
				commit(ttEventClause)
			case 138:
//line query/tokeniser.rl:183
				setText(ttAttributeSelector)
//line query/tokeniser.rl:184
				commit(ttAttributeSelector)
//line query/tokeniser.rl:189
				commit(ttPredicate)
			case 135, 137:
//line query/tokeniser.rl:199
				setText(ttDuration)
//line query/tokeniser.rl:200
				commit(ttDuration)
//line query/tokeniser.rl:204
				commit(ttWithinClause)
			case 139:
//line query/tokeniser.rl:112
				setText(ttEventDeclAlias)
//line query/tokeniser.rl:113
				commit(ttEventDeclAlias)
//line query/tokeniser.rl:117
				commit(ttEventDecl)
//line query/tokeniser.rl:157
				commit(ttEventClause)
			case 133:
//line query/tokeniser.rl:182
				propose(ttAttributeSelector)
//line query/tokeniser.rl:183
				setText(ttAttributeSelector)
//line query/tokeniser.rl:184
				commit(ttAttributeSelector)
//line query/tokeniser.rl:189
				commit(ttPredicate)
//line query/tokeniser.go:4756
			}
		}

	_out:
		{
		}
	}

//line query/tokeniser.rl:216

	if cs < sase_first_final {
		if p == pe {
			return nil, fmt.Errorf("Unexpected EOF")
		} else {
			return nil, fmt.Errorf("Error at position %d", p)
		}
	}

	tokens = tokens[1:] // Remove root node
	return tokens, nil
}
