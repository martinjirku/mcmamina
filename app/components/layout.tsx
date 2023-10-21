import { NavLink } from "@remix-run/react";
import type { LinkProps } from "@remix-run/react";
import { useReducer } from "react";
import type { FC, PropsWithChildren } from "react";

import { AnimatedHamburger } from "./animatedHamburger";
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
  const defaultClasses =
    "rounded-md transition-colors ease-out-in delay-50 duration-200 text-indigo-50 py-4 px-4 inline-block w-full h-full";
  return (
    <li className="inline-block align-middle relative">
      <NavLink
        className={({ isActive }) =>
          `${defaultClasses} ${
            isActive
              ? "bg-indigo-500 opacity-70 pointer-events-none"
              : "bg-indigo-400 hover:bg-indigo-500 hover:opacity-90"
          }`
        }
        to={to}
      >
        {children}
      </NavLink>
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
  const [isOpen, toggleOpen] = useReducer((v: boolean) => !v, false);
  return (
    <div className="h-screen flex flex-col">
      <header className="w-full bg-indigo-400 sticky shadow-lg -top-5 text-cyan-50">
        <nav className="hidden md:flex justify-around py-6" aria-label="Hlavné">
          <ul className="flex gap-6">
            <MenuItem to="/">Domov</MenuItem>
            <MenuItem to="/o-nas">O nás</MenuItem>
            <MenuItem to="/kontakt">Kontakt</MenuItem>
            <MenuItem to="/podpora">Podporte nás</MenuItem>
          </ul>
        </nav>
        <nav className="w-full md:hidden left-0 flex-col" aria-label="Hlavné">
          <div className="w-full flex relative bg-indigo-400 z-30">
            <h1 className="font-light p-2 py-2 text-2xl">Mc Mamina</h1>
            <button
              className="stroke-indigo-100 absolute right-0 pr-2"
              onClick={toggleOpen}
              aria-label="Open menu"
            >
              <AnimatedHamburger open={isOpen} dimension={46} />
            </button>
          </div>
          <ul
            className={`${
              isOpen ? "h-auto" : "h-0 overflow-hidden"
            } flex z-10 flex-col gap-6 absolute top-12 bg-indigo-400 w-full text-cyan-50 transition-transform ease-in-out duration-500 transform ${
              isOpen
                ? "translate-y-0 opacity-100"
                : "-translate-y-full opacity-0"
            }`}
          >
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
      <Footer />
    </div>
  );
};

const Footer: FC = () => (
  <footer className="w-full flex pt-10 px-6 pb-16 shadow-xl content-center justify-around bg-neutral-900 h-min-40 text-indigo-100 lg:sticky lg:-bottom-48 ">
    <div className="flex flex-col flex-wrap sm:flex-row gap-10 md:gap-x-32 justify-between xl:flex-nowrap md:max-w-5xl">
      <address className="flex-grow not-italic font-thin text-sm">
        <h2 className="font-bold leading-10 underline underline-offset-4">
          Materské centrum MAMINA o.z.
        </h2>
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
        <span>Martina Cabanová</span>
        <br />
        <span>
          <a href="tel:+421904102740">
            <Phone className="fill-indigo-100 mr-2" dimension={16} />
            +421948 523 493 - štatutárka
          </a>
        </span>
        <br />
        <span>
          <a href="mailto:mcmamina@mcmamina.sk">
            <Mail className="fill-indigo-100 mr-2" dimension={16} />
            mcmamina@mcmamina.sk
          </a>
        </span>
        <br />
        <span>
          <a href="mailto:info@mcmamina.sk">
            <Mail className="fill-indigo-100 mr-2" dimension={16} />
            info@mcmamina.sk
          </a>
        </span>
      </address>
      <div className="flex-grow font-thin text-sm">
        <OpeningHours />
        <Contacts />
      </div>
      <div className="flex-grow flex-shrink font-thin text-sm flex w-full md:w-auto justify-start items-center flex-col">
        <div className="flex-grow-0 flex justify-around items-center pb-5">
          <Socials />
        </div>
        <div className="w-full flex-grow flex items-center justify-center mt-10 lg:mt-0">
          <a
            href="https://linkedin.com/in/martin-j-65786267"
            target="_blank"
            rel="noreferrer"
            className="px-5 py-3 bg-slate-800 rounded-md  hover:animate-spin hover:cursor-pointer"
            style={{ animationIterationCount: 1 }}
          >
            <Signature className="fill-indigo-100 mr-2" dimension={16} />{" "}
            Vytvoril MJ
          </a>
        </div>
      </div>
    </div>
  </footer>
);

