%option noyywrap

%{
#include "token.h"
%}

%%
[Cc][Rr][Ee][Aa][Tt][Ee]                                        {return CREATE;}
[Tt][Aa][Bb][Ll][Ee]                                            {return TABLE;}

[Dd][Rr][Oo][Pp]                                                {return DROP;}

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
[ \t]                                                           {/*ignore space*/}
%%