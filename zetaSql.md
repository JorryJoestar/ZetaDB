ast
    ddl
        createTableStmt
        dropTableStmt
        alterTableAddStmt
        alterTableDropStmt
        createAssertStmt
        dropAssertStmt
        createViewStmt
        dropViewStmt
        createIndexStmt
        dropIndexStmt
        createTriggerStmt
        dropTriggerStmt
        createPsmStmt
        dropPsmStmt
    dml
        deleteStmt
            DELETE FROM ID WHERE condition SEMICOLON
        insertStmt
            INSERTINTO ID LPAREN attriNameList RPAREN VALUES LPAREN elementaryValueList RPAREN SEMICOLON
        updateStmt
            UPDATE ID SET updateList WHERE condition SEMICOLON
    dcl
        BEGINTOKEN SEMICOLON
        START TRANSACTION SEMICOLON
        COMMIT SEMICOLON
        ROLLBACK SEMICOLON
        SHOW TABLES SEMICOLON
        SHOW ASSERTIONS SEMICOLON
        SHOW VIEWS SEMICOLON
        SHOW INDEXS SEMICOLON
        SHOW TRIGGERS SEMICOLON
        SHOW FUNCTIONS SEMICOLON
        SHOW PROCEDURES SEMICOLON
        createUserStmt
            CREATE USER ID PASSWORD PASSWORDS SEMICOLON
        logUserStmt
            CONNECT AS USER ID PASSWORD PASSWORDS SEMICOLON
        psmCallStmt
    dql
        dqlEntry SEMICOLON



dqlEntry
    dqlEntry UNION dqlEntry
    dqlEntry DIFFERENCE dqlEntry
    dqlEntry INTERSECTION dqlEntry
    LPAREN dqlEntry RPAREN
    query
        selectStmt fromStmt WHERE condition GROUPBY attriNameOptionTableNameList HAVING condition ORDERBY orderByList LIMIT INTVALUE COMMA INTVALUE
