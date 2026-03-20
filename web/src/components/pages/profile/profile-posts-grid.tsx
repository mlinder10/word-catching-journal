import api from "@/api";
import { useInfiniteQuery } from "@tanstack/react-query";
import { FeedView } from "@/components/feed-view";
import type { RefObject } from "react";

const useUserPostsQuery = (userId: string) =>
  useInfiniteQuery({
    queryKey: ["posts", userId],
    queryFn: async ({ pageParam }) =>
      api.fetchProfilePosts({ userId, pageParam }),
    getNextPageParam: (lastPage) => {
      if (lastPage.currentPage < lastPage.totalPages) {
        return lastPage.currentPage + 1;
      }
      return undefined;
    },
    initialPageParam: 1,
  });

type ProfilePostsGridProps = {
  userId: string;
  scrollRef: RefObject<HTMLElement | null>;
};

export default function ProfilePostsGrid({
  userId,
  scrollRef,
}: ProfilePostsGridProps) {
  const { data, isPending, error, isFetchingNextPage, fetchNextPage } =
    useUserPostsQuery(userId);

  return (
    <section className="flex-1 px-4">
      <FeedView
        posts={data?.pages.flatMap((p) => p.posts) ?? []}
        isLoading={isPending}
        error={error}
        isFetchingNextPage={isFetchingNextPage}
        scrollRef={scrollRef}
        fetchNextPage={fetchNextPage}
      />
    </section>
  );
}
