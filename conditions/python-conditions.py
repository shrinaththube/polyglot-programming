

sun_is_bright = True
moon_is_cool = True

# in km
earth_diameter = 12742
moon_diameter = 3475

if sun_is_bright:
    print("Shine Shine and Shine")

if moon_is_cool:
    print("chill wibes")

if earth_diameter > moon_diameter:
    print("Mother earch is definitely bigger than moon")
    if sun_is_bright and moon_is_cool:
        print("Blinded by the light")
        sun_is_bright = False
else:
    print("Hmmm. Something looks odd")

if sun_is_bright:
    print("Hotness overloaded")
else:
    print("I am white dwarf now")
