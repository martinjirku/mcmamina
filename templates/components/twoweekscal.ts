import {
  computePosition,
  offset,
  autoPlacement,
  shift,
  arrow,
} from "@floating-ui/dom";

const getShowTooltip = (tooltip: HTMLElement) => () => {
  tooltip.classList.remove("opacity-0");
  tooltip.classList.add("opacity-100");
};
const getHideTooltip = (tooltip: HTMLElement) => () => {
  tooltip.classList.remove("opacity-100");
  tooltip.classList.add("opacity-0");
};

document
  .querySelectorAll('.two-weeks-calendar .day-wrapper[tabindex="0"]')
  .forEach((el) => {
    const identifier = new Array(...el.classList).findLast((c) =>
      c.startsWith("day-")
    );
    if (!identifier) return;
    const tooltip: HTMLElement = document.querySelector(
      `.event-tooltip.${identifier}`
    );
    if (!tooltip) return;
    const arrowRef: HTMLElement = tooltip.querySelector(".arrow");

    computePosition(el, tooltip, {
      placement: "bottom-start",
      strategy: "absolute",
      middleware: [
        offset(7),
        autoPlacement(),
        shift(),
        arrow({ element: arrowRef, padding: 3 }),
      ],
    }).then(({ x, y, middlewareData: { arrow }, placement }) => {
      Object.assign(tooltip.style, {
        left: `${x}px`,
        top: `${y}px`,
      });
      if (arrow !== undefined) {
        Object.assign(arrowRef.style, {
          left:
            placement === "right"
              ? "-3px"
              : arrow?.x != null
              ? `${arrow.x}px`
              : "",
          right: placement === "left" ? "-3px" : "",
          top:
            placement === "bottom"
              ? "-3px"
              : arrow?.y != null
              ? `${arrow.y}px`
              : "",
          bottom: placement === "top" ? "-3px" : "",
        });
      }
    });
    let isHovering = false;
    const showTooltip = () => {
      tooltip.classList.remove("opacity-0");
      tooltip.classList.remove("pointer-events-none");
      tooltip.classList.add("opacity-100");
    };
    const hideTooltip = () => {
      if (isHovering) return;
      tooltip.classList.remove("opacity-100");
      tooltip.classList.add("pointer-events-none");
      tooltip.classList.add("opacity-0");
    };
    el.addEventListener("mouseenter", () => {
      isHovering = true;
      showTooltip();
    });
    el.addEventListener("mouseleave", () => {
      isHovering = false;
      setTimeout(hideTooltip, 30);
    });
    el.addEventListener("focus", () => {
      isHovering = true;
      showTooltip();
    });
    el.addEventListener("blur", () => {
      isHovering = false;
      setTimeout(hideTooltip, 30);
    });
    tooltip.addEventListener("mouseenter", () => {
      isHovering = true;
    });
    tooltip.addEventListener("mouseleave", () => {
      isHovering = false;
      hideTooltip();
    });
  });
