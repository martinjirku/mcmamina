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
        arrow({ element: arrowRef }),
      ],
    }).then(({ x, y, middlewareData: { arrow } }) => {
      Object.assign(tooltip.style, {
        left: `${x}px`,
        top: `${y}px`,
      });
      if (arrow !== undefined) {
        Object.assign(arrowRef.style, {
          left: `${arrow.x}px`,
          top: `${arrow.y}px`,
        });
      }
    });
    const showTooltip = getShowTooltip(tooltip);
    const hideTooltip = getHideTooltip(tooltip);
    [
      ["mouseenter", showTooltip] as const,
      ["mouseleave", hideTooltip] as const,
      ["focus", showTooltip] as const,
      ["blur", hideTooltip] as const,
    ].forEach(([event, listener]) => {
      el.addEventListener(event, listener);
    });
  });
