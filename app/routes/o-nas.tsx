import { Layout } from "~/components/layout";
import backgroundImage from "~/images/crayons-1445053_640.jpg";
import { FullWidthCard, CardContent } from "~/components/fullWidthCard";

// https://www.mcmamina.sk/nova_stranka/vtour/
export default function ONasPage() {
  return (
    <Layout
      className="w-full bg-cover bg-center text-indigo-800 font-light"
      style={{ backgroundImage: `url(${backgroundImage})` }}
    >
      <FullWidthCard
        background="bg-indigo-950"
        padding="px-0 md:px-5"
        margin="mt-0"
      >
        <iframe
          src="https://www.mcmamina.sk/nova_stranka/vtour/"
          title="Virtuálna prehliadka"
          className="w-full lg:w-4/5 h-96"
          allowFullScreen
        ></iframe>
      </FullWidthCard>
      <FullWidthCard
        background="bg-indigo-200"
        padding="px-0 pt-12 pb-24 md:px-5"
        margin="m-0"
      >
        <CardContent className="text-indigo-800 px-2 sm:px-5 md:px-0 bg-">
          <h1 className="text-4xl font-semibold mb-10 text-center">
            Niečo o nás
          </h1>
          <p className="text-xl text-justify md:text-left">
            Materské centrum Mamina je nezisková organizácia, ktorá vznikla v
            roku 2004. Vznikla ako potreba stretávať sa v bezpečnom prostredí a
            zmysluplne tráviť čas so svojimi deťmi na materskej, či rodičovskej
            dovolenke. Naše Materské centrum tento rok oslávilo 19té narodeniny
            Za ten čas sme my, dobrovoľníčky, vybudovali nádhernú herňu,
            organizujeme krúžky, veľké akcie ako karneval, Míľa pre mamu,
            Mikuláš, prednášky, podporné skupiny,...Tvoríme komunitu, kde si
            každý nájde, to, čo potrebuje: zábavu, oddych, sebarealizáciu,
            informácie a hlavne krásny čas so svojim dieťaťom.
          </p>
        </CardContent>
      </FullWidthCard>
      <FullWidthCard
        background="bg-teal-900"
        padding="px-0 pt-12"
        margin="mt-0 mb-24"
      >
        <CardContent className="text-slate-200 px-2 sm:px-5 font-thin md:px-0 md:pb-10 text-xl text-justify md:text-left">
          <h1 className="text-4xl font-semibold mb-10 text-center">Projekty</h1>
          <p className="pb-10">
            V roku 2013 sme získali ocenenie mesta BB za dobrovoľnícku činnosť
            „Zrnká pomoci“, za zmysluplnú a tvorivú činnosť, prejavenie úcty k
            materstvu a úlohe matky a dosiahnuté výsledky v tejto oblasti.
          </p>
          <p className="pb-10">
            V roku 2016 sme získali podporu mesta BB a v spolupráci s OZ Duly
            sme rozbehli pilotný projekt sociálneho charakteru „Sprevádzanie
            tehotenstvom a príprava na pôrod pre sociálne znevýhodnené matky“. V
            rámci projektu sme viackrát zorganizovali zbierku šatstva a
            potravinovú pomoc pre rodiny v núdzi. V roku 2017 sa nám podarilo
            úspešne pokračovať v projekte.
          </p>
          <p className="pb-10">
            V roku 2019 sme sa stali súčasťou projektu „ Rovnosť šancí pre
            každého“, ktorý sa zaoberá problematikou aktívneho sociálneho
            začleňovania sa a svojimi aktivitami poukazuje a oboznamuje
            návštevníkov s rôznymi formami diskriminácie, ponúka bezplatné
            právne a kariérne poradenstvo.
          </p>
          <p className="pb-10">
            V roku 2020 sme boli nominovaní na ocenenie v ,,Srdce na dlani“, za
            naše dobrovoľné aktivity pre rozvoj komunity.
          </p>
          <p className="pb-10">
            V roku 2020 a 2021 sme pokračovali v spolupráci s InEnregy v
            projekte „ Rovnosť šancí pre každého“ v online priestore.
          </p>
          <p className="pb-10">
            Zapojili sme sa do kampane organizovanej Úniou materských centier s
            názvom ,,Láskavo do života“, ktorej nosnou témou bola starostlivosť
            o ženy v predpôrodnom a popôrodnom období v krajinách V4.
          </p>
          <p className="pb-10">
            Ďalej sme v roku 2021 rozbehli projekt ,, Učenie pre život“ v
            spolupráci s Úniou materských centier . Tento projekt pomáha
            znevýhodneným rodinám, jednorodičom, rodinám s adoptovaným dieťaťom
            a rodinám v hmotnej núdzi.
          </p>
          <p className="pb-10">
            V roku 2022 v spolupráci s Komunitnou nadáciou Zdravé mesto a
            Nadáciou Orange organizujeme kurzy slovenčiny pre odídencov z
            Ukrajiny. O tieto kurzy majú záujem prevažne ženy, ktoré majú malé
            deti. Počas kurzu slovenčiny je pre deti účastníčok pripravený
            zaujímavý kreatívny program. Vďaka odstráneniu jazykovej bariéry sa
            môžu ženy lepšie zapojiť do života na Slovensku a nájsť tu
            uplatnenie v práci a tiež bezpečie pre seba a svoje deti.
          </p>
          <p className="pb-10">
            V roku 2023 sme sa zapojili do projektu „Hrajme a učme sa bez
            hraníc“ v spolupráci s Unicefom. Tento projekt zahŕňa zriadenie
            hracích a vzdelávacích skupín (zamerané na deti vo veku 0-6 rokov).
            A zabezpečuje aj podporu pre rodičov, ktorí sa kvôli vojne ocitli v
            cudzom prostredí. Hracie skupiny slúžili hlavne ako integrácia detí
            a opatrovateľov ukrajinských odídencov. Naše materské centrum v
            rámci tohto projektu organizovalo bezplatnú - Montessori herničku,
            voľnočasové aktivity a komunitné akcie.
          </p>
          <p className="pb-10">
            V tomto roku nás podporila Komunitná nadácia Zdravé mesto v projekte
            Podpora a pomoc každej mame cez ktorý sme zakúpili kancelárske
            potreby a odbornú literatúru.
          </p>
          <p className="pb-10">
            Ďalej sme v roku 2023 dostali podporu od Nadácie SPP v projekte Ženy
            ženám – podpora a pomoc v komunite, vďaka ktorému sme sa tento raz
            zamerali na ženy/matky a organizovali sme pre ne workshopy, podporné
            skupiny a kurzy.
          </p>
        </CardContent>
      </FullWidthCard>
    </Layout>
  );
}
