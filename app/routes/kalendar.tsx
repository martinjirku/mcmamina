import { Outlet } from "@remix-run/react";

import { Layout } from "~/components/layout";
import backgroundImage from "~/images/crayons-1445053_640.jpg";

export default function Index() {
  return (
    <Layout
      className="w-full bg-cover bg-center text-indigo-800 font-light"
      style={{ backgroundImage: `url(${backgroundImage})` }}
    >
      <Outlet />
    </Layout>
  );
}
