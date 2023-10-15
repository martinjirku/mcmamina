import type { FC } from "react";

import type { SvgProps } from "./icons";

import "./animatedHamburger.css";

type AnimatedHamburgerProps = SvgProps & {
  open?: boolean | undefined;
};

export const AnimatedHamburger: FC<AnimatedHamburgerProps> = ({
  className,
  dimension,
  open,
}) => (
  <svg
    height={dimension ?? 24}
    width={dimension ?? 24}
    className={`${className ?? ""} fill-none mc-animated-hamburger ${
      open ? "active" : ""
    }`}
    viewBox="0 0 60 40"
  >
    <g transform="matrix(1,0,0,1,-389.5,-264.004)">
      <g>
        <g transform="matrix(1,0,0,1,0,5)">
          <path
            id="top"
            className="className"
            d="M390,270L420,270L420,270C420,270 420.195,250.19 405,265C389.805,279.81 390,279.967 390,279.967"
          />
        </g>
        <g transform="matrix(1,1.22465e-16,1.22465e-16,-1,0.00024296,564.935)">
          <path
            id="bottom"
            className="className"
            d="M390,270L420,270L420,270C420,270 420.195,250.19 405,265C389.805,279.81 390,279.967 390,279.967"
          />
        </g>
        <path className="className" id="middle" d="M390,284.967L420,284.967" />
      </g>
    </g>
  </svg>
);
