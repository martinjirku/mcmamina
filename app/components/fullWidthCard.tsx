import { FC, PropsWithChildren } from "react";

interface Props {
  className?: string;
}
export const FullWidthCard: FC<PropsWithChildren<Props>> = ({
  className,
  children,
}) => {
  let classes = `w-full bg-indigo-100 bg-opacity-80 text-xl leading-10 flex justify-around py-12 px-5 my-12 md:py-16 md:my-16 xl:py-20 xl:my-28`;
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
    <div className={`w-full md:max-w-2xl ${className ? className : ""}`}>
      {children}
    </div>
  );
};
