// Generated from /Users/faraujo/workspace/research/sysflow/sf-processor/plugins/sfpe/lang/Sfpl.g4 by ANTLR 4.7.1
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class SfplParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.7.1", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		RULE=1, FILTER=2, MACRO=3, LIST=4, NAME=5, ITEMS=6, COND=7, DESC=8, ACTION=9, 
		OUTPUT=10, PRIORITY=11, TAGS=12, AND=13, OR=14, NOT=15, LT=16, LE=17, 
		GT=18, GE=19, EQ=20, NEQ=21, IN=22, CONTAINS=23, ICONTAINS=24, STARTSWITH=25, 
		PMATCH=26, EXISTS=27, LBRACK=28, RBRACK=29, LPAREN=30, RPAREN=31, LISTSEP=32, 
		DECL=33, DEF=34, SEVERITY=35, ID=36, NUMBER=37, PATH=38, STRING=39, WS=40, 
		NL=41, COMMENT=42, ANY=43;
	public static final int
		RULE_policy = 0, RULE_prule = 1, RULE_pfilter = 2, RULE_pmacro = 3, RULE_plist = 4, 
		RULE_expression = 5, RULE_or_expression = 6, RULE_and_expression = 7, 
		RULE_term = 8, RULE_items = 9, RULE_variable = 10, RULE_atom = 11, RULE_text = 12, 
		RULE_binary_operator = 13, RULE_unary_operator = 14;
	public static final String[] ruleNames = {
		"policy", "prule", "pfilter", "pmacro", "plist", "expression", "or_expression", 
		"and_expression", "term", "items", "variable", "atom", "text", "binary_operator", 
		"unary_operator"
	};

	private static final String[] _LITERAL_NAMES = {
		null, "'rule'", "'filter'", "'macro'", "'list'", "'name'", "'items'", 
		"'condition'", "'desc'", "'action'", "'output'", "'priority'", "'tags'", 
		"'and'", "'or'", "'not'", "'<'", "'<='", "'>'", "'>='", "'='", "'!='", 
		"'in'", "'contains'", "'icontains'", "'startswith'", "'pmatch'", "'exists'", 
		"'['", "']'", "'('", "')'", "','", "'-'"
	};
	private static final String[] _SYMBOLIC_NAMES = {
		null, "RULE", "FILTER", "MACRO", "LIST", "NAME", "ITEMS", "COND", "DESC", 
		"ACTION", "OUTPUT", "PRIORITY", "TAGS", "AND", "OR", "NOT", "LT", "LE", 
		"GT", "GE", "EQ", "NEQ", "IN", "CONTAINS", "ICONTAINS", "STARTSWITH", 
		"PMATCH", "EXISTS", "LBRACK", "RBRACK", "LPAREN", "RPAREN", "LISTSEP", 
		"DECL", "DEF", "SEVERITY", "ID", "NUMBER", "PATH", "STRING", "WS", "NL", 
		"COMMENT", "ANY"
	};
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}

	@Override
	public String getGrammarFileName() { return "Sfpl.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public SfplParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}
	public static class PolicyContext extends ParserRuleContext {
		public TerminalNode EOF() { return getToken(SfplParser.EOF, 0); }
		public List<PruleContext> prule() {
			return getRuleContexts(PruleContext.class);
		}
		public PruleContext prule(int i) {
			return getRuleContext(PruleContext.class,i);
		}
		public List<PfilterContext> pfilter() {
			return getRuleContexts(PfilterContext.class);
		}
		public PfilterContext pfilter(int i) {
			return getRuleContext(PfilterContext.class,i);
		}
		public List<PmacroContext> pmacro() {
			return getRuleContexts(PmacroContext.class);
		}
		public PmacroContext pmacro(int i) {
			return getRuleContext(PmacroContext.class,i);
		}
		public List<PlistContext> plist() {
			return getRuleContexts(PlistContext.class);
		}
		public PlistContext plist(int i) {
			return getRuleContext(PlistContext.class,i);
		}
		public PolicyContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_policy; }
	}

	public final PolicyContext policy() throws RecognitionException {
		PolicyContext _localctx = new PolicyContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_policy);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(34); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				setState(34);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,0,_ctx) ) {
				case 1:
					{
					setState(30);
					prule();
					}
					break;
				case 2:
					{
					setState(31);
					pfilter();
					}
					break;
				case 3:
					{
					setState(32);
					pmacro();
					}
					break;
				case 4:
					{
					setState(33);
					plist();
					}
					break;
				}
				}
				setState(36); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( _la==DECL );
			setState(38);
			match(EOF);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class PruleContext extends ParserRuleContext {
		public TerminalNode DECL() { return getToken(SfplParser.DECL, 0); }
		public TerminalNode RULE() { return getToken(SfplParser.RULE, 0); }
		public List<TerminalNode> DEF() { return getTokens(SfplParser.DEF); }
		public TerminalNode DEF(int i) {
			return getToken(SfplParser.DEF, i);
		}
		public List<TextContext> text() {
			return getRuleContexts(TextContext.class);
		}
		public TextContext text(int i) {
			return getRuleContext(TextContext.class,i);
		}
		public TerminalNode DESC() { return getToken(SfplParser.DESC, 0); }
		public TerminalNode COND() { return getToken(SfplParser.COND, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode PRIORITY() { return getToken(SfplParser.PRIORITY, 0); }
		public TerminalNode SEVERITY() { return getToken(SfplParser.SEVERITY, 0); }
		public TerminalNode TAGS() { return getToken(SfplParser.TAGS, 0); }
		public ItemsContext items() {
			return getRuleContext(ItemsContext.class,0);
		}
		public TerminalNode ACTION() { return getToken(SfplParser.ACTION, 0); }
		public TerminalNode OUTPUT() { return getToken(SfplParser.OUTPUT, 0); }
		public PruleContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_prule; }
	}

	public final PruleContext prule() throws RecognitionException {
		PruleContext _localctx = new PruleContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_prule);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(40);
			match(DECL);
			setState(41);
			match(RULE);
			setState(42);
			match(DEF);
			setState(43);
			text();
			setState(44);
			match(DESC);
			setState(45);
			match(DEF);
			setState(46);
			text();
			setState(47);
			match(COND);
			setState(48);
			match(DEF);
			setState(49);
			expression();
			setState(50);
			_la = _input.LA(1);
			if ( !(_la==ACTION || _la==OUTPUT) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			setState(51);
			match(DEF);
			setState(52);
			text();
			setState(53);
			match(PRIORITY);
			setState(54);
			match(DEF);
			setState(55);
			match(SEVERITY);
			setState(56);
			match(TAGS);
			setState(57);
			match(DEF);
			setState(58);
			items();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class PfilterContext extends ParserRuleContext {
		public TerminalNode DECL() { return getToken(SfplParser.DECL, 0); }
		public TerminalNode FILTER() { return getToken(SfplParser.FILTER, 0); }
		public List<TerminalNode> DEF() { return getTokens(SfplParser.DEF); }
		public TerminalNode DEF(int i) {
			return getToken(SfplParser.DEF, i);
		}
		public TerminalNode ID() { return getToken(SfplParser.ID, 0); }
		public TerminalNode COND() { return getToken(SfplParser.COND, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public PfilterContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_pfilter; }
	}

	public final PfilterContext pfilter() throws RecognitionException {
		PfilterContext _localctx = new PfilterContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_pfilter);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(60);
			match(DECL);
			setState(61);
			match(FILTER);
			setState(62);
			match(DEF);
			setState(63);
			match(ID);
			setState(64);
			match(COND);
			setState(65);
			match(DEF);
			setState(66);
			expression();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class PmacroContext extends ParserRuleContext {
		public TerminalNode DECL() { return getToken(SfplParser.DECL, 0); }
		public TerminalNode MACRO() { return getToken(SfplParser.MACRO, 0); }
		public List<TerminalNode> DEF() { return getTokens(SfplParser.DEF); }
		public TerminalNode DEF(int i) {
			return getToken(SfplParser.DEF, i);
		}
		public TerminalNode ID() { return getToken(SfplParser.ID, 0); }
		public TerminalNode COND() { return getToken(SfplParser.COND, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public PmacroContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_pmacro; }
	}

	public final PmacroContext pmacro() throws RecognitionException {
		PmacroContext _localctx = new PmacroContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_pmacro);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(68);
			match(DECL);
			setState(69);
			match(MACRO);
			setState(70);
			match(DEF);
			setState(71);
			match(ID);
			setState(72);
			match(COND);
			setState(73);
			match(DEF);
			setState(74);
			expression();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class PlistContext extends ParserRuleContext {
		public TerminalNode DECL() { return getToken(SfplParser.DECL, 0); }
		public TerminalNode LIST() { return getToken(SfplParser.LIST, 0); }
		public List<TerminalNode> DEF() { return getTokens(SfplParser.DEF); }
		public TerminalNode DEF(int i) {
			return getToken(SfplParser.DEF, i);
		}
		public TerminalNode ID() { return getToken(SfplParser.ID, 0); }
		public TerminalNode ITEMS() { return getToken(SfplParser.ITEMS, 0); }
		public ItemsContext items() {
			return getRuleContext(ItemsContext.class,0);
		}
		public PlistContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_plist; }
	}

	public final PlistContext plist() throws RecognitionException {
		PlistContext _localctx = new PlistContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_plist);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(76);
			match(DECL);
			setState(77);
			match(LIST);
			setState(78);
			match(DEF);
			setState(79);
			match(ID);
			setState(80);
			match(ITEMS);
			setState(81);
			match(DEF);
			setState(82);
			items();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ExpressionContext extends ParserRuleContext {
		public Or_expressionContext or_expression() {
			return getRuleContext(Or_expressionContext.class,0);
		}
		public ExpressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expression; }
	}

	public final ExpressionContext expression() throws RecognitionException {
		ExpressionContext _localctx = new ExpressionContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_expression);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(84);
			or_expression();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Or_expressionContext extends ParserRuleContext {
		public List<And_expressionContext> and_expression() {
			return getRuleContexts(And_expressionContext.class);
		}
		public And_expressionContext and_expression(int i) {
			return getRuleContext(And_expressionContext.class,i);
		}
		public List<TerminalNode> OR() { return getTokens(SfplParser.OR); }
		public TerminalNode OR(int i) {
			return getToken(SfplParser.OR, i);
		}
		public Or_expressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_or_expression; }
	}

	public final Or_expressionContext or_expression() throws RecognitionException {
		Or_expressionContext _localctx = new Or_expressionContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_or_expression);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(86);
			and_expression();
			setState(91);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==OR) {
				{
				{
				setState(87);
				match(OR);
				setState(88);
				and_expression();
				}
				}
				setState(93);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class And_expressionContext extends ParserRuleContext {
		public List<TermContext> term() {
			return getRuleContexts(TermContext.class);
		}
		public TermContext term(int i) {
			return getRuleContext(TermContext.class,i);
		}
		public List<TerminalNode> AND() { return getTokens(SfplParser.AND); }
		public TerminalNode AND(int i) {
			return getToken(SfplParser.AND, i);
		}
		public And_expressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_and_expression; }
	}

	public final And_expressionContext and_expression() throws RecognitionException {
		And_expressionContext _localctx = new And_expressionContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_and_expression);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(94);
			term();
			setState(99);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==AND) {
				{
				{
				setState(95);
				match(AND);
				setState(96);
				term();
				}
				}
				setState(101);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class TermContext extends ParserRuleContext {
		public VariableContext variable() {
			return getRuleContext(VariableContext.class,0);
		}
		public TerminalNode NOT() { return getToken(SfplParser.NOT, 0); }
		public TermContext term() {
			return getRuleContext(TermContext.class,0);
		}
		public List<AtomContext> atom() {
			return getRuleContexts(AtomContext.class);
		}
		public AtomContext atom(int i) {
			return getRuleContext(AtomContext.class,i);
		}
		public Unary_operatorContext unary_operator() {
			return getRuleContext(Unary_operatorContext.class,0);
		}
		public Binary_operatorContext binary_operator() {
			return getRuleContext(Binary_operatorContext.class,0);
		}
		public TerminalNode LPAREN() { return getToken(SfplParser.LPAREN, 0); }
		public TerminalNode RPAREN() { return getToken(SfplParser.RPAREN, 0); }
		public TerminalNode IN() { return getToken(SfplParser.IN, 0); }
		public TerminalNode PMATCH() { return getToken(SfplParser.PMATCH, 0); }
		public List<ItemsContext> items() {
			return getRuleContexts(ItemsContext.class);
		}
		public ItemsContext items(int i) {
			return getRuleContext(ItemsContext.class,i);
		}
		public List<TerminalNode> LISTSEP() { return getTokens(SfplParser.LISTSEP); }
		public TerminalNode LISTSEP(int i) {
			return getToken(SfplParser.LISTSEP, i);
		}
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TermContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_term; }
	}

	public final TermContext term() throws RecognitionException {
		TermContext _localctx = new TermContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_term);
		int _la;
		try {
			setState(135);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,7,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(102);
				variable();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(103);
				match(NOT);
				setState(104);
				term();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(105);
				atom();
				setState(106);
				unary_operator();
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(108);
				atom();
				setState(109);
				binary_operator();
				setState(110);
				atom();
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(112);
				atom();
				setState(113);
				_la = _input.LA(1);
				if ( !(_la==IN || _la==PMATCH) ) {
				_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(114);
				match(LPAREN);
				setState(117);
				_errHandler.sync(this);
				switch (_input.LA(1)) {
				case LT:
				case GT:
				case ID:
				case NUMBER:
				case PATH:
				case STRING:
					{
					setState(115);
					atom();
					}
					break;
				case LBRACK:
					{
					setState(116);
					items();
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				setState(126);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==LISTSEP) {
					{
					{
					setState(119);
					match(LISTSEP);
					setState(122);
					_errHandler.sync(this);
					switch (_input.LA(1)) {
					case LT:
					case GT:
					case ID:
					case NUMBER:
					case PATH:
					case STRING:
						{
						setState(120);
						atom();
						}
						break;
					case LBRACK:
						{
						setState(121);
						items();
						}
						break;
					default:
						throw new NoViableAltException(this);
					}
					}
					}
					setState(128);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(129);
				match(RPAREN);
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(131);
				match(LPAREN);
				setState(132);
				expression();
				setState(133);
				match(RPAREN);
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ItemsContext extends ParserRuleContext {
		public TerminalNode LBRACK() { return getToken(SfplParser.LBRACK, 0); }
		public TerminalNode RBRACK() { return getToken(SfplParser.RBRACK, 0); }
		public List<AtomContext> atom() {
			return getRuleContexts(AtomContext.class);
		}
		public AtomContext atom(int i) {
			return getRuleContext(AtomContext.class,i);
		}
		public List<TerminalNode> LISTSEP() { return getTokens(SfplParser.LISTSEP); }
		public TerminalNode LISTSEP(int i) {
			return getToken(SfplParser.LISTSEP, i);
		}
		public ItemsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_items; }
	}

	public final ItemsContext items() throws RecognitionException {
		ItemsContext _localctx = new ItemsContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_items);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(137);
			match(LBRACK);
			setState(146);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << LT) | (1L << GT) | (1L << ID) | (1L << NUMBER) | (1L << PATH) | (1L << STRING))) != 0)) {
				{
				setState(138);
				atom();
				setState(143);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==LISTSEP) {
					{
					{
					setState(139);
					match(LISTSEP);
					setState(140);
					atom();
					}
					}
					setState(145);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				}
			}

			setState(148);
			match(RBRACK);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class VariableContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(SfplParser.ID, 0); }
		public VariableContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_variable; }
	}

	public final VariableContext variable() throws RecognitionException {
		VariableContext _localctx = new VariableContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_variable);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(150);
			match(ID);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class AtomContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(SfplParser.ID, 0); }
		public TerminalNode PATH() { return getToken(SfplParser.PATH, 0); }
		public TerminalNode NUMBER() { return getToken(SfplParser.NUMBER, 0); }
		public TerminalNode STRING() { return getToken(SfplParser.STRING, 0); }
		public AtomContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_atom; }
	}

	public final AtomContext atom() throws RecognitionException {
		AtomContext _localctx = new AtomContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_atom);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(152);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << LT) | (1L << GT) | (1L << ID) | (1L << NUMBER) | (1L << PATH) | (1L << STRING))) != 0)) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class TextContext extends ParserRuleContext {
		public TextContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_text; }
	}

	public final TextContext text() throws RecognitionException {
		TextContext _localctx = new TextContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_text);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(156); 
			_errHandler.sync(this);
			_alt = 1;
			do {
				switch (_alt) {
				case 1:
					{
					{
					setState(154);
					if (!(!(p.GetCurrentToken().GetText() == "desc" ||
						      p.GetCurrentToken().GetText() == "condition" ||
						      p.GetCurrentToken().GetText() == "action" ||
						      p.GetCurrentToken().GetText() == "output" ||
						      p.GetCurrentToken().GetText() == "priority" ||
						      p.GetCurrentToken().GetText() == "tags"))) throw new FailedPredicateException(this, "!(p.GetCurrentToken().GetText() == \"desc\" ||\n\t      p.GetCurrentToken().GetText() == \"condition\" ||\n\t      p.GetCurrentToken().GetText() == \"action\" ||\n\t      p.GetCurrentToken().GetText() == \"output\" ||\n\t      p.GetCurrentToken().GetText() == \"priority\" ||\n\t      p.GetCurrentToken().GetText() == \"tags\")");
					setState(155);
					matchWildcard();
					}
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				setState(158); 
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,10,_ctx);
			} while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER );
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Binary_operatorContext extends ParserRuleContext {
		public TerminalNode LT() { return getToken(SfplParser.LT, 0); }
		public TerminalNode LE() { return getToken(SfplParser.LE, 0); }
		public TerminalNode GT() { return getToken(SfplParser.GT, 0); }
		public TerminalNode GE() { return getToken(SfplParser.GE, 0); }
		public TerminalNode EQ() { return getToken(SfplParser.EQ, 0); }
		public TerminalNode NEQ() { return getToken(SfplParser.NEQ, 0); }
		public TerminalNode CONTAINS() { return getToken(SfplParser.CONTAINS, 0); }
		public TerminalNode ICONTAINS() { return getToken(SfplParser.ICONTAINS, 0); }
		public TerminalNode STARTSWITH() { return getToken(SfplParser.STARTSWITH, 0); }
		public Binary_operatorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_binary_operator; }
	}

	public final Binary_operatorContext binary_operator() throws RecognitionException {
		Binary_operatorContext _localctx = new Binary_operatorContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_binary_operator);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(160);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << LT) | (1L << LE) | (1L << GT) | (1L << GE) | (1L << EQ) | (1L << NEQ) | (1L << CONTAINS) | (1L << ICONTAINS) | (1L << STARTSWITH))) != 0)) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Unary_operatorContext extends ParserRuleContext {
		public TerminalNode EXISTS() { return getToken(SfplParser.EXISTS, 0); }
		public Unary_operatorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_unary_operator; }
	}

	public final Unary_operatorContext unary_operator() throws RecognitionException {
		Unary_operatorContext _localctx = new Unary_operatorContext(_ctx, getState());
		enterRule(_localctx, 28, RULE_unary_operator);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(162);
			match(EXISTS);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public boolean sempred(RuleContext _localctx, int ruleIndex, int predIndex) {
		switch (ruleIndex) {
		case 12:
			return text_sempred((TextContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean text_sempred(TextContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return !(p.GetCurrentToken().GetText() == "desc" ||
			      p.GetCurrentToken().GetText() == "condition" ||
			      p.GetCurrentToken().GetText() == "action" ||
			      p.GetCurrentToken().GetText() == "output" ||
			      p.GetCurrentToken().GetText() == "priority" ||
			      p.GetCurrentToken().GetText() == "tags");
		}
		return true;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3-\u00a7\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\3\2\3\2\3\2\3\2\6\2"+
		"%\n\2\r\2\16\2&\3\2\3\2\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3"+
		"\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\5"+
		"\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\7\3\7\3"+
		"\b\3\b\3\b\7\b\\\n\b\f\b\16\b_\13\b\3\t\3\t\3\t\7\td\n\t\f\t\16\tg\13"+
		"\t\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\5\nx\n"+
		"\n\3\n\3\n\3\n\5\n}\n\n\7\n\177\n\n\f\n\16\n\u0082\13\n\3\n\3\n\3\n\3"+
		"\n\3\n\3\n\5\n\u008a\n\n\3\13\3\13\3\13\3\13\7\13\u0090\n\13\f\13\16\13"+
		"\u0093\13\13\5\13\u0095\n\13\3\13\3\13\3\f\3\f\3\r\3\r\3\16\3\16\6\16"+
		"\u009f\n\16\r\16\16\16\u00a0\3\17\3\17\3\20\3\20\3\20\2\2\21\2\4\6\b\n"+
		"\f\16\20\22\24\26\30\32\34\36\2\6\3\2\13\f\4\2\30\30\34\34\5\2\22\22\24"+
		"\24&)\4\2\22\27\31\33\2\u00a8\2$\3\2\2\2\4*\3\2\2\2\6>\3\2\2\2\bF\3\2"+
		"\2\2\nN\3\2\2\2\fV\3\2\2\2\16X\3\2\2\2\20`\3\2\2\2\22\u0089\3\2\2\2\24"+
		"\u008b\3\2\2\2\26\u0098\3\2\2\2\30\u009a\3\2\2\2\32\u009e\3\2\2\2\34\u00a2"+
		"\3\2\2\2\36\u00a4\3\2\2\2 %\5\4\3\2!%\5\6\4\2\"%\5\b\5\2#%\5\n\6\2$ \3"+
		"\2\2\2$!\3\2\2\2$\"\3\2\2\2$#\3\2\2\2%&\3\2\2\2&$\3\2\2\2&\'\3\2\2\2\'"+
		"(\3\2\2\2()\7\2\2\3)\3\3\2\2\2*+\7#\2\2+,\7\3\2\2,-\7$\2\2-.\5\32\16\2"+
		"./\7\n\2\2/\60\7$\2\2\60\61\5\32\16\2\61\62\7\t\2\2\62\63\7$\2\2\63\64"+
		"\5\f\7\2\64\65\t\2\2\2\65\66\7$\2\2\66\67\5\32\16\2\678\7\r\2\289\7$\2"+
		"\29:\7%\2\2:;\7\16\2\2;<\7$\2\2<=\5\24\13\2=\5\3\2\2\2>?\7#\2\2?@\7\4"+
		"\2\2@A\7$\2\2AB\7&\2\2BC\7\t\2\2CD\7$\2\2DE\5\f\7\2E\7\3\2\2\2FG\7#\2"+
		"\2GH\7\5\2\2HI\7$\2\2IJ\7&\2\2JK\7\t\2\2KL\7$\2\2LM\5\f\7\2M\t\3\2\2\2"+
		"NO\7#\2\2OP\7\6\2\2PQ\7$\2\2QR\7&\2\2RS\7\b\2\2ST\7$\2\2TU\5\24\13\2U"+
		"\13\3\2\2\2VW\5\16\b\2W\r\3\2\2\2X]\5\20\t\2YZ\7\20\2\2Z\\\5\20\t\2[Y"+
		"\3\2\2\2\\_\3\2\2\2][\3\2\2\2]^\3\2\2\2^\17\3\2\2\2_]\3\2\2\2`e\5\22\n"+
		"\2ab\7\17\2\2bd\5\22\n\2ca\3\2\2\2dg\3\2\2\2ec\3\2\2\2ef\3\2\2\2f\21\3"+
		"\2\2\2ge\3\2\2\2h\u008a\5\26\f\2ij\7\21\2\2j\u008a\5\22\n\2kl\5\30\r\2"+
		"lm\5\36\20\2m\u008a\3\2\2\2no\5\30\r\2op\5\34\17\2pq\5\30\r\2q\u008a\3"+
		"\2\2\2rs\5\30\r\2st\t\3\2\2tw\7 \2\2ux\5\30\r\2vx\5\24\13\2wu\3\2\2\2"+
		"wv\3\2\2\2x\u0080\3\2\2\2y|\7\"\2\2z}\5\30\r\2{}\5\24\13\2|z\3\2\2\2|"+
		"{\3\2\2\2}\177\3\2\2\2~y\3\2\2\2\177\u0082\3\2\2\2\u0080~\3\2\2\2\u0080"+
		"\u0081\3\2\2\2\u0081\u0083\3\2\2\2\u0082\u0080\3\2\2\2\u0083\u0084\7!"+
		"\2\2\u0084\u008a\3\2\2\2\u0085\u0086\7 \2\2\u0086\u0087\5\f\7\2\u0087"+
		"\u0088\7!\2\2\u0088\u008a\3\2\2\2\u0089h\3\2\2\2\u0089i\3\2\2\2\u0089"+
		"k\3\2\2\2\u0089n\3\2\2\2\u0089r\3\2\2\2\u0089\u0085\3\2\2\2\u008a\23\3"+
		"\2\2\2\u008b\u0094\7\36\2\2\u008c\u0091\5\30\r\2\u008d\u008e\7\"\2\2\u008e"+
		"\u0090\5\30\r\2\u008f\u008d\3\2\2\2\u0090\u0093\3\2\2\2\u0091\u008f\3"+
		"\2\2\2\u0091\u0092\3\2\2\2\u0092\u0095\3\2\2\2\u0093\u0091\3\2\2\2\u0094"+
		"\u008c\3\2\2\2\u0094\u0095\3\2\2\2\u0095\u0096\3\2\2\2\u0096\u0097\7\37"+
		"\2\2\u0097\25\3\2\2\2\u0098\u0099\7&\2\2\u0099\27\3\2\2\2\u009a\u009b"+
		"\t\4\2\2\u009b\31\3\2\2\2\u009c\u009d\6\16\2\2\u009d\u009f\13\2\2\2\u009e"+
		"\u009c\3\2\2\2\u009f\u00a0\3\2\2\2\u00a0\u009e\3\2\2\2\u00a0\u00a1\3\2"+
		"\2\2\u00a1\33\3\2\2\2\u00a2\u00a3\t\5\2\2\u00a3\35\3\2\2\2\u00a4\u00a5"+
		"\7\35\2\2\u00a5\37\3\2\2\2\r$&]ew|\u0080\u0089\u0091\u0094\u00a0";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}