const OpeningHours: FC = () => (
  <>
    <h1 className="font-bold leading-10 underline underline-offset-4">
      Otváracie hodiny
    </h1>
    <table>
      <thead>
        <tr className="hidden">
          <th>Dni</th>
          <th>Čas</th>
        </tr>
      </thead>
      <tbody>
        <tr>
          <td className="w-44 sm:w-32 md:w-36 lg:w-40">Pondelok</td>
          <td>
            <time dateTime="09:00-2" aria-label="09:00">
              9:00
            </time>{" "}
            -{" "}
            <time dateTime="12:30-2" aria-label="12:30">
              12:30
            </time>{" "}
            |{" "}
            <time dateTime="16:00-2" aria-label="16:00">
              16:00
            </time>{" "}
            -{" "}
            <time dateTime="19:00-2" aria-label="19:00">
              19:00
            </time>
          </td>
        </tr>
        <tr>
          <td>Utorok</td>
          <td>
            <time dateTime="09:00-2" aria-label="09:00">
              9:00
            </time>{" "}
            -{" "}
            <time dateTime="12:30-2" aria-label="12:30">
              12:30
            </time>{" "}
            |{" "}
            <time dateTime="16:00-2" aria-label="16:00">
              16:00
            </time>{" "}
            -{" "}
            <time dateTime="19:00-2" aria-label="19:00">
              19:00
            </time>
          </td>
        </tr>
        <tr>
          <td>Streda</td>
          <td>
            <time dateTime="09:00-2" aria-label="09:00">
              9:00
            </time>{" "}
            -{" "}
            <time dateTime="12:30-2" aria-label="12:30">
              12:30
            </time>{" "}
            |{" "}
            <time dateTime="16:00-2" aria-label="16:00">
              16:00
            </time>{" "}
            -{" "}
            <time dateTime="19:00-2" aria-label="19:00">
              19:00
            </time>
          </td>
        </tr>
        <tr>
          <td>Štvrtok</td>
          <td>
            <time dateTime="09:00-2" aria-label="09:00">
              9:00
            </time>{" "}
            -{" "}
            <time dateTime="12:30-2" aria-label="12:30">
              12:30
            </time>{" "}
            |{" "}
            <time dateTime="16:00-2" aria-label="16:00">
              16:00
            </time>{" "}
            -{" "}
            <time dateTime="19:00-2" aria-label="19:00">
              19:00
            </time>
          </td>
        </tr>
        <tr>
          <td>Piatok</td>
          <td>
            <time dateTime="09:00-2" aria-label="09:00">
              9:00
            </time>{" "}
            -{" "}
            <time dateTime="12:30-2" aria-label="12:30">
              12:30
            </time>
          </td>
        </tr>
      </tbody>
    </table>
  </>
);

const Contacts: FC = () => (
  <>
    <h1 className="font-bold leading-10 underline underline-offset-4 inline-block mt-5">
      Kontakty
    </h1>
    <br />
    <table>
      <tbody>
        <tr>
          <td
            className="w-44 sm:w-32 md:w-36 lg:w-40 align-text-top"
            rowSpan={2}
          >
            Oslavy
          </td>
          <td>
            <a href="mailto:oslavy@mcmamina.sk">
              <Mail className="fill-indigo-100 mr-2" dimension={16} />
              oslavy@mcmamina.sk
            </a>
          </td>
        </tr>
        <tr>
          <td>
            <a href="tel:+421904102740">
              <Phone className="fill-indigo-100 mr-2" dimension={16} />
              +421904 102 740
            </a>
          </td>
        </tr>
        <tr>
          <td rowSpan={1}>Akcie</td>
          <td>
            <a href="mailto:akcie@mcmamina.sk">
              <Mail className="fill-indigo-100 mr-2" dimension={16} />
              akcie@mcmamina.sk
            </a>
          </td>
        </tr>
        <tr>
          <td>Predpôrodné kurzy</td>
          <td>
            <a href="tel:+421950492901">
              <Phone className="fill-indigo-100 mr-2" dimension={16} />
              +421950 492 901
            </a>
          </td>
        </tr>
        <tr>
          <td rowSpan={1}>Burza</td>
          <td>
            <a href="mailto:burza@mcmamina.sk">
              <Mail className="fill-indigo-100 mr-2" dimension={16} />
              burza@mcmamina.sk
            </a>
          </td>
        </tr>
      </tbody>
    </table>
  </>
);

const Socials: FC = () => (
  <>
    <a
      href="https://www.facebook.com/MaterskeCentrumMamina/"
      aria-label="Facebook stránka Materského centra MAMINA"
      target="_blank"
      className="p-2"
      rel="noreferrer"
      title="Facebook stránka Materského centra MAMINA"
    >
      <Facebook
        className="fill-indigo-100 hover:animate-pulse"
        dimension={42}
      />
    </a>
    <a
      href="https://www.facebook.com/MaterskeCentrumMamina/"
      aria-label="Instagram stránka Materského centra MAMINA"
      target="_blank"
      rel="noreferrer"
      className="p-2"
      title="Instagram stránka Materského centra MAMINA"
    >
      <Instagram
        className="fill-indigo-100 hover:animate-pulse"
        dimension={42}
      />
    </a>
  </>
);
