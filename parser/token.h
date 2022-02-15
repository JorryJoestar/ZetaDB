enum {

//constraint
    UNIQUE            = 231,
    PRIMARYKEY        = 230,  //PRIMARY KEY
    CHECK             = 289,
    FOREIGNKEY        = 291,
    REFERENCES        = 290,
    NOT_DEFERRABLE           = 292,  //NOT DEFERRABLE
    DEFERED_DEFERRABLE       = 293,  //DEFERRABLE INITIALLY DEFERRED
    IMMEDIATE_DEFERRABLE     = 294,  //DEFERRABLE INITIALLY IMMEDIATE
    UPDATE_NULL      = 295,  //ON UPDATE SET NULL
    UPDATE_CASCADE   = 296,  //ON UPDATE SET CASCADE
    DELETE_NULL      = 297,  //ON DELETE SET NULL
    DELETE_CASCADE   = 298,  //ON DELETE SET CASCADE
    CONSTRAINT      = 288,
    DEFAULT         = 229,

//condition
    AND             = 243,
    OR              = 244,

//predicate
    LIKE            = 258,
    NOT             = 252,
    IN              = 253,
    ALL             = 254,
    ANY             = 255,
    IS              = 256,

    NOTEQUAL        = 246,  // <>
    LESS            = 247,  // <
    GREATER         = 248,  // >
    LESSEQUAL       = 249,  // <=
    GREATEREQUAL    = 250,  // >=
    EQUAL           = 251,  // =

//elementary value
    INTVALUE          = 204,  //int value
    FLOATVALUE        = 205,  //float value
    STRINGVALUE       = 206,  //string surrounded by single or double quotation marks
    BOOLVALUE         = 551,  //BOOLVALUE true/false

//attriNameOptionTableName
    DOT               = 241,  //.

//domain
    CHAR            = 224,
    VARCHAR         = 210,  //VARCHAR(n)
    BIT             = 211,  //BIT(n)
    BITVARYING      = 212,  //BITVARYING(n)
    BOOLEAN         = 213,
    INT             = 214,
    INTEGER         = 215,
    SHORTINT        = 216,
    FLOAT           = 217,
    REAL            = 218,
    DOUBLEPRECISION = 219,
    DECIMAL         = 220,  //DECIMAL(n,d)
    NUMERIC         = 221,  //NUMERIC(n,d)
    DATE            = 222,
    TIME            = 223,

//public
    NULLMARK        = 257,  //NULL
    LPAREN            = 200,  //(
    RPAREN            = 201,  //)
    SEMICOLON         = 202,  //;
    COMMA             = 203,  //,
    ID                = 207,  //begin with non-number character
    PASSWORD          = 208,
    TABLEDOTATTRIBUTE = 209, //tableName.attributeName





    
    TABLE           = 225,
    DROP            = 226,
    ALTER           = 227,
    ADD             = 228,

    CREATE          = 232,

    FROM            = 233,
    ON              = 234,
    WHERE           = 235,
    GROUPBY         = 236,  //GROUP BY
    HAVING          = 237,
    ORDERBY         = 238,  //ORDER BY
    DISTINCT        = 239,
    SELECT          = 240,


    AS              = 242,
    STAR            = 245,  //*
    AVG             = 259,
    MIN             = 260,
    MAX             = 261,
    COUNT           = 262,
    SUM             = 263,

    JOIN            = 264,
    NATURAL         = 265,
    FULL            = 266,
    LEFT            = 267,
    RIGHT           = 268,
    OUTER           = 269,
    CROSS           = 270,

    SUBTRACT        = 271,  //-
    DIVISION        = 272,  ///
    CONCATENATION   = 273,  //||
    PLUS            = 274,  //+

    DESC            = 275,
    ASC             = 276,
    EXISTS          = 277,

    DELETE          = 278,
    UPDATE          = 279,
    SET             = 280,
    INSERTINTO      = 281,  //INSERT INTO
    VALUES          = 282,

    BEGINTOKEN      = 283,  //BEGIN
    START           = 284,
    TRANSACTION     = 285,
    COMMIT          = 286,
    ROLLBACK        = 287,

    ASSERTION       = 301, 
    TRIGGER         = 302,
    BEFORE          = 303,
    OF              = 304,
    AFTER           = 305,
    INSERT          = 306,
    OLD             = 307,
    ROW             = 308,
    NEW             = 309,
    FOR             = 310,
    EACH            = 311,
    STATEMENT       = 312,
    WHEN            = 313,
    END             = 314,
    REFERENCING     = 315,
    INSTEAD         = 316,
    VIEW            = 317,
    INDEX           = 318,

    CONNECT         = 319, 
    TO              = 320,
    AUTHORIZATION   = 321,
    DISCONNECT      = 322,
    CONNECTION      = 323,

    CALL            = 550,

    FUNCTION        = 552,
    RETURNS         = 553,
    OUT             = 554,
    INOUT           = 555,
    PROCEDURE       = 556,
    DECLARE         = 557,
    CONDITION       = 558,
    SQLSTATE        = 559,
    CURSOR          = 560,
    RETURN          = 561,
    OPEN            = 562,
    CLOSE           = 563,
    COLON           = 564,  //:
    LOOP            = 565,
    LEAVE           = 566,
    FETCH           = 567,
    INTO            = 568,
    DO              = 569,
    CONTINUE        = 570,
    EXIT            = 571,
    UNDO            = 572,
    HANDLER         = 573,
    IF              = 574,
    THEN            = 575,
    ELSE            = 576,
    ELSEIF          = 577

};