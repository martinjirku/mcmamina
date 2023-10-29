import {
  useFloating,
  useHover,
  useFocus,
  useInteractions,
  arrow,
  FloatingArrow,
  offset,
  autoPlacement,
  safePolygon,
  shift,
} from "@floating-ui/react";
import { FC, useState, useRef, useEffect } from "react";

export interface Event {
  title: string;
  start?: string;
  end?: string;
}
interface DayDto {
  title: string;
  abbr: string;
  date: Date;
  isToday?: boolean;
  isBeforeToday?: boolean;
  events?: Event[];
}

const capitalizeString = (str: string) => {
  return str.charAt(0).toUpperCase() + str.slice(1);
};

const isSameDay = (date1: Date, date2: Date) => {
  return date1.toLocaleDateString() === date2.toLocaleDateString();
};

const getDays = (events: Event[]): DayDto[] => {
  const todayDate = new Date();
  const day = todayDate.getDay();
  const days = [...Array(15).keys()];
  const dayDtos = days.map<DayDto>((d) => {
    const currentDay = new Date(todayDate);
    currentDay.setDate(todayDate.getDate() - day + d);
    return {
      title: capitalizeString(
        currentDay.toLocaleDateString("sk-SK", { weekday: "long" }),
      ),
      abbr: capitalizeString(
        currentDay.toLocaleDateString("sk-SK", { weekday: "short" }),
      ),
      date: currentDay,
      isToday: currentDay.toDateString() === todayDate.toDateString(),
      isBeforeToday: currentDay < todayDate,
      events: events?.filter((e) =>
        e.start ? isSameDay(new Date(e.start), currentDay) : false,
      ),
    };
  });
  return dayDtos;
};

export const TwoWeeksCalendar: FC<{ events: Event[] }> = ({ events }) => {
  const days = getDays(events);

  return (
    <div className="sm:w-full scroll-m-0">
      <div className="w-full grid grid-cols-5 gap-2 text-indigo-500">
        {days.map((day, i) => (
          <Day
            key={`${i}-${day.date.toString().split(" ").join("-")}`}
            day={day}
          />
        ))}
      </div>
    </div>
  );
};

const Day: FC<{ day: DayDto }> = ({ day }) => {
  const [isOpen, setIsOpen] = useState(false);
  const arrowRef = useRef(null);
  const { refs, floatingStyles, context, update } = useFloating({
    open: isOpen,
    onOpenChange: setIsOpen,
    placement: "bottom-start",
    strategy: "absolute",
    middleware: [
      offset(7),
      autoPlacement(),
      shift(),
      arrow({ element: arrowRef }),
    ],
  });
  const hover = useHover(context, { move: true, handleClose: safePolygon() });
  const focus = useFocus(context, { visibleOnly: true });
  const { getReferenceProps, getFloatingProps } = useInteractions([
    hover,
    focus,
  ]);
  useEffect(() => {
    function handleResize() {
      update();
    }
    window.addEventListener("resize", handleResize);
    return () => {
      window.removeEventListener("resize", handleResize);
    };
  }, [update]);
  let tooltip = null;
  const hasEvents = day.events?.length ?? 0 > 0;
  if (hasEvents) {
    tooltip = (
      <div
        ref={refs.setFloating}
        className={`${
          isOpen ? "opacity-100" : "opacity-0 pointer-events-none"
        } transition-all duration-500 w-fit text-left text-sm bg-indigo-500 text-indigo-100 p-2 z-10 rounded-md`}
        style={floatingStyles}
        {...getFloatingProps()}
      >
        <h3 className="text-md font-semibold pb-2">
          {day.date?.toLocaleDateString("sk-SK", {})} - {day.title}
        </h3>
        <ul>
          {day.events?.map((v, idx) => (
            <li key={`${v.title}-${v.start}`}>
              <span
                className={`w-2 h-2 md:w-3 md:h-3 rounded-full inline-block mr-2 ${
                  idx === 0
                    ? "bg-pink-500"
                    : idx === 1
                    ? "bg-green-500"
                    : "bg-yellow-500"
                }`}
              />
              {new Date(v.start ?? "")?.toLocaleTimeString("sk-SK", {
                hour12: false,
                hour: "numeric",
                minute: "numeric",
                second: undefined,
              })}{" "}
              -{" "}
              {v.end
                ? new Date(v.end ?? "")?.toLocaleTimeString("sk-SK", {
                    hour12: false,
                    hour: "numeric",
                    minute: "numeric",
                    second: undefined,
                  })
                : ""}{" "}
              - {v.title}
            </li>
          ))}
        </ul>
        <FloatingArrow
          ref={arrowRef}
          context={context}
          className="fill-indigo-500"
        />
      </div>
    );
  }
  return (
    <>
      <div
        ref={refs.setReference}
        className={`group h-16 flex flex-col border hover:scale-105 scale-100 transition-transform duration-500 ease-in-out border-indigo-300 rounded-lg ${
          day.isToday ? "border-indigo-500" : ""
        } ${day.isBeforeToday ? "opacity-40 pointer-events-none" : ""} ${
          day.date.getDay() % 6 === 0 || day.date.getDay() % 7 === 0
            ? "bg-indigo-200"
            : ""
        } `}
        {...getReferenceProps()}
        // eslint-disable-next-line jsx-a11y/no-noninteractive-tabindex
        tabIndex={hasEvents ? 0 : undefined}
      >
        <div className={`text-sm flex flex-grow-0 text px-1 pt-1`}>
          <span className="text-ellipsis font-semibold">{day.abbr}</span>
        </div>
        <div className={`text-sm flex flex-grow px-1 pb-1`}>
          <div
            className={`flex flex-grow flex-shrink gap-1 flex-wrap content-start`}
          >
            {day.events?.map((event, idx) => (
              <span
                className={`w-2 h-2 md:w-3 md:h-3 rounded-full  ${
                  idx === 0
                    ? "bg-pink-500"
                    : idx === 1
                    ? "bg-green-500"
                    : "bg-yellow-500"
                }`}
                key={`${day}-${idx}`}
              ></span>
            ))}
          </div>
          <div
            className={`flex-grow-0 flex flex-shrink-0 justify-end self-end`}
          >
            <span className="border-radius rounded-full border border-indigo-300 text-center w-6 h-6 flex justify-center items-center">
              {day.date.getDate()}
            </span>
          </div>
        </div>
      </div>
      {tooltip}
    </>
  );
};
