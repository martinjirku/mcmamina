interface DayDto {
  title: string;
  abbr: string;
  date: Date;
  isToday?: boolean;
  events?: number[];
}

const capitalizeString = (str: string) => {
  return str.charAt(0).toUpperCase() + str.slice(1);
};

const getDays = (): DayDto[] => {
  const todayDate = new Date();
  const day = todayDate.getDay();
  const days = [...Array(14).keys()];
  const dayDtos = days.map<DayDto>((d) => {
    const currentDay = new Date(todayDate);
    currentDay.setDate(todayDate.getDate() - day + d + 1);
    return {
      title: capitalizeString(
        currentDay.toLocaleDateString("sk-SK", { weekday: "long" }),
      ),
      abbr: capitalizeString(
        currentDay.toLocaleDateString("sk-SK", { weekday: "short" }),
      ),
      date: currentDay,
      isToday: currentDay.toDateString() === todayDate.toDateString(),
      events: d === 4 ? [1, 2] : d === 8 ? [1] : undefined,
    };
  });
  return dayDtos;
};

export const TwoWeeksCalendar = () => {
  const days = getDays();
  return (
    <div className="w-full grid grid-cols-7 gap-2 text-indigo-500">
      {days.map((day, i) => (
        <div
          key={`${i}-${day}`}
          className={`h-16 flex flex-col border hover:scale-105 scale-100 transition-transform duration-500 ease-in-out border-indigo-300 rounded-lg ${
            day.isToday
              ? "bg-indigo-400 text-indigo-100"
              : day.date.getDay() % 6 === 0 || day.date.getDay() % 7 === 0
              ? "bg-indigo-100"
              : "bg-indigo-200"
          } `}
        >
          <div className={`text-sm flex flex-grow-0 text px-1 pt-1`}>
            <span className="text-ellipsis font-semibold">{day.abbr}</span>
          </div>
          <div className={`text-sm flex flex-grow px-1 pb-1`}>
            <div className={`flex flex-grow gap-1`}>
              {day.events?.map((event, idx) => (
                <span
                  className={`w-3 h-3 rounded-full bg-slate-200 ${
                    idx === 0
                      ? "bg-pink-500"
                      : idx === 1
                      ? "bg-green-500"
                      : "bg-yellow-500"
                  }`}
                  key={`${i}-${day}-${idx}`}
                ></span>
              ))}
            </div>
            <div className={`flex-grow-0 flex justify-end self-end`}>
              <span className="border-radius rounded-full border border-indigo-300 text-center w-6 h-6 flex justify-center items-center">
                {day.date.getDate()}
              </span>
            </div>
          </div>
        </div>
      ))}
    </div>
  );
};
