%option noyywrap

%{
#include "token.h"
%}

%%
[Cc][Rr][Ee][Aa][Tt][Ee]                                        {return CREATE;}
[Tt][Aa][Bb][Ll][Ee]                                            {return TABLE;}

[Dd][Rr][Oo][Pp]                                                {return DROP;}

[Aa][Ll][Tt][Ee][Rr]                                            {return ALTER;}
[Aa][Dd][Dd]                                                    {return ADD;}

[Aa][Ss][Ss][Ee][Rr][Tt][Ii][Oo][Nn]                            {return ASSERTION;}

[Vv][Ii][Ee][Ww]                                                {return VIEW;}
[Aa][Ss]                                                        {return AS;}

[Ii][Nn][Dd][Ee][Xx]                                            {return INDEX;}
[Oo][Nn]                                                        {return ON;}

[Tt][Rr][Ii][Gg][Gg][Ee][Rr]                                    {return TRIGGER;}
[Rr][Ee][Ff][Ee][Rr][Ee][Nn][Cc][Ii][Nn][Gg]                    {return REFERENCING;}
[Bb][Ee][Ff][Oo][Rr][Ee]                                        {return BEFORE;}
[Uu][Pp][Dd][Aa][Tt][Ee]                                        {return UPDATE;}
[Oo][Ff]                                                        {return OF;}
[Aa][Ff][Tt][Ee][Rr]                                            {return AFTER;}
[Ii][Nn][Ss][Tt][Ee][Aa][Dd]                                    {return INSTEAD;}
[Ii][Nn][Ss][Ee][Rr][Tt]                                        {return INSERT;}
[Dd][Ee][Ll][Ee][Tt][Ee]                                        {return DELETE;}
[Oo][Ll][Dd]                                                    {return OLD;}
[Rr][Oo][Ww]                                                    {return ROW;}
[Nn][Ee][Ww]                                                    {return NEW;}
[Ff][Oo][Rr]                                                    {return FOR;}
[Ee][Aa][Cc][Hh]                                                {return EACH;}
[Ss][Tt][Aa][Tt][Ee][Mm][Ee][Nn][Tt]                            {return STATEMENT;}
[Ww][Hh][Ee][Nn]                                                {return WHEN;}
[Bb][Ee][Gg][Ii][Nn]                                            {return BEGINTOKEN;}
[Ee][Nn][Dd]                                                    {return END;}

[Uu][Nn][Ii][Qq][Uu][Ee]                                        {return UNIQUE;}
[Pp][Rr][Ii][Mm][Aa][Rr][Yy][ ][Kk][Ee][Yy]                     {return PRIMARYKEY;}
[Cc][Hh][Ee][Cc][Kk]                                            {return CHECK;}
[Ff][Oo][Rr][Ee][Ii][Gg][Nn][ ][Kk][Ee][Yy]                     {return FOREIGNKEY;}
[Rr][Ee][Ff][Ee][Rr][Ee][Nn][Cc][Ee][Ss]                        {return REFERENCES;}
[Nn][Oo][Tt][ ][Dd][Ee][Ff][Ee][Rr][Rr][Aa][Bb][Ll][Ee]         {return NOT_DEFERRABLE;}
[Dd][Ee][Ff][Ee][Rr][Rr][Aa][Bb][Ll][Ee][ ][Ii][Nn][Ii][Tt][Ii][Aa][Ll][Ll][Yy][ ][Dd][Ee][Ff][Ee][Rr][Rr][Ee][Dd] {return DEFERED_DEFERRABLE;}
[Dd][Ee][Ff][Ee][Rr][Rr][Aa][Bb][Ll][Ee][ ][Ii][Nn][Ii][Tt][Ii][Aa][Ll][Ll][Yy][ ][Ii][Mm][Mm][Ee][Dd][Ii][Aa][Tt][Ee] {return IMMEDIATE_DEFERRABLE;}
[Oo][Nn][ ][Uu][Pp][Dd][Aa][Tt][Ee][ ][Ss][Ee][Tt][ ][Nn][Uu][Ll][Ll] {return UPDATE_NULL;}
[Oo][Nn][ ][Uu][Pp][Dd][Aa][Tt][Ee][ ][Ss][Ee][Tt][ ][Cc][Aa][Ss][Cc][Aa][Dd][Ee] {return UPDATE_CASCADE;}
[Oo][Nn][ ][Dd][Ee][Ll][Ee][Tt][Ee][ ][Ss][Ee][Tt][ ][Nn][Uu][Ll][Ll] {return DELETE_NULL;}
[Oo][Nn][ ][Dd][Ee][Ll][Ee][Tt][Ee][ ][Ss][Ee][Tt][ ][Cc][Aa][Ss][Cc][Aa][Dd][Ee] {return DELETE_CASCADE;}
[Cc][Oo][Nn][Ss][Tt][Rr][Aa][Ii][Nn][Tt]                        {return CONSTRAINT;}
[Dd][Ee][Ff][Aa][Uu][Ll][Tt]                                    {return DEFAULT;}

