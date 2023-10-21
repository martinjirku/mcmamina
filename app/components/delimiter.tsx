import { FC } from "react";

type Spacing = "none" | "xs" | "sm" | "md";

interface Props {
  space?: Spacing | undefined;
}

const spacing = {
  none: "my-0",
  xs: "my-1",
  sm: "my-2",
  md: "my-4",
} as Record<Spacing, string>;

export const HorizontalDelimiter: FC<Props> = ({ space = "md" }) => (
  <hr
    className={`h-px border-t-0 bg-gradient-to-r from-indigo-500 to-pink-500 ${spacing[space]}`}
  />
);
