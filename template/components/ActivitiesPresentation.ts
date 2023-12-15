import "./ActivitiesPresentation.css";

const currentClasses = ["scale-125"];
const inactiveClasses = ["filter", "grayscale", "grayscale-100"];
const currentContentClasses = ["translate-x-0", "opacity-1"];
const inactiveContentClasses = [
  "translate-x-full",
  "opacity-0",
  "mouse-events-none",
];
const imageSelector = ".image";
const contentsSelector = ".activities-presentation .activity-content";

const activitiesButtonsElement = document.querySelectorAll(
  ".activities-presentation .image-wrapper"
);
interface ActiveItem {
  el?: HTMLElement;
  [key: string | symbol]: any; // Adding an index signature
  timeoutID?: number;
}

let activeActivity: ActiveItem = new Proxy<ActiveItem>(
  {},
  {
    get: function (target, property) {
      return target[property];
    },
    set: function (target, property, value) {
      if (target.el === value) {
        return true;
      }
      // remove from active
      if (property === "el" && target.el !== undefined) {
        const imgEl = target.el.querySelector(imageSelector);
        if (imgEl !== null) {
          imgEl.classList.remove(...currentClasses);
          imgEl.classList.add(...inactiveClasses);
        }
        const activityID = target.el.dataset.activity as string;
        const activityContent = document.querySelector(
          `${contentsSelector}[data-activity="${activityID}"]`
        );
        activityContent?.classList.remove(...currentContentClasses);
        activityContent?.classList.add(...inactiveContentClasses);
      }
      target.el = value;
      // set the new active
      if (target.el !== undefined) {
        const imgEl = target.el.querySelector(imageSelector);
        if (imgEl !== null) {
          imgEl.classList.add(...currentClasses);
          imgEl.classList.remove(...inactiveClasses);
        }
        const activityID = target.el.dataset.activity as string;
        // let's set correct contentSelector
        const activityContent = document.querySelector(
          `${contentsSelector}[data-activity="${activityID}"]`
        );
        activityContent?.classList.add(...currentContentClasses);
        activityContent?.classList.remove(...inactiveContentClasses);
        clearTimeout(target.timeoutID);
        target.timeoutID = window.setTimeout(() => {
          activeActivity.el =
            (activeActivity.el.nextSibling as HTMLElement) ??
            (activitiesButtonsElement[0] as HTMLElement) ??
            undefined;
        }, 5e3);
      }
      return true;
    },
  }
);

for (var index in activitiesButtonsElement) {
  if (activitiesButtonsElement.hasOwnProperty(index)) {
    var element = activitiesButtonsElement[index];
    const active = element.querySelector(
      currentClasses.map((c) => `.${c}`).join(" ")
    )?.parentElement;
    if (active) {
      activeActivity.el = active;
    }
    element.addEventListener("click", (e) => {
      activeActivity.el = e.currentTarget as HTMLElement;
    });
  }
}
