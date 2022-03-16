#include "isogram.h"
#include <stdio.h>

bool is_isogram(const char *phrase)
{
    if (!phrase)
    {
        return false;
    }

    int repeated = 0;
    while (*phrase != '\0')
    {
        char c = *phrase;
        if (c == ' ' || c == '-')
        {
            phrase++;
            continue;
        }
        if (repeated & (1 << c))
        {
            return false;
        }
        repeated |= (1 << c);
        phrase++;
    }
    return true;
}
