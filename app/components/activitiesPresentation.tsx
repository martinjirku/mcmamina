import { FC, KeyboardEventHandler, useState, useRef, useEffect } from "react";

import dancingKid from "~/images/caleb-woods-jYaImw-FQNI-unsplash.jpg";
import dielnicky from "~/images/dielnicky.jpg";
import hraveCvicenie from "~/images/hrave-cvicenie.jpg";
import tinka from "~/images/TINKA.jpg";
import classes from "./activitiesPresentation.module.css";
import { Phone, Mail, Facebook } from "./icons";
interface ImageProps {
  src: string;
  active?: boolean;
  onClick: () => void;
}
const Image: FC<ImageProps> = ({ src, active = false, onClick }) => {
  const wrapperClasses =
    "h-44 outline-none flex-grow cursor-pointer overflow-hidden";
  const bgClasses = `${classes.image} h-full w-full hover:scale-125 bg-cover transform transition duration-500`;
  const activeClasses = active ? "scale-125" : "filter grayscale grayscale-100";
  const btn = useRef<HTMLButtonElement>(null);
  return (
    <button
      ref={btn}
      className={wrapperClasses}
      onClick={onClick}
      tabIndex={-1}
    >
      <div
        className={`${bgClasses} ${activeClasses}`}
        style={{ backgroundImage: `url(${src})` }}
      ></div>
    </button>
  );
};

interface ActivityProps {
  title: string;
  children: React.ReactNode;
  contact: { email?: string; phone?: string; fb?: string };
  time: string;
}
export const ActivityContent: FC<ActivityProps> = (props) => {
  let email: React.ReactNode, phone: React.ReactNode, fb: React.ReactNode;
  if (props.contact.email) {
    email = (
      <div className="w-full">
        <a href={`mailto:${props.contact.email}`}>
          <Mail className="fill-indigo-100 mr-2" dimension={16} />
          {props.contact.email}
        </a>
      </div>
    );
  }
  if (props.contact.phone) {
    phone = (
      <div className="w-full">
        <a href={`tel:${props.contact.phone}`}>
          <Phone className="fill-indigo-100 mr-2" dimension={16} />
          {props.contact.phone}
        </a>
      </div>
    );
  }
  if (props.contact.fb) {
    fb = (
      <div className="w-full">
        <a href={props.contact.fb}>
          <Facebook className="fill-indigo-100 mr-2" dimension={16} />
          Prihlasovanie na FB
        </a>
      </div>
    );
  }
  return (
    <div className="w-full relative flex flex-row gap-4">
      <div className="flex-grow flex-shrink">
        <h1 className="text-2xl underline-offset-1 underline pb-5">
          {props.title}
        </h1>
        <p>{props.children}</p>
      </div>
      <div className="w-72 flex-grow-0 flex-shrink-0">
        <h1 className="text-lg underline-offset-1 underline pb-2">Kontakt</h1>
        {email}
        {phone}
        {fb}
        <h1 className="text-lg underline-offset-1 underline pb-2 pt-3">
          Aktivity
        </h1>
        <div className="w-full">{props.time}</div>
      </div>
    </div>
  );
};

interface ContentItemProps {
  children: React.ReactNode;
  active: boolean;
}
const Content: FC<ContentItemProps> = ({ children, active }) => {
  return (
    <div
      className={`w-full overflow-none transition-all duration-500 ease-in-out transform absolute ${
        active ? "translate-x-0 opacity-1" : "translate-x-full opacity-0"
      }`}
    >
      {children}
    </div>
  );
};

export const ActivitiesPresentation = () => {
  const [active, setActive] = useState(0);
  useEffect(() => {
    const interval = window.setInterval(() => {
      setActive((active + 1) % 4);
    }, 5e3);
    return () => {
      window.clearInterval(interval);
    };
  }, [active]);
  const images = [dielnicky, tinka, hraveCvicenie, dancingKid];
  const handleKeyDown: KeyboardEventHandler<HTMLDivElement> = (e) => {
    if (e.key === "ArrowRight") setActive((active + 1) % images.length);
    if (e.key === "ArrowLeft")
      setActive((active - 1 + images.length) % images.length);
  };
  return (
    <div className="w-full flex flex-col relative">
      <div
        role="menu"
        className="relative flex flex-row"
        tabIndex={0}
        onKeyDown={handleKeyDown}
      >
        {images.map((src, idx) => (
          <Image
            key={src}
            src={src}
            active={active === idx}
            onClick={() => setActive(idx)}
          />
        ))}
      </div>
      <div className="w-full relative  text-white text-lg my-8 h-48 overflow-hidden">
        <Content active={active === 0}>
          <ActivityContent
            title="Montessori hernička"
            contact={{
              phone: "+421 948 523 493",
              fb: "https://www.facebook.com/MaterskeCentrumMamina/",
            }}
            time="Každý piatok o 9:30."
          >
            Zážitkový a vzdelávací program pre najmenších od 2 do 4 rokov
            inšpirovaný princípmi Montessori pedagogiky.
          </ActivityContent>
        </Content>
        <Content active={active === 1}>
          <ActivityContent
            title="Angličtina s Tinkou"
            contact={{
              phone: "+421 907 948 207",
              email: "anglictinamcmamina@gmail.com",
            }}
            time="Utorok o 16:30 v 3 skupinách a štvrtok o 10:00"
          >
            Tinka vedie krúžok angličtiny hravou a prirodzenou cestou.
          </ActivityContent>
        </Content>
        <Content active={active === 2}>
          <ActivityContent
            title="Happy gym"
            contact={{
              phone: "+421 907 228 779",
              email: "happygymzv@gmail.com ",
            }}
            time="Utorok o 8:45 v troch skupinkách"
          >
            Cvičenie pre najmenších zamerané na psychomotorický, sociálny,
            citový a rozumový vývoj dieťaťa.
          </ActivityContent>
        </Content>
        <Content active={active === 3}>
          <ActivityContent
            title="Tanečno - pohybová príprava"
            contact={{
              email: "tkelement@tkelement.com",
            }}
            time="Pondelok o 16:00 v dvoch skupinkách"
          >
            Cieľom u detí je získať hravou formou - správne držanie tela,
            hudobno-pohybové cítenie, zamerať sa na rytmus, tempo, takt,
            dynamiku, frázovanie a iné.
          </ActivityContent>
        </Content>
      </div>
    </div>
  );
};
