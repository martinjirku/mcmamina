import type { MetaFunction } from "@remix-run/node";

import { Layout } from "~/components/layout";
import backgroundImage from "~/images/crayons-1445053_640.jpg";
// import { useOptionalUser } from "~/utils";

export const meta: MetaFunction = () => [{ title: "Remix Notes" }];

export default function Index() {
  // const user = useOptionalUser();
  return (
    <Layout
      className="w-full bg-cover bg-center"
      style={{ backgroundImage: `url(${backgroundImage})` }}
    >
      content
    </Layout>
  );
}
