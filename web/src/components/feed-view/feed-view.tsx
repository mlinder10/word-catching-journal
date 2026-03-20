import type { Post } from "@/types";
import { PostView, PostViewSkeleton } from "../post-view";
import useFeed from "./use-feed";
import type { RefObject } from "react";
import useAlert from "@/hooks/use-alert";

type FeedViewProps = {
  posts: Post[];
  isLoading: boolean;
  error: Error | null;
  scrollRef: RefObject<HTMLElement | null>;
  isFetchingNextPage: boolean;
  fetchNextPage: () => void;
};

export function FeedView({
  posts,
  isLoading,
  error,
  scrollRef,
  isFetchingNextPage,
  fetchNextPage,
}: FeedViewProps) {
  const { endRef } = useFeed(scrollRef, fetchNextPage);
  useAlert(error, "Error loading posts");

  if (isLoading || error) return <FeedViewSkeleton />;

  return (
    <ul className="gap-4 grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 pb-4">
      {posts.map((p) => (
        <li key={p.id} ref={p.id === posts.at(-1)?.id ? endRef : undefined}>
          <PostView post={p} />
        </li>
      ))}
      {isFetchingNextPage &&
        Array.from({ length: 3 + (3 - (posts.length % 3)) }).map((_, i) => (
          <PostViewSkeleton key={i} />
        ))}
    </ul>
  );
}

function FeedViewSkeleton() {
  return (
    <ul className="gap-4 grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3">
      {Array.from({ length: 9 }).map((_, i) => (
        <li key={i}>
          <PostViewSkeleton />
        </li>
      ))}
    </ul>
  );
}
