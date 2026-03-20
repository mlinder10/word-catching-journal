import { useCallback, useEffect, useRef, type RefObject } from "react";

export default function useFeed(
  scrollRef: RefObject<HTMLElement | null>,
  fetchNextPage: () => void,
) {
  const endRef = useRef<HTMLLIElement>(null);

  const handleScroll = useCallback(
    (e: Event) => {
      if (!endRef.current) return;

      console.log("fetch more");
      if (
        e.target instanceof HTMLElement &&
        e.target.scrollTop + e.target.clientHeight >= endRef.current.offsetTop
      ) {
        fetchNextPage();
      }
    },
    [endRef, fetchNextPage],
  );

  useEffect(() => {
    if (!scrollRef?.current) return;

    const ref = scrollRef.current;
    ref.addEventListener("scroll", handleScroll);
    return () => ref.removeEventListener("scroll", handleScroll);
  }, [scrollRef, handleScroll]);

  return { endRef };
}
