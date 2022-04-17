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
CONNECT AS USER [ uer_name ] PASSWORD [ password ];

CREATE USER [ user_name ] PASSWORD [ password ];

DROP USER [ user_name ];

QUIT;

HALT;
```

``DROP USER [ user_name ];`` and ``HALT;`` can only be executed a specific administrator
  
#### Modification On Table & Query

```
SELECT * FROM [ table_name ] WHERE [ condition ];

INSERT INTO [ table_name ] VALUES ( [ elementary_value_list ] );

UPDATE [ table_name ] SET [ update_list ] WHERE [ condition ];

DELETE FROM [ table_name ] WHERE [ condition ];
```

``[ condition ]`` is a nested structure, it can be:

- ``[ attribute_name ] [ compare_mark ] [ elementary_value ]``
- ``[ condition ] AND [ condition ]``
- ``[ condition ] OR [ condition ]``

``[ compare_mark ]`` can be:

- ``=``
- ``<>``
- ``<``
- ``>``
- ``<=``
- ``>=``

``[ elementary_value_list ]`` is a list, it can be:

- ``[ elementary_value ]``
- ``[ elementary_value_list ] , [ elementary_value ]``

``[ update_list ]`` is a list, it can be:

- ``[ update_list_entry ]``
- ``[ update_list ] , [ update_list_entry ]``
			
``[ update_list_entry ]`` is ``[ attribute_name ] = [ elementary_value ]``


#### Create & Drop Table

```
CREATE TABLE [ table_name ] ( [ attribute_declaration_list ] );
DROP TABLE [ table_name ];
```

``[ attribute_declaration_list ]`` is a list, it can be:

- ``[ attribute_declaration ]``
- ``[ attribute_declaration_list ] , [ attribute_declaration ]``
