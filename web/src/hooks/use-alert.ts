import { parseNetworkError } from "@/lib/api";
import { useEffect } from "react";
import { toast } from "sonner";

export default function useAlert(error: Error | null, message?: string) {
  useEffect(() => {
    if (error) toast.error(parseNetworkError(error), { richColors: true });
  }, [message, error]);
}
