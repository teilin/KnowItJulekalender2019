# Piratsneglen Sneglulf og skattekartet hans

Av: Hugo Wallenburg

Piratsneglen Sneglulf har funnet et skattekart! Skaperen av kartet var en slu snegle og laget kompliserte instruksjoner for å finne fram til skatten heller enn å tegne et enkelt kart.

Skattekartet definerer to tallrekker tilsvarende x- og y-koordinater i den sagnomsuste rutete ørkenen Kvadratahara som alle snegler kjenner til. Kvadratahara er nøyaktig 1000 meter lang og 1000 meter bred. Koordinatene fra tallrekken på skattekartet peker til ruter i Kvadratahara. Skattekartet sier at den sneglen som skal finne skatten må lete på hvert koordinat i rekken og vil til slutt finne skatten.

[Her er koordinatene.](./coords.txt)

Sneglulf følger koordinatene fra starten og sjekker hvert koordinat for skatt. Han beveger seg først horisontalt helt bort til X og deretter vertikalt opp eller ned til Y. Han bruker ett minutt på å bevege seg én rute, men på grunn av slimet han legger igjen bak seg bruker han ett minutt ekstra hver gang han er innom en rute om igjen (andre gang han er innom ruta bruker han 2 minutter, deretter 3 minutter, etc).

Etter å ha sjekket alle koordinatene finner Sneglulf skatten! Hvor lang tid brukte han på reisen?

## Eksempel

Gitt koordinatene [(1, 3), (1, 0), (3, 2)] får vi følgende rute: 1. Sneglulf beveger seg fra (0, 0) til (1, 0) (horisontalt) og deretter til (1, 3) (vertikalt). Han har brukt 4 minutter. Det ligger nå slim på koordinatene [(0, 0), (1, 0), (1, 1), (1, 2)]. 2. Sneglulf beveger seg fra (1, 3) til (1, 0). Han bruker ett minutt ekstra på rutene han allerede har vært innom Han har nå brukt 10 minutter. 3. Sneglulf beveger seg fra (1, 0) til (3, 0) og deretter til (3, 2). Han finner skatten! Han brukte totalt 14 minutter.

Oppgi svaret i minutter.