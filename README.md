# ZetaDB

ZetaDB is a relational database that based on Golang and lex&yacc.

## Set Up

ZetaDB is composed by server part and client part. In order to use and test it, compiled clients for multiple platforms and architectures are provided.

- darwin_amd64
- darwin_m1
- linux_amd64
- windows_amd64

Them can be found here:
[executable_clients](https://github.com/JorryJoestar/ZetaDBclient/tree/master/%20executable_file)

## Supported Sql Currently

Quite a few of sql commands are supported according to the design, and all of them can be parsed if the syntax is correct. Howerer, only a limited amount of commands are realised and can function well currently. These sql commands are lists below:

#### User Sign In, Sign Up, Drop User, Log Out

```
CONNECT AS USER [ user_name ] PASSWORD [ password ];

CREATE USER [ user_name ] PASSWORD [ password ];

DROP USER [ user_name ];

QUIT;

HALT;
```

Attention: ``DROP USER [ user_name ];`` and ``HALT;`` can only be executed a specific administrator
  
#### Modification On Table & Query

```
SELECT * FROM [ table_name ] WHERE [ condition ];

INSERT INTO [ table_name ] VALUES ( [ elementary_value_list ] );

UPDATE [ table_name ] SET [ update_list ] WHERE [ condition ];

DELETE FROM [ table_name ] WHERE [ condition ];
```

Notice:

``[ condition ]`` can be ``[ attribute_name ] [ compare_mark ] [ elementary_value ]``, ``[ condition ] AND [ condition ]`` or ``[ condition ] OR [ condition ]``

``[ compare_mark ]`` can be ``=``, ``<>``, ``<``, ``>``, ``<=``, ``>=``

``[ elementary_value ]`` can be int value, float value or string

#### Create & Drop Table

```
CREATE TABLE [ table_name ] ( [ attribute_declaration_list ] );
DROP TABLE [ table_name ];
```

Notice:

``[ attribute_declaration ]`` is ``[ attribute_name ] [ data_type ]``

#### Supported Data Type

- CHAR
- VARCHAR
- BIT
- BITVARYING
- BOOLEAN
- INT
- INTEGER
- SHORTINT
- FLOAT
- REAL
- DOUBLEPRECISION
- DECIMAL
- NUMERIC
- DATE
- TIME 