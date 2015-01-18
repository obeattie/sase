package parser

import (
    "fmt"
    
    log "github.com/cihub/seelog"
)

%% machine sase;
%% write data;

func Parse(data string) (*Query, error) {
    result := &Query{}
    
    var (
        cs int             // current state
        p  int = 0         // data offset
        pe int = len(data) // data end offset
        eof int = pe // eof offset
        marks []int = make([]int, 100) // stacked marked positions
        tokens = []*token{&token{
            Tt:ttRoot,
        }}
    )
    
    mark := func() {
        marks = append(marks, p)
    }
    
    unmark := func() {
        marks = marks[:len(marks)-1]
    }
    
    lastMark := func() int {
        return marks[len(marks)-1]
    }
    
    markedText := func() string {
        return data[lastMark():p]
    }
    
    startToken := func(tt tokenType) *token {
        log.Tracef("Start token %v", tt)
        t := &token{
            Tt: tt,
        }
        tokens = append(tokens, t)
        return t
    }
    
    endToken := func(tt tokenType) {
        lastT := tokens[len(tokens)-1]
        if lastT.Tt != tt {
            panic(fmt.Sprintf("%v != %v", tt, lastT.Tt))
        }
        log.Tracef("End token %v (val: %s)", tt, lastT.Value)
        tokens = tokens[:len(tokens)-1]
    }
    
    // Suppress "variable not used"
    var (
        _ = cs
        _ = p
        _ = pe
        _ = eof
        _ = marks
        _ = tokens
        _ = mark
        _ = unmark
        _ = lastMark
        _ = startToken
        _ = endToken
    )
    
    %%{
        action mark { mark() }
        action unmark { unmark() }
        
        Identifier = [A-Za-z_] [A-Za-z0-9_]+;
        Selector = Identifier ("." Identifier)*;
        EscapedChar = "\\" any >mark %{
            
        } %unmark;
        
        # Durations
        Decimal = [\-+]? digit+ ("." digit+)?;
        Duration = (Decimal ("ns"i | "us"i | "ms"i | "s"i | "m"i | "h"i)) >mark %{
            t := startToken(ttDuration)
            t.Value = markedText()
            endToken(ttDuration)
        } %unmark;
        
        # Logical connectives
        Conjunction = "AND"i | "∧";
        Disjunction = "OR"i | "∨";
        Negation = "NOT"i | "~" | "¬"; # TODO: Does this work?
        LogicalConnective = (Conjunction | Disjunction | Negation);
        
        # Comparison operators
        Eq = "=";
        Ne = "!=" | "≠";
        Gt = ">";
        Lt = "<";
        Ge = ">=";
        Le = "<=";
        ComparisonOperator = (Eq | Ne | Gt | Lt | Ge | Le);
        
        # Predicates
        PredicateClause = Selector space* ComparisonOperator space* Selector;
        Predicate = PredicateClause (space* LogicalConnective space+ PredicateClause)*;
        
        # Event definitions
        EventDeclaration = Identifier space+ Identifier; # Type followed by variable name;
        AnyEventDeclaration = "ANY(" space* EventDeclaration (space* "," space* EventDeclaration)* space* ")";
        NotEventDeclaration = "!" space* "(" space* EventDeclaration space* ")";
        SeqArg = EventDeclaration | AnyEventDeclaration | NotEventDeclaration;
        SeqEventDeclaration = "SEQ(" (SeqArg (space* "," space* SeqArg)*)? ")";
        EventCapture = SeqArg | SeqEventDeclaration;
        
        EventClause = "EVENT"i space+ EventCapture;
        WhereClause = "WHERE"i space+ Predicate;
        WithinClause = "WITHIN"i space+ Duration;
        Query = space* EventClause (space+ WhereClause)? (space+ WithinClause)? space* (';' space*)?;
        
        Main := Query;
        write init;
        write exec;
    }%%
    
    if cs < sase_first_final {
        if p == pe {
            return nil, fmt.Errorf("Unexpected EOF")
        } else {
            return nil, fmt.Errorf("Error at position %d", p)
        }
    }
    
    return result, nil
}
