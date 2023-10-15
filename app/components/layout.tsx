import { Link } from "@remix-run/react";
import type { LinkProps } from "@remix-run/react";
import type { FC, PropsWithChildren } from "react";

import {
  Phone,
  Mail,
  Bank,
  Home,
  Signature,
  Facebook,
  Instagram,
} from "./icons";

interface MenuItemProps {
  to: LinkProps["to"];
}
const MenuItem: FC<PropsWithChildren<MenuItemProps>> = ({ children, to }) => {
  return (
    <li className="inline-block align-middle relative rounded-md transition-colors ease-out-in delay-50 duration-200 bg-indigo-400 hover:bg-indigo-500 hover:opacity-90">
      <Link className="py-4 px-4 inline-block w-full h-full" to={to}>
        {children}
      </Link>
    </li>
  );
};

interface LyoutProps {
  className?: string | undefined;
  style?: React.CSSProperties | undefined;
}
export const Layout: FC<PropsWithChildren<LyoutProps>> = ({
  className,
  style,
  children,
}) => {
  return (
    <div className="h-screen flex flex-col">
      <header className="w-full bg-indigo-400 sticky top-0 shadow-lg">
        <nav className="flex justify-around py-6" aria-label="Hlavné">
          <ul className="flex gap-6 text-cyan-50">
            <MenuItem to="/">Domov</MenuItem>
            <MenuItem to="/o-nas">O nás</MenuItem>
            <MenuItem to="/kontakt">Kontakt</MenuItem>
            <MenuItem to="/podpora">Podporte nás</MenuItem>
          </ul>
        </nav>
      </header>
      <main className={`${className} flex-grow`} style={style}>
        {children}
      </main>
      <footer className="w-full flex pt-10 px-6 pb-16 content-center justify-around bg-neutral-900 h-min-40 text-indigo-200">
        <div className="grid grid-cols-4 gap-4">
          <address className="col-span-4 md:col-span-2 not-italic font-thin text-sm">
            <h2 className="font-bold leading-10 underline underline-offset-4">
              Materské centrum MAMINA o.z.
            </h2>
            <span className="inline-block">
              Tatranská 10, 97411 Banská Bystrica{" "}
            </span>
            <br />
            <span className="inline-block">
              <a href="tel:+421950492901">
                <Phone className="fill-indigo-100 mr-2" dimension={16} />
                +421950 492 901 - Darina Baksová (predpôrodné kurzy)
              </a>
            </span>
            <br />
            <span>
              <a href="mailto:akcie@mcmamina.sk">
                <Mail className="fill-indigo-100 mr-2" dimension={16} />
                akcie@mcmamina.sk
              </a>
            </span>
            <br />
            <h1 className="font-bold leading-10 underline underline-offset-4 inline-block mt-5">
              Oslavy
            </h1>
            <br />
            <span>
              <a href="mailto:oslavy@mcmamina.sk">
                <Mail className="fill-indigo-100 mr-2" dimension={16} />
                oslavy@mcmamina.sk
              </a>
            </span>
            <br />
            <span>
              <a href="tel:+421904102740">
                <Phone className="fill-indigo-100 mr-2" dimension={16} />
                +421904 102 740 - Zuzka Zemanová
              </a>
            </span>
          </address>
          <div className="col-span-4 md:col-span-1 font-thin text-sm">
            <h1 className="font-bold leading-10 underline underline-offset-4">
              Fakturačné údaje
            </h1>
            <span>
              <Home className="fill-indigo-100 mr-2" dimension={16} />
              Tatranská 10, 97411 Banská Bystrica
            </span>
            <br />
            <span>IČO: 37956825</span>
            <br />
            <span>DIČ: 2022358239</span>
            <br />
            <span>
              <Bank className="fill-indigo-100 mr-2" dimension={16} /> SK62 8330
              0000 0023 0190 0933
            </span>
            <br />
            <h1 className="inline-block mt-3 font-bold leading-10 underline underline-offset-4">
              Štatutárka
            </h1>
            <br />
            <span>Mária Krajčová</span>
            <br />
            <span>
              <a href="mailto:info@mcmamina.sk">
                <Mail className="fill-indigo-100 mr-2" dimension={16} />
                info@mcmamina.sk
              </a>
            </span>
          </div>
          <div className="col-span-4 md:col-span-1 mt-2 md:mt-0 font-thin text-sm flex justify-start items-center flex-col">
            <div className="flex-grow-0 flex justify-around items-center pb-5">
              <a
                href="https://www.facebook.com/MaterskeCentrumMamina/"
                aria-label="Facebook stránka Materského centra MAMINA"
                target="_blank"
                rel="noreferrer"
                title="Facebook stránka Materského centra MAMINA"
              >
                <Facebook
                  className="fill-indigo-100 mr-2 hover:animate-pulse"
                  dimension={42}
                />
              </a>
              <a
                href="https://www.facebook.com/MaterskeCentrumMamina/"
                aria-label="Instagram stránka Materského centra MAMINA"
                target="_blank"
                rel="noreferrer"
                title="Instagram stránka Materského centra MAMINA"
              >
                <Instagram
                  className="fill-indigo-100 mr-2 hover:animate-pulse"
                  dimension={42}
                />
              </a>
            </div>
            <div className="w-full flex-grow flex items-center justify-center mt-10 md:mt-0">
              <div
                className="px-5 py-3 bg-slate-800 rounded-md  hover:animate-spin hover:cursor-pointer"
                style={{ animationIterationCount: 1 }}
              >
                <Signature className="fill-indigo-100 mr-2" dimension={16} />{" "}
                Vytvoril MJ
              </div>
            </div>
          </div>
        </div>
      </footer>
    </div>
  );
};
