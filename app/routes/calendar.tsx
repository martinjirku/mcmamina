import { json } from "@remix-run/node";
import type { LoaderFunctionArgs } from "@remix-run/node";
import { useLoaderData } from "@remix-run/react";

import { getGoogleCalendarClient } from "~/gateway/google.server";

// eslint-disable-next-line @typescript-eslint/no-unused-vars
export const loader = async (loader: LoaderFunctionArgs) => {
  const calendars = await getGoogleCalendarClient().events.list({
    calendarId: "n4bgt6kl18u5ueku1g38f5kic8@group.calendar.google.com",
    timeMin: new Date().toISOString(),
  });
  return json(calendars);
};

export default function CalendarPage() {
  const actionData = useLoaderData<typeof loader>();

  return (
    <div>
      <h1>Calendar</h1>
      <code>{JSON.stringify(actionData.data, undefined, 2)}</code>
    </div>
  );
}
