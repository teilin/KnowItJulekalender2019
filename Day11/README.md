# Sledestans

Av: Didrik Pemmer Aalen

For å rekke frem til alle barn hele verden går sleden til julenissen i 10703437km/t, og fordi nissen er magisk akslererer sleden til maksfart, samt stopper på null komma niks. Dessverre er det noe gale med magien som stopper sleden i det han er ferdig med runden sin. I et forsøk på å stoppe sleden finner nissen en kjempelang landingsstripe som består av forskjellige terreng.

De forskjellige terrengene har ulik innvirkning på farten til sleden, og er som følger:

```
G = gress. Senker farten til sleden med 27km/t
I = is. Øker farten til sleden med 12km/t * antall isområder på rad
A = asfalt. Senker farten til sleden med 59km/t
S = skog. Senker farten til sleden med 212km/t
F = fjell. Kommer alltid i par, der den første er oppover som senker farten med 70km/t mens den andre er nedoverbakke som øker farten med 35km/t.
```

Hvert terreng forekommer i lengder på 1 km.

## Eksempel

Hvis sledens fart er på 300km/t og landingsstripen ser slik ut: IIGGFFAIISGIFFSGFAAASS, vil farten påvirkes som følger:

```
 1: 300 + 12 = 312
 2: 312 + 24 = 336
 3: 336 - 27 = 309
 4: 309 - 27 = 282
 5: 282 - 70 = 212
 6: 212 + 35 = 247
 7: 247 - 59 = 188
 8: 188 + 12 = 200
 9: 200 + 24 = 224
10: 224 - 212 = 12
11: 12 - 15 = -3
````

Sleden stopper etter å ha kjørt 11 kilometer.

Oppgave
Hvor mange kilometer kjører nissen med sleden sin før den stopper gitt [følgende landingsstripe](./terreng.txt)?

## Løsning

Sleden stopper etter 202128 km.

Diskusjonstråd: https://gist.github.com/knowitkodekalender/ba6dc0c72f1656677786ad96fb0ecd4c
