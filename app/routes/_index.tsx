import type { MetaFunction } from "@remix-run/node";

import { Layout } from "~/components/layout";
import backgroundImage from "~/images/crayons-1445053_640.jpg";
// import { useOptionalUser } from "~/utils";

export const meta: MetaFunction = () => [{ title: "MC Mamina" }];

export default function Index() {
  // const user = useOptionalUser();
  return (
    <Layout
      className="w-full bg-cover bg-center text-indigo-800 font-light"
      style={{ backgroundImage: `url(${backgroundImage})` }}
    >
      <div className="w-full bg-indigo-100 bg-opacity-80 text-xl leading-10 flex justify-around py-12 my-12 md:py-16 md:my-16 xl:py-20 xl:my-28">
        <div className="w-full px-2 md:max-w-lg">
          Naše centrum je pre Vás otvorené každý pracovný deň v čase 9:00 -
          12:30 a 16:00 - 19:00 (v piatky iba doobeda).
        </div>
      </div>
      <div className="w-full bg-indigo-100 bg-opacity-80 text-xl leading-10 flex justify-around py-12 my-12 md:py-16 md:my-16 xl:py-20 xl:my-28">
        <div className="w-full md:max-w-2xl px-2">
          <h1 className="text-xxl font-semibold">Príspevky na chod</h1>
          <hr className="h-px border-t-0 my-4 bg-gradient-to-r from-indigo-500 to-pink-500 " />
          <div className="grid grid-cols-2 gap-10">
            <div className="col-span-2 sm:col-span-1">
              <h2 className="">Do herne</h2>
              <table className="text-sm">
                <tbody>
                  <tr>
                    <td className="p-1">0 do 1 rok</td>
                    <td className="p-1">1€</td>
                  </tr>
                  <tr>
                    <td className="p-1">1 rok a viac</td>
                    <td className="p-1">3€</td>
                  </tr>
                  <tr>
                    <td className="p-1">(Najviac však za 2 deti)</td>
                  </tr>
                  <tr>
                    <td className="p-1">10-vstupová pernamentka</td>
                    <td className="p-1">20€</td>
                  </tr>
                </tbody>
              </table>
            </div>
            <div className="col-span-2 sm:col-span-1">
              <h2 className="">Požičiavanie kníh</h2>
              <p className="text-sm">0,50 € / kniha / 2 týždne</p>
              <h2 className="mt-6">Narodeninové oslavy</h2>
              <p className="text-sm">50 € / oslava</p>
              <p className="text-sm">storno poplatok 10 €</p>
            </div>
          </div>
        </div>
      </div>
    </Layout>
  );
}
