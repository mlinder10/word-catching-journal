import { useEffect, useState } from "react";

export default function useLocalStorage<T>(key: string, initialValue: T) {
  const rawStored = localStorage.getItem(key);
  const parsedStored = JSON.parse(rawStored ?? "{}");
  const [value, setValue] = useState<T>(parsedStored ?? initialValue);

  useEffect(() => {
    localStorage.setItem(key, JSON.stringify(value));
  }, [key, value]);

  return [value, setValue] as const;
}
