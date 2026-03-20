import { useEffect } from "react";
import { toast } from "sonner";

export default function useAlert(error: Error | null, message?: string) {
  useEffect(() => {
    if (error) toast.error(message ?? error.message, { richColors: true });
  }, [message, error]);
}
