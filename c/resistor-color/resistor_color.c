#include "resistor_color.h"

int color_code(resistor_band_t color)
{
    return color;
}

resistor_band_t *colors()
{
    static resistor_band_t colors[WHITE];
    for (resistor_band_t color = BLACK; color <= WHITE; color++)
    {
        colors[color] = color;
    }
    return colors;
}
