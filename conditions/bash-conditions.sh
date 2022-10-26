#!/bin/bash

# https://linuxize.com/post/bash-if-else-statement/

sun_is_bright=true
moon_is_cool=true

if $sun_is_bright; then
    echo "Shine Shine and Shine"
fi

if $moon_is_cool; then
    echo "chill vibes"
fi

# in km
earth_diameter=12742
moon_diameter=3475

if [ $earth_diameter -gt $moon_diameter ]; then
    echo "Mother Earth is definitely bigger than moon"
    if $sun_is_bright && $moon_is_cool; then
        echo "Blinded by the light"
        sun_is_bright=false
    fi
else
    echo "Hmmm. Something looks odd"
fi

if $sun_is_bright; then
    echo "Hotness overloaded"
else
    echo "I am white dwarf now"
fi

