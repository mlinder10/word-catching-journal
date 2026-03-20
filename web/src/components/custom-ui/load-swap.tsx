import type { ReactNode } from "react";
import Spinner from "./spinner";

type LoadSwapProps = {
  isLoading: boolean;
  label?: string;
  children: ReactNode;
};

export default function LoadSwap({
  isLoading,
  label,
  children,
}: LoadSwapProps) {
  return isLoading ? (
    <>
      <Spinner />
      {label && <span>{label}</span>}
    </>
  ) : (
    children
  );
}
