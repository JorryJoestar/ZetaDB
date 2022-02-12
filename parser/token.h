enum {
    /**
     * general element
     * 0 is reserved as the notation of statement end in Lex function in lexer.go
     */
    LPAREN          = 1,    //Left parenthesis  (
    RPAREN          = 2,    //right parenthesis )
    SQLEND          = 3,    //semicolon(;) + arbitrary number of white space(" \t") + return(\n);
    COMMA           = 4,    //comma ,
    INTVALUE        = 5,    //int values
    FLOATVALUE      = 6,    //float vaules
    STRINGVALUE     = 7,    //string surrounded by single or double quotation marks
    ID              = 8,    //table_name, attribute_name

    //data types 
    CHAR            = 100,  //CHAR
    VARCHAR         = 101,  //VARCHAR(n)
    BIT             = 102,  //BIT(n)
    BITVARYING      = 103,  //BITVARYING(n)
    BOOLEAN         = 104,  //BOOLEAN
    INT             = 105,  //INT
    INTEGER         = 106,  //INTEGER
    SHORTINT        = 107,  //SHORTINT
    FLOAT           = 108,  //FLOAT
    REAL            = 109,  //REAL
    DOUBLEPRECISION = 110,  //DOUBLE PRECISION
    DECIMAL         = 111,  //DECIMAL(n,d)
    NUMERIC         = 112,  //NUMERIC(n,d)
    DATE            = 113,  //DATE
    TIME            = 114,  //TIME

    /**
     * table modification language
     */
    CREATE          = 200,  //CREATE
    TABLE           = 201,  //TABLE
    DROP            = 202,  //DROP
    ALTER           = 203,  //ALTER
    ADD             = 204,  //ADD
    DEFAULT         = 205,  //DEFAULT
    PRIMARYKEY      = 206,  //PRIMARY KEY
    UNIQUE          = 207,  //UNIQUE

    /**
     * selection language
     */
    SELECT          = 300,  //SELECT
    FROM            = 301,  //FROM
    ON              = 302,  //ON
    WHERE           = 303,  //WHERE
    GROUPBY         = 304,  //GROUP BY
    HAVING          = 305,  //HAVING
    ORDERBY         = 306,  //ORDER BY
    DISTINCT        = 307,  //DISTINCT
    
    STAR            = 308,  // *
    DOT             = 309,  // .
    AS              = 310,  //AS
    AND             = 311,  //AND
    OR              = 312,  //OR
    
    EQUAL           = 313,  // =
    NOTEQUAL        = 314,  // <>
    LESS            = 315,  // <
    GREATER         = 316,  // >
    LESSEQUAL       = 317,  // <=
    GREATEREQUAL    = 318,  // >=
    
    LIKE            = 319,  //LIKE
    NOT             = 320,  //NOT
    IN              = 321,  //IN
    ALL             = 322,  //ALL
    ANY             = 323,  //ANY
    IS              = 324,  //IS
    NULLMARK        = 325,  //NULL
    
    SUM             = 326,  //SUM
    AVG             = 327,  //AVG
    MIN             = 328,  //MIN
    MAX             = 329,  //MAX
    COUNT           = 330,  //COUNT
    
    CROSS           = 331,  //CROSS
    JOIN            = 332,  //JOIN
    NATURAL         = 333,  //NATURAL
    FULL            = 334,  //FULL
    LEFT            = 335,  //LEFT
    RIGHT           = 336,  //RIGHT
    OUTER           = 337,  //OUTER
    
    PLUS            = 338,  // +
    SUBTRACT        = 339,  // -
    DIVISION        = 340,  // /
    CONCATENATION   = 341,  // ||

    ASC             = 342,  //ASC
    DESC            = 343,  //DESC

    TABATTRI        = 344,  //tableName.attributeName

    EXISTS          = 345,  //EXISTS

    /**
     * deletion language
     * update language
     * insertion language
     * transaction
     * constraint
     * assertion
     * trigger
     * view
     * index
     */
    DELETE          = 400,  //DELETE

    UPDATE          = 401,  //UPDATE
    SET             = 402,  //SET
    INSERTINTO      = 403,  //INSERT INTO
    VALUES          = 404,  //VALUES

    BEGINTOKEN      = 405,  //BEGIN
    START           = 406,  //START
    TRANSACTION     = 407,  //TRANSACTION
    COMMIT          = 408,  //COMMIT
    ROLLBACK        = 409,  //ROLLBACK
    READ            = 410,  //READ
    WRITE           = 411,  //WRITE
    ONLY            = 412,  //ONLY
    LEVEL           = 413,  //LEVEL
    UNCOMMITTED     = 414,  //UNCOMMITTED
    COMMITTED       = 415,  //COMMITTED
    REPEATABLE      = 416,  //REPEATABLE
    SERIALIZABLE    = 417,  //SERIALIZABLE
    ISOLATION       = 418,  //ISOLATION

    CONSTRAINT      = 420,  //CONSTRAINT
    CHECK           = 421,  //CHECK
    REFERENCES      = 422,  //REFERENCES
    FOREIGNKEY      = 423,  //FOREIGNKEY
    
    DEFERRABLENOT           = 424,  //NOT DEFERRABLE
    DEFERRABLEDEFERED       = 425,  //DEFERRABLE INITIALLY DEFERRED
    DEFERRABLEIMMEDIATE     = 426,  //DEFERRABLE INITIALLY IMMEDIATE
    UPDATENULL      = 427,  //ON UPDATE SET NULL
    UPDATECASCADE   = 428,  //ON UPDATE SET CASCADE
    DELETENULL      = 429,  //ON DELETE SET NULL
    DELETECASCADE   = 430,  //ON DELETE SET CASCADE
    DEFERRED        = 431,  //DEFERRED
    IMMEDIATE       = 432,  //IMMEDIATE

    ASSERTION       = 440,  //ASSERTION

    TRIGGER         = 441,  //TRIGGER
    BEFORE          = 442,  //BEFORE
    OF              = 443,  //OF
    AFTER           = 444,  //AFTER
    INSERT          = 445,  //INSERT
    OLD             = 446,  //OLD
    ROW             = 447,  //ROW
    NEW             = 448,  //NEW
    FOR             = 449,  //FOR
    EACH            = 450,  //EACH
    STATEMENT       = 451,  //STATEMENT
    WHEN            = 452,  //WHEN
    END             = 453,  //END
    REFERENCING     = 454,  //REFERENCING
    INSTEAD         = 455,  //INSTEAD
    VIEW            = 456,  //VIEW
    INDEX           = 457,  //INDEX

    /**
     * server environment (schemas, catalogs, connections)
     * 
     */
    SCHEMA          = 500,  //SCHEMA
    CATALOG         = 501,  //CATALOG
    CONNECT         = 502,  //CONNECT
    TO              = 503,  //TO
    AUTHORIZATION   = 504,  //AUTHORIZATION
    DISCONNECT      = 505,  //DISCONNECT
    PASSWORD        = 506,  //password for connect to, begin with number, 
                            //password for connect can be PASSWORD or ID or INTVALUE
    CONNECTION      = 507,  //CONNECTION

    /**
     * persistent stored modules
     */
    CALL            = 550,  //CALL
    BOOLVALUE       = 551,  //BOOLVALUE true/false
    FUNCTION        = 552,  //FUNCTION
    RETURNS         = 553,  //RETURNS
    OUT             = 554,  //OUT
    INOUT           = 555,  //INOUT
    PROCEDURE       = 556,  //PROCEDURE
    DECLARE         = 557,  //DECLARE
    CONDITION       = 558,  //CONDITION
    SQLSTATE        = 559,  //SQLSTATE
    CURSOR          = 560,  //CURSOR
    RETURN          = 561,  //RETURN
    OPEN            = 562,  //OPEN
    CLOSE           = 563,  //CLOSE
    COLON           = 564,  //COLON :
    LOOP            = 565,  //LOOP
    LEAVE           = 566,  //LEAVE
    FETCH           = 567,  //FETCH
    INTO            = 568,  //INTO
    DO              = 569,  //DO
    CONTINUE        = 570,  //CONTINUE
    EXIT            = 571,  //EXIT
    UNDO            = 572,  //UNDO
    HANDLER         = 573,  //HANDLER
    IF              = 574,  //IF
    THEN            = 575,  //THEN
    ELSE            = 576,  //ELSE
    ELSEIF          = 577   //ELSEIF
};