# Use the hash, you must!

Av: Sverre Johann Bjørke

Det er desember, og det medfører jo sjølvsagt ein ny Star Wars-film. I det høve har Luka bestemt seg for å lage eit program for å finne ut kva han og kollegaene sine Star Wars-namn er.

I [denne fila](./names.txt) ligg det 4 lister, seprarert med ---. Ei med mannsnamn, ei med kvinnenamn, og ei liste med del en og del to av etternamn. For å plukke frå listene blir følgande gjort:

* ASCII-verdien for kvar bokstav i fornamnet blir lagt saman. Denne verdien modulo antall namn i respektivt manne- eller kvinnelista gir indeksen som brukast for å hente ut fornamn frå riktig liste.
* Etternamn delast på 2, der første halvdel blir den lengste i tilfeller der etternamna har odde antal bokstavar.
  * Posisjonen i alfabetet for kvar bokstav i første halvdel leggast saman. Denne summen modulo antal element i lista over første del av etternavn blir indeksen som brukast for å plukke frå denne lista.
  * ASCII-verdien for kvar bokstav i andre halvdel av etternamnet blir ganga saman. Dette produktet blir om personen er kvinne ganga med antall bokstavar i heile namnet. Om personen er mann blir produktet ganga med antall bokstavar i fornamnet.
    * I det resulterande talet sorterast alle siffera i synkande rekkefølge. Dette talet modulo antal element i del-to-lista blir indeksen for å hente ut siste del av starwars-namnet.

Gitt [ansattlista](./employees.csv), kva for eit Star Wars-namn endar flest ansatte opp med?

## Eksempel

For Jan,Johannsen,M:

* Jan = 74 + 97 + 110 = 281
* 281 % 36 = 29. Plass 29 i lista over mannsnamn (nullindeksert) = Poe
* Johannsen -> Johan, nsen
  * Alfabetverdiane (a = 1, b = 2 osv) av Johan = 10, 15, 8, 1, 14, som summert blir 48. 48 % 24 = 0. Indeks 0 i lista = Light
  * ASCII-verdiane av nsen = 110, 115, 101, 110. Produktet av desse blir 140541500. Ganga med 3 blir det 421624500.
  * 421624500 sortert blir 654422100. 654422100 modulo 26 = 20. Indeks 20 gir verse

Jan Johannsen blir dermed Poe Lightverse.

## Løsning

Svar: Mest brukte Star Wars navnet blir **Malkili Deathfire**

Diskusjonstråd: https://gist.github.com/knowitkodekalender/428527fe837f3513c8861cf7eb6ff8a3
