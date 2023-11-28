import type { MetaFunction } from "@remix-run/node";
import { json } from "@remix-run/node";
import { Link, useLoaderData } from "@remix-run/react";

import { ActivitiesPresentation } from "~/components/activitiesPresentation";
import { Logo } from "~/components/animatedLogo";
import { TwoWeeksCalendar, Event } from "~/components/calendar2Weeks";
import { HorizontalDelimiter } from "~/components/delimiter";
import { FullWidthCard, CardContent } from "~/components/fullWidthCard";
import { Layout } from "~/components/layout";
import { getEvents } from "~/gateway/google.server";
import backgroundImage from "~/images/crayons-1445053_640.jpg";

export const meta: MetaFunction = () => [{ title: "MC Mamina" }];

export const loader = async () => {
  const eventItems = await getEvents();
  const events = eventItems?.items
    ?.filter((i) => i.summary)
    .map<Event>((i) => ({
      title: i.summary ?? "",
      start: i.start?.dateTime ?? undefined,
      end: i.end?.dateTime ?? undefined,
    }));
  return json({ events });
};

export default function Index() {
  // const user = useOptionalUser();
  const actionData = useLoaderData<typeof loader>();
  return (
    <Layout
      className="w-full bg-cover bg-center text-indigo-800 font-light"
      style={{ backgroundImage: `url(${backgroundImage})` }}
    >
      <FullWidthCard margin="mb-0 mt-12">
        <CardContent className="flex flex-col md:flex-row gap-3 md:gap-4 lg:gap-10 justify-center">
          <div className="flex-grow md:max-w-xs flex justify-around items-center">
            <Logo className="w-48 lg:w-64" animated />
          </div>
          <div className="flex-grow md:max-w-xl gap-6 flex flex-col items-center justify-evenly text-center">
            <span>
              Naše centrum je pre Vás otvorené každý pracovný deň v čase 9:00 -
              12:30 a 16:00 - 19:00 (v piatky iba doobeda). Najbližšie akcie
              nájdete v našom{" "}
              <Link className="underline underline-offset-4" to="/kalendar">
                kalendári
              </Link>
              . Najbližšie dni nás čaká:
            </span>
            <div className="w-full w-grow">
              <TwoWeeksCalendar events={actionData.events as Event[]} />
            </div>
          </div>
        </CardContent>
      </FullWidthCard>
      <FullWidthCard
        background="bg-indigo-950"
        padding="px-0 md:px-5"
        margin="mt-0"
      >
        <CardContent className="flex columns-2 text-indigo-800">
          <ActivitiesPresentation />
        </CardContent>
      </FullWidthCard>
      <FullWidthCard margin="mt-0 mb-12">
        <CardContent>
          <h1 className="text-xxl font-semibold">Príspevky na chod</h1>
          <HorizontalDelimiter />
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
                    <td className="p-1">25€</td>
                  </tr>
                </tbody>
              </table>
            </div>
            <div className="col-span-2 sm:col-span-1">
              <h2 className="">Požičiavanie kníh</h2>
              <p className="text-sm">0,50 € / kniha / mesiac</p>
              <p className="text-sm">5 € ročná permanentka</p>
            </div>
          </div>
          <br />
          <h1 className="text-xxl font-semibold">Podporte nás aj...</h1>
          <HorizontalDelimiter />
          <div className="grid grid-cols-2 gap-10">
            <div className="col-span-2 sm:col-span-1">
              <h2 className="">
                <Link to="./podpora/2-percenta-z-dane">2% z daní</Link>
              </h2>
            </div>
            <div className="col-span-2 sm:col-span-1">
              <h2 className="">
                <Link to="./podpora/dobrovolnictvo">Dobrovoľníctvom</Link>
              </h2>
            </div>
          </div>
        </CardContent>
      </FullWidthCard>
    </Layout>
  );
}
