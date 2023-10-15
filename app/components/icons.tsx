import type { FC } from "react";

export interface SvgProps {
  className?: string | undefined;
  dimension?: number | undefined;
}
export const Phone: FC<SvgProps> = ({ className, dimension }) => (
  <svg
    xmlns="http://www.w3.org/2000/svg"
    height={dimension ?? 24}
    viewBox="0 -960 960 960"
    width={dimension ?? 24}
    className={`${className} inline-block`}
  >
    <title></title>
    <path d="M798-120q-125 0-247-54.5T329-329Q229-429 174.5-551T120-798q0-18 12-30t30-12h162q14 0 25 9.5t13 22.5l26 140q2 16-1 27t-11 19l-97 98q20 37 47.5 71.5T387-386q31 31 65 57.5t72 48.5l94-94q9-9 23.5-13.5T670-390l138 28q14 4 23 14.5t9 23.5v162q0 18-12 30t-30 12ZM241-600l66-66-17-94h-89q5 41 14 81t26 79Zm358 358q39 17 79.5 27t81.5 13v-88l-94-19-67 67ZM241-600Zm358 358Z" />
  </svg>
);

export const Mail: FC<SvgProps> = ({ className, dimension }) => (
  <svg
    xmlns="http://www.w3.org/2000/svg"
    height={dimension ?? 24}
    viewBox="0 -960 960 960"
    width={dimension ?? 24}
    className={`${className} inline-block`}
  >
    <title></title>
    <path d="M160-160q-33 0-56.5-23.5T80-240v-480q0-33 23.5-56.5T160-800h640q33 0 56.5 23.5T880-720v480q0 33-23.5 56.5T800-160H160Zm320-280L160-640v400h640v-400L480-440Zm0-80 320-200H160l320 200ZM160-640v-80 480-400Z" />
  </svg>
);

export const Bank: FC<SvgProps> = ({ className, dimension }) => (
  <svg
    xmlns="http://www.w3.org/2000/svg"
    height={dimension ?? 24}
    viewBox="0 -960 960 960"
    width={dimension ?? 24}
    className={`${className} inline-block`}
  >
    <title></title>
    <path d="M200-280v-280h80v280h-80Zm240 0v-280h80v280h-80ZM80-120v-80h800v80H80Zm600-160v-280h80v280h-80ZM80-640v-80l400-200 400 200v80H80Zm178-80h444-444Zm0 0h444L480-830 258-720Z" />
  </svg>
);

export const Home: FC<SvgProps> = ({ className, dimension }) => (
  <svg
    xmlns="http://www.w3.org/2000/svg"
    height={dimension ?? 24}
    viewBox="0 -960 960 960"
    width={dimension ?? 24}
    className={`${className} inline-block`}
  >
    <title></title>
    <path d="M240-200h120v-240h240v240h120v-360L480-740 240-560v360Zm-80 80v-480l320-240 320 240v480H520v-240h-80v240H160Zm320-350Z" />
  </svg>
);

export const Signature: FC<SvgProps> = ({ className, dimension }) => (
  <svg
    xmlns="http://www.w3.org/2000/svg"
    height={dimension ?? 24}
    viewBox="0 -960 960 960"
    width={dimension ?? 24}
    className={`${className} inline-block`}
  >
    <title></title>
    <path d="M563-491q73-54 114-118.5T718-738q0-32-10.5-47T679-800q-47 0-83 79.5T560-541q0 14 .5 26.5T563-491ZM120-120v-80h80v80h-80Zm160 0v-80h80v80h-80Zm160 0v-80h80v80h-80Zm160 0v-80h80v80h-80Zm160 0v-80h80v80h-80ZM136-280l-56-56 64-64-64-64 56-56 64 64 64-64 56 56-64 64 64 64-56 56-64-64-64 64Zm482-40q-30 0-55-11.5T520-369q-25 14-51.5 25T414-322l-28-75q28-10 53.5-21.5T489-443q-5-22-7.5-48t-2.5-56q0-144 57-238.5T679-880q52 0 85 38.5T797-734q0 86-54.5 170T591-413q7 7 14.5 10.5T621-399q26 0 60.5-33t62.5-87l73 34q-7 17-11 41t1 42q10-5 23.5-17t27.5-30l63 49q-26 36-60 58t-63 22q-21 0-37.5-12.5T733-371q-28 25-57 38t-58 13Z" />
  </svg>
);

export const Facebook: FC<SvgProps> = ({ className, dimension }) => (
  <svg
    xmlns="http://www.w3.org/2000/svg"
    height={dimension ?? 24}
    viewBox="0 0 24 24"
    width={dimension ?? 24}
    className={`${className} inline-block`}
  >
    <rect fill="none" height="24" width="24" />
    <path d="M22,12c0-5.52-4.48-10-10-10S2,6.48,2,12c0,4.84,3.44,8.87,8,9.8V15H8v-3h2V9.5C10,7.57,11.57,6,13.5,6H16v3h-2 c-0.55,0-1,0.45-1,1v2h3v3h-3v6.95C18.05,21.45,22,17.19,22,12z" />
  </svg>
);

export const Instagram: FC<SvgProps> = ({ className, dimension }) => (
  <svg
    xmlns="http://www.w3.org/2000/svg"
    viewBox="0 0 50 50"
    height={dimension ?? 24}
    width={dimension ?? 24}
    className={`${className} inline-block`}
  >
    <path d="M 16 3 C 8.83 3 3 8.83 3 16 L 3 34 C 3 41.17 8.83 47 16 47 L 34 47 C 41.17 47 47 41.17 47 34 L 47 16 C 47 8.83 41.17 3 34 3 L 16 3 z M 37 11 C 38.1 11 39 11.9 39 13 C 39 14.1 38.1 15 37 15 C 35.9 15 35 14.1 35 13 C 35 11.9 35.9 11 37 11 z M 25 14 C 31.07 14 36 18.93 36 25 C 36 31.07 31.07 36 25 36 C 18.93 36 14 31.07 14 25 C 14 18.93 18.93 14 25 14 z M 25 16 C 20.04 16 16 20.04 16 25 C 16 29.96 20.04 34 25 34 C 29.96 34 34 29.96 34 25 C 34 20.04 29.96 16 25 16 z" />
  </svg>
);
