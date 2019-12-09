# Lotto-Lotte feirer jul i Vegas

Av: Hugo Wallenburg

I familien til Lotto-Lotte er det tradisjon å dra til Las Vegas for å feire jul. De er alle glad i kasinoleker og pengespill og finner stor glede i å telle kort (i hemmelighet såklart).

På årets tur finner familien en ny type maskin: en programmatisk spilleautomat. Spilleautomaten har ti hjul nummerert 0-9 med 4 operasjoner på hvert hjul. Det siste sifferet i det gjeldende gevinstbeløpet bestemmer hvilket hjul neste operasjon hentes fra; operasjonen derfra utføres og det valgte hjulet snur ett hakk. Dette fortsetter helt til man utfører operasjonen “STOPP”, og gevinsten blir betalt ut.

Tabell over mulige operasjoner:

```
+----------------+------------------------------------------------------------------------+
| Operasjonsnavn | Beskrivelse                                                            |
+----------------+------------------------------------------------------------------------+
| PLUSS4         | Legg til 4                                                             |
| PLUSS101       | Legg til 101                                                           |
| MINUS9         | Trekk fra 9                                                            |
| MINUS1         | Trekk fra 1                                                            |
| REVERSERSIFFER | Snu rekkefølgen på tallets sifre. ex: 123 = 321, -123 = -321, 012 = 21 |
| OPP7           | Pluss på 1 til det siste sifferet er 7. ex: 21 = 27, -13 = -7, 17 = 17 |
| GANGEMSD       | Gang tallet med dets øverste siffer. ex: 23 = 46, -31 = -93            |
| DELEMSD        | Del tallet på dets øverste siffer. ex: 23 = 11, -31 = -11              |
| PLUSS1TILPAR   | Legg til 1 på alle partallssiffer. ex: 120 = 131, -1234 = -1335        |
| TREKK1FRAODDE  | Trekk fra 1 på alle oddetallssiffer. ex: 1234 = 224, -1234 = -224      |
| ROTERPAR       | Roter alle partallshjul i maskinen ett hakk                            |
| ROTERODDE      | Roter alle oddetallshjul i maskinen ett hakk                           |
| ROTERALLE      | Roter alle hjulene i maskinen                                          |
| STOPP          | Stopp maskinen                                                         |
+----------------+------------------------------------------------------------------------+
```

Merk at divisjon er [Evklids heltallsdivisjon](https://en.wikipedia.org/wiki/Euclidean_division). Resultatet fra divisjon av negative tall rundes nedover (-31 / 3 = -11).

Startposisjon for hjulene finnes i [denne filen](./wheels.txt).

## Eksempel

Dersom Lotto-Lotte putter 6 mynter i maskinen vil de første operasjonene skje som følger:

* Fra hjul 6: MINUS1: 6 - 1 = 5
* Fra hjul 5: GANGEMSD: 5 * 5 = 25
* Fra hjul 5: PLUSS4: 25 + 4 = 29
* Fra hjul 9: PLUSS4: 29 + 4 = 33
* Fra hjul 3: MINUS9: 33 - 9 = 24
* Fra hjul 4: ROTERODDE: Alle oddetallshjul (1, 3, 5, 7, 9) roteres ett hakk
* Fra hjul 4: MINUS1: 24 - 1 = 23
* Fra hjul 3: TREKK1FRAODDE: Trekker 1 fra oddetallssifferet 3 i 23: 23 -> 22
* etc.

Etter de åtte første operasjonene beskrevet over er Lotto-Lottes gevinst 22.

Lotto-Lotte har 10 mynter på seg og kan velge å putte hvilket antall hun vil på automaten før hun starter den. Dessverre har hun dårlig tid og rekker kun å spille én gang. Hva er den største gevinsten Lotte kan få på ett spill?