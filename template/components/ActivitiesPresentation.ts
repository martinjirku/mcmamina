import "./ActivitiesPresentation.css";

const currentClasses = ["scale-125"];
const inactiveClasses = ["filter", "grayscale", "grayscale-100"];
const imageSelector = ".iamge";
interface ActiveItem {
  el?: HTMLElement;
  [key: string | symbol]: any; // Adding an index signature
}

let activeActivity: ActiveItem = new Proxy<ActiveItem>(
  {},
  {
    get: function (target, property) {
      return target[property];
    },
    set: function (target, property, value) {
      // remove from active
      if (property === "el" && target.el !== undefined) {
        const imgEl = target.el.querySelector(imageSelector);
        if (imgEl !== null) {
          imgEl.classList.remove(...currentClasses);
          imgEl.classList.add(...inactiveClasses);
        }
      }
      // set the new active
      if (value !== undefined) {
        target.el = value;
        const imgEl = target.el.querySelector(imageSelector);
        if (imgEl !== null) {
          imgEl.classList.add(...currentClasses);
          imgEl.classList.remove(...inactiveClasses);
        }
      }
      return true;
    },
  }
);

const activitiesButtonsElement = document.querySelectorAll(
  ".activities-presentation .image-wrapper"
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
      console.log(">>>>");
      activeActivity.el = e.currentTarget as HTMLElement;
    });
  }
}
