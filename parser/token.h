enum {
//createTableStmt
    CREATE          = 200,
    TABLE           = 201,

//dropTableStmt
    DROP            = 202,

//alterTableStmt
    ALTER           = 203,
    ADD             = 204,

//createAssertStmt
    ASSERTION       = 205,

//createViewStmt
    VIEW            = 206,
    AS              = 207,

//createIndexStmt
    INDEX           = 208,
    ON              = 209,

//createTriggerStmt
    TRIGGER         = 210,
    REFERENCING     = 211,
    BEFORE          = 212,
    UPDATE          = 213,
    OF              = 214,
    AFTER           = 215,
    INSTEAD         = 216,
    INSERT          = 217,
    DELETE          = 218,
    OLD             = 219,
    ROW             = 220,
    NEW             = 221,
    FOR             = 222,
    EACH            = 223,
    STATEMENT       = 224,
    WHEN            = 225,
    BEGINTOKEN      = 226,  //BEGIN
    END             = 227,

//aggregation
    STAR            = 228,  //*
    AVG             = 229,
    MIN             = 230,
    MAX             = 231,
    COUNT           = 232,
    SUM             = 233,
    DISTINCT        = 234,

//constraint
    UNIQUE            = 235,
    PRIMARYKEY        = 236,  //PRIMARY KEY
    CHECK             = 237,
    FOREIGNKEY        = 238,
    REFERENCES        = 239,
    NOT_DEFERRABLE           = 240,  //NOT DEFERRABLE
    DEFERED_DEFERRABLE       = 241,  //DEFERRABLE INITIALLY DEFERRED
    IMMEDIATE_DEFERRABLE     = 242,  //DEFERRABLE INITIALLY IMMEDIATE
    UPDATE_NULL      = 243,  //ON UPDATE SET NULL
    UPDATE_CASCADE   = 244,  //ON UPDATE SET CASCADE
    DELETE_NULL      = 245,  //ON DELETE SET NULL
    DELETE_CASCADE   = 246,  //ON DELETE SET CASCADE
    CONSTRAINT      = 247,
    DEFAULT         = 248,

//condition
    AND             = 249,
    OR              = 250,

//predicate
    LIKE            = 251,
    NOT             = 252,
    IN              = 253,
    ALL             = 254,
    ANY             = 255,
    IS              = 256,
    EXISTS          = 257,

//expression
    SUBTRACT        = 258,  //-
    DIVISION        = 259,  ///
    CONCATENATION   = 260,  //||
    PLUS            = 261,  //+

    NOTEQUAL        = 262,  // <>
    LESS            = 263,  // <
    GREATER         = 264,  // >
    LESSEQUAL       = 265,  // <=
    GREATEREQUAL    = 266,  // >=
    EQUAL           = 267,  // =

//elementary value
    INTVALUE          = 268,  //int value
    FLOATVALUE        = 269,  //float value
    STRINGVALUE       = 270,  //string surrounded by single or double quotation marks
    BOOLVALUE         = 271,  //BOOLVALUE true/false

//attriNameOptionTableName
    DOT               = 272,  //.

//domain
    CHAR            = 273,
    VARCHAR         = 274,  //VARCHAR(n)
    BIT             = 275,  //BIT(n)
    BITVARYING      = 276,  //BITVARYING(n)
    BOOLEAN         = 277,
    INT             = 278,
    INTEGER         = 279,
    SHORTINT        = 280,
    FLOAT           = 281,
    REAL            = 282,
    DOUBLEPRECISION = 283,
    DECIMAL         = 284,  //DECIMAL(n,d)
    NUMERIC         = 285,  //NUMERIC(n,d)
    DATE            = 286,
    TIME            = 287,

//psm
    CALL            = 288,
    ELSEIF          = 289,
    THEN            = 290,
    IF              = 291,
    ELSE            = 292,
    CURSOR          = 293,
    DO              = 294,
    RETURN          = 295,
    SET             = 296,
    OUT             = 297,
    INOUT           = 298,
    DECLARE         = 299,
    FUNCTION        = 300,
    RETURNS         = 301,
    PROCEDURE       = 302,

//delete
    FROM            = 303,
    WHERE           = 304,

//insert
    INSERTINTO      = 305,  //INSERT INTO
    VALUES          = 306,

//dcl
    START           = 307,
    TRANSACTION     = 308,
    COMMIT          = 309,
    ROLLBACK        = 310,
    SHOW            = 311,
    TABLES          = 312,
    ASSERTIONS      = 313,
    VIEWS           = 314,
    INDEXS          = 315,
    TRIGGERS        = 316,
    FUNCTIONS       = 317,
    PROCEDURES      = 318,
    USER            = 319,
    PASSWORD        = 320,
    PASSWORDS       = 321,
    CONNECT         = 322, 
    INITIALIZE      = 323,
    HALT            = 324,

//dql
    JOIN            = 325,
    NATURAL         = 326,
    FULL            = 327,
    LEFT            = 328,
    RIGHT           = 329,
    OUTER           = 330,
    CROSS           = 331,
    SELECT          = 332,
    GROUPBY         = 333,  //GROUP BY
    HAVING          = 334,
    ORDERBY         = 335,  //ORDER BY
    LIMIT           = 336,
    UNION           = 337,
    DIFFERENCE      = 338,
    INTERSECTION    = 339,
    DESC            = 340,
    ASC             = 341,

//public
    RPAREN            = 342,  //)
    SEMICOLON         = 343,  //;
    COMMA             = 344,  //,
    ID                = 345,  //begin with non-number character
    TABLEDOTATTRIBUTE = 346, //tableName.attributeName
    NULLMARK          = 347,  //NULL
    LPAREN            = 348,  //(
};