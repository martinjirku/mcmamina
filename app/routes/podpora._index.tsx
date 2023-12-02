import { FullWidthCard, CardContent } from "~/components/fullWidthCard";
import { getSponsors } from "~/models/sponsor.server";
import { MetaFunction, json } from "@remix-run/node";
import { useLoaderData } from "@remix-run/react";

export const loader = async () => {
  const sponsors = await getSponsors();
  return json({ sponsors });
};

export const meta: MetaFunction = () => [{ title: "Podporili nás" }];

export default function Index() {
  const { sponsors } = useLoaderData<typeof loader>();
  return (
    <FullWidthCard>
      <CardContent className="leading-normal">
        <h1 className="text-4xl font-bold text-center">Podporili nás</h1>
        <p className="mt-4 text-center">
          Ďakujeme všetkým, ktorí nám pomohli pri organizovaní materského
          centra.
        </p>
        <div className="mt-8 flex flex-wrap justify-center">
          {sponsors.map((sponsor) => (
            <div
              key={sponsor.url}
              className="flex flex-col items-center justify-center w-1/2 sm:w-1/3 md:w-1/4 lg:w-1/6 p-4"
            >
              <a href={sponsor.url} target="_blank" rel="noreferrer">
                <img src={sponsor.img} alt={"sponzor"} className="w-full" />
              </a>
            </div>
          ))}
        </div>
      </CardContent>
    </FullWidthCard>
  );
}
