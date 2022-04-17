# ZetaDB

ZetaDB is a relational database

## Supported Sql Currently

---
Quite a few of sql commands are supported according to the design, and all of them can be parsed if the syntax is correct.

Howerer, only a limited amount of commands are realised and can function well currently.

These sql commands are lists below:

### User Sign In, Sign Up, Drop User, Log Out

- CONNECT AS USER [ uer_name ] PASSWORD [ password ];
- CREATE USER [ user_name ] PASSWORD [ password ];
- DROP USER [ user_name ];
- QUIT;
- HALT;
  
### Modification On Table & Query

- SELECT * FROM [ table_name ] WHERE [ condition ];
- INSERT INTO [ table_name ] VALUES ( elementaryValueList );
- UPDATE [ table_name ] SET [ updateList ] WHERE [ condition ];
- DELETE FROM [ table_name ] WHERE [ condition ];

### Create & Drop Table

- CREATE TABLE [ table_name ] ( [ attributeDeclarationList ] );
- DROP TABLE [ table_name ];
  
