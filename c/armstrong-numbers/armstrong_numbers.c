#include <stdio.h>
#include "armstrong_numbers.h"

bool is_armstrong_number(int candidate)
{
    int count, tmp, sum;

    sum = 0;
    count = 0;
    tmp = candidate;
    while (tmp != 0)
    {
        tmp /= 10;
        count++;
    }

    for (int i = candidate; i != 0; i /= 10)
    {
        int n = i % 10;
        int local = 1;
        for (int j = 0; j < count; j++)
        {
            local *= n;
        }
        sum += local;
    }
    return candidate == sum;
}
