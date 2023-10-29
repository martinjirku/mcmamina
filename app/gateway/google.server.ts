import { google } from "googleapis";
import type { calendar_v3 } from "googleapis";
import { MemoryCache } from "memory-cache-node";

export const getGoogleCalendarClient = () => {
  const googleKey = process.env.GOOGLE_API_KEY ?? "";
  const calendar = google.calendar({ version: "v3", auth: googleKey });
  return calendar;
};

const itemsExpirationCheckIntervalInSecs = 10 * 60;
const maxItemCount = 10;
const cachedEvents = new MemoryCache<string, calendar_v3.Schema$Events>(
  itemsExpirationCheckIntervalInSecs,
  maxItemCount,
);

export const getEvents = async () => {
  const calendarId =
    process.env.GOOGLE_CALENDAR_ID ??
    "n4bgt6kl18u5ueku1g38f5kic8@group.calendar.google.com";
  if (cachedEvents.hasItem(calendarId)) {
    return Promise.resolve(cachedEvents.retrieveItemValue(calendarId));
  }
  const result = await getGoogleCalendarClient().events.list({
    calendarId,
    timeMin: new Date().toISOString(),
  });
  cachedEvents.storeExpiringItem(
    calendarId,
    result.data,
    itemsExpirationCheckIntervalInSecs,
  );
  return result.data;
};
