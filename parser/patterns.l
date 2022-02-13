%option noyywrap

%{
#include "token.h"
%}

%%

[1-9]+[0-9]*|[0]|-[1-9]+[0-9]*                                  {return INTVALUE;}
[1-9]+[0-9]*[.][0-9]+|0[.][0-9]+                                {return FLOATVALUE;}
["][a-z_A-Z0-9]*["]|['][a-z_A-Z0-9]*[']                         {return STRINGVALUE;}
[Tt][Rr][Uu][Ee]|[Ff][Aa][Ll][Ss][Ee]                           {return BOOLVALUE;}

%%