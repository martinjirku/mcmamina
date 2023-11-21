import { FC, KeyboardEventHandler, useState, useRef, useEffect } from "react";

import dancingKid from "~/images/caleb-woods-jYaImw-FQNI-unsplash.jpg";
import dielnicky from "~/images/dielnicky.jpg";
import hraveCvicenie from "~/images/hrave-cvicenie.jpg";
import tinka from "~/images/TINKA.jpg";
import classes from "./activitiesPresentation.module.css";

interface ImageProps {
  src: string;
  active?: boolean;
  onClick: () => void;
  onNext?: () => void;
  onPrev?: () => void;
}
const Image: FC<ImageProps> = ({
  src,
  active = false,
  onClick,
  onNext,
  onPrev,
}) => {
  const wrapperClasses =
    "h-44 outline-none flex-grow cursor-pointer overflow-hidden";
  const bgClasses = `${classes.image} h-full w-full hover:scale-125 bg-cover transform transition duration-500`;
  const activeClasses = active ? "scale-125" : "filter grayscale grayscale-100";
  const handleKeyDown: KeyboardEventHandler<HTMLButtonElement> = (e) => {
    if (e.key === "ArrowRight") onNext?.();
    if (e.key === "ArrowLeft") onPrev?.();
    if (e.key === "enter") onClick();
  };
  const btn = useRef<HTMLButtonElement>(null);
  useEffect(() => {
    if (active) btn.current?.focus();
  }, [active]);
  return (
    // eslint-disable-next-line jsx-a11y/no-noninteractive-element-interactions
    <button
      ref={btn}
      className={wrapperClasses}
      tabIndex={active ? 0 : -1}
      onKeyDown={handleKeyDown}
      onClick={onClick}
    >
      <div
        className={`${bgClasses} ${activeClasses}`}
        style={{ backgroundImage: `url(${src})` }}
      ></div>
    </button>
  );
};

export const ActivitiesPresentation = () => {
  const [active, setActive] = useState(0);
  let content = null;
  if (active === 0) {
    content = (
      <div className="">
        <h1 className="text-2xl underline-offset-1 underline pb-5">
          Montessori hernička
        </h1>
        <p>
          Zážitkový a vzdelávací program pre najmenších od 2 do 4 rokov
          inšpirovaný princípmi Montessori pedagogiky.
        </p>
      </div>
    );
  }
  if (active === 1) {
    content = (
      <div className="">
        <h1 className="text-2xl underline-offset-1 underline pb-5">
          Angličtina s Tinkou
        </h1>
        <p>Tinka vedie krúžok angličtiny hravou a prirodzenou cestou.</p>
      </div>
    );
  }
  if (active === 2) {
    content = (
      <div className="">
        <h1 className="text-2xl underline-offset-1 underline pb-5">
          Happy gym
        </h1>
        <p>
          Cvičenie pre najmenších každý utorok od 8:45 hod. v troch skupinkách.
          Cvičenie je zamerané na psychomotorický, sociálny, citový a rozumový
          vývoj dieťaťa.
        </p>
      </div>
    );
  }
  if (active === 3) {
    content = (
      <div className="">
        <h1 className="text-2xl underline-offset-1 underline pb-5">
          Tanečno - pohybová príprava
        </h1>
        <p>
          Cieľom u detí je získať hravou formou - správne držanie tela,
          hudobno-pohybové cítenie, zamerať sa na rytmus, tempo, takt, dynamiku,
          frázovanie a iné.
        </p>
      </div>
    );
  }
  return (
    <div className="w-full flex flex-col">
      <div className="w-full flex flex-row">
        {[dielnicky, tinka, hraveCvicenie, dancingKid].map((src, idx, arr) => (
          <Image
            key={src}
            src={src}
            active={active === idx}
            onClick={() => setActive(idx)}
            onNext={() => setActive((idx + 1) % arr.length)}
            onPrev={() => setActive((idx - 1 + arr.length) % arr.length)}
          />
        ))}
      </div>
      <div className="w-full text-white text-lg py-10">{content}</div>
    </div>
  );
};

interface ArrowProps {
  color: string;
  children: React.ReactNode;
}
const Arrow: FC<ArrowProps> = ({ color, children }) => {
  return (
    <div className={`group p-2 h-14 pr-14 ${color} cursor-pointer relative`}>
      {children}
      <div className="absolute right-2 top-2 group-hover:translate-x-7 transition-transform duration-200 ease-in-out ">
        <div className={`${color} h-10 w-10 transform rotate-45`}></div>
      </div>
    </div>
  );
};
export const ActivitiesPresentationOld = () => {
  return (
    <>
      <div className="flex flex-col relative">
        <Arrow color="bg-cyan-300">Montessori hernička</Arrow>
        <Arrow color="bg-lime-300">Angličtina s Tinkou</Arrow>
        <Arrow color="bg-orange-300">Happy Gym</Arrow>
        <Arrow color="bg-fuchsia-300">Cvičenie pre tehotné</Arrow>
      </div>
      <div>The rest</div>
    </>
  );
};
