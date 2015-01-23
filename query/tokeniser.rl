package query

import (
    "fmt"
    "time"

    log "github.com/cihub/seelog"
)

%% machine sase;
%% write data;

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
    
    %%{
        action mark { mark = p }
        
        Letter = [A-Za-z_];
        Digit = [0-9];
        Decimal = [\-+]? Digit+ ('.' Digit+)?;
        Space = space;
        Identifier = Letter (Letter | Digit)*;
        
        # -- EVENT clause
            # Basic event declaration (type and variable name)
            EventDeclType =
                Identifier
                >mark
                >{ propose(ttEventDeclType) }
                %{ setText(ttEventDeclType) }
                %{ commit(ttEventDeclType) };
            EventDeclAlias =
                Identifier
                >mark
                >{ propose(ttEventDeclAlias) }
                %{ setText(ttEventDeclAlias) }
                %{ commit(ttEventDeclAlias) };
            EventDecl =
                (EventDeclType Space+ EventDeclAlias)
                >{ propose(ttEventDecl) }
                %{ commit(ttEventDecl) };
            
            # ANY declaration: matches any of the contained event types
            # Eg: ANY(EventType e1, AnotherEventType e2, YetAnotherEventType e3)
            AnyDeclArg = EventDecl;
            AnyDecl =
                "ANY"i
                >{ propose(ttAnyDecl) }
                Space? "(" Space*
                (AnyDeclArg (Space* "," Space* AnyDeclArg)*)?
                Space* ")"
                %{ commit(ttAnyDecl) };
            
            # NOT declaration: terminates a match if the matched event is found
            # Eg: !(ShouldNotHappen e1)
            #     !(ANY(ShouldNotHappen e1, ShouldNotHappenEither e2))
            NegatedDeclArg = EventDecl | AnyDecl;
            NegatedDecl =
                ("!" | "NOT"i)
                >{ propose(ttNegatedDecl) }
                Space? "(" Space*
                (NegatedDeclArg (Space* "," Space* NegatedDeclArg)*)?
                Space* ")"
                %{ commit(ttNegatedDecl) };

            # SEQ declaration: matches a sequence of events
            # Eg: SEQ(EventType e1, !(ShouldNotHappen e2), AnotherEventType e3)
            SeqDeclArg = EventDecl | AnyDecl | NegatedDecl;
            SeqDecl =
                "SEQ"i
                >{ propose(ttSeqDecl) }
                Space? "("
                (SeqDeclArg (Space* "," Space* SeqDeclArg)*)?
                Space* ")"
                %{ commit(ttSeqDecl) };
            
            EventClauseArg = NegatedDecl | AnyDecl | SeqDecl | EventDecl;
            EventClause =
                "EVENT"i Space+ EventClauseArg
                >{ propose(ttEventClause) }
                %{ commit(ttEventClause) };
        
        # -- WHERE clause
            Eq = "==" >{ propose(ttEq) } %{ commit(ttEq) };
            Ne = "!=" >{ propose(ttEq) } %{ commit(ttEq) };
            Gt = ">" >{ propose(ttEq) } %{ commit(ttEq) };
            Lt = "<" >{ propose(ttEq) } %{ commit(ttEq) };
            Ge = ">=" >{ propose(ttEq) } %{ commit(ttEq) };
            Le = "<=" >{ propose(ttEq) } %{ commit(ttEq) };
            Comparison = (Eq | Ne | Gt | Lt | Ge | Le);
            
            Conjunction =
                ("AND"i | "^" | "&&")
                >{ propose(ttConjunction) }
                %{ commit(ttConjunction) };
            Disjunction =
                ("OR"i | "∨" | "||")
                >{ propose(ttDisjunction) }
                %{ commit(ttDisjunction) };
            Connective = Conjunction | Disjunction;
            
            AttributeSelector =
                Identifier
                >mark
                ("." Identifier)*
                >{ propose(ttAttributeSelector) }
                %{ setText(ttAttributeSelector) }
                %{ commit(ttAttributeSelector) };
            PredicateArg =
                AttributeSelector
                >{ propose(ttPredicate) }
                Space* Comparison Space* AttributeSelector
                %{ commit(ttPredicate) };
            Predicate = PredicateArg (Space+ Connective Space+ PredicateArg)*;
            
            WhereClause = "WHERE"i Space+ Predicate;
        
        # -- WITHIN clause
            Duration =
                (Decimal ("ns"i | "us"i | "ms"i | "s"i | "m"i | "h"i))
                >mark
                >{ propose(ttDuration) }
                %{ setText(ttDuration) }
                %{ commit(ttDuration) };
            WithinClause =
                "WITHIN"i Space+ Duration
                >{ propose(ttWithinClause) }
                %{ commit(ttWithinClause) };
        
        Query =
            Space*
            EventClause
            (Space+ WhereClause)?
            (Space+ WithinClause)?
            (Space* ";")? Space*;
        
        main := Query;
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
    
    tokens = tokens[1:] // Remove root node
    return tokens, nil
}
