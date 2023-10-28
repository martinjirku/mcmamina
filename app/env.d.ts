/// <reference types="@remix-run/dev" />
/// <reference types="@remix-run/node" />

declare global {
  interface Window {
    ENV: {
      GOOGLE_API_KEY: string;
    };
  }
}
