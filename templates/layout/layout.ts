import "./layout.css";

const classes = [
  "h-auto",
  "translate-y-0",
  "opacity-100",
  "h-0",
  "overflow-hidden",
  "-translate-y-full",
  "opacity-0",
];
const el = document.querySelector(".mobile-menu");
const btn = el.querySelector(".hamburger");
const ul = el.querySelector("ul");
btn.addEventListener("click", () => {
  classes.forEach((c) => ul.classList.toggle(c));
  btn.classList.toggle("active");
});