[*]                                                             {return STAR;}
[Ss][Uu][Mm]                                                    {return SUM;}
[Aa][Vv][Gg]                                                    {return AVG;}
[Mm][Ii][Nn]                                                    {return MIN;}
[Mm][Aa][Xx]                                                    {return MAX;}
[Cc][Oo][Uu][Nn][Tt]                                            {return COUNT;}
[Dd][Ii][Ss][Tt][Ii][Nn][Cc][Tt]                                {return DISTINCT;}

[+]                                                             {return PLUS;}
[-]                                                             {return SUBTRACT;}
[/]                                                             {return DIVISION;}
[|][|]                                                          {return CONCATENATION;}

[Cc][Hh][Aa][Rr]                                                {return CHAR;}
[Vv][Aa][Rr][Cc][Hh][Aa][Rr]                                    {return VARCHAR;}
[Bb][Ii][Tt]                                                    {return BIT;}
[Bb][Ii][Tt][Vv][Aa][Rr][Yy][Ii][Nn][Gg]                        {return BITVARYING;}
[Bb][Oo][Oo][Ll][Ee][Aa][Nn]                                    {return BOOLEAN;}
[Ii][Nn][Tt]                                                    {return INT;}
[Ii][Nn][Tt][Ee][Gg][Ee][Rr]                                    {return INTEGER;}
[Ss][Hh][Oo][Rr][Tt][Ii][Nn][Tt]                                {return SHORTINT;}
[Ff][Ll][Oo][Aa][Tt]                                            {return FLOAT;}
[Rr][Ee][Aa][Ll]                                                {return REAL;}
[Dd][Oo][Uu][Bb][Ll][Ee][ ][Pp][Rr][Ee][Cc][Ii][Ss][Ii][Oo][Nn] {return DOUBLEPRECISION;}
[Dd][Ee][Cc][Ii][Mm][Aa][Ll]                                    {return DECIMAL;}
[Nn][Uu][Mm][Ee][Rr][Ii][Cc]                                    {return NUMERIC;}
[Dd][Aa][Tt][Ee]                                                {return DATE;}
[Tt][Ii][Mm][Ee]                                                {return TIME;}

[Ll][Ii][Kk][Ee]                                                {return LIKE;}
[Ii][Nn]                                                        {return IN;}
[Aa][Ll][Ll]                                                    {return ALL;}
[Aa][Nn][Yy]                                                    {return ANY;}
[Ii][Ss]                                                        {return IS;}
[Ee][Xx][Ii][Ss][Tt][Ss]                                        {return EXISTS;}

[=]                                                             {return EQUAL;}
[<][>]                                                          {return NOTEQUAL;}
[<]                                                             {return LESS;}
[>]                                                             {return GREATER;}
[<][=]                                                          {return LESSEQUAL;}
[>][=]                                                          {return GREATEREQUAL;}

[Aa][Nn][Dd]                                                    {return AND;}
[Oo][Rr]                                                        {return OR;}

[.]                                                             {return DOT;}

[Ee][Ll][Ss][Ee][Ii][Ff]                                        {return ELSEIF;}
[Cc][Aa][Ll][Ll]                                                {return CALL;}
[Tt][Hh][Ee][Nn]                                                {return THEN;}
[Ii][Ff]                                                        {return IF;}
[Ee][Ll][Ss][Ee]                                                {return ELSE;}
[Cc][Uu][Rr][Ss][Oo][Rr]                                        {return CURSOR;}
[Dd][Oo]                                                        {return DO;}
[Rr][Ee][Tt][Uu][Rr][Nn]                                        {return RETURN;}
[Ss][Ee][Tt]                                                    {return SET;}
[Oo][Uu][Tt]                                                    {return OUT;}
[Ii][Nn][Oo][Uu][Tt]                                            {return INOUT;}
[Dd][Ee][Cc][Ll][Aa][Rr][Ee]                                    {return DECLARE;}
[Pp][Rr][Oo][Cc][Ee][Dd][Uu][Rr][Ee]                            {return PROCEDURE;}
[Ff][Uu][Nn][Cc][Tt][Ii][Oo][Nn]                                {return FUNCTION;}
[Rr][Ee][Tt][Uu][Rr][Nn][Ss]                                    {return RETURNS;}

[Ff][Rr][Oo][Mm]                                                {return FROM;}
[Ww][Hh][Ee][Rr][Ee]                                            {return WHERE;}

