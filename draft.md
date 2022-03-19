storage/indexPage.go
    mode1 page: Modification on elements and pointerPageId would not influent pointerNum, always remember to update pointerNum if necessary.
    
    mode2 page: In order to insert a new indexRecord, use IndexPageInsertNewIndexRecord(), this will increase the length of records automatically. when calling IndexPageSetIndexRecordAt(), it is invalid to modify a position that have not contained an indexRecord before.