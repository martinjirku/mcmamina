import { FC, PropsWithChildren } from "react";

interface Props {
  background?: string;
  padding?: string;
  margin?: string;
  className?: string;
}
export const FullWidthCard: FC<PropsWithChildren<Props>> = ({
  className,
  background = "bg-indigo-100",
  padding = "py-12 px-5 md:py-16 xl:py-20",
  margin = "my-12 md:my-16 xl:my-28",
  children,
}) => {
  let classes = `w-full ${background} bg-opacity-95 text-xl leading-10 flex justify-around ${padding} ${margin}`;
  if (className) {
    classes += ` ${className}`;
  }
  return <div className={classes}>{children}</div>;
};

export const CardContent: FC<PropsWithChildren<Props>> = ({
  className,
  children,
}) => {
  return (
    <div
      className={`w-full md:max-w-3xl lg:max-w-4xl ${
        className ? className : ""
      }`}
    >
      {children}
    </div>
  );
};
