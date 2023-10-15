import { Layout } from "~/components/layout";
import backgroundImage from "~/images/crayons-1445053_640.jpg";

// https://www.mcmamina.sk/nova_stranka/vtour/
export default function ONasPage() {
  return (
    <Layout
      className="w-full bg-cover bg-center text-indigo-800 font-light"
      style={{ backgroundImage: `url(${backgroundImage})` }}
    >
      <div className="w-full bg-indigo-100 bg-opacity-80 text-xl leading-10 flex justify-around py-12 my-12 md:py-16 md:my-16 xl:py-20 xl:my-28">
        <iframe
          src="https://www.mcmamina.sk/nova_stranka/vtour/"
          title="Virtuálna prehliadka"
          className="w-4/5 h-96"
          allowFullScreen
        ></iframe>
      </div>
    </Layout>
  );
}
