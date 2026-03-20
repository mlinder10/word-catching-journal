import { useMutation } from "@tanstack/react-query";
import { useState } from "react";

type ToggleFn = (id: string, currentState: boolean) => Promise<void>;

export function usePostAction(
  postId: string,
  initialMarked: boolean,
  initialCount: number,
  toggleFn: ToggleFn,
) {
  const [isMarked, setIsMarked] = useState(initialMarked);
  const { mutateAsync, isPending } = useMutation({
    mutationKey: [postId, isMarked],
    mutationFn: () => toggleFn(postId, isMarked),
  });

  let displayCount = initialCount;
  if (isMarked && !initialMarked) displayCount++;
  if (!isMarked && initialMarked) displayCount--;

  async function handleToggle() {
    const previousState = isMarked;
    try {
      setIsMarked(!previousState);
      await mutateAsync();
    } catch {
      setIsMarked(previousState);
    }
  }

  return { isMarked, displayCount, isPending, handleToggle };
}