[Ii][Nn][Ss][Ee][Rr][Tt][ ][Ii][Nn][Tt][Oo]                     {return INSERTINTO;}
[Vv][Aa][Ll][Uu][Ee][Ss]                                        {return VALUES;}

[Ss][Tt][Aa][Rr][Tt]                                            {return START;}
[Tt][Rr][Aa][Nn][Ss][Aa][Cc][Tt][Ii][Oo][Nn]                    {return TRANSACTION;}
[Cc][Oo][Mm][Mm][Ii][Tt]                                        {return COMMIT;}
[Rr][Oo][Ll][Ll][Bb][Aa][Cc][Kk]                                {return ROLLBACK;}
[Ss][Hh][Oo][Ww]                                                {return SHOW;}
[Tt][Aa][Bb][Ll][Ee][Ss]                                        {return TABLES;}
[Aa][Ss][Ss][Ee][Rr][Tt][Ii][Oo][Nn][Ss]                        {return ASSERTIONS;}
[Vv][Ii][Ee][Ww][Ss]                                            {return VIEWS;}
[Ii][Nn][Dd][Ee][Xx][Ss]                                        {return INDEXS;}
[Tt][Rr][Ii][Gg][Gg][Ee][Rr][Ss]                                {return TRIGGERS;}
[Ff][Uu][Nn][Cc][Tt][Ii][Oo][Nn][Ss]                            {return FUNCTIONS;}
[Pp][Rr][Oo][Cc][Ee][Dd][Uu][Rr][Ee][Ss]                        {return PROCEDURES;}
[Uu][Ss][Ee][Rr]                                                {return USER;}
[Pp][Aa][Ss][Ss][Ww][Oo][Rr][Dd]                                {return PASSWORD;}
[Cc][Oo][Nn][Nn][Ee][Cc][Tt]                                    {return CONNECT;}
[Ii][Nn][Ii][Tt][Ii][Aa][Ll][Ii][Zz][Ee]                        {return INITIALIZE;}
[Hh][Aa][Ll][Tt]                                                {return HALT;}

[Aa][Ss][Cc]                                                    {return ASC;}
[Dd][Ee][Ss][Cc]                                                {return DESC;}
[Cc][Rr][Oo][Ss][Ss]                                            {return CROSS;}
[Jj][Oo][Ii][Nn]                                                {return JOIN;}
[Nn][Aa][Tt][Uu][Rr][Aa][Ll]                                    {return NATURAL;}
[Ff][Uu][Ll][Ll]                                                {return FULL;}
[Oo][Uu][Tt][Ee][Rr]                                            {return OUTER;}
[Ll][Ee][Ff][Tt]                                                {return LEFT;}
[Rr][Ii][Gg][Hh][Tt]                                            {return RIGHT;}
[Ss][Ee][Ll][Ee][Cc][Tt]                                        {return SELECT;}
[Gg][Rr][Oo][Uu][Pp][ ][Bb][Yy]                                 {return GROUPBY;}
[Hh][Aa][Vv][Ii][Nn][Gg]                                        {return HAVING;}
[Oo][Rr][Dd][Ee][Rr][ ][Bb][Yy]                                 {return ORDERBY;}
[Ll][Ii][Mm][Ii][Tt]                                            {return LIMIT;}
[Uu][Nn][Ii][Oo][Nn]                                            {return UNION;}
[Dd][Ii][Ff][Ff][Ee][Rr][Ee][Nn][Cc][Ee]                        {return DIFFERENCE;}
[Ii][Nn][Tt][Ee][Rr][Ss][Ee][Cc][Tt][Ii][Oo][Nn]                {return INTERSECTION;}

[1-9]+[0-9]*|[0]|-[1-9]+[0-9]*                                  {return INTVALUE;}
[1-9]+[0-9]*[.][0-9]+|0[.][0-9]+                                {return FLOATVALUE;}
["][a-z_A-Z0-9]*["]|['][a-z_A-Z0-9]*[']                         {return STRINGVALUE;}
[Tt][Rr][Uu][Ee]|[Ff][Aa][Ll][Ss][Ee]                           {return BOOLVALUE;}

[;]                                                             {return SEMICOLON;}
[(]                                                             {return LPAREN;}
[)]                                                             {return RPAREN;}
[,]                                                             {return COMMA;}
[Nn][Oo][Tt]                                                    {return NOT;}
[Nn][Uu][Ll][Ll]                                                {return NULLMARK;}
[a-z_A-Z][a-z_A-Z0-9]*                                          {return ID;}
[a-z_A-Z0-9!]+                                                  {return PASSWORDS;}
[ \t\n]                                                           {/*ignore space*/}
%%