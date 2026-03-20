import api from "@/api";
import type { Profile } from "@/types";
import { useInfiniteQuery } from "@tanstack/react-query";
import { Link } from "@tanstack/react-router";
import ProfilePicture from "@/components/custom-ui/profile-picture";
import { Separator } from "@/components/ui/separator";
import { Skeleton } from "@/components/ui/skeleton";
import useAlert from "@/hooks/use-alert";

type FollowListProps = {
  userId: string;
};

const useFollowersQuery = (userId: string) =>
  useInfiniteQuery({
    queryKey: ["followers", userId],
    queryFn: async ({ pageParam }) => api.fetchFollowers({ userId, pageParam }),
    getNextPageParam: (lastPage) => {
      if (lastPage.currentPage < lastPage.totalPages) {
        return lastPage.currentPage + 1;
      }
      return undefined;
    },
    initialPageParam: 1,
  });

export function FollowersList({ userId }: FollowListProps) {
  const { data, isPending, error } = useFollowersQuery(userId);
  useAlert(error, "Failed to load users");

  if (isPending || error) return <UsersListSkeleton />;
  return <UsersList users={data?.pages.flatMap((p) => p.users) ?? []} />;
}

const useFollowingQuery = (userId: string) =>
  useInfiniteQuery({
    queryKey: ["following", userId],
    queryFn: async ({ pageParam }) => api.fetchFollowing({ userId, pageParam }),
    getNextPageParam: (lastPage) => {
      if (lastPage.currentPage < lastPage.totalPages) {
        return lastPage.currentPage + 1;
      }
      return undefined;
    },
    initialPageParam: 1,
  });

export function FollowingList({ userId }: FollowListProps) {
  const { data, isPending, error } = useFollowingQuery(userId);
  useAlert(error, "Failed to load users");

  if (isPending || error) return <UsersListSkeleton />;
  return <UsersList users={data?.pages.flatMap((p) => p.users) ?? []} />;
}

type UsersListProps = {
  users: Profile[];
};

function UsersList({ users }: UsersListProps) {
  return (
    <ul className="space-y-2 px-4 h-full overflow-y-auto">
      {users.map((u, i) => (
        <li key={u.id} className="space-y-2">
          <Link
            to="/app/profile/$profileId"
            params={{ profileId: u.id }}
            className="flex items-center gap-2"
          >
            <ProfilePicture
              username={u.username}
              imageUrl={u.imageUrl}
              color={u.color}
              size={32}
            />
            <div>
              <p className="text-muted-foreground text-sm">@{u.username}</p>
              <p>{u.displayname}</p>
            </div>
          </Link>
          {i !== users.length - 1 && <Separator />}
        </li>
      ))}
    </ul>
  );
}

function UsersListSkeleton() {
  return (
    <ul className="space-y-2 px-4">
      {Array.from({ length: 10 }).map((_, i) => (
        <li key={i} className="space-y-2">
          <div className="flex items-center gap-2">
            <Skeleton className="rounded-full w-8 aspect-square" />
            <div className="space-y-1">
              <Skeleton className="w-20 h-4" />
              <Skeleton className="w-32 h-4" />
            </div>
          </div>
          <Separator />
        </li>
      ))}
    </ul>
  );
}
