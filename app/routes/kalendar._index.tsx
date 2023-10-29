import { FullWidthCard, CardContent } from "~/components/fullWidthCard";

export default function Index() {
  return (
    <FullWidthCard>
      <CardContent className="leading-normal">
        <iframe
          title="Kalendár"
          className="w-full"
          height="600"
          src="https://calendar.google.com/calendar/embed?src=n4bgt6kl18u5ueku1g38f5kic8%40group.calendar.google.com&ctz=Europe%2FPrague"
        />
      </CardContent>
    </FullWidthCard>
  );
}
