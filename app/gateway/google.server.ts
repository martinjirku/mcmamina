import { google } from "googleapis";

export const getGoogleCalendarClient = () => {
  const googleKey = process.env.GOOGLE_API_KEY ?? "";
  const calendar = google.calendar({ version: "v3", auth: googleKey });
  return calendar;
};
