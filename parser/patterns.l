%option noyywrap

%{
#include "token.h"
%}

%%

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

[1-9]+[0-9]*|[0]|-[1-9]+[0-9]*                                  {return INTVALUE;}
[1-9]+[0-9]*[.][0-9]+|0[.][0-9]+                                {return FLOATVALUE;}
["][a-z_A-Z0-9]*["]|['][a-z_A-Z0-9]*[']                         {return STRINGVALUE;}
[Tt][Rr][Uu][Ee]|[Ff][Aa][Ll][Ss][Ee]                           {return BOOLVALUE;}

%%