# Alvesekvens

Av: Didrik Pemmer Aalen

En alvesekvens er en selvlagende/selvbeskrivende sekvens som er definert på følgende måte.

Gitt et alfabet som består av tall: {x, y…, z} starter vi sekvensen ved å sette inn like mange x’er som verdien av x.
Videre skal vi legge til de andre tallene i alfabetet i den rekkefølgen de er definert i, og når vi kommer til slutten av alfabetet starter man på begynnelsen av alfabetet igjen.
Måten vi legger til tallene i sekvensen er ved å vite hvor mange ganger vi har lagt til tall i sekvensen tidligere, og bruke det tallet til å hente ut et tall fra sekvensen. Verdien på tallet vi henter ut fra sekvensen er hvor mange man skal legge til av neste tall i alfabetet.

## Eksempel

For eksempel gitt alfabetet {2, 3} vil de syv første operasjonene være:

* [2, 2] - Legg til 2 av 2.
* [2, 2, 3, 3] - Vi har nå lagt til tall én gang, så da henter vi tallet med indeks 1 i sekvensen (i bold over), også legger vi til verdien (2) 3ere til sekvensen.
* [2, 2, 3, 3, 2, 2, 2] - Vi har nå lagt til tall to ganger, så da henter vi tallet med indeks 2 i sekvensen (i bold over), også legger vi til verdien (3) 2ere til sekvensen.
* [2, 2, 3, 3, 2, 2, 2, 3, 3, 3] Og slik fortsetter vi nedover.
* [2, 2, 3, 3, 2, 2, 2, 3, 3, 3, 2, 2]
* [2, 2, 3, 3, 2, 2, 2, 3, 3, 3, 2, 2, 3, 3]
* [2, 2, 3, 3, 2, 2, 2, 3, 3, 3, 2, 2, 3, 3, 2, 2]

## Oppgave

Gitt et alfabet {2, 3, 5, 7, 11}, hva er summen av alle 7ere i sekvensen når lengden på sekvensen = 217532235?

## Løsning

Summen av alle 7er tall: 304552150

Diskusjonstråd: https://gist.github.com/knowitkodekalender/49ede6886d68cc21d385d574e5050fec