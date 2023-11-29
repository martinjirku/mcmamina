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
        margin="mt-0 mb-24"
      >
        <CardContent className="text-indigo-800">
          <h1 className="text-4xl font-semibold mb-10 text-center">
            Niečo o nás
          </h1>
          <p className="text-xl text-left">
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
    </Layout>
  );
}
