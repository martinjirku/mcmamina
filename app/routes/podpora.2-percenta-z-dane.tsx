import { HorizontalDelimiter } from "~/components/delimiter";
import { FullWidthCard, CardContent } from "~/components/fullWidthCard";

export default function Index() {
  return (
    <FullWidthCard>
      <CardContent className="leading-normal">
        <h1 className="text-2xl font-semibold pb-3">2% z dane</h1>
        <p className="pb-3">
          Aj tento rok Vás chceme poprosiť o podporu nášho materského centra
          darovaním 2% z Vašich daní. Ako na to?
        </p>
        <h2 className="text-2xl pt-4 font-light">
          Postup pre zamestnancov, za ktorých daňové priznanie podáva
          zamestnávateľ
        </h2>
        <HorizontalDelimiter space="sm" />
        <ul className="list-inside list-disc pl-5 pt-5">
          <li>
            Do 15.02.2023 požiadajte zamestnávateľa o vykonanie ročného
            zúčtovania zaplatených preddavkov na daň.
          </li>
          <li>
            Potom požiadajte zamestnávateľa, aby Vám vystavil tlačivo Potvrdenie
            o zaplatení dane.
          </li>
          <li>
            Z tohto Potvrdenia si viete zistiť dátum zaplatenia dane a
            vypočítať:
            <ol className="list-inside list-decimal pl-5">
              <li>
                2% z Vašej zaplatenej dane - to je maximálna suma, ktorú môžete
                v prospech prijímateľa poukázať, ak ste neboli dobrovoľníkom,
                alebo dobrovoľnícky odpracovali menej ako 40 hodín. Táto suma
                však musí byť minimálne 3 €.
              </li>
              <li>
                3% z Vašej zaplatenej dane, ak ste odpracovali dobrovoľnícky
                minimálne 40 hodín a získate o tom Potvrdenie od
                organizácie/organizácií, pre ktoré ste dobrovoľnícky pracovali.
              </li>
            </ol>
          </li>
          <li>
            Vyplňte Vyhlásenie o poukázaní podielu zaplatenej dane z príjmov
            fyzickej osoby (naše údaje tu už máte predvyplnené). Vyhlásenie
            môžete podať aj online.
          </li>
          <li>
            Obe tieto tlačivá, teda Vyhlásenie spolu s Potvrdením, doručte do
            30.04.2023 na daňový úrad (ak ste poukázali 3% z dane, povinnou
            prílohou k Vyhláseniu a Potvrdeniu o zaplatení dane je aj Potvrdenie
            o odpracovaní minimálne 40 hodín dobrovoľníckej činnosti).
          </li>
        </ul>
        <h2 className="text-2xl pt-6 font-light">
          Postup pre SZČO a zamestnancov, ktorí si daňové priznanie podávajú
          sami
        </h2>
        <HorizontalDelimiter space="sm" />
        <p className="pt-5">
          Údaje, ktoré potrebujete do daňového priznania uviesť, nájdete tu:
        </p>
        <p className="pt-3 text-md">
          <span className="font-bold">Obchodné meno (Názov)</span>: Materské
          centrum MAMINA
        </p>
        <p className="text-md">
          <span className="font-bold">Právna forma</span>: Občianske združenie
        </p>
        <p className="text-md">
          <span className="font-bold">IČO/SID</span>: 37956825
        </p>
        <p className="text-md">
          <span className="font-bold">Sídlo</span>: 97411 Banská Bystrica,
          Tatranská 10
        </p>
        <p className="pt-3">Vypočítajte si:</p>
        <ul className="list-inside list-disc pl-5">
          <li>
            2% z Vašej zaplatenej dane - to je maximálna suma, ktorú môžete v
            prospech prijímateľa poukázať, ak ste neboli dobrovoľníkom, alebo
            dobrovoľnícky odpracovali menej ako 40 hodín. Táto suma však musí
            byť minimálne 3 €.
          </li>
          <li>
            3% z Vašej zaplatenej dane, ak ste v roku 2022 odpracovali
            dobrovoľnícky minimálne 40 hodín a získate o tom Potvrdenie od
            organizácie/organizácií, pre ktoré ste v roku 2022 dobrovoľnícky
            pracovali.
          </li>
        </ul>
        <p className="pt-2">
          V daňovom priznaní pre fyzické osoby sú už uvedené kolónky na
          poukázanie 2% (3%) z dane v prospech 1 prijímateľa. Riadne vyplnené
          daňové priznanie doručte v lehote, ktorú máte na podanie daňového
          priznania (zvyčajne do 31.03.) na Váš daňový úrad a v tomto termíne aj
          zaplaťte daň z príjmov. Ak ste poukázali 3% z dane, povinnou prílohou
          k Daňovému priznaniu je aj Potvrdenie o odpracovaní minimálne 40 hodín
          dobrovoľníckej činnosti.
        </p>
        <h2 className="text-2xl pt-6 font-light">Postup pre právnické osoby</h2>
        <HorizontalDelimiter space="sm" />
        <p className="pt-5">
          Právnické osoby môžu poukázať 1,0% (2%) z dane aj viacerým
          prijímateľom, minimálna výška v prospech jedného prijímateľa je 8€.
        </p>
        <ul className="list-inside list-disc pl-5">
          <li>
            Ak právnická osoba (firma) v roku 2022 až do termínu na podanie
            daňového priznania a zaplatenie dane v roku 2023 (zvyčajne do
            31.3.2023) nedarovala financie vo výške minimálne 0,5% z dane na
            verejnoprospešný účel (aj inej organizácii, nemusí byť iba
            prijímateľovi), tak môže poukázať iba 1,0% z dane – vyznačí v
            daňovom priznaní, že poukazuje iba 1,0% z dane.
          </li>
          <li>
            Ak právnická osoba (firma) v roku 2022 až do termínu na podanie
            daňového priznania a zaplatenie dane v roku 2023 (zvyčajne do
            31.3.2023) darovala financie vo výške minimálne 0,5% z dane na
            verejnoprospešný účel (aj inej organizácii, nemusí byť iba
            prijímateľovi), tak môže poukázať 2% z dane – označí v daňovom
            priznaní, že poukazuje 2% z dane.
          </li>
        </ul>
        <p className="pt-5">
          Údaje, ktoré potrebujete do daňového priznania uviesť, nájdete tu:
        </p>
        <p className="pt-3 text-md">
          <span className="font-bold">Obchodné meno (Názov)</span>: Materské
          centrum MAMINA
        </p>
        <p className="text-md">
          <span className="font-bold">Právna forma</span>: Občianske združenie
        </p>
        <p className="text-md">
          <span className="font-bold">IČO/SID</span>: 37956825
        </p>
        <p className="text-md">
          <span className="font-bold">Sídlo</span>: 97411 Banská Bystrica,
          Tatranská 10
        </p>
        <p className="pt-3">
          Vypočítajte si Vaše 1,0% (2%) z dane z príjmov právnickej osoby– to je
          maximálna suma, ktorú môžete poukázať v prospech
          prijímateľa/prijímateľov, poukázať môžete aj menej ako 1,0% (2%), musí
          však byť splnená podmienka minimálne 8 € na jedného prijímateľa.
        </p>
        <p>
          V daňovom priznaní pre právnické osoby– časť VI. sú už uvedené kolónky
          na poukázanie 1,0% (2%) z dane v prospech 1 prijímateľa. Pokiaľ ste si
          vybrali viac prijímateľov, vložte do daňového priznania ďalší list
          papiera ako prílohu (je uvedená na poslednej strane DP) a uveďte tam
          analogicky všetky potrebné identifikačné údaje o prijímateľoch a sumu,
          ktorú chcete v ich prospech poukázať. V kolónke 5 uveďte, koľkým
          prijímateľom chcete podiel zaplatenej dane poukázať.
        </p>
        <p>
          Riadne vyplnené daňové priznanie doručte v lehote, ktorú máte na
          podanie daňového priznania, na Váš daňový úrad a v tomto termíne aj
          zaplaťte daň z príjmov.
        </p>
      </CardContent>
    </FullWidthCard>
  );
}
