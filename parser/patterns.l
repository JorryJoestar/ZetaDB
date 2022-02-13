%option noyywrap

%{
#include "token.h"
%}

%%

[(]     {return LPAREN;}


%%