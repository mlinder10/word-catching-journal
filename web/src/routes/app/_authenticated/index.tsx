import api from "@/api";
import { FeedView } from "@/components/feed-view";
import PageTitle from "@/components/custom-ui/page-title";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";
import { useInfiniteQuery } from "@tanstack/react-query";
import { createFileRoute } from "@tanstack/react-router";
import { useRef } from "react";

export const Route = createFileRoute("/app/_authenticated/")({
  component: FeedPage,
});

function FeedPage() {
  return (
    <main className="flex flex-col h-full">
      <Tabs defaultValue="recent" className="flex flex-col flex-1">
        <PageTitle
          title="Feed"
          ActionComponent={
            <TabsList>
              <TabsTrigger value="recent">Recent</TabsTrigger>
              <TabsTrigger value="following">Following</TabsTrigger>
            </TabsList>
          }
        />
        <TabsContent value="recent">
          <RecentFeed />
        </TabsContent>
        <TabsContent value="following">
          <FollowingFeed />
        </TabsContent>
      </Tabs>
    </main>
  );
}

const useRecentFeed = () =>
  useInfiniteQuery({
    queryKey: ["posts", "recent"],
    queryFn: (args) => api.fetchRecentPosts(args),
    getNextPageParam: (lastPage) => {
      if (lastPage.currentPage < lastPage.totalPages) {
        return lastPage.currentPage + 1;
      }
      return undefined;
    },
    initialPageParam: 1,
  });

function RecentFeed() {
  const scrollRef = useRef<HTMLElement | null>(null);
  const { data, isPending, error, isFetchingNextPage, fetchNextPage } =
    useRecentFeed();

  return (
    <section className="flex-1 px-4 pt-4" ref={scrollRef}>
      <FeedView
        posts={data?.pages.flatMap((p) => p.posts) ?? []}
        isLoading={isPending}
        error={error}
        scrollRef={scrollRef}
        isFetchingNextPage={isFetchingNextPage}
        fetchNextPage={fetchNextPage}
      />
    </section>
  );
}

const useFollowingFeed = () =>
  useInfiniteQuery({
    queryKey: ["posts", "following"],
    queryFn: (args) => api.fetchFollowingPosts(args),
    getNextPageParam: (lastPage) => {
      if (lastPage.currentPage < lastPage.totalPages) {
        return lastPage.currentPage + 1;
      }
      return undefined;
    },
    initialPageParam: 1,
  });

function FollowingFeed() {
  const scrollRef = useRef<HTMLElement | null>(null);
  const { data, isPending, error, isFetchingNextPage, fetchNextPage } =
    useFollowingFeed();

  return (
    <section className="flex-1 px-4 pt-4" ref={scrollRef}>
      <FeedView
        posts={data?.pages.flatMap((p) => p.posts) ?? []}
        isLoading={isPending}
        error={error}
        scrollRef={scrollRef}
        isFetchingNextPage={isFetchingNextPage}
        fetchNextPage={fetchNextPage}
      />
    </section>
  );
}
