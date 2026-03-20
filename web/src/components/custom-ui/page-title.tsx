import type { ReactNode } from "react";

type PageTitleProps = {
  title: string;
  ActionComponent?: ReactNode;
};

export default function PageTitle({ title, ActionComponent }: PageTitleProps) {
  return (
    <header className="flex justify-between p-4 border-b">
      <h2 className="font-bold text-2xl">{title}</h2>
      {ActionComponent && ActionComponent}
    </header>
  );
}
