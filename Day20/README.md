# Alvene ved det runde bord

Av: Hugo Wallenburg

Gavemakerne på Nordpolen har en høyt optimalisert arbeidsflyt. Alvene sitter i grupper på 5 rundt et bord hvor hver alv har en veldig spesifikk oppgave for gaven de lager.

Etter hver alv har jobbet på gaven er det noen enkle regler for hvem som skal gjøre neste steg.

## Forklaring

Alvene rundt bordet er nummerert 1–5 med klokka. Alv 1 starter med å utføre steg 1. Etter Alv 1 er ferdig med steg 1 går gaven videre med klokka til alv 2 som utfører steg 2.

Kompliserte gaver har kompliserte regler for bygging. Alv 1 utfører det første steget, deretter gjelder følgende regler for dagens gave:

1. Dersom neste steg er et primtall, skal alven som til nå har gjort færrest oppgaver utføre neste steg.
2. Dersom neste steg er delelig på 28, snur retningen gaven blir gitt videre. Dersom gjeldende retning er med klokka vil gaven nå gå mot klokka fram til retningen snur igjen. Om alv 1 skulle ha gitt gaven videre til alv 2, vil gaven nå istedenfor gå til alv 5.
3. Dersom neste steg er partall og neste alv har utført flest oppgaver, hopp over den alven og gi gaven til den neste.
4. Dersom neste steg er delelig på 7, skal alv 5 gjøre neste steg.
5. Gi gaven videre til neste alv i gjeldende klokkeretning.

For hvert steg er det kun den første gyldige regelen som skal brukes. Dersom det er flere kandidater for neste alv (f.eks. om flere alver har gjort like mange oppgaver) dropper vi regelen og prøver neste regel.

Gaven er ferdig etter **1 000 740 steg**.

## Eksempel

De første stegene blir som følger:

1. Alv 1 gjør steg 1. (start)
2. Alv 2 gjør steg 2. (regel 1 og regel 3 bortgår pga. flere kandidater, regel 5 gjelder)
3. Alv 3 gjør steg 3. (regel 1 bortgår, regel 5 gjelder)
4. Alv 4 gjør steg 4. (regel 3 bortgår, regel 5 gjelder)
5. Alv 5 gjør steg 5. (regel 1 gjelder)
6. Alv 1 gjør steg 6. (regel 3 bortgår, regel 5 gjelder)
7. Alv 5 gjør steg 7. (regel 4 gjelder)
8. Alv 1 gjør steg 8. (regel 3 bortgår, regel 5 gjelder)
9. osv.

## Oppgave

Alvenes fagorganisasjon for gavemakere er bekymret for ujevn arbeidsfordeling i gavemakerprossessen. Hva er differansen i antall steg utført mellom den alven som gjorde minst og den alven som gjorde mest?

## Løsning

Laget et program som gikk igjennom og gjorde det som den skulle på hvert steg. Stegene stemmer fint overens med det som er i eksempelet. Svaret jeg får ut er ikke rett. Debugget og debugget, men finner ikke ut hva jeg gjør feil. Antar jeg muligens tolker noe i oppgaven på feil måte. Dette var bare maksimalt irriterende.

Jeg fikk ikke submittet noen rett løsning, men leste over koden andre hadde gjort dagen derpå for å finne ut hva jeg gjorde feil. Det som er i filen nå er en løsning en annen hadde som jeg ønsket å teste for å se og lære av.

Løsning: 22766

Diskusjonstråd: https://gist.github.com/knowitkodekalender/7a113e547bdd9caafb018b79800975d9