import { isAxiosError } from "axios";

export function parseNetworkError(err: unknown) {
  if (!isAxiosError<{ error: string }>(err)) return "An unknown error occurred";
  return err.response?.data.error ?? "An unknown error occurred";
}